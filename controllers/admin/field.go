package admin

import (
	"net/http"

	"github.com/agile-work/srv-shared/sql-builder/builder"

	"github.com/agile-work/srv-shared/util"

	"github.com/agile-work/srv-mdl-core/models/field"

	"github.com/go-chi/chi"

	mdlShared "github.com/agile-work/srv-mdl-shared"
	mdlSharedModels "github.com/agile-work/srv-mdl-shared/models"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// PostField sends the request to model creating a new field
func PostField(res http.ResponseWriter, req *http.Request) {
	mdlSharedModels.TranslationFieldsRequestLanguageCode = req.Header.Get("Content-Language")
	field := &field.Field{}
	response := mdlShared.NewResponse()

	if err := response.Load(req, field); err != nil {
		response.NewError(http.StatusInternalServerError, "PostField response load", err.Error())
		response.Render(res, req)
		return
	}

	field.SchemaCode = chi.URLParam(req, "schema_code")

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "PostField field new transaction", err.Error())
		response.Render(res, req)
		return
	}

	mdlSharedModels.TranslationFieldsRequestLanguageCode = "all"
	if err := field.Create(trs); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "PostField "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()

	response.Data = field
	response.Render(res, req)
}

// GetAllFields return all field instances from the model
func GetAllFields(res http.ResponseWriter, req *http.Request) {
	mdlSharedModels.TranslationFieldsRequestLanguageCode = req.Header.Get("Content-Language")
	response := mdlShared.NewResponse()

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "GetAllFields field new transaction", err.Error())
		response.Render(res, req)
		return
	}

	metaData := mdlShared.Metadata{}
	metaData.Load(req)
	opt := metaData.GenerateDBOptions()
	opt.AddCondition(builder.Equal("schema_code", chi.URLParam(req, "schema_code")))
	fields := &field.Fields{}
	if err := fields.LoadAll(trs, opt); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "GetAllFields "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Data = fields
	response.Metadata = metaData
	response.Render(res, req)
}

// GetField return only one field from the model
func GetField(res http.ResponseWriter, req *http.Request) {
	mdlSharedModels.TranslationFieldsRequestLanguageCode = req.Header.Get("Content-Language")
	response := mdlShared.NewResponse()

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "GetField field new transaction", err.Error())
		response.Render(res, req)
		return
	}

	field := &field.Field{SchemaCode: chi.URLParam(req, "schema_code"), Code: chi.URLParam(req, "field_code")}
	if err := field.Load(trs); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "GetField "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Data = field
	response.Render(res, req)
}

// UpdateField sends the request to model updating a field
func UpdateField(res http.ResponseWriter, req *http.Request) {
	mdlSharedModels.TranslationFieldsRequestLanguageCode = req.Header.Get("Content-Language")
	field := &field.Field{}
	response := mdlShared.NewResponse()

	if err := response.Load(req, field); err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateField field new transaction", err.Error())
		response.Render(res, req)
		return
	}

	field.SchemaCode = chi.URLParam(req, "schema_code")
	field.Code = chi.URLParam(req, "field_code")

	body, err := util.GetBody(req)
	if err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateField "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}

	columns, translations, err := util.GetColumnsFromBody(body, field)
	if err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateField "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateField field new transaction", err.Error())
		response.Render(res, req)
		return
	}

	if err := field.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "UpdateField "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Data = field
	response.Render(res, req)
}

// DeleteField sends the request to model deleting a field
func DeleteField(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "DeleteField field new transaction", err.Error())
		response.Render(res, req)
		return
	}

	field := &field.Field{SchemaCode: chi.URLParam(req, "schema_code"), Code: chi.URLParam(req, "field_code")}
	if err := field.Delete(trs); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "DeleteField "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Render(res, req)
}
