package admin

import (
	"net/http"

	"github.com/agile-work/srv-shared/util"

	"github.com/agile-work/srv-mdl-shared/models/user"

	"github.com/go-chi/chi"

	mdlShared "github.com/agile-work/srv-mdl-shared"
	mdlSharedModels "github.com/agile-work/srv-mdl-shared/models"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// PostUser sends the request to model creating a new user
func PostUser(res http.ResponseWriter, req *http.Request) {
	mdlSharedModels.TranslationFieldsRequestLanguageCode = req.Header.Get("Content-Language")
	user := &user.User{}
	response := mdlShared.NewResponse()

	if err := response.Load(req, user); err != nil {
		response.NewError(http.StatusInternalServerError, "PostUser response load", err.Error())
		response.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "PostUser user new transaction", err.Error())
		response.Render(res, req)
		return
	}

	mdlSharedModels.TranslationFieldsRequestLanguageCode = "all"
	if err := user.Create(trs); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "PostUser "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()

	response.Data = user
	response.Render(res, req)
}

// GetAllUsers return all user instances from the model
func GetAllUsers(res http.ResponseWriter, req *http.Request) {
	mdlSharedModels.TranslationFieldsRequestLanguageCode = req.Header.Get("Content-Language")
	response := mdlShared.NewResponse()

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "GetAllUsers user new transaction", err.Error())
		response.Render(res, req)
		return
	}

	metaData := mdlShared.Metadata{}
	metaData.Load(req)
	opt := metaData.GenerateDBOptions()
	users := &user.Users{}
	if err := users.LoadAll(trs, opt); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "GetAllUsers "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Data = users
	response.Metadata = metaData
	response.Render(res, req)
}

// GetUser return only one user from the model
func GetUser(res http.ResponseWriter, req *http.Request) {
	mdlSharedModels.TranslationFieldsRequestLanguageCode = req.Header.Get("Content-Language")
	response := mdlShared.NewResponse()

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "GetUser user new transaction", err.Error())
		response.Render(res, req)
		return
	}

	user := &user.User{Username: chi.URLParam(req, "username")}
	if err := user.Load(trs); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "GetUser "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Data = user
	response.Render(res, req)
}

// UpdateUser sends the request to model updating a user
func UpdateUser(res http.ResponseWriter, req *http.Request) {
	mdlSharedModels.TranslationFieldsRequestLanguageCode = req.Header.Get("Content-Language")
	user := &user.User{}
	response := mdlShared.NewResponse()

	if err := response.Load(req, user); err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateUser user new transaction", err.Error())
		response.Render(res, req)
		return
	}

	user.Username = chi.URLParam(req, "username")

	body, err := util.GetBody(req)
	if err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateUser "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}

	columns, translations, err := util.GetColumnsFromBody(body, user)
	if err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateUser "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateUser user new transaction", err.Error())
		response.Render(res, req)
		return
	}

	if err := user.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "UpdateUser "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Data = user
	response.Render(res, req)
}

// DeleteUser sends the request to model deleting a user
func DeleteUser(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "DeleteUser user new transaction", err.Error())
		response.Render(res, req)
		return
	}

	user := &user.User{Username: chi.URLParam(req, "username")}
	if err := user.Delete(trs); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "DeleteUser "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Render(res, req)
}

// AddGroupInUser sends the request to service deleting an user
func AddGroupInUser(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// RemoveGroupFromUser sends the request to service deleting an user
func RemoveGroupFromUser(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// GetAllUsersByGroup return all user instances by group from the service
func GetAllUsersByGroup(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// GetAllPermissionsByUser return all user instances by group from the service
func GetAllPermissionsByUser(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}
