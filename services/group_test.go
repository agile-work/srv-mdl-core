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

type ServiceGroupTestSuite struct {
	suite.Suite
	GroupInstanceID      string
	PermissionInstanceID string
	UserInstanceID       string
	UserID               string
}

func (suite *ServiceGroupTestSuite) SetupTest() {
	config, _ := config.NewConfig("..\\config.toml")
	db.Connect(config.Host, config.Port, config.User, config.Password, config.DBName, false)
	suite.UserID = "307e481c-69c5-11e9-96a0-06ea2c43bb20"
}

func (suite *ServiceGroupTestSuite) Test00001CreateGroup() {
	createUserToGroup(suite)

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

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)

	groupValue := reflect.ValueOf(response.Data).Elem()
	suite.GroupInstanceID = groupValue.FieldByName("ID").Interface().(string)
}

func (suite *ServiceGroupTestSuite) Test00002LoadAllGroups() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/groups", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := LoadAllGroups(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceGroupTestSuite) Test00003LoadGroup() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/groups", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("group_id", suite.GroupInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadGroup(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceGroupTestSuite) Test00004UpdateGroup() {
	data := map[string]interface{}{
		"description": "Descrição do Grupo Teste 01 Updated",
		"active":      false,
	}
	jsonData, _ := json.Marshal(&data)

	req, _ := http.NewRequest("PATCH", "http://localhost:3333/api/v1/admin/groups", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("group_id", suite.GroupInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := UpdateGroup(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceGroupTestSuite) Test00005InsertUserInGroup() {
	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/groups", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("group_id", suite.GroupInstanceID)
	rctx.URLParams.Add("user_id", suite.UserInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := InsertUserInGroup(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceGroupTestSuite) Test00006LoadAllUsersByGroup() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/groups", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("group_id", suite.GroupInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadAllUsersByGroup(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceGroupTestSuite) Test00007RemoveUserFromGroup() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/groups", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("group_id", suite.GroupInstanceID)
	rctx.URLParams.Add("user_id", suite.UserInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := RemoveUserFromGroup(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceGroupTestSuite) Test00008InsertPermission() {
	data := map[string]interface{}{
		"group_id":       suite.GroupInstanceID,
		"structure_type": "user",
		"structure_id":   suite.UserInstanceID,
		"type":           100,
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/groups", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := InsertPermission(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)

	permissionValue := reflect.ValueOf(response.Data).Elem()
	suite.PermissionInstanceID = permissionValue.FieldByName("ID").Interface().(string)
}

func (suite *ServiceGroupTestSuite) Test00009LoadAllPermissionsByGroup() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/groups", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("group_id", suite.GroupInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadAllPermissionsByGroup(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceGroupTestSuite) Test00010RemovePermission() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/groups", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("permission_id", suite.PermissionInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := RemovePermission(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceGroupTestSuite) Test00011DeleteGroup() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/groups", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("group_id", suite.GroupInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := DeleteGroup(req)

	deleteUserToGroup(suite)

	assert.Equal(suite.T(), 200, response.Code)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestServiceGroupSuite(t *testing.T) {
	suite.Run(t, new(ServiceGroupTestSuite))
}

func createUserToGroup(suite *ServiceGroupTestSuite) {
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

	userValue := reflect.ValueOf(response.Data).Elem()
	suite.UserInstanceID = userValue.FieldByName("ID").Interface().(string)
}

func deleteUserToGroup(suite *ServiceGroupTestSuite) {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/users", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("user_id", suite.UserInstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	DeleteUser(req)
}
