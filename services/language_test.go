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

type ServiceLanguageTestSuite struct {
	suite.Suite
	InstanceID string
	UserID     string
}

func (suite *ServiceLanguageTestSuite) SetupTest() {
	config, _ := config.NewConfig("..\\config.toml")
	db.Connect(config.Host, config.Port, config.User, config.Password, config.DBName, false)
	suite.UserID = "307e481c-69c5-11e9-96a0-06ea2c43bb20"
}

func (suite *ServiceLanguageTestSuite) Test00001CreateLanguage() {
	data := map[string]interface{}{
		"name":        "Language Teste 01",
		"code":        "languageteste01",
		"active":      false,
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/languages", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := CreateLanguage(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)

	languageValue := reflect.ValueOf(response.Data).Elem()
	suite.InstanceID = languageValue.FieldByName("ID").Interface().(string)
}

func (suite *ServiceLanguageTestSuite) Test00002LoadAllLanguages() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/languages", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := LoadAllLanguages(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceLanguageTestSuite) Test00003LoadLanguage() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/languages", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("language_id", suite.InstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadLanguage(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceLanguageTestSuite) Test00004UpdateLanguage() {
	data := map[string]interface{}{
		"name": "Language Teste 01 Updated",
	}
	jsonData, _ := json.Marshal(&data)

	req, _ := http.NewRequest("PATCH", "http://localhost:3333/api/v1/admin/languages", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("language_id", suite.InstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := UpdateLanguage(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceLanguageTestSuite) Test00005DeleteLanguage() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/languages", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("language_id", suite.InstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := DeleteLanguage(req)

	assert.Equal(suite.T(), 200, response.Code)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestServiceLanguageSuite(t *testing.T) {
	suite.Run(t, new(ServiceLanguageTestSuite))
}
