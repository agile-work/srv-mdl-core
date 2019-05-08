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

type ServiceTabTestSuite struct {
	suite.Suite
	TabInstanceID     string
	SchemaInstanceID  string
	PageInstanceID    string
	SectionInstanceID string
	UserID            string
}

func (suite *ServiceTabTestSuite) SetupTest() {
	config, _ := config.NewConfig("..\\config.toml")
	db.Connect(config.Host, config.Port, config.User, config.Password, config.DBName, false)
	suite.UserID = "307e481c-69c5-11e9-96a0-06ea2c43bb20"
}

func (suite *ServiceTabTestSuite) Test00001CreateTab() {
	createSchemaToTab(suite)
	createPageToTab(suite)
	createSectionToTab(suite)

	data := map[string]interface{}{
		"code":        "tabteste01",
		"name":        "Tab Teste 01",
		"description": "Descrição do Tab Teste 01",
		"schema_id":   suite.SchemaInstanceID,
		"page_id":     suite.PageInstanceID,
		"section_id":  suite.SectionInstanceID,
		"tab_order":   1,
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/tabs", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := CreateTab(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)

	tabValue := reflect.ValueOf(response.Data).Elem()
	suite.TabInstanceID = tabValue.FieldByName("ID").Interface().(string)
}

func (suite *ServiceTabTestSuite) Test00002LoadAllTabs() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/tabs", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("section_id", suite.SectionInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadAllTabs(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceTabTestSuite) Test00003LoadTab() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/tabs", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("tab_id", suite.TabInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadTab(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceTabTestSuite) Test00004UpdateTab() {
	data := map[string]interface{}{
		"type": "chart",
	}
	jsonData, _ := json.Marshal(&data)

	req, _ := http.NewRequest("PATCH", "http://localhost:3333/api/v1/admin/tabs", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("tab_id", suite.TabInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := UpdateTab(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceTabTestSuite) Test00005DeleteTab() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/tabs", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("tab_id", suite.TabInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := DeleteTab(req)

	deleteSectionToTab(suite)
	deletePageToTab(suite)
	deleteSchemaToTab(suite)

	assert.Equal(suite.T(), 200, response.Code)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestServiceTabSuite(t *testing.T) {
	suite.Run(t, new(ServiceTabTestSuite))
}

func createSchemaToTab(suite *ServiceTabTestSuite) {
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

func deleteSchemaToTab(suite *ServiceTabTestSuite) {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/schemas", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("schema_id", suite.SchemaInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	DeleteSchema(req)
}

func createPageToTab(suite *ServiceTabTestSuite) {
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

func deletePageToTab(suite *ServiceTabTestSuite) {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/pages", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("page_id", suite.PageInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	DeletePage(req)
}

func createSectionToTab(suite *ServiceTabTestSuite) {
	data := map[string]interface{}{
		"code":        "sectionteste01",
		"name":        "Section Teste 01",
		"description": "Descrição do Section Teste 01",
		"schema_id":   suite.SchemaInstanceID,
		"page_id":     suite.PageInstanceID,
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/sections", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := CreateSection(req)

	sectionValue := reflect.ValueOf(response.Data).Elem()
	suite.SectionInstanceID = sectionValue.FieldByName("ID").Interface().(string)
}

func deleteSectionToTab(suite *ServiceTabTestSuite) {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/sections", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("section_id", suite.SectionInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	DeleteSection(req)
}
