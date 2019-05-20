package services

import (
	"fmt"
	"net/http"

	moduleShared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-mdl-shared/db"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/agile-work/srv-mdl-core/models"
)

const (
	instancesTablePrefix string = "cst_"
)

// CreateInstance persists the request body creating a new object in the database
func CreateInstance(r *http.Request) *moduleShared.Response {
	instance := models.Instance{}
	schemaCode := chi.URLParam(r, "schema_code")

	return db.Create(r, &instance, "CreateInstance", fmt.Sprintf("%s%s", instancesTablePrefix, schemaCode))
}

// LoadAllInstances return all instances from the object
func LoadAllInstances(r *http.Request) *moduleShared.Response {
	instances := []models.Instance{}
	schemaCode := chi.URLParam(r, "schema_code")

	return db.Load(r, &instances, "LoadAllInstances", fmt.Sprintf("%s%s", instancesTablePrefix, schemaCode), nil)
}

// LoadInstance return only one object from the database
func LoadInstance(r *http.Request) *moduleShared.Response {
	instance := models.Instance{}
	schemaCode := chi.URLParam(r, "schema_code")
	instanceID := chi.URLParam(r, "instance_id")
	table := fmt.Sprintf("%s%s", instancesTablePrefix, schemaCode)
	instanceIDColumn := fmt.Sprintf("%s.id", table)
	condition := builder.Equal(instanceIDColumn, instanceID)

	return db.Load(r, &instance, "LoadInstance", table, condition)
}

// UpdateInstance updates object data in the database
func UpdateInstance(r *http.Request) *moduleShared.Response {
	schemaCode := chi.URLParam(r, "schema_code")
	instanceID := chi.URLParam(r, "instance_id")
	table := fmt.Sprintf("%s%s", instancesTablePrefix, schemaCode)
	instanceIDColumn := fmt.Sprintf("%s.id", table)
	condition := builder.Equal(instanceIDColumn, instanceID)

	instance := models.Instance{
		ID: instanceID,
	}

	return db.Update(r, &instance, "UpdateInstance", table, condition)
}

// DeleteInstance deletes object from the database
func DeleteInstance(r *http.Request) *moduleShared.Response {
	schemaCode := chi.URLParam(r, "schema_code")
	instanceID := chi.URLParam(r, "instance_id")
	table := fmt.Sprintf("%s%s", instancesTablePrefix, schemaCode)
	instanceIDColumn := fmt.Sprintf("%s.id", table)
	condition := builder.Equal(instanceIDColumn, instanceID)

	return db.Remove(r, "DeleteInstance", table, condition)
}
