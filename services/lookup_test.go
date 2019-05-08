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

type ServiceLookupTestSuite struct {
	suite.Suite
	LookupInstanceID       string
	LookupOptionInstanceID string
	UserID                 string
}

func (suite *ServiceLookupTestSuite) SetupTest() {
	config, _ := config.NewConfig("..\\config.toml")
	db.Connect(config.Host, config.Port, config.User, config.Password, config.DBName, false)
	suite.UserID = "307e481c-69c5-11e9-96a0-06ea2c43bb20"
}

func (suite *ServiceLookupTestSuite) Test00001CreateLookup() {
	data := map[string]interface{}{
		"name":         "Lookup Teste 01",
		"description":  "Descrição do Lookup Teste 01",
		"code":         "lookupteste01",
		"type":         "static",
		"value":        "id",
		"label":        "value",
		"autocomplete": "value",
		"active":       true,
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/lookups", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := CreateLookup(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)

	lookupValue := reflect.ValueOf(response.Data).Elem()
	suite.LookupInstanceID = lookupValue.FieldByName("ID").Interface().(string)
}

func (suite *ServiceLookupTestSuite) Test00002LoadAllLookups() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/lookups", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := LoadAllLookups(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceLookupTestSuite) Test00003LoadLookup() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/lookups", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("lookup_id", suite.LookupInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadLookup(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceLookupTestSuite) Test00004UpdateLookup() {
	data := map[string]interface{}{
		"description": "Descrição do Lookup Teste 01 Updated",
		"active":      false,
	}
	jsonData, _ := json.Marshal(&data)

	req, _ := http.NewRequest("PATCH", "http://localhost:3333/api/v1/admin/lookups", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("lookup_id", suite.LookupInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := UpdateLookup(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceLookupTestSuite) Test00005CreateLookupOption() {
	data := map[string]interface{}{
		"lookup_id": suite.LookupInstanceID,
		"code":      "valorteste01",
		"value":     "valorteste01",
		"label":     "Valor Teste 01",
		"active":    true,
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/lookups", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("lookup_id", suite.LookupInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := CreateLookupOption(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)

	lookupOptionValue := reflect.ValueOf(response.Data).Elem()
	suite.LookupOptionInstanceID = lookupOptionValue.FieldByName("ID").Interface().(string)
}

func (suite *ServiceLookupTestSuite) Test00006LoadAllLookupOptions() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/lookups", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("lookup_id", suite.LookupInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadAllLookupOptions(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceLookupTestSuite) Test00007LoadLookupOption() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/lookups", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("lookup_id", suite.LookupInstanceID)
	rctx.URLParams.Add("lookup_option_id", suite.LookupOptionInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadLookupOption(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceLookupTestSuite) Test00008UpdateLookupOption() {
	data := map[string]interface{}{
		"label":  "Valor Teste 01 Updated",
		"active": true,
	}
	jsonData, _ := json.Marshal(&data)

	req, _ := http.NewRequest("PATCH", "http://localhost:3333/api/v1/admin/lookups", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("lookup_id", suite.LookupInstanceID)
	rctx.URLParams.Add("lookup_option_id", suite.LookupOptionInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := UpdateLookupOption(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceLookupTestSuite) Test00009DeleteLookupOption() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/lookups", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("lookup_id", suite.LookupInstanceID)
	rctx.URLParams.Add("lookup_option_id", suite.LookupOptionInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := DeleteLookupOption(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceLookupTestSuite) Test00010DeleteLookup() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/lookups", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("lookup_id", suite.LookupInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := DeleteLookup(req)

	assert.Equal(suite.T(), 200, response.Code)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestServiceLookupSuite(t *testing.T) {
	suite.Run(t, new(ServiceLookupTestSuite))
}
