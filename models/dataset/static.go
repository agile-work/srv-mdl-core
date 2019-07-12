package dataset

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/agile-work/srv-mdl-shared/models/customerror"
	"github.com/agile-work/srv-mdl-shared/models/translation"

	"github.com/agile-work/srv-shared/constants"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/agile-work/srv-shared/sql-builder/db"
	"github.com/agile-work/srv-shared/util"
)

// StaticDefinition define specific fields for the dataset definition
type StaticDefinition struct {
	Options   map[string]Option `json:"options"`
	OrderType string            `json:"order_type,omitempty"`
	Order     []string          `json:"order,omitempty"`
}

// Option defines the struct of a static option
type Option struct {
	Code      string                  `json:"code" updatable:"false" validate:"required"`
	Label     translation.Translation `json:"label" validate:"required"`
	Active    bool                    `json:"active"`
	CreatedBy string                  `json:"created_by"`
	CreatedAt time.Time               `json:"created_at"`
	UpdatedBy string                  `json:"updated_by"`
	UpdatedAt time.Time               `json:"updated_at"`
}

// AddOption insert a new option to the static dataset
func (def *StaticDefinition) AddOption(trs *db.Transaction, option *Option, datasetCode string) error {
	ds := Dataset{Code: datasetCode}
	if err := ds.Load(); err != nil {
		return customerror.New(http.StatusBadRequest, "load dataset", err.Error())
	}

	if ds.Type != constants.DatasetStatic {
		return customerror.New(http.StatusBadRequest, "load dataset", "invalid dataset type")
	}

	genericDef, err := ds.getDefinition()
	if err != nil {
		return customerror.New(http.StatusBadRequest, "get definition", err.Error())
	}
	def = genericDef.(*StaticDefinition)
	if util.Contains(def.Order, option.Code) {
		return customerror.New(http.StatusBadRequest, "add option", "code already exists")
	}

	def.Options[option.Code] = *option
	def.Order = append(def.Order, option.Code)

	defJSON, err := json.Marshal(def)
	if err != nil {
		return customerror.New(http.StatusBadRequest, "marshal definition", err.Error())
	}

	if err := trs.Exec(builder.Update(constants.TableCoreDatasets, "definitions").Values(defJSON).Where(builder.Equal("code", datasetCode))); err != nil {
		return customerror.New(http.StatusBadRequest, "add option", err.Error())
	}

	return nil
}

// UpdateOption change option values in dataset
func (def *StaticDefinition) UpdateOption(trs *db.Transaction, optionCode, datasetCode string, columns map[string]interface{}) error {
	for col, value := range columns {
		path := fmt.Sprintf("{options, %s, %s}", optionCode, col)
		if err := db.UpdateJSONAttributeTx(trs.Tx, constants.TableCoreDatasets, "definitions", path, value, builder.Equal("code", datasetCode)); err != nil {
			return err
		}
	}
	return nil
}

// UpdateOptionsOrder update options order
func (def *StaticDefinition) UpdateOptionsOrder(trs *db.Transaction, datasetCode string, order []string) error {
	ds := Dataset{Code: datasetCode}
	if err := ds.Load(); err != nil {
		return customerror.New(http.StatusBadRequest, "load dataset", err.Error())
	}

	if ds.Type != constants.DatasetStatic {
		return customerror.New(http.StatusBadRequest, "load dataset", "invalid dataset type")
	}

	genericDef, err := ds.getDefinition()
	if err != nil {
		return customerror.New(http.StatusBadRequest, "get definition", err.Error())
	}
	def = genericDef.(*StaticDefinition)

	if len(def.Options) != len(order) {
		return customerror.New(http.StatusBadRequest, "update order", "invalid request order")
	}

	for key := range def.Options {
		if !util.Contains(order, key) {
			return customerror.New(http.StatusBadRequest, "update order", fmt.Sprintf("code %s not found", key))
		}
	}

	if err := db.UpdateJSONAttributeTx(trs.Tx, constants.TableCoreDatasets, "definitions", "{order}", order, builder.Equal("code", datasetCode)); err != nil {
		return customerror.New(http.StatusBadRequest, "update order", err.Error())
	}
	return nil
}

// DeleteOption delete a option from dataset
func (def *StaticDefinition) DeleteOption(trs *db.Transaction, optionCode, datasetCode string) error {
	ds := Dataset{Code: datasetCode}
	if err := ds.Load(); err != nil {
		return customerror.New(http.StatusBadRequest, "load dataset", err.Error())
	}

	if ds.Type != constants.DatasetStatic {
		return customerror.New(http.StatusBadRequest, "load dataset", "invalid dataset type")
	}

	genericDef, err := ds.getDefinition()
	if err != nil {
		return customerror.New(http.StatusBadRequest, "get definition", err.Error())
	}
	def = genericDef.(*StaticDefinition)
	if !util.Contains(def.Order, optionCode) {
		return customerror.New(http.StatusBadRequest, "delete option", "code not found")
	}

	delete(def.Options, optionCode)
	def.Order = util.Remove(def.Order, optionCode)

	defJSON, err := json.Marshal(def)
	if err != nil {
		return customerror.New(http.StatusBadRequest, "marshal definition", err.Error())
	}

	if err := trs.Exec(builder.Update(constants.TableCoreDatasets, "definitions").Values(defJSON).Where(builder.Equal("code", datasetCode))); err != nil {
		return customerror.New(http.StatusBadRequest, "delete option", err.Error())
	}

	return nil
}

func (def *StaticDefinition) getInstances() []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, code := range def.Order {
		item := map[string]interface{}{}
		item["code"] = code
		if option, ok := def.Options[code]; ok && option.Active {
			item["label"] = option.Label
		}
		result = append(result, item)
	}
	return result
}
