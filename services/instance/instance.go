package instance

import (
	"fmt"
	"net/http"

	moduleShared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-mdl-shared/db"
	moduleSharedModels "github.com/agile-work/srv-mdl-shared/models"
	shared "github.com/agile-work/srv-shared"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	sql "github.com/agile-work/srv-shared/sql-builder/db"
	"github.com/go-chi/chi"

	"github.com/agile-work/srv-mdl-core/models"
)

// CreateSchemaInstance persists the request body creating a new object in the database
func CreateSchemaInstance(r *http.Request) *moduleShared.Response {
	instance := models.Instance{}
	schemaCode := chi.URLParam(r, "schema_code")

	return db.Create(r, &instance, "CreateInstance", fmt.Sprintf("%s%s", shared.InstancesTablePrefix, schemaCode))
}

// LoadAllSchemaInstances return all instances from the object
func LoadAllSchemaInstances(r *http.Request) *moduleShared.Response {
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}

	schemaCode := chi.URLParam(r, "schema_code")
	user := moduleSharedModels.User{}
	err := user.Load(r.Header.Get("userID"))
	if err != nil {
		response.Code = http.StatusForbidden
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, "LoadAllInstances Load user", err.Error()))
		return response
	}

	results, err := user.GetSecurityInstances(schemaCode)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, "LoadAllInstances LoadAllInstances loading instances", err.Error()))
		return response
	}

	response.Data = results

	return response
}

// LoadSchemaInstance return only one object from the database
func LoadSchemaInstance(r *http.Request) *moduleShared.Response {
	return nil
}

// UpdateSchemaInstance updates object data in the database
func UpdateSchemaInstance(r *http.Request) *moduleShared.Response {
	schemaCode := chi.URLParam(r, "schema_code")
	instanceID := chi.URLParam(r, "instance_id")
	table := fmt.Sprintf("%s%s", shared.InstancesTablePrefix, schemaCode)
	instanceIDColumn := fmt.Sprintf("%s.id", table)
	condition := builder.Equal(instanceIDColumn, instanceID)

	instance := models.Instance{
		ID: instanceID,
	}

	return db.Update(r, &instance, "UpdateInstance", table, condition)
}

// DeleteSchemaInstance deletes object from the database
func DeleteSchemaInstance(r *http.Request) *moduleShared.Response {
	schemaCode := chi.URLParam(r, "schema_code")
	instanceID := chi.URLParam(r, "instance_id")
	table := fmt.Sprintf("%s%s", shared.InstancesTablePrefix, schemaCode)
	instanceIDColumn := fmt.Sprintf("%s.id", table)
	condition := builder.Equal(instanceIDColumn, instanceID)

	return db.Remove(r, "DeleteInstance", table, condition)
}

// LoadLookupInstance return only one instance from the object in the database
func LoadLookupInstance(r *http.Request) *moduleShared.Response {
	moduleSharedModels.TranslationFieldsRequestLanguageCode = r.Header.Get("Content-Language")
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}

	lookup := &models.Lookup{}
	lookupCode := chi.URLParam(r, "lookup_code")
	lookupIDColumn := fmt.Sprintf("%s.code", shared.TableCoreLookups)
	condition := builder.Equal(lookupIDColumn, lookupCode)

	err := sql.SelectStruct(shared.TableCoreLookups, lookup, condition)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingData, "LoadLookupInstance", err.Error()))

		return response
	}

	results, err := lookup.GetInstances()
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingData, "LoadLookupInstance load instances", err.Error()))

		return response
	}

	response.Data = results
	return response
}
