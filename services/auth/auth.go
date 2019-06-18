package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	moduleShared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-mdl-shared/models"
	shared "github.com/agile-work/srv-shared"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/agile-work/srv-shared/sql-builder/db"
	"github.com/agile-work/srv-shared/token"
)

// Login validate credentials and return user token
func Login(r *http.Request) *moduleShared.Response {
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}
	body, _ := ioutil.ReadAll(r.Body)

	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(body, &jsonMap)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorParsingRequest, "Login unmarshal body", err.Error()))
		return response
	}

	_, emailOk := jsonMap["email"]
	_, passwordOk := jsonMap["password"]
	if !emailOk || !passwordOk {
		err = errors.New("Invalid credentials body")
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorParsingRequest, "Login parsing body", err.Error()))
		return response
	}

	user := models.User{}
	emailColumn := fmt.Sprintf("%s.email", shared.TableCoreUsers)
	opt := db.Options{
		Conditions: builder.Equal(emailColumn, jsonMap["email"]),
	}
	err = db.SelectStruct(shared.TableCoreUsers, &user, &opt)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingData, "Login load user", err.Error()))

		return response
	}

	if user.ID == "" {
		err = errors.New("Invalid user email")
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLogin, "Login validation", err.Error()))
		return response
	}

	if user.Password != jsonMap["password"].(string) {
		err = errors.New("Invalid user password")
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLogin, "Login validation", err.Error()))
		return response
	}

	user.Password = ""

	payload := make(map[string]interface{})
	payload["code"] = user.Username
	payload["scope"] = "user"
	payload["user_id"] = user.ID
	payload["language_code"] = user.LanguageCode

	tokenString, err := token.New(payload)
	if err != nil {
		err = errors.New("error generation token")
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLogin, "Login validation", err.Error()))
		return response
	}

	jsonMap["user"] = user
	jsonMap["token"] = tokenString
	delete(jsonMap, "password")
	delete(jsonMap, "email")
	response.Data = jsonMap

	return response
}
