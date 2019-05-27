package instance

import (
	"fmt"
	"net/http"

	"github.com/agile-work/srv-shared/sql-builder/builder"

	moduleShared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-mdl-shared/db"
	"github.com/agile-work/srv-mdl-shared/models"
	shared "github.com/agile-work/srv-shared"
	sql "github.com/agile-work/srv-shared/sql-builder/db"
	"github.com/go-chi/chi"
)

// LoadAllResources return all instances from the object
func LoadAllResources(r *http.Request) *moduleShared.Response {
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}
	userID := r.Header.Get("userID")
	securityFields, err := db.GetUserAvailableFields(userID, "resources", shared.SecurityStructureField)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, "Loading user fields permission", err.Error()))
		return response
	}
	fields := []string{}
	treeJoin := make(map[string]string)
	columns := []string{}

	for _, f := range securityFields {
		if f.StructureClass == shared.FieldLookupTree {
			columns = append(columns, f.StructureCode)
			table := fmt.Sprintf("jsonb_array_elements(%s.data->'trees') %s", shared.TableCustomResources, f.StructureCode)
			treeJoin[table] = fmt.Sprintf("%s->>'field' = '%s'", f.StructureCode, f.StructureCode)
		} else {
			fields = append(fields, f.StructureCode)
		}
	}

	columns = append(columns, models.GetUserSelectableFields()...)
	columns = append(columns, shared.TableCustomResources+".id")

	on := fmt.Sprintf("%s.id = %s.parent_id", shared.TableCoreUsers, shared.TableCustomResources)
	statement := builder.Select(columns...).JSON("data", fields...).From(shared.TableCustomResources).Join(shared.TableCoreUsers, on)
	for table, on := range treeJoin {
		statement.Join(table, on)
	}

	rows, err := sql.Query(statement)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, "LoadAllResources", err.Error()))
		return response
	}

	results, err := sql.MapScan(rows)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, "LoadAllResources Parsing query rows to map", err.Error()))
		return response
	}

	response.Data = results
	return response
}

// LoadResource return one instance from the object
func LoadResource(r *http.Request) *moduleShared.Response {
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}
	userID := r.Header.Get("userID")
	resourceID := chi.URLParam(r, "resource_id")

	securityFields, err := db.GetUserAvailableFields(userID, "resources", shared.SecurityStructureField)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, "Loading user fields permission", err.Error()))
		return response
	}
	fields := []string{}
	treeJoin := make(map[string]string)
	columns := []string{}

	for _, f := range securityFields {
		if f.StructureClass == shared.FieldLookupTree {
			columns = append(columns, f.StructureCode)
			table := fmt.Sprintf("jsonb_array_elements(%s.data->'trees') %s", shared.TableCustomResources, f.StructureCode)
			treeJoin[table] = fmt.Sprintf("%s->>'field' = '%s'", f.StructureCode, f.StructureCode)
		} else {
			fields = append(fields, f.StructureCode)
		}
	}

	columns = append(columns, models.GetUserSelectableFields()...)
	columns = append(columns, shared.TableCustomResources+".id")

	on := fmt.Sprintf("%s.id = %s.parent_id", shared.TableCoreUsers, shared.TableCustomResources)
	statement := builder.Select(columns...).JSON("data", fields...).From(shared.TableCustomResources).Join(shared.TableCoreUsers, on)
	for table, on := range treeJoin {
		statement.Join(table, on)
	}
	statement.Where(builder.Equal("id", resourceID))

	rows, err := sql.Query(statement)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, "LoadResource", err.Error()))
		return response
	}

	results, err := sql.MapScan(rows)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingInstances, "LoadResource Parsing query rows to map", err.Error()))
		return response
	}

	if len(results) > 0 {
		response.Data = results[0]
	}
	return response
}

// UpdateResource update an instance from the object
func UpdateResource(r *http.Request) *moduleShared.Response {
	resourceID := chi.URLParam(r, "resource_id")

	resourceMap := map[string]interface{}{}

	response := db.GetResponse(r, &resourceMap, "UpdateResource")
	if response.Code != http.StatusOK {
		return response
	}

	// TODO: validate resource fields at resourceMap before update

	sql.UpdateStructToJSON(resourceID, "data", "cst_resources", &resourceMap)
	return response
}
