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

type ServiceWidgetTestSuite struct {
	suite.Suite
	InstanceID string
	UserID     string
}

func (suite *ServiceWidgetTestSuite) SetupTest() {
	config, _ := config.NewConfig("..\\config.toml")
	db.Connect(config.Host, config.Port, config.User, config.Password, config.DBName, false)
	suite.UserID = "307e481c-69c5-11e9-96a0-06ea2c43bb20"
}

func (suite *ServiceWidgetTestSuite) Test00001CreateWidget() {
	data := map[string]interface{}{
		"type": "table",
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/admin/widgets", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := CreateWidget(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)

	widgetValue := reflect.ValueOf(response.Data).Elem()
	suite.InstanceID = widgetValue.FieldByName("ID").Interface().(string)
}

func (suite *ServiceWidgetTestSuite) Test00002LoadAllWidgets() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/widgets", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	response := LoadAllWidgets(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceWidgetTestSuite) Test00003LoadWidget() {
	req, _ := http.NewRequest("GET", "http://localhost:3333/api/v1/admin/widgets", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("widget_id", suite.InstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := LoadWidget(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceWidgetTestSuite) Test00004UpdateWidget() {
	data := map[string]interface{}{
		"type": "chart",
	}
	jsonData, _ := json.Marshal(&data)

	req, _ := http.NewRequest("PATCH", "http://localhost:3333/api/v1/admin/widgets", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("widget_id", suite.InstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := UpdateWidget(req)

	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceWidgetTestSuite) Test00005DeleteWidget() {
	req, _ := http.NewRequest("DELETE", "http://localhost:3333/api/v1/admin/widgets", nil)
	req.Header.Set("Content-Language", "pt-br")
	req.Header.Set("userID", suite.UserID)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("widget_id", suite.InstanceID)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	response := DeleteWidget(req)

	assert.Equal(suite.T(), 200, response.Code)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestServiceWidgetSuite(t *testing.T) {
	suite.Run(t, new(ServiceWidgetTestSuite))
}
