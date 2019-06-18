package admin

import (
	"encoding/json"
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

// CreateCurrency persists the request body creating a new object in the database
func CreateCurrency(r *http.Request) *moduleShared.Response {
	currency := models.Currency{}

	return db.Create(r, &currency, "CreateCurrency", shared.TableCoreCurrencies)
}

// LoadAllCurrencies return all instances from the object
func LoadAllCurrencies(r *http.Request) *moduleShared.Response {
	// TODO: Make a way of limit the columns for the get all. Passing fields to LoadStruct.
	currencies := []models.Currency{}

	return db.Load(r, &currencies, "LoadAllCurrencies", shared.TableCoreCurrencies, nil)
}

// LoadCurrency return only one object from the database
func LoadCurrency(r *http.Request) *moduleShared.Response {
	// TODO: Limit the total of rates to +/- 100 records if has no filter
	currency := models.Currency{}
	currencyCode := chi.URLParam(r, "currency_code")
	currencyCodeColumn := fmt.Sprintf("%s.code", shared.TableCoreCurrencies)
	condition := builder.Equal(currencyCodeColumn, currencyCode)

	return db.Load(r, &currency, "LoadCurrency", shared.TableCoreCurrencies, condition)
}

// UpdateCurrency updates object data in the database
func UpdateCurrency(r *http.Request) *moduleShared.Response {
	currencyCode := chi.URLParam(r, "currency_code")
	currencyCodeColumn := fmt.Sprintf("%s.code", shared.TableCoreCurrencies)
	condition := builder.Equal(currencyCodeColumn, currencyCode)
	currency := models.Currency{
		ID: currencyCode,
	}

	return db.Update(r, &currency, "UpdateCurrency", shared.TableCoreCurrencies, condition)
}

// DeleteCurrency deletes object from the database
func DeleteCurrency(r *http.Request) *moduleShared.Response {
	currencyCode := chi.URLParam(r, "currency_code")
	currencyCodeColumn := fmt.Sprintf("%s.code", shared.TableCoreCurrencies)
	condition := builder.Equal(currencyCodeColumn, currencyCode)

	return db.Remove(r, "DeleteCurrency", shared.TableCoreCurrencies, condition)
}

// AddCurrencyRate persists the request body creating a new object in the database
func AddCurrencyRate(r *http.Request) *moduleShared.Response {
	currencyCode := chi.URLParam(r, "currency_code")
	toCode := chi.URLParam(r, "to_currency_code")

	rate := models.CurrencyRate{}

	response := db.GetResponse(r, &rate, "AddCurrencyRate")
	if response.Code != http.StatusOK {
		return response
	}

	cols := db.GetBodyColumns(r)
	if !util.Contains(cols, "start_at", "end_at") {
		trs, err := sql.NewTransaction()
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorInsertingRecord, "AddCurrencyRate new transaction", err.Error()))

			return response
		}
		endAt := time.Now().Format(time.RFC3339)
		querySQL := fmt.Sprintf(`update %s 
		set rates = jsonb_set(
			rates, 
			('{%s,-1}') ::text[], 
			rates::jsonb#>('{%s,-1}')::text[] || '{"end_at": "%s"}'
			,true
		) where code = '%s'`, shared.TableCoreCurrencies, toCode, toCode, endAt, currencyCode)
		trs.Add(builder.Raw(querySQL))

		t := time.Now()
		rate.Start = &t
		rateBytes, _ := json.Marshal(rate)
		querySQL = fmt.Sprintf(`update %s set rates = jsonb_insert(
		rates, '{%s,-1}', '%s', true) 
		where code = '%s'`, shared.TableCoreCurrencies, toCode, string(rateBytes), currencyCode)
		trs.Add(builder.Raw(querySQL))

		err = trs.Exec()
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorInsertingRecord, "AddCurrencyRate", err.Error()))

			return response
		}
	} else {
		// TODO: insert with start_at and end_at adjusting all rates
		response.Code = http.StatusNotImplemented
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorInsertingRecord, "AddCurrencyRate", ""))
		return response
	}
	return response
}
