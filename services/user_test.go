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

type ServiceUserTestSuite struct {
	suite.Suite
	UserInstanceID  string
	GroupInstanceID string
	UserID          string
}

func (suite *ServiceUserTestSuite) SetupTest() {
	config, _ := config.NewConfig("..\\config.toml")
	db.Connect(config.Host, config.Port, config.User, config.Password, config.DBName, false)
	suite.UserID = "307e481c-69c5-11e9-96a0-06ea2c43bb20"
}

func (suite *ServiceUserTestSuite) Test00001CreateUser() {
	createGroupToUser(suite)

	data := map[string]interface{}{
		"username":      "usuarioteste01",
		"first_name":    "Usuário",
		"last_name":     "Teste 01",
		"email":         "usuarioteste01@domain.com",
		"password":      "123456",
		"language_code": "pt-br",
		"active":        true,
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/users", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := CreateUser(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)

	userValue := reflect.ValueOf(response.Data).Elem()
	suite.UserInstanceID = userValue.FieldByName("ID").Interface().(string)
}

func (suite *ServiceUserTestSuite) Test00002LoadAllUsers() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/users", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := LoadAllUsers(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceUserTestSuite) Test00003LoadUser() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/users", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("user_id", suite.UserInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadUser(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceUserTestSuite) Test00004UpdateUser() {
	data := map[string]interface{}{
		"last_name": "User 01 Updated",
		"active":    false,
	}
	jsonData, _ := json.Marshal(&data)

	req, _ := http.NewRequest("PATCH", "http://localhost:3333/api/v1/admin/users", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("user_id", suite.UserInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := UpdateUser(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceUserTestSuite) Test00005LoadAllGroupsByUser() {
	insertUserInGroup(suite)

	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/users", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("user_id", suite.UserInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadAllGroupsByUser(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceUserTestSuite) Test00006DeleteUser() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/users", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("user_id", suite.UserInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := DeleteUser(req)

	deleteGroupToUser(suite)

	assert.Equal(suite.T(), 200, response.Code)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestServiceUserSuite(t *testing.T) {
	suite.Run(t, new(ServiceUserTestSuite))
}

func createGroupToUser(suite *ServiceUserTestSuite) {
	data := map[string]interface{}{
		"name":        "Grupo Teste 01",
		"description": "Descrição do Grupo Teste 01",
		"code":        "grupoteste01",
		"active":      true,
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/groups", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := CreateGroup(req)

	groupValue := reflect.ValueOf(response.Data).Elem()
	suite.GroupInstanceID = groupValue.FieldByName("ID").Interface().(string)
}

func insertUserInGroup(suite *ServiceUserTestSuite) {
	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/groups", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("group_id", suite.GroupInstanceID)
	rctx.URLParams.Add("user_id", suite.UserInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	InsertUserInGroup(req)
}

func deleteGroupToUser(suite *ServiceUserTestSuite) {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/groups", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("group_id", suite.GroupInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	DeleteGroup(req)
}
