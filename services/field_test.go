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

type ServiceFieldTestSuite struct {
	suite.Suite
	FieldInstanceID           string
	FieldValidationInstanceID string
	SchemaInstanceID          string
	UserID                    string
}

func (suite *ServiceFieldTestSuite) SetupTest() {
	config, _ := config.NewConfig("..\\config.toml")
	db.Connect(config.Host, config.Port, config.User, config.Password, config.DBName, false)
	suite.UserID = "307e481c-69c5-11e9-96a0-06ea2c43bb20"
}

func (suite *ServiceFieldTestSuite) Test00001CreateField() {
	createSchemaToField(suite)

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

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)

	fieldValue := reflect.ValueOf(response.Data).Elem()
	suite.FieldInstanceID = fieldValue.FieldByName("ID").Interface().(string)
}

func (suite *ServiceFieldTestSuite) Test00002LoadAllFields() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/fields", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("schema_id", suite.SchemaInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadAllFields(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceFieldTestSuite) Test00003LoadField() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/fields", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("field_id", suite.FieldInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadField(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceFieldTestSuite) Test00004UpdateField() {
	data := map[string]interface{}{
		"description": "Descrição do Field Teste 01 Updated",
		"active":      false,
	}
	jsonData, _ := json.Marshal(&data)

	req, _ := http.NewRequest("PATCH", "http://localhost:3333/api/v1/admin/fields", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("field_id", suite.FieldInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := UpdateField(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceFieldTestSuite) Test00005CreateFieldValidation() {
	data := map[string]interface{}{
		"code":       "fieldvalidationteste01",
		"schema_id":  suite.SchemaInstanceID,
		"field_id":   suite.FieldInstanceID,
		"validation": "$require$",
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/fields", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := CreateFieldValidation(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)

	fieldValue := reflect.ValueOf(response.Data).Elem()
	suite.FieldValidationInstanceID = fieldValue.FieldByName("ID").Interface().(string)
}

func (suite *ServiceFieldTestSuite) Test00006LoadAllFieldValidations() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/fields", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("field_id", suite.FieldInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadAllFieldValidations(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceFieldTestSuite) Test00007LoadFieldValidation() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/fields", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("field_validation_id", suite.FieldValidationInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadFieldValidation(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceFieldTestSuite) Test00008UpdateFieldValidation() {
	data := map[string]interface{}{
		"validation": "$readOnly$",
	}
	jsonData, _ := json.Marshal(&data)

	req, _ := http.NewRequest("PATCH", "http://localhost:3333/api/v1/admin/fields", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("field_validation_id", suite.FieldValidationInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := UpdateFieldValidation(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceFieldTestSuite) Test00009DeleteFieldValidation() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/fields", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("field_validation_id", suite.FieldValidationInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := DeleteFieldValidation(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceFieldTestSuite) Test00010DeleteField() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/fields", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("field_id", suite.FieldInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := DeleteField(req)

	deleteSchemaToField(suite)

	assert.Equal(suite.T(), 200, response.Code)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestServiceFieldSuite(t *testing.T) {
	suite.Run(t, new(ServiceFieldTestSuite))
}

func createSchemaToField(suite *ServiceFieldTestSuite) {
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

func deleteSchemaToField(suite *ServiceFieldTestSuite) {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/schemas", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("schema_id", suite.SchemaInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	DeleteSchema(req)
}
