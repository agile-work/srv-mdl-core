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

type ServiceSectionTestSuite struct {
	suite.Suite
	SectionInstanceID string
	SchemaInstanceID  string
	PageInstanceID    string
	UserID            string
}

func (suite *ServiceSectionTestSuite) SetupTest() {
	config, _ := config.NewConfig("..\\config.toml")
	db.Connect(config.Host, config.Port, config.User, config.Password, config.DBName, false)
	suite.UserID = "307e481c-69c5-11e9-96a0-06ea2c43bb20"
}

func (suite *ServiceSectionTestSuite) Test00001CreateSection() {
	createSchemaToSection(suite)
	createPageToSection(suite)

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

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)

	sectionValue := reflect.ValueOf(response.Data).Elem()
	suite.SectionInstanceID = sectionValue.FieldByName("ID").Interface().(string)
}

func (suite *ServiceSectionTestSuite) Test00002LoadAllSections() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/sections", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("page_id", suite.PageInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadAllSections(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceSectionTestSuite) Test00003LoadSection() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/sections", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("section_id", suite.SectionInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadSection(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceSectionTestSuite) Test00004UpdateSection() {
	data := map[string]interface{}{
		"description": "Descrição do Section Teste 01 Updated",
	}
	jsonData, _ := json.Marshal(&data)

	req, _ := http.NewRequest("PATCH", "http://localhost:3333/api/v1/admin/sections", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("section_id", suite.SectionInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := UpdateSection(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceSectionTestSuite) Test00005DeleteSection() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/sections", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("section_id", suite.SectionInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := DeleteSection(req)

	deletePageToSection(suite)
	deleteSchemaToSection(suite)

	assert.Equal(suite.T(), 200, response.Code)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestServiceSectionSuite(t *testing.T) {
	suite.Run(t, new(ServiceSectionTestSuite))
}

func createSchemaToSection(suite *ServiceSectionTestSuite) {
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

func deleteSchemaToSection(suite *ServiceSectionTestSuite) {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/schemas", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("schema_id", suite.SchemaInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	DeleteSchema(req)
}

func createPageToSection(suite *ServiceSectionTestSuite) {
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

func deletePageToSection(suite *ServiceSectionTestSuite) {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/pages", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("page_id", suite.PageInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	DeletePage(req)
}
