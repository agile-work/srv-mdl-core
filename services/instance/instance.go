package instance

import (
	"fmt"
	"net/http"

	moduleShared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-mdl-shared/db"
	shared "github.com/agile-work/srv-shared"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	sql "github.com/agile-work/srv-shared/sql-builder/db"
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
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}

	schemaCode := chi.URLParam(r, "schema_code")
	userID := r.Header.Get("userID")
	languageCode := r.Header.Get("Content-Language")

	// TODO: check user instances permission
	columns, err := getUserSchemaColumns(schemaCode, userID, languageCode)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, "Loading user fields permission", err.Error()))
		return response
	}

	instanceTable := fmt.Sprintf("%s%s", instancesTablePrefix, schemaCode)
	//TODO: Paging and order by
	statement := builder.Select("id").JSON("data", columns...).From(instanceTable)

	rows, err := sql.Query(statement)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, "Loading instances", err.Error()))
		return response
	}

	results, err := sql.MapScan(rows)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, "Parsing query rows to map", err.Error()))
		return response
	}

	response.Data = results

	return response
}

// LoadInstance return only one object from the database
func LoadInstance(r *http.Request) *moduleShared.Response {
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}

	schemaCode := chi.URLParam(r, "schema_code")
	instanceID := chi.URLParam(r, "instance_id")
	userID := r.Header.Get("userID")
	languageCode := r.Header.Get("Content-Language")

	// TODO: check user instances permission
	columns, err := getUserSchemaColumns(schemaCode, userID, languageCode)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, "Loading user fields permission", err.Error()))
		return response
	}

	instanceTable := fmt.Sprintf("%s%s", instancesTablePrefix, schemaCode)

	statement := builder.Select("id").JSON("data", columns...).From(instanceTable).Where(builder.Equal("id", instanceID))

	rows, err := sql.Query(statement)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, "Loading instances", err.Error()))
		return response
	}

	results, err := sql.MapScan(rows)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, "Parsing query rows to map", err.Error()))
		return response
	}

	response.Data = results[0]
	return response
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

func getUserSchemaColumns(schemaCode, userID, languageCode string) ([]string, error) {
	// TODO: review the need for the language_code
	statement := builder.Raw("select code from core_v_fields_by_permission where schema_code=$1 and user_id=$2 and language_code=$3", schemaCode, userID, languageCode)
	rows, err := sql.Query(statement)
	if err != nil {
		return nil, err
	}
	results, err := sql.MapScan(rows)
	if err != nil {
		return nil, err
	}
	columns := []string{}
	for _, result := range results {
		columns = append(columns, result["code"].(string))
	}
	return columns, nil
}
