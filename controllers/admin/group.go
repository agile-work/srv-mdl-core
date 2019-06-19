package admin

import (
	"net/http"

	"github.com/agile-work/srv-shared/util"

	"github.com/agile-work/srv-mdl-core/models/group"

	"github.com/go-chi/chi"

	mdlShared "github.com/agile-work/srv-mdl-shared"
	mdlSharedModels "github.com/agile-work/srv-mdl-shared/models"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// PostGroup sends the request to model creating a new group
func PostGroup(res http.ResponseWriter, req *http.Request) {
	mdlSharedModels.TranslationFieldsRequestLanguageCode = req.Header.Get("Content-Language")
	group := &group.Group{}
	response := mdlShared.NewResponse()

	if err := response.Load(req, group); err != nil {
		response.NewError(http.StatusInternalServerError, "PostGroup response load", err.Error())
		response.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "PostGroup group new transaction", err.Error())
		response.Render(res, req)
		return
	}

	mdlSharedModels.TranslationFieldsRequestLanguageCode = "all"
	if err := group.Create(trs); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "PostGroup "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()

	response.Data = group
	response.Render(res, req)
}

// GetAllGroups return all group instances from the model
func GetAllGroups(res http.ResponseWriter, req *http.Request) {
	mdlSharedModels.TranslationFieldsRequestLanguageCode = req.Header.Get("Content-Language")
	response := mdlShared.NewResponse()

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "GetAllGroups group new transaction", err.Error())
		response.Render(res, req)
		return
	}

	metaData := mdlShared.Metadata{}
	metaData.Load(req)
	opt := metaData.GenerateDBOptions()
	groups := &group.Groups{}
	if err := groups.LoadAll(trs, opt); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "GetAllGroups "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Data = groups
	response.Metadata = metaData
	response.Render(res, req)
}

// GetGroup return only one group from the model
func GetGroup(res http.ResponseWriter, req *http.Request) {
	mdlSharedModels.TranslationFieldsRequestLanguageCode = req.Header.Get("Content-Language")
	response := mdlShared.NewResponse()

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "GetGroup group new transaction", err.Error())
		response.Render(res, req)
		return
	}

	group := &group.Group{Code: chi.URLParam(req, "group_code")}
	if err := group.Load(trs); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "GetGroup "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Data = group
	response.Render(res, req)
}

// UpdateGroup sends the request to model updating a group
func UpdateGroup(res http.ResponseWriter, req *http.Request) {
	mdlSharedModels.TranslationFieldsRequestLanguageCode = req.Header.Get("Content-Language")
	group := &group.Group{}
	response := mdlShared.NewResponse()

	if err := response.Load(req, group); err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateGroup group new transaction", err.Error())
		response.Render(res, req)
		return
	}

	group.Code = chi.URLParam(req, "group_code")

	body, err := util.GetBody(req)
	if err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateGroup "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}

	columns, translations, err := util.GetColumnsFromBody(body, group)
	if err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateGroup "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateGroup group new transaction", err.Error())
		response.Render(res, req)
		return
	}

	if err := group.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "UpdateGroup "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Data = group
	response.Render(res, req)
}

// DeleteGroup sends the request to model deleting a group
func DeleteGroup(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "DeleteGroup group new transaction", err.Error())
		response.Render(res, req)
		return
	}

	group := &group.Group{Code: chi.URLParam(req, "group_code")}
	if err := group.Delete(trs); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "DeleteGroup "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Render(res, req)
}

// AddUserInGroup sends the request to service deleting an user
func AddUserInGroup(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// DeleteGroupUser sends the request to service deleting a user from a group
func DeleteGroupUser(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// GetAllGroupPermissions return all group instances from the service
func GetAllGroupPermissions(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// PostGroupPermission sends the request to service creating a permission in a group
func PostGroupPermission(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// DeleteGroupPermission sends the request to service deleting a permission from a group
func DeleteGroupPermission(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// GetAllGroupsByUser sends the request to service deleting a permission from a group
func GetAllGroupsByUser(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}
