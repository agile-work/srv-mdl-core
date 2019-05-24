package services

import (
	"fmt"
	"net/http"

	"github.com/agile-work/srv-mdl-shared/db"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/agile-work/srv-mdl-core/models"
	moduleShared "github.com/agile-work/srv-mdl-shared"
	shared "github.com/agile-work/srv-shared"
)

// CreateField persists the request body creating a new object in the database
func CreateField(r *http.Request) *moduleShared.Response {
	schemaID := chi.URLParam(r, "schema_id")
	field := models.Field{
		SchemaID: schemaID,
	}

	return db.Create(r, &field, "CreateField", shared.TableCoreSchemaFields)
}

// LoadAllFields return all instances from the object
func LoadAllFields(r *http.Request) *moduleShared.Response {
	fields := []models.Field{}
	schemaID := chi.URLParam(r, "schema_id")
	schemaIDColumn := fmt.Sprintf("%s.schema_id", shared.TableCoreSchemaFields)
	condition := builder.Equal(schemaIDColumn, schemaID)

	return db.Load(r, &fields, "LoadAllFields", shared.TableCoreSchemaFields, condition)
}

// LoadField return only one object from the database
func LoadField(r *http.Request) *moduleShared.Response {
	field := models.Field{}
	fieldID := chi.URLParam(r, "field_id")
	fieldIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchemaFields)
	condition := builder.Equal(fieldIDColumn, fieldID)

	return db.Load(r, &field, "LoadField", shared.TableCoreSchemaFields, condition)
}

// UpdateField updates object data in the database
func UpdateField(r *http.Request) *moduleShared.Response {
	fieldID := chi.URLParam(r, "field_id")
	fieldIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchemaFields)
	condition := builder.Equal(fieldIDColumn, fieldID)
	field := models.Field{
		ID: fieldID,
	}

	return db.Update(r, &field, "UpdateField", shared.TableCoreSchemaFields, condition)
}

// DeleteField deletes object from the database
func DeleteField(r *http.Request) *moduleShared.Response {
	fieldID := chi.URLParam(r, "field_id")
	fieldIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchemaFields)
	condition := builder.Equal(fieldIDColumn, fieldID)

	return db.Remove(r, "DeleteField", shared.TableCoreSchemaFields, condition)
}

// CreateFieldValidation persists the request body creating a new object in the database
func CreateFieldValidation(r *http.Request) *moduleShared.Response {
	fieldValidation := models.FieldValidation{}

	return db.Create(r, &fieldValidation, "CreateFieldValidation", shared.TableCoreSchemaFldValidations)
}

// LoadAllFieldValidations return all instances from the object
func LoadAllFieldValidations(r *http.Request) *moduleShared.Response {
	fieldValidations := []models.FieldValidation{}
	fieldID := chi.URLParam(r, "field_id")
	fieldIDColumn := fmt.Sprintf("%s.field_id", shared.TableCoreSchemaFldValidations)
	condition := builder.Equal(fieldIDColumn, fieldID)

	return db.Load(r, &fieldValidations, "LoadAllFieldValidations", shared.TableCoreSchemaFldValidations, condition)
}

// LoadFieldValidation return only one object from the database
func LoadFieldValidation(r *http.Request) *moduleShared.Response {
	fieldValidation := models.FieldValidation{}
	fieldValidationID := chi.URLParam(r, "field_validation_id")
	fieldValidationIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchemaFldValidations)
	condition := builder.Equal(fieldValidationIDColumn, fieldValidationID)

	return db.Load(r, &fieldValidation, "LoadFieldValidation", shared.TableCoreSchemaFldValidations, condition)
}

// UpdateFieldValidation updates object data in the database
func UpdateFieldValidation(r *http.Request) *moduleShared.Response {
	fieldValidationID := chi.URLParam(r, "field_validation_id")
	fieldValidationIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchemaFldValidations)
	condition := builder.Equal(fieldValidationIDColumn, fieldValidationID)
	fieldValidation := models.FieldValidation{
		ID: fieldValidationID,
	}

	return db.Update(r, &fieldValidation, "UpdateFieldValidation", shared.TableCoreSchemaFldValidations, condition)
}

// DeleteFieldValidation deletes object from the database
func DeleteFieldValidation(r *http.Request) *moduleShared.Response {
	fieldValidationID := chi.URLParam(r, "field_validation_id")
	fieldValidationIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchemaFldValidations)
	condition := builder.Equal(fieldValidationIDColumn, fieldValidationID)

	return db.Remove(r, "DeleteFieldValidation", shared.TableCoreSchemaFldValidations, condition)
}
