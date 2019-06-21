package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-mdl-shared/util"

	"github.com/agile-work/srv-mdl-core/models/group"

	"github.com/go-chi/chi"

	"github.com/agile-work/srv-shared/sql-builder/db"
)

// PostGroup sends the request to model creating a new group
func PostGroup(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	group := &group.Group{}
	resp := response.New()

	if err := resp.Parse(req, group); err != nil {
		resp.NewError("PostGroup response load", err)
		resp.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("PostGroup group new transaction", err)
		resp.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"
	if err := group.Create(trs); err != nil {
		trs.Rollback()
		resp.NewError("PostGroup", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = group
	resp.Render(res, req)
}

// GetAllGroups return all group instances from the model
func GetAllGroups(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("GetAllGroups group new transaction", err)
		resp.Render(res, req)
		return
	}

	metaData := response.Metadata{}
	metaData.Load(req)
	opt := metaData.GenerateDBOptions()
	groups := &group.Groups{}
	if err := groups.LoadAll(trs, opt); err != nil {
		trs.Rollback()
		resp.NewError("GetAllGroups", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = groups
	resp.Metadata = metaData
	resp.Render(res, req)
}

// GetGroup return only one group from the model
func GetGroup(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("GetGroup group new transaction", err)
		resp.Render(res, req)
		return
	}

	group := &group.Group{Code: chi.URLParam(req, "group_code")}
	if err := group.Load(trs); err != nil {
		trs.Rollback()
		resp.NewError("GetGroup", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = group
	resp.Render(res, req)
}

// UpdateGroup sends the request to model updating a group
func UpdateGroup(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	group := &group.Group{}
	resp := response.New()

	if err := resp.Parse(req, group); err != nil {
		resp.NewError("UpdateGroup group new transaction", err)
		resp.Render(res, req)
		return
	}

	group.Code = chi.URLParam(req, "group_code")

	body, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdateGroup", err)
		resp.Render(res, req)
		return
	}

	columns, translations, err := util.GetColumnsFromBody(body, group)
	if err != nil {
		resp.NewError("UpdateGroup", err)
		resp.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateGroup group new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := group.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		resp.NewError("UpdateGroup", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = group
	resp.Render(res, req)
}

// DeleteGroup sends the request to model deleting a group
func DeleteGroup(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeleteGroup group new transaction", err)
		resp.Render(res, req)
		return
	}

	group := &group.Group{Code: chi.URLParam(req, "group_code")}
	if err := group.Delete(trs); err != nil {
		trs.Rollback()
		resp.NewError("DeleteGroup", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Render(res, req)
}

// AddUserInGroup sends the request to service deleting an user
func AddUserInGroup(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// DeleteGroupUser sends the request to service deleting a user from a group
func DeleteGroupUser(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetAllGroupPermissions return all group instances from the service
func GetAllGroupPermissions(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// PostGroupPermission sends the request to service creating a permission in a group
func PostGroupPermission(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// DeleteGroupPermission sends the request to service deleting a permission from a group
func DeleteGroupPermission(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetAllGroupsByUser sends the request to service deleting a permission from a group
func GetAllGroupsByUser(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}
