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

// CreateCurrency persists the request body creating a new object in the database
func CreateCurrency(r *http.Request) *moduleShared.Response {
	currency := models.Currency{}

	return db.Create(r, &currency, "CreateCurrency", shared.TableCoreCurrencies)
}

// LoadAllCurrencies return all instances from the object
func LoadAllCurrencies(r *http.Request) *moduleShared.Response {
	currencies := []models.Currency{}

	return db.Load(r, &currencies, "LoadAllCurrencies", shared.TableCoreCurrencies, nil)
}

// LoadCurrency return only one object from the database
func LoadCurrency(r *http.Request) *moduleShared.Response {
	currency := models.Currency{}
	currencyID := chi.URLParam(r, "currency_id")
	currencyIDColumn := fmt.Sprintf("%s.id", shared.TableCoreCurrencies)
	condition := builder.Equal(currencyIDColumn, currencyID)

	return db.Load(r, &currency, "LoadCurrency", shared.TableCoreCurrencies, condition)
}

// UpdateCurrency updates object data in the database
func UpdateCurrency(r *http.Request) *moduleShared.Response {
	currencyID := chi.URLParam(r, "currency_id")
	currencyIDColumn := fmt.Sprintf("%s.id", shared.TableCoreCurrencies)
	condition := builder.Equal(currencyIDColumn, currencyID)
	currency := models.Currency{
		ID: currencyID,
	}

	return db.Update(r, &currency, "UpdateCurrency", shared.TableCoreCurrencies, condition)
}

// DeleteCurrency deletes object from the database
func DeleteCurrency(r *http.Request) *moduleShared.Response {
	currencyID := chi.URLParam(r, "currency_id")
	currencyIDColumn := fmt.Sprintf("%s.id", shared.TableCoreCurrencies)
	condition := builder.Equal(currencyIDColumn, currencyID)

	return db.Remove(r, "DeleteCurrency", shared.TableCoreCurrencies, condition)
}

// CreateCurrencyRate persists the request body creating a new object in the database
func CreateCurrencyRate(r *http.Request) *moduleShared.Response {
	currencyID := chi.URLParam(r, "currency_id")
	currencyRate := models.CurrencyRate{
		FromCurrencyID: currencyID,
	}

	return db.Create(r, &currencyRate, "CreateCurrencyRate", shared.TableCoreCryRates)
}

// LoadAllCurrencyRates return all instances from the object
func LoadAllCurrencyRates(r *http.Request) *moduleShared.Response {
	currencies := []models.CurrencyRate{}
	fromCurrencyCode := chi.URLParam(r, "currency_id")
	fromCurrencyCodeColumn := fmt.Sprintf("%s.from_currency_id", shared.TableCoreCryRates)
	condition := builder.Equal(fromCurrencyCodeColumn, fromCurrencyCode)

	return db.Load(r, &currencies, "LoadAllCurrencyRates", shared.TableCoreCryRates, condition)
}

// LoadCurrencyRate return only one object from the database
func LoadCurrencyRate(r *http.Request) *moduleShared.Response {
	currencyRate := models.CurrencyRate{}
	currencyRateID := chi.URLParam(r, "currency_rate_id")
	currencyRateIDColumn := fmt.Sprintf("%s.id", shared.TableCoreCryRates)
	condition := builder.Equal(currencyRateIDColumn, currencyRateID)

	return db.Load(r, &currencyRate, "LoadCurrencyRate", shared.TableCoreCryRates, condition)
}

// UpdateCurrencyRate updates object data in the database
func UpdateCurrencyRate(r *http.Request) *moduleShared.Response {
	currencyRateID := chi.URLParam(r, "currency_rate_id")
	currencyRateIDColumn := fmt.Sprintf("%s.id", shared.TableCoreCryRates)
	condition := builder.Equal(currencyRateIDColumn, currencyRateID)
	currencyRate := models.CurrencyRate{
		ID: currencyRateID,
	}

	return db.Update(r, &currencyRate, "UpdateCurrencyRate", shared.TableCoreCryRates, condition)
}

// DeleteCurrencyRate deletes object from the database
func DeleteCurrencyRate(r *http.Request) *moduleShared.Response {
	currencyRateID := chi.URLParam(r, "currency_rate_id")
	currencyRateIDColumn := fmt.Sprintf("%s.id", shared.TableCoreCryRates)
	condition := builder.Equal(currencyRateIDColumn, currencyRateID)

	return db.Remove(r, "DeleteCurrencyRate", shared.TableCoreCryRates, condition)
}
