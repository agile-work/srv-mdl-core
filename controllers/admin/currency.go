package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-mdl-shared/util"
	sharedUtil "github.com/agile-work/srv-shared/util"

	"github.com/agile-work/srv-mdl-core/models/currency"

	"github.com/go-chi/chi"

	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// PostCurrency sends the request to model creating a new currency
func PostCurrency(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	currency := &currency.Currency{}
	resp := response.New()

	if err := resp.Parse(req, currency); err != nil {
		resp.NewError("PostCurrency response load", err)
		resp.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("PostCurrency user new transaction", err)
		resp.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"
	if err := currency.Create(trs); err != nil {
		trs.Rollback()
		resp.NewError("PostCurrency", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = currency
	resp.Render(res, req)
}

// GetAllCurrencies return all currency instances from the model
func GetAllCurrencies(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	metadata := response.Metadata{}
	if err := metadata.Load(req); err != nil {
		resp.NewError("GetLookupInstance metadata parse", err)
		resp.Render(res, req)
		return
	}
	opt := metadata.GenerateDBOptions()
	currencies := &currency.Currencies{}
	if err := currencies.LoadAll(opt); err != nil {
		resp.NewError("GetAllCurrencies", err)
		resp.Render(res, req)
		return
	}
	resp.Data = currencies
	resp.Metadata = metadata
	resp.Render(res, req)
}

// GetCurrency return only one currency from the model
func GetCurrency(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	currency := &currency.Currency{Code: chi.URLParam(req, "currency_code")}
	if err := currency.Load(); err != nil {
		resp.NewError("GetCurrency", err)
		resp.Render(res, req)
		return
	}
	resp.Data = currency
	resp.Render(res, req)
}

// UpdateCurrency sends the request to model updating a currency
func UpdateCurrency(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	currency := &currency.Currency{}
	resp := response.New()

	if err := resp.Parse(req, currency); err != nil {
		resp.NewError("UpdateCurrency currency new transaction", err)
		resp.Render(res, req)
		return
	}

	currency.Code = chi.URLParam(req, "currency_code")

	body, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdateCurrency currency get body", err)
		resp.Render(res, req)
		return
	}

	columns, translations := util.GetColumnsFromBody(body, currency)

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateCurrency currency new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := currency.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		resp.NewError("UpdateCurrency", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = currency
	resp.Render(res, req)
}

// DeleteCurrency sends the request to model deleting a currency
func DeleteCurrency(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeleteCurrency currency new transaction", err)
		resp.Render(res, req)
		return
	}

	currency := &currency.Currency{Code: chi.URLParam(req, "currency_code")}
	if err := currency.Delete(trs); err != nil {
		trs.Rollback()
		resp.NewError("DeleteCurrency", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Render(res, req)
}

// AddRate sends the request to model creating a new rate
func AddRate(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	rate := &currency.Rate{}
	resp := response.New()

	if err := resp.Parse(req, rate); err != nil {
		resp.NewError("AddRate response load", err)
		resp.Render(res, req)
		return
	}

	fromCode := chi.URLParam(req, "currency_code")
	toCode := chi.URLParam(req, "to_currency_code")

	body, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("AddRate get body", err)
		resp.Render(res, req)
		return
	}

	if !sharedUtil.Contains(util.GetBodyColumns(body), "start_at", "end_at") {
		trs, err := db.NewTransaction()
		if err != nil {
			resp.NewError("AddRate new transaction", err)
			resp.Render(res, req)
			return
		}

		if err := rate.AddRate(trs, fromCode, toCode); err != nil {
			trs.Rollback()
			resp.NewError("AddRate", err)
			resp.Render(res, req)
			return
		}
		trs.Commit()
	} else {
		// TODO: insert with start_at and end_at adjusting all rates
		resp.NewError("AddRate", err)
		resp.Render(res, req)
	}
}
