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

type ServicePageTestSuite struct {
	suite.Suite
	PageInstanceID   string
	SchemaInstanceID string
	UserID           string
}

func (suite *ServicePageTestSuite) SetupTest() {
	config, _ := config.NewConfig("..\\config.toml")
	db.Connect(config.Host, config.Port, config.User, config.Password, config.DBName, false)
	suite.UserID = "307e481c-69c5-11e9-96a0-06ea2c43bb20"
}

func (suite *ServicePageTestSuite) Test00001CreatePage() {
	createSchemaToPage(suite)

	data := map[string]interface{}{
		"code":        "pageteste01",
		"name":        "Page Teste 01",
		"description": "Descrição do Page Teste 01",
		"schema_id":   suite.SchemaInstanceID,
		"type":        "Form",
		"active":      true,
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/pages", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := CreatePage(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)

	pageValue := reflect.ValueOf(response.Data).Elem()
	suite.PageInstanceID = pageValue.FieldByName("ID").Interface().(string)
}

func (suite *ServicePageTestSuite) Test00002LoadAllPages() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/pages", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("schema_id", suite.SchemaInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadAllPages(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServicePageTestSuite) Test00003LoadPage() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/pages", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("page_id", suite.PageInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadPage(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServicePageTestSuite) Test00004UpdatePage() {
	data := map[string]interface{}{
		"description": "Descrição do Page Teste 01 Updated",
		"active":      false,
	}
	jsonData, _ := json.Marshal(&data)

	req, _ := http.NewRequest("PATCH", "http://localhost:3333/api/v1/admin/pages", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("page_id", suite.PageInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := UpdatePage(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServicePageTestSuite) Test00005DeletePage() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/pages", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("page_id", suite.PageInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := DeletePage(req)

	deleteSchemaToPage(suite)

	assert.Equal(suite.T(), 200, response.Code)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestServicePageSuite(t *testing.T) {
	suite.Run(t, new(ServicePageTestSuite))
}

func createSchemaToPage(suite *ServicePageTestSuite) {
	data := map[string]interface{}{
		"name":        "Schema Teste 01",
		"description": "Descrição do Schema Teste 01",
		"code":        "schemateste01",
		"plugin":      false,
		"active":      true,
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/schemas", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := CreateSchema(req)

	schemaValue := reflect.ValueOf(response.Data).Elem()
	suite.SchemaInstanceID = schemaValue.FieldByName("ID").Interface().(string)
}

func deleteSchemaToPage(suite *ServicePageTestSuite) {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/schemas", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("schema_id", suite.SchemaInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	DeleteSchema(req)
}
