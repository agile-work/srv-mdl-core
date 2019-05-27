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

	userID := r.Header.Get("userID")
	schemaCode := chi.URLParam(r, "schema_code")

	securityFields, err := db.GetUserAvailableFields(userID, schemaCode, shared.SecurityStructureField)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, "Loading user fields permission", err.Error()))
		return response
	}
	fields := []string{}
	treeJoin := make(map[string]string)
	columns := []string{}

	instanceTable := fmt.Sprintf("%s%s", instancesTablePrefix, schemaCode)

	for _, f := range securityFields {
		if f.StructureType == shared.FieldLookupTree {
			columns = append(columns, f.StructureCode)
			table := fmt.Sprintf("jsonb_array_elements(%s.data->'trees') %s", shared.TableCustomResources, f.StructureCode)
			treeJoin[table] = fmt.Sprintf("on %s->>'field' = '%s'", f.StructureCode, f.StructureCode)
		} else {
			fields = append(fields, f.StructureCode)
		}
	}

	columns = append(columns, instanceTable+".id")
	statement := builder.Select(columns...).JSON("data", fields...).From(instanceTable)
	for table, on := range treeJoin {
		statement.Join(table, on)
	}

	rows, err := sql.Query(statement)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, "LoadAllInstances", err.Error()))
		return response
	}

	results, err := sql.MapScan(rows)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, "LoadAllInstances Parsing query rows to map", err.Error()))
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

	userID := r.Header.Get("userID")
	schemaCode := chi.URLParam(r, "schema_code")

	securityFields, err := db.GetUserAvailableFields(userID, schemaCode, shared.SecurityStructureField)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, "Loading user fields permission", err.Error()))
		return response
	}
	fields := []string{}
	treeJoin := make(map[string]string)
	columns := []string{}

	instanceTable := fmt.Sprintf("%s%s", instancesTablePrefix, schemaCode)

	for _, f := range securityFields {
		if f.StructureType == shared.FieldLookupTree {
			columns = append(columns, f.StructureCode)
			table := fmt.Sprintf("jsonb_array_elements(%s.data->'trees') %s", shared.TableCustomResources, f.StructureCode)
			treeJoin[table] = fmt.Sprintf("on %s->>'field' = '%s'", f.StructureCode, f.StructureCode)
		} else {
			fields = append(fields, f.StructureCode)
		}
	}

	columns = append(columns, instanceTable+".id")
	statement := builder.Select(columns...).JSON("data", fields...).From(instanceTable)
	for table, on := range treeJoin {
		statement.Join(table, on)
	}

	rows, err := sql.Query(statement)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, "LoadInstance", err.Error()))
		return response
	}

	results, err := sql.MapScan(rows)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, "LoadInstance Parsing query rows to map", err.Error()))
		return response
	}

	if len(results) > 0 {
		response.Data = results[0]
	}
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
