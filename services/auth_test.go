package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/agile-work/srv-shared/sql-builder/db"
	"github.com/agile-work/srv-mdl-core/config"
	moduleShared "github.com/agile-work/srv-mdl-shared"
shared "github.com/agile-work/srv-shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServiceAuthTestSuite struct {
	suite.Suite
	InstanceID string
}

func (suite *ServiceAuthTestSuite) SetupTest() {
	config, _ := config.NewConfig("..\\config.toml")
	db.Connect(config.Host, config.Port, config.User, config.Password, config.DBName, false)
}

func (suite *ServiceAuthTestSuite) Test00001LoginSuccess() {
	data := map[string]interface{}{
		"email":    "admin@domain.com",
		"password": "123456",
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/auth/login", bytes.NewBuffer(jsonData))

	response := Login(req)

	assert.NotNil(suite.T(), response.Data != nil, "response.Data should not be null")
	assert.Equal(suite.T(), 200, response.Code)
}

func (suite *ServiceAuthTestSuite) Test00002LoginErrorUnmarshal() {
	data := "teste"
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/auth/login", bytes.NewBuffer(jsonData))

	res := Login(req)

	response := moduleShared.Response{}
	jsonMap, _ := json.Marshal(res)
	json.Unmarshal(jsonMap, &response)

	assert.Equal(suite.T(), "Login unmarshal body", response.Errors[0].Scope)
}

func (suite *ServiceAuthTestSuite) Test00002LoginErrorParsing() {
	data := map[string]interface{}{
		"email": "admin@domain.com",
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/auth/login", bytes.NewBuffer(jsonData))

	res := Login(req)

	response := moduleShared.Response{}
	jsonMap, _ := json.Marshal(res)
	json.Unmarshal(jsonMap, &response)

	assert.Equal(suite.T(), "Invalid credentials body", response.Errors[0].Error)
}

func (suite *ServiceAuthTestSuite) Test00003LoginErrorEmailInvalid() {
	data := map[string]interface{}{
		"email":    "admin@domain",
		"password": "123456",
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/auth/login", bytes.NewBuffer(jsonData))

	res := Login(req)

	response := moduleShared.Response{}
	jsonMap, _ := json.Marshal(res)
	json.Unmarshal(jsonMap, &response)

	assert.Equal(suite.T(), "Invalid user email", response.Errors[0].Error)
}

func (suite *ServiceAuthTestSuite) Test00004LoginErrorPasswordInvalid() {
	data := map[string]interface{}{
		"email":    "admin@domain.com",
		"password": "1234567",
	}
	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", "http://localhost:3333/api/v1/auth/login", bytes.NewBuffer(jsonData))

	res := Login(req)

	response := moduleShared.Response{}
	jsonMap, _ := json.Marshal(res)
	json.Unmarshal(jsonMap, &response)

	assert.Equal(suite.T(), "Invalid user password", response.Errors[0].Error)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestServiceAuthSuite(t *testing.T) {
	suite.Run(t, new(ServiceAuthTestSuite))
}
