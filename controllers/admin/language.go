package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-shared/util"

	"github.com/agile-work/srv-mdl-core/models/language"

	"github.com/go-chi/chi"

	mdlShared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// PostLanguage sends the request to model creating a new language
func PostLanguage(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	language := &language.Language{}
	response := mdlShared.NewResponse()

	if err := response.Load(req, language); err != nil {
		response.NewError(http.StatusInternalServerError, "PostLanguage response load", err.Error())
		response.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "PostLanguage language new transaction", err.Error())
		response.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"
	if err := language.Create(trs); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "PostLanguage "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()

	response.Data = language
	response.Render(res, req)
}

// GetAllLanguages return all language instances from the model
func GetAllLanguages(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	response := mdlShared.NewResponse()

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "GetAllLanguages language new transaction", err.Error())
		response.Render(res, req)
		return
	}

	metaData := mdlShared.Metadata{}
	metaData.Load(req)
	opt := metaData.GenerateDBOptions()
	languages := &language.Languages{}
	if err := languages.LoadAll(trs, opt); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "GetAllLanguages "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Data = languages
	response.Metadata = metaData
	response.Render(res, req)
}

// GetLanguage return only one language from the model
func GetLanguage(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	response := mdlShared.NewResponse()

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "GetLanguage language new transaction", err.Error())
		response.Render(res, req)
		return
	}

	language := &language.Language{Code: chi.URLParam(req, "language_code")}
	if err := language.Load(trs); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "GetLanguage "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Data = language
	response.Render(res, req)
}

// UpdateLanguage sends the request to model updating a language
func UpdateLanguage(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	language := &language.Language{}
	response := mdlShared.NewResponse()

	if err := response.Load(req, language); err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateLanguage language new transaction", err.Error())
		response.Render(res, req)
		return
	}

	language.Code = chi.URLParam(req, "language_code")

	body, err := util.GetBodyMap(req)
	if err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateLanguage "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}

	columns, translations, err := util.GetColumnsFromBody(body, language)
	if err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateLanguage "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateLanguage language new transaction", err.Error())
		response.Render(res, req)
		return
	}

	if err := language.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "UpdateLanguage "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Data = language
	response.Render(res, req)
}

// DeleteLanguage sends the request to model deleting a language
func DeleteLanguage(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "DeleteLanguage language new transaction", err.Error())
		response.Render(res, req)
		return
	}

	language := &language.Language{Code: chi.URLParam(req, "language_code")}
	if err := language.Delete(trs); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "DeleteLanguage "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Render(res, req)
}
