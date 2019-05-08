package services

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/agile-work/srv-shared/sql-builder/db"
	"github.com/agile-work/srv-mdl-core/config"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServiceCurrencyTestSuite struct {
	suite.Suite
	CurrencyInstanceID string
	pluginInstanceID   string
	UserID             string
}

func (suite *ServiceCurrencyTestSuite) SetupTest() {
	config, _ := config.NewConfig("..\\config.toml")
	db.Connect(config.Host, config.Port, config.User, config.Password, config.DBName, false)
	suite.UserID = "307e481c-69c5-11e9-96a0-06ea2c43bb20"
}

func (suite *ServiceCurrencyTestSuite) Test00001CreateCurrency() {

	data := map[string]interface{}{
		"name":   "Currency Teste 01",
		"code":   "currencyteste01",
		"active": true,
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/currencies", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := CreateCurrency(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)

	currencyValue := reflect.ValueOf(response.Data).Elem()
	suite.CurrencyInstanceID = currencyValue.FieldByName("ID").Interface().(string)
}

func (suite *ServiceCurrencyTestSuite) Test00002LoadAllCurrencies() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/currencies", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := LoadAllCurrencies(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceCurrencyTestSuite) Test00003LoadCurrency() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/currencies", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("currency_id", suite.CurrencyInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadCurrency(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceCurrencyTestSuite) Test00004UpdateCurrency() {
	data := map[string]interface{}{
		"active": false,
	}
	jsonData, _ := json.Marshal(&data)

	req, _ := http.NewRequest("PATCH", "http://localhost:3333/api/v1/admin/currencies", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("currency_id", suite.CurrencyInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := UpdateCurrency(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceCurrencyTestSuite) Test00005DeleteCurrency() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/currencies", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("currency_id", suite.CurrencyInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := DeleteCurrency(req)

	assert.Equal(suite.T(), 200, response.Code)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestServiceCurrencySuite(t *testing.T) {
	suite.Run(t, new(ServiceCurrencyTestSuite))
}
