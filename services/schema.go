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
	schema := models.Schema{}

	return db.Create(r, &schema, "CreateSchema", shared.TableCoreSchemas)
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
		response.Errors = append(response.Errors, moduleShared.NewResponseError(moduleShared.ErrorInsertingRecord, "InsertModuleSchema", err.Error()))

		return response
	}

	return response
}

// LoadAllModulesBySchema return all instances from the object
func LoadAllModulesBySchema(r *http.Request) *moduleShared.Response {
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}

	modules := []models.Schema{}
	schemaID := chi.URLParam(r, "schema_id")
	tblTranslationName := fmt.Sprintf("%s AS %s_name", shared.TableCoreTranslations, shared.TableCoreTranslations)
	tblTranslationDescription := fmt.Sprintf("%s AS %s_description", shared.TableCoreTranslations, shared.TableCoreTranslations)
	languageCode := r.Header.Get("Content-Language")

	statemant := builder.Select(
		"core_schemas.id",
		"core_schemas.code",
		"core_translations_name.value AS name",
		"core_translations_description.value AS description",
		"core_schemas.module",
		"core_schemas.active",
		"core_schemas.created_by",
		"core_schemas.created_at",
		"core_schemas.updated_by",
		"core_schemas.updated_at",
	).From(shared.TableCoreSchemas).Join(
		tblTranslationName, "core_translations_name.structure_id = core_schemas.id and core_translations_name.structure_field = 'name'",
	).Join(
		tblTranslationDescription, "core_translations_description.structure_id = core_schemas.id and core_translations_description.structure_field = 'description'",
	).Join(
		shared.TableCoreSchemasModels, "core_schemas_modules.module_id = core_schemas.id",
	).Where(
		builder.And(
			builder.Equal("core_schemas_modules.schema_id", schemaID),
			builder.Equal("core_translations_name.language_code", languageCode),
			builder.Equal("core_translations_description.language_code", languageCode),
		),
	)

	err := sql.QueryStruct(statemant, &modules)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(moduleShared.ErrorLoadingData, "LoadAllModulesBySchema", err.Error()))

		return response
	}

	response.Data = modules

	return response
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
		response.Errors = append(response.Errors, moduleShared.NewResponseError(moduleShared.ErrorDeletingData, "RemoveModuleFromSchema", err.Error()))

		return response
	}

	return response
}
