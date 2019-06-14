package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-core/models/field"
	services "github.com/agile-work/srv-mdl-core/services/admin"

	"github.com/go-chi/chi"

	mdlShared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-mdl-shared/db"
	sharedModels "github.com/agile-work/srv-mdl-shared/models"
)

// PostField sends the request to service creating a new field
func PostField(w http.ResponseWriter, r *http.Request) {
	fld := &field.Field{}
	sharedModels.TranslationFieldsRequestLanguageCode = r.Header.Get("Content-Language")
	response := db.GetResponse(r, fld, "CreateField")
	if response.Code != http.StatusOK {
		response.Render(w, r)
		return
	}

	fld.SchemaCode = chi.URLParam(r, "schema_code")

	sharedModels.TranslationFieldsRequestLanguageCode = "all"
	scope, err := services.CreateField(fld)
	if err != nil {
		response.NewError(http.StatusInternalServerError, scope, err.Error())
		response.Render(w, r)
		return
	}

	response.Data = fld
	response.Render(w, r)
}

// GetAllFields return all field instances from the service
func GetAllFields(w http.ResponseWriter, r *http.Request) {
	sharedModels.TranslationFieldsRequestLanguageCode = r.Header.Get("Content-Language")
	response := mdlShared.Response{
		Code: http.StatusOK,
	}
	flds, err := field.LoadAll(chi.URLParam(r, "schema_code"))
	if err != nil {
		response.NewError(http.StatusInternalServerError, "GetAllFields load all", err.Error())
		response.Render(w, r)
		return
	}
	response.Data = flds
	response.Render(w, r)
}

// GetField return only one field from the service
func GetField(w http.ResponseWriter, r *http.Request) {
	sharedModels.TranslationFieldsRequestLanguageCode = r.Header.Get("Content-Language")
	response := mdlShared.Response{
		Code: http.StatusOK,
	}
	fld := &field.Field{
		Code:       chi.URLParam(r, "field_code"),
		SchemaCode: chi.URLParam(r, "schema_code"),
	}
	err := fld.Load()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "GetField load", err.Error())
		response.Render(w, r)
		return
	}
	response.Data = fld
	response.Render(w, r)
}

// UpdateField sends the request to service updating a field
func UpdateField(w http.ResponseWriter, r *http.Request) {
	fld := &field.Field{}
	sharedModels.TranslationFieldsRequestLanguageCode = r.Header.Get("Content-Language")
	response := db.GetResponse(r, fld, "UpdateField")
	if response.Code != http.StatusOK {
		response.Render(w, r)
		return
	}

	fld.Code = chi.URLParam(r, "field_code")
	fld.SchemaCode = chi.URLParam(r, "schema_code")

	body, err := mdlShared.GetBody(r)
	if err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateField get body", err.Error())
		response.Render(w, r)
		return
	}

	columns, translations, err := mdlShared.GetColumnsFromBody(body, fld)
	if err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateField get columns", err.Error())
		response.Render(w, r)
		return
	}

	err = fld.Update(columns, translations)
	if err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateField load", err.Error())
		response.Render(w, r)
		return
	}
	response.Data = fld
	response.Render(w, r)
}

// DeleteField sends the request to service deleting a field
func DeleteField(w http.ResponseWriter, r *http.Request) {
	response := mdlShared.Response{
		Code: http.StatusOK,
	}
	fld := &field.Field{
		Code:       chi.URLParam(r, "field_code"),
		SchemaCode: chi.URLParam(r, "schema_code"),
	}
	err := fld.Delete()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "DeleteField delete", err.Error())
		response.Render(w, r)
		return
	}
	response.Render(w, r)
}
