package instance

import (
	"fmt"
	"net/http"
	"reflect"

	moduleShared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-mdl-shared/db"
	shared "github.com/agile-work/srv-shared"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	sql "github.com/agile-work/srv-shared/sql-builder/db"
	"github.com/go-chi/chi"
	"github.com/tidwall/gjson"

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
	userID := r.Header.Get("userID")
	schemaCode := chi.URLParam(r, "schema_code")

	return loadInstances(userID, schemaCode, "LoadAllInstances", "")
}

// LoadInstance return only one object from the database
func LoadInstance(r *http.Request) *moduleShared.Response {
	userID := r.Header.Get("userID")
	schemaCode := chi.URLParam(r, "schema_code")
	instanceID := chi.URLParam(r, "instance_id")

	response := loadInstances(userID, schemaCode, "LoadInstance", instanceID)

	dataType := reflect.TypeOf(response.Data).Kind()
	if dataType == reflect.Slice {
		dataValue := reflect.ValueOf(response.Data)
		if dataValue.Len() > 0 {
			response.Data = dataValue.Index(0).Interface()
		}
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

func loadInstances(userID, schemaCode, scope, instanceID string) *moduleShared.Response {
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}

	securityFields, err := db.GetUserAvailableFields(userID, schemaCode, shared.SecurityStructureField, instanceID)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, fmt.Sprintf("%s loading user fields permission", scope), err.Error()))
		return response
	}

	securityInstances, err := db.GetUserInstancePermissions(userID, schemaCode)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, fmt.Sprintf("%s loading user instance permissions", scope), err.Error()))
		return response
	}

	treeFields, err := db.GetTreeSecurityFieldsFromSchema(schemaCode)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, fmt.Sprintf("%s loading tree fields from schema", scope), err.Error()))
		return response
	}

	fields := []string{}
	treeJoin := make(map[string]string)
	columns := []string{}

	instanceTable := fmt.Sprintf("%s%s", instancesTablePrefix, schemaCode)

	for _, f := range securityFields {
		if f.StructureClass == shared.FieldLookup &&
			gjson.Get(string(f.StructureDefinitions), "definitions.lookup_type").String() == shared.FieldLookupTree {
			columns = append(columns, f.StructureCode)
			table := fmt.Sprintf("jsonb_array_elements(%s.data->'trees') AS %s", instanceTable, f.StructureCode)
			treeJoin[table] = fmt.Sprintf("%s->>'field' = '%s'", f.StructureCode, f.StructureCode)
		} else {
			fields = append(fields, f.StructureCode)
		}
	}

	columns = append(columns, fmt.Sprintf("%s.id", instanceTable))
	statement := builder.Select(columns...).JSON("data", fields...).From(instanceTable)

	for table, on := range treeJoin {
		statement.LeftJoin(table, on)
	}

	conditions := []builder.Builder{}
	hasGroupGlobal := false

LoopField:
	for _, rowField := range treeFields {
		for _, rowSecurity := range securityInstances {
			switch rowSecurity.Scope {
			case shared.SecurityPermissionScopeGroup:
				hasGroupGlobal = true
				break LoopField
			case shared.SecurityPermissionScopeGroupUnit:
			case shared.SecurityPermissionScopeUnit:
				if rowField["tree"] == rowSecurity.TreeCode {
					conditions = append(
						conditions,
						builder.Raw(
							fmt.Sprintf(`
								'%s' @> (
									SELECT
										unit.path
									FROM
										core_tree_units AS unit
									JOIN 
										core_trees AS tree
									ON
										tree.id = unit.tree_id
									WHERE
										unit.code = %s->>'tree_unit'
										AND tree.code = %s->>'tree'
								)`,
								rowSecurity.TreeUnitPath,
								rowField["code"],
								rowField["code"],
							),
						),
					)
				}
			}
		}
	}

	instanceIDSchemaCondition := builder.Equal("1", 1)
	if !hasGroupGlobal {
		instanceIDCondition := ""
		if instanceID != "" {
			instanceIDCondition = fmt.Sprintf("AND intp.instance_id = '%s'", instanceID)
			instanceIDSchemaCondition = builder.Equal(fmt.Sprintf("%s.id", instanceTable), instanceID)
		}
		conditions = append(
			conditions,
			builder.Raw(
				fmt.Sprintf(`
					%s.id IN (
						SELECT
							intp.instance_id
						FROM
							core_instance_premissions AS intp
						WHERE
							intp.user_id = '%s'
							AND intp.instance_type = '%s'
							%s
					)`,
					instanceTable,
					userID,
					schemaCode,
					instanceIDCondition,
				),
			),
		)
	}

	if len(conditions) > 0 {
		statement.Where(
			builder.And(
				instanceIDSchemaCondition,
				builder.Or(
					conditions...,
				),
			),
		)
	} else {
		statement.Where(instanceIDSchemaCondition)
	}

	rows, err := sql.Query(statement)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, fmt.Sprintf("%s", scope), err.Error()))
		return response
	}

	results, err := sql.MapScan(rows)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, fmt.Sprintf("%s parsing query rows to map", scope), err.Error()))
		return response
	}

	response.Data = results
	return response
}
