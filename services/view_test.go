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

type ServiceViewTestSuite struct {
	suite.Suite
	SchemaInstanceID string
	PageInstanceID   string
	ViewInstanceID   string
	UserID           string
}

func (suite *ServiceViewTestSuite) SetupTest() {
	config, _ := config.NewConfig("..\\config.toml")
	db.Connect(config.Host, config.Port, config.User, config.Password, config.DBName, false)
	suite.UserID = "307e481c-69c5-11e9-96a0-06ea2c43bb20"
}

func (suite *ServiceViewTestSuite) Test00001CreateView() {
	createSchemaToView(suite)
	createPageToView(suite)

	data := map[string]interface{}{
		"code":        "viewteste01",
		"name":        "View Teste 01",
		"description": "Descrição do View Teste 01",
		"schema_id":   suite.SchemaInstanceID,
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/views", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := CreateView(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)

	viewValue := reflect.ValueOf(response.Data).Elem()
	suite.ViewInstanceID = viewValue.FieldByName("ID").Interface().(string)
}

func (suite *ServiceViewTestSuite) Test00002LoadAllViews() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/views", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("schema_id", suite.SchemaInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadAllViews(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceViewTestSuite) Test00003LoadView() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/views", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("view_id", suite.ViewInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadView(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceViewTestSuite) Test00004UpdateView() {
	data := map[string]interface{}{
		"description": "Descrição do View Teste 01 Updated",
	}
	jsonData, _ := json.Marshal(&data)

	req, _ := http.NewRequest("PATCH", "http://localhost:3333/api/v1/admin/views", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("view_id", suite.ViewInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := UpdateView(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceViewTestSuite) Test00005InsertPageInView() {
	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/views", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("view_id", suite.ViewInstanceID)
	rctx.URLParams.Add("page_id", suite.PageInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := InsertPageInView(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceViewTestSuite) Test00006LoadAllPagesByView() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/views", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("view_id", suite.ViewInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadAllPagesByView(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceViewTestSuite) Test00007RemovePageFromView() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/views", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("view_id", suite.ViewInstanceID)
	rctx.URLParams.Add("page_id", suite.PageInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := RemovePageFromView(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceViewTestSuite) Test00008DeleteView() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/views", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("view_id", suite.ViewInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := DeleteView(req)

	deletePageToView(suite)
	deleteSchemaToView(suite)

	assert.Equal(suite.T(), 200, response.Code)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestServiceViewSuite(t *testing.T) {
	suite.Run(t, new(ServiceViewTestSuite))
}

func createSchemaToView(suite *ServiceViewTestSuite) {
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

func deleteSchemaToView(suite *ServiceViewTestSuite) {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/schemas", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("schema_id", suite.SchemaInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	DeleteSchema(req)
}

func createPageToView(suite *ServiceViewTestSuite) {
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

	pageValue := reflect.ValueOf(response.Data).Elem()
	suite.PageInstanceID = pageValue.FieldByName("ID").Interface().(string)
}

func deletePageToView(suite *ServiceViewTestSuite) {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/pages", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("page_id", suite.PageInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	DeletePage(req)
}
