package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-core/models"
	moduleShared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-mdl-shared/db"
	sharedModels "github.com/agile-work/srv-mdl-shared/models"
	shared "github.com/agile-work/srv-shared"
	"github.com/go-chi/chi"
)

// CreateField insert a new field in the database
func CreateField(r *http.Request) *moduleShared.Response {
	field := &models.Field{}
	languageCode := r.Header.Get("Content-Language")
	sharedModels.TranslationFieldsRequestLanguageCode = languageCode
	response := db.GetResponse(r, field, "CreateField")
	if response.Code != http.StatusOK {
		return response
	}

	field.SchemaCode = chi.URLParam(r, "schema_code")

	err := field.ProcessDefinitions(languageCode, r.Method)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorInsertingRecord, "CreateField processing definitions", err.Error()))

		return response
	}
	response.Data = field
	return response
}

// LoadAllFields returns all fields from a schema
func LoadAllFields(r *http.Request) *moduleShared.Response {
	return nil
}

// LoadField return an specific field from a schema
func LoadField(r *http.Request) *moduleShared.Response {
	return nil
}

// UpdateField updates the field attributes in the database
func UpdateField(r *http.Request) *moduleShared.Response {
	return nil
}

// DeleteField deletes an specific field in the database
func DeleteField(r *http.Request) *moduleShared.Response {
	return nil
}

// AddFieldValidation include a new validation to a field
func AddFieldValidation(r *http.Request) *moduleShared.Response {
	return nil
}

// UpdateFieldValidation update the validation attributes
func UpdateFieldValidation(r *http.Request) *moduleShared.Response {
	return nil
}

// DeleteFieldValidation delete a validation from the database
func DeleteFieldValidation(r *http.Request) *moduleShared.Response {
	return nil
}
