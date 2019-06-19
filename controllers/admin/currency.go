package admin

import (
	"net/http"

	"github.com/agile-work/srv-shared/util"

	"github.com/agile-work/srv-mdl-core/models/currency"

	"github.com/go-chi/chi"

	mdlShared "github.com/agile-work/srv-mdl-shared"
	mdlSharedModels "github.com/agile-work/srv-mdl-shared/models"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// PostCurrency sends the request to model creating a new currency
func PostCurrency(res http.ResponseWriter, req *http.Request) {
	mdlSharedModels.TranslationFieldsRequestLanguageCode = req.Header.Get("Content-Language")
	currency := &currency.Currency{}
	response := mdlShared.NewResponse()

	if err := response.Load(req, currency); err != nil {
		response.NewError(http.StatusInternalServerError, "PostCurrency response load", err.Error())
		response.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "PostCurrency currency new transaction", err.Error())
		response.Render(res, req)
		return
	}

	mdlSharedModels.TranslationFieldsRequestLanguageCode = "all"
	if err := currency.Create(trs); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "PostCurrency "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()

	response.Data = currency
	response.Render(res, req)
}

// GetAllCurrencies return all currency instances from the model
func GetAllCurrencies(res http.ResponseWriter, req *http.Request) {
	mdlSharedModels.TranslationFieldsRequestLanguageCode = req.Header.Get("Content-Language")
	response := mdlShared.NewResponse()

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "GetAllCurrencies currency new transaction", err.Error())
		response.Render(res, req)
		return
	}

	metaData := mdlShared.Metadata{}
	metaData.Load(req)
	opt := metaData.GenerateDBOptions()
	currencies := &currency.Currencies{}
	if err := currencies.LoadAll(trs, opt); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "GetAllCurrencies "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Data = currencies
	response.Metadata = metaData
	response.Render(res, req)
}

// GetCurrency return only one currency from the model
func GetCurrency(res http.ResponseWriter, req *http.Request) {
	mdlSharedModels.TranslationFieldsRequestLanguageCode = req.Header.Get("Content-Language")
	response := mdlShared.NewResponse()

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "GetCurrency currency new transaction", err.Error())
		response.Render(res, req)
		return
	}

	currency := &currency.Currency{Code: chi.URLParam(req, "currency_code")}
	if err := currency.Load(trs); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "GetCurrency "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Data = currency
	response.Render(res, req)
}

// UpdateCurrency sends the request to model updating a currency
func UpdateCurrency(res http.ResponseWriter, req *http.Request) {
	mdlSharedModels.TranslationFieldsRequestLanguageCode = req.Header.Get("Content-Language")
	currency := &currency.Currency{}
	response := mdlShared.NewResponse()

	if err := response.Load(req, currency); err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateCurrency currency new transaction", err.Error())
		response.Render(res, req)
		return
	}

	currency.Code = chi.URLParam(req, "currency_code")

	body, err := util.GetBody(req)
	if err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateCurrency "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}

	columns, translations, err := util.GetColumnsFromBody(body, currency)
	if err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateCurrency "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateCurrency currency new transaction", err.Error())
		response.Render(res, req)
		return
	}

	if err := currency.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "UpdateCurrency "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Data = currency
	response.Render(res, req)
}

// DeleteCurrency sends the request to model deleting a currency
func DeleteCurrency(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "DeleteCurrency currency new transaction", err.Error())
		response.Render(res, req)
		return
	}

	currency := &currency.Currency{Code: chi.URLParam(req, "currency_code")}
	if err := currency.Delete(trs); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "DeleteCurrency "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Render(res, req)
}

// AddRate sends the request to model creating a new rate
func AddRate(res http.ResponseWriter, req *http.Request) {
	mdlSharedModels.TranslationFieldsRequestLanguageCode = req.Header.Get("Content-Language")
	rate := &currency.Rate{}
	response := mdlShared.NewResponse()

	if err := response.Load(req, rate); err != nil {
		response.NewError(http.StatusInternalServerError, "AddRate rate new transaction", err.Error())
		response.Render(res, req)
		return
	}

	fromCode := chi.URLParam(req, "currency_code")
	toCode := chi.URLParam(req, "to_currency_code")

	body, err := util.GetBody(req)
	if err != nil {
		response.NewError(http.StatusInternalServerError, "AddRate "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}

	cols, err := util.GetBodyColumns(body)
	if err != nil {
		response.NewError(http.StatusInternalServerError, "AddRate "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}

	if !util.Contains(cols, "start_at", "end_at") {
		trs, err := db.NewTransaction()
		if err != nil {
			response.NewError(http.StatusInternalServerError, "AddRate currency new transaction", err.Error())
			response.Render(res, req)
			return
		}

		if err := rate.AddRate(trs, fromCode, toCode); err != nil {
			trs.Rollback()
			response.NewError(http.StatusInternalServerError, "AddRate "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
			response.Render(res, req)
			return
		}
		trs.Commit()
	} else {
		// TODO: insert with start_at and end_at adjusting all rates
		response.NewError(http.StatusNotImplemented, "AddRate "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
	}
}
