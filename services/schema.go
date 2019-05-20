package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/agile-work/srv-mdl-shared/db"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	sql "github.com/agile-work/srv-shared/sql-builder/db"
	"github.com/go-chi/chi"

	"github.com/agile-work/srv-mdl-core/models"
	moduleShared "github.com/agile-work/srv-mdl-shared"
	shared "github.com/agile-work/srv-shared"
)

// CreateSchema persists the request body creating a new object in the database
func CreateSchema(r *http.Request) *moduleShared.Response {
	schema := models.Schema{
		Status: shared.SchemaStatusProcessing,
	}
	response := db.Create(r, &schema, "CreateSchema", shared.TableCoreSchemas)

	params := map[string]interface{}{
		"schema_id":   schema.ID,
		"schema_code": schema.Code,
	}

	id, err := moduleShared.CreateJobInstance(schema.CreatedBy, shared.JobSystemCreateSchema, params)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorJobExecution, "CreateSchema job execution", err.Error()))

		return response
	}

	schema.JobID = id
	schemaIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchemas)
	condition := builder.Equal(schemaIDColumn, schema.ID)

	err = sql.UpdateStruct(shared.TableCoreSchemas, &schema, condition, "job_id")
	if err != nil {
		response := &moduleShared.Response{}
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorInsertingRecord, "CreateSchema update job id", err.Error()))

		return response
	}

	return response
}

// LoadAllSchemas return all instances from the object
func LoadAllSchemas(r *http.Request) *moduleShared.Response {
	schemas := []models.Schema{}

	return db.Load(r, &schemas, "LoadAllSchemas", shared.TableCoreSchemas, nil)
}

// LoadSchema return only one object from the database
func LoadSchema(r *http.Request) *moduleShared.Response {
	schema := models.Schema{}
	schemaID := chi.URLParam(r, "schema_id")
	schemaIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchemas)
	condition := builder.Equal(schemaIDColumn, schemaID)

	return db.Load(r, &schema, "LoadSchema", shared.TableCoreSchemas, condition)
}

// UpdateSchema updates object data in the database
func UpdateSchema(r *http.Request) *moduleShared.Response {
	schemaID := chi.URLParam(r, "schema_id")
	schemaIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchemas)
	condition := builder.Equal(schemaIDColumn, schemaID)
	schema := models.Schema{
		ID: schemaID,
	}

	return db.Update(r, &schema, "UpdateSchema", shared.TableCoreSchemas, condition)
}

// DeleteSchema deletes object from the database
func DeleteSchema(r *http.Request) *moduleShared.Response {
	schemaID := chi.URLParam(r, "schema_id")
	schemaIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchemas)
	condition := builder.Equal(schemaIDColumn, schemaID)

	return db.Remove(r, "DeleteSchema", shared.TableCoreSchemas, condition)
}

// CallDeleteSchema deletes object from the database
func CallDeleteSchema(r *http.Request) *moduleShared.Response {
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}
	schema := models.Schema{
		Status: shared.SchemaStatusDeleting,
		Active: false,
	}
	schemaID := chi.URLParam(r, "schema_id")
	schemaIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchemas)
	condition := builder.Equal(schemaIDColumn, schemaID)

	err := sql.UpdateStruct(shared.TableCoreSchemas, &schema, condition, "status", "active")
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorInsertingRecord, "CallDeleteSchema update status", err.Error()))

		return response
	}

	params := map[string]interface{}{
		"schema_id": schemaID,
	}

	_, err = moduleShared.CreateJobInstance(schema.UpdatedBy, shared.JobSystemDeleteSchema, params)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorJobExecution, "DeleteSchema job execution", err.Error()))

		return response
	}

	return response
}

// InsertModuleInSchema persists the request creating a new object in the database
func InsertModuleInSchema(r *http.Request) *moduleShared.Response {
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}

	schemaID := chi.URLParam(r, "schema_id")
	moduleID := chi.URLParam(r, "module_id")

	userID := r.Header.Get("userID")
	now := time.Now()

	statemant := builder.Insert(
		shared.TableCoreSchemasModels,
		"schema_id",
		"module_id",
		"created_by",
		"created_at",
		"updated_by",
		"updated_at",
	).Values(
		schemaID,
		moduleID,
		userID,
		now,
		userID,
		now,
	)

	err := sql.Exec(statemant)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorInsertingRecord, "InsertModuleSchema", err.Error()))

		return response
	}

	return response
}

// LoadAllModulesBySchema return all instances from the object
func LoadAllModulesBySchema(r *http.Request) *moduleShared.Response {
	modules := []models.Schema{}
	schemaID := chi.URLParam(r, "schema_id")
	schemaIDColumn := fmt.Sprintf("%s.schema_id", shared.ViewCoreSchModules)
	languageCode := r.Header.Get("Content-Language")
	languageCodeColumn := fmt.Sprintf("%s.language_code", shared.ViewCoreSchModules)
	condition := builder.And(
		builder.Equal(schemaIDColumn, schemaID),
		builder.Equal(languageCodeColumn, languageCode),
	)

	return db.Load(r, &modules, "LoadAllModulesBySchema", shared.ViewCoreSchModules, condition)
}

// RemoveModuleFromSchema deletes object from the database
func RemoveModuleFromSchema(r *http.Request) *moduleShared.Response {
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}

	schemaID := chi.URLParam(r, "schema_id")
	moduleID := chi.URLParam(r, "module_id")

	statemant := builder.Delete(shared.TableCoreSchemasModels).Where(
		builder.And(
			builder.Equal("schema_id", schemaID),
			builder.Equal("module_id", moduleID),
		),
	)

	err := sql.Exec(statemant)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorDeletingData, "RemoveModuleFromSchema", err.Error()))

		return response
	}

	return response
}
