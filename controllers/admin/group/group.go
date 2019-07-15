package group

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
	resp := response.New()

	metadata := response.Metadata{}
	if err := metadata.Load(req); err != nil {
		resp.NewError("GetLookupInstance metadata parse", err)
		resp.Render(res, req)
		return
	}
	opt := metadata.GenerateDBOptions()
	groups := &group.Groups{}
	if err := groups.LoadAll(opt); err != nil {
		resp.NewError("GetAllGroups", err)
		resp.Render(res, req)
		return
	}
	translation.SetSliceTranslationsLanguage(groups, req.Header.Get("Content-Language"))
	resp.Data = groups
	resp.Metadata = metadata
	resp.Render(res, req)
}

// GetGroup return only one group from the model
func GetGroup(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	group := &group.Group{Code: chi.URLParam(req, "group_code")}
	if err := group.Load(); err != nil {
		resp.NewError("GetGroup", err)
		resp.Render(res, req)
		return
	}
	translation.SetStructTranslationsLanguage(group, req.Header.Get("Content-Language"))
	resp.Data = group
	resp.Render(res, req)
}

// UpdateGroup sends the request to model updating a group
func UpdateGroup(res http.ResponseWriter, req *http.Request) {
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

	columns, translations := util.GetColumnsFromBody(body, group)

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

// UpdateTree sends the request to model updating a tree in group
func UpdateTree(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	grp := &group.Group{}
	tree := &group.SecurityTree{}
	resp := response.New()

	if err := resp.Parse(req, tree); err != nil {
		resp.NewError("UpdateGroup group new transaction", err)
		resp.Render(res, req)
		return
	}

	grp.Code = chi.URLParam(req, "group_code")
	grp.Tree = *tree

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateGroup group new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := grp.UpdateTree(trs); err != nil {
		trs.Rollback()
		resp.NewError("UpdateGroup", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = grp
	resp.Render(res, req)
}

// UpdateUsers sends the request to model updating a users in group
func UpdateUsers(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	grp := &group.Group{}
	users := &group.SecurityUser{}
	resp := response.New()

	if err := resp.Parse(req, users); err != nil {
		resp.NewError("UpdateUsers group new transaction", err)
		resp.Render(res, req)
		return
	}

	body, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdateUsers group", err)
	}

	grp.Code = chi.URLParam(req, "group_code")
	grp.Users = *users

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateUsers group new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := grp.UpdateUsers(trs, req.Header.Get("Username"), util.GetBodyColumns(body)); err != nil {
		trs.Rollback()
		resp.NewError("UpdateUsers", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = grp
	resp.Render(res, req)
}
