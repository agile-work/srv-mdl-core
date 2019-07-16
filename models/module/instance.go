package module

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/customerror"
	"github.com/agile-work/srv-shared/constants"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// Instance defines each service instance for this module
type Instance struct {
	Host string `json:"host" validate:"required"`
	Port int    `json:"port" validate:"required"`
}

// Add insert a new instance to serve request for this module
func (i *Instance) Add(trs *db.Transaction, moduleCode string) error {
	mdl := Module{Code: moduleCode}
	if err := mdl.Load(); err != nil {
		return customerror.New(http.StatusBadRequest, "load dataset", err.Error())
	}

	if getIndex(*i, mdl.Definitions.Instances) != -1 {
		return customerror.New(http.StatusBadGateway, "add", "instance already exists")
	}

	mdl.Definitions.Instances = append(mdl.Definitions.Instances, *i)

	if err := db.UpdateJSONAttributeTx(trs.Tx, constants.TableCoreModules, "definitions", "{instances}", mdl.Definitions.Instances, builder.Equal("code", moduleCode)); err != nil {
		return err
	}

	return nil
}

// Delete removes a instance from this module
func (i *Instance) Delete(trs *db.Transaction, moduleCode string) error {
	mdl := Module{Code: moduleCode}
	if err := mdl.Load(); err != nil {
		return customerror.New(http.StatusBadRequest, "load dataset", err.Error())
	}

	if len(mdl.Definitions.Instances) <= 1 {
		return customerror.New(http.StatusForbidden, "delete", "at least one instance is required")
	}

	index := getIndex(*i, mdl.Definitions.Instances)
	if index == -1 {
		return customerror.New(http.StatusNotFound, "delete", "instance not found")
	}

	mdl.Definitions.Instances = append(mdl.Definitions.Instances[:index], mdl.Definitions.Instances[index+1:]...)

	if err := db.UpdateJSONAttributeTx(trs.Tx, constants.TableCoreModules, "definitions", "{instances}", mdl.Definitions.Instances, builder.Equal("code", moduleCode)); err != nil {
		return err
	}

	return nil
}

func getIndex(i Instance, instances []Instance) int {
	for index, instance := range instances {
		if instance.Host == i.Host && instance.Port == i.Port {
			return index
		}
	}
	return -1
}
