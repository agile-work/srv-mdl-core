package group

import (
	"net/http"

	"github.com/agile-work/srv-mdl-core/models/group"
	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-mdl-shared/util"
	"github.com/agile-work/srv-shared/sql-builder/db"
	"github.com/go-chi/chi"
)

// UpdateTree sends the request to model updating a tree in group
func UpdateTree(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateTree new transaction", err)
		resp.Render(res, req)
		return
	}

	columns, err := util.GetBodyUpdatableJSONColumns(req, false, &group.SecurityTree{}, req.Header.Get("username"), req.Header.Get("Content-Language"))
	if err != nil {
		resp.NewError("UpdateTree request parse", err)
		resp.Render(res, req)
		return
	}

	groupDefinition := &group.Definitions{}
	if err := groupDefinition.UpdateTree(trs, chi.URLParam(req, "group_code"), columns); err != nil {
		trs.Rollback()
		resp.NewError("UpdateTree", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = groupDefinition
	resp.Render(res, req)
}

// UpdateUsers sends the request to model updating a users in group
func UpdateUsers(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateUsers new transaction", err)
		resp.Render(res, req)
		return
	}

	columns, err := util.GetBodyUpdatableJSONColumns(req, false, &group.SecurityUser{}, req.Header.Get("username"), req.Header.Get("Content-Language"))
	if err != nil {
		resp.NewError("UpdateUsers request parse", err)
		resp.Render(res, req)
		return
	}

	groupDefinition := &group.Definitions{}
	if err := groupDefinition.UpdateUsers(trs, chi.URLParam(req, "group_code"), req.Header.Get("username"), columns); err != nil {
		trs.Rollback()
		resp.NewError("UpdateUsers", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = groupDefinition
	resp.Render(res, req)
}

// UpdatePermissionWidgets sends the request to model updating a users in group
func UpdatePermissionWidgets(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdatePermissionWidgets new transaction", err)
		resp.Render(res, req)
		return
	}

	columns, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdatePermissionWidgets request parse", err)
		resp.Render(res, req)
		return
	}

	groupDefinition := &group.Definitions{}
	if err := groupDefinition.UpdatePermissionWidgets(trs, chi.URLParam(req, "group_code"), columns); err != nil {
		trs.Rollback()
		resp.NewError("UpdatePermissionWidgets", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = groupDefinition
	resp.Render(res, req)
}

// DeletePermissionWidgets sends the request to model updating a users in group
func DeletePermissionWidgets(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeletePermissionWidgets new transaction", err)
		resp.Render(res, req)
		return
	}

	groupDefinition := &group.Definitions{}
	if err := groupDefinition.DeletePermissionWidgets(trs, chi.URLParam(req, "group_code"), chi.URLParam(req, "widget_code")); err != nil {
		trs.Rollback()
		resp.NewError("DeletePermissionWidgets", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = groupDefinition
	resp.Render(res, req)
}

// UpdatePermissionProcesses sends the request to model updating a users in group
func UpdatePermissionProcesses(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdatePermissionProcesses new transaction", err)
		resp.Render(res, req)
		return
	}

	columns, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdatePermissionProcesses request parse", err)
		resp.Render(res, req)
		return
	}

	groupDefinition := &group.Definitions{}
	if err := groupDefinition.UpdatePermissionProcesses(trs, chi.URLParam(req, "group_code"), columns); err != nil {
		trs.Rollback()
		resp.NewError("UpdatePermissionProcesses", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = groupDefinition
	resp.Render(res, req)
}

// DeletePermissionProcesses sends the request to model updating a users in group
func DeletePermissionProcesses(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeletePermissionProcesses new transaction", err)
		resp.Render(res, req)
		return
	}

	groupDefinition := &group.Definitions{}
	if err := groupDefinition.DeletePermissionProcesses(trs, chi.URLParam(req, "group_code"), chi.URLParam(req, "process_code")); err != nil {
		trs.Rollback()
		resp.NewError("DeletePermissionProcesses", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = groupDefinition
	resp.Render(res, req)
}
