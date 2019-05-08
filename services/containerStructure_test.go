package services

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/agile-work/srv-mdl-core/models"

	"github.com/agile-work/srv-shared/sql-builder/db"
	"github.com/agile-work/srv-mdl-core/config"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServiceContainerStructureTestSuite struct {
	suite.Suite
	ContainerStructureInstanceID string
	SchemaInstanceID             string
	PageInstanceID               string
	SectionInstanceID            string
	FieldInstanceID              string
	UserID                       string
}

func (suite *ServiceContainerStructureTestSuite) SetupTest() {
	config, _ := config.NewConfig("..\\config.toml")
	db.Connect(config.Host, config.Port, config.User, config.Password, config.DBName, false)
	suite.UserID = "307e481c-69c5-11e9-96a0-06ea2c43bb20"
}

func (suite *ServiceContainerStructureTestSuite) Test00001CreateContainerStructure() {
	createSchemaToContainerStructure(suite)
	createPageToContainerStructure(suite)
	createSectionToContainerStructure(suite)
	createFieldToContainerStructure(suite)

	data := map[string]interface{}{
		"schema_id":       suite.SchemaInstanceID,
		"page_id":         suite.PageInstanceID,
		"container_id":    suite.SectionInstanceID,
		"container_type":  shared.TableCoreSchPagSections,
		"structure_id":    suite.FieldInstanceID,
		"structure_type":  shared.TableCoreSchFields,
		"position_row":    1,
		"position_column": 1,
		"width":           200,
		"height":          10,
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/container_structures", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := CreateContainerStructure(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)

	containerStructureValue := reflect.ValueOf(response.Data).Elem()
	suite.ContainerStructureInstanceID = containerStructureValue.FieldByName("ID").Interface().(string)
}

func (suite *ServiceContainerStructureTestSuite) Test00002LoadAllContainerStructures() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/container_structures", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("container_id", suite.SectionInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadAllContainerStructures(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceContainerStructureTestSuite) Test00003LoadContainerStructure() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/container_structures", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("container_structure_id", suite.ContainerStructureInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadContainerStructure(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceContainerStructureTestSuite) Test00004UpdateContainerStructure() {
	data := map[string]interface{}{
		"type": "chart",
	}
	jsonData, _ := json.Marshal(&data)

	req, _ := http.NewRequest("PATCH", "http://localhost:3333/api/v1/admin/container_structures", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("container_structure_id", suite.ContainerStructureInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := UpdateContainerStructure(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceContainerStructureTestSuite) Test00005DeleteContainerStructure() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/container_structures", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("container_structure_id", suite.ContainerStructureInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := DeleteContainerStructure(req)

	deleteFieldToContainerStructure(suite)
	deleteSectionToContainerStructure(suite)
	deletePageToContainerStructure(suite)
	deleteSchemaToContainerStructure(suite)

	assert.Equal(suite.T(), 200, response.Code)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestServiceContainerStructureSuite(t *testing.T) {
	suite.Run(t, new(ServiceContainerStructureTestSuite))
}

func createSchemaToContainerStructure(suite *ServiceContainerStructureTestSuite) {
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

func deleteSchemaToContainerStructure(suite *ServiceContainerStructureTestSuite) {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/schemas", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("schema_id", suite.SchemaInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	DeleteSchema(req)
}

func createPageToContainerStructure(suite *ServiceContainerStructureTestSuite) {
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

func deletePageToContainerStructure(suite *ServiceContainerStructureTestSuite) {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/pages", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("page_id", suite.PageInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	DeletePage(req)
}

func createSectionToContainerStructure(suite *ServiceContainerStructureTestSuite) {
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

func deleteSectionToContainerStructure(suite *ServiceContainerStructureTestSuite) {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/sections", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("section_id", suite.SectionInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	DeleteSection(req)
}

func createFieldToContainerStructure(suite *ServiceContainerStructureTestSuite) {
	data := map[string]interface{}{
		"code":        "fieldteste01",
		"schema_id":   suite.SchemaInstanceID,
		"name":        "Field Teste 01",
		"description": "Descrição do Field Teste 01",
		"field_type":  "text",
		"active":      true,
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/fields", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := CreateField(req)

	fieldValue := reflect.ValueOf(response.Data).Elem()
	suite.FieldInstanceID = fieldValue.FieldByName("ID").Interface().(string)
}

func deleteFieldToContainerStructure(suite *ServiceContainerStructureTestSuite) {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/fields", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("field_id", suite.FieldInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	DeleteField(req)
}
