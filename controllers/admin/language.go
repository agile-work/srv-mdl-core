package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-mdl-shared/util"

	"github.com/agile-work/srv-mdl-core/models/language"

	"github.com/go-chi/chi"

	"github.com/agile-work/srv-shared/sql-builder/db"
)

// PostLanguage sends the request to model creating a new language
func PostLanguage(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	language := &language.Language{}
	resp := response.New()

	if err := resp.Parse(req, language); err != nil {
		resp.NewError("PostLanguage response load", err)
		resp.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("PostLanguage language new transaction", err)
		resp.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"
	if err := language.Create(trs); err != nil {
		trs.Rollback()
		resp.NewError("PostLanguage", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = language
	resp.Render(res, req)
}

// GetAllLanguages return all language instances from the model
func GetAllLanguages(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("GetAllLanguages language new transaction", err)
		resp.Render(res, req)
		return
	}

	metaData := response.Metadata{}
	metaData.Load(req)
	opt := metaData.GenerateDBOptions()
	languages := &language.Languages{}
	if err := languages.LoadAll(trs, opt); err != nil {
		trs.Rollback()
		resp.NewError("GetAllLanguages", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = languages
	resp.Metadata = metaData
	resp.Render(res, req)
}

// GetLanguage return only one language from the model
func GetLanguage(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("GetLanguage language new transaction", err)
		resp.Render(res, req)
		return
	}

	language := &language.Language{Code: chi.URLParam(req, "language_code")}
	if err := language.Load(trs); err != nil {
		trs.Rollback()
		resp.NewError("GetLanguage", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = language
	resp.Render(res, req)
}

// UpdateLanguage sends the request to model updating a language
func UpdateLanguage(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	language := &language.Language{}
	resp := response.New()

	if err := resp.Parse(req, language); err != nil {
		resp.NewError("UpdateLanguage language new transaction", err)
		resp.Render(res, req)
		return
	}

	language.Code = chi.URLParam(req, "language_code")

	body, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdateLanguage", err)
		resp.Render(res, req)
		return
	}

	columns, translations, err := util.GetColumnsFromBody(body, language)
	if err != nil {
		resp.NewError("UpdateLanguage", err)
		resp.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateLanguage language new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := language.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		resp.NewError("UpdateLanguage", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = language
	resp.Render(res, req)
}

// DeleteLanguage sends the request to model deleting a language
func DeleteLanguage(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeleteLanguage language new transaction", err)
		resp.Render(res, req)
		return
	}

	language := &language.Language{Code: chi.URLParam(req, "language_code")}
	if err := language.Delete(trs); err != nil {
		trs.Rollback()
		resp.NewError("DeleteLanguage", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Render(res, req)
}
