package group

import (
	"net/http"

	"github.com/agile-work/srv-mdl-core/models/group"
	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-mdl-shared/util"
	"github.com/agile-work/srv-shared/sql-builder/db"
	"github.com/go-chi/chi"
)

// UpdatePermissionModule sends the request to model updating a users in group
func UpdatePermissionModule(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdatePermissionModule new transaction", err)
		resp.Render(res, req)
		return
	}

	columns, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdatePermissionModule request parse", err)
		resp.Render(res, req)
		return
	}

	groupDefinition := &group.Definitions{}
	if err := groupDefinition.UpdatePermissionModule(trs, chi.URLParam(req, "group_code"), chi.URLParam(req, "schema_code"), chi.URLParam(req, "module_code"), columns); err != nil {
		trs.Rollback()
		resp.NewError("UpdatePermissionModule", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = groupDefinition
	resp.Render(res, req)
}

// DeletePermissionModule sends the request to model updating a users in group
func DeletePermissionModule(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeletePermissionModule new transaction", err)
		resp.Render(res, req)
		return
	}

	groupDefinition := &group.Definitions{}
	if err := groupDefinition.DeletePermissionModule(trs, chi.URLParam(req, "group_code"), chi.URLParam(req, "schema_code"), chi.URLParam(req, "module_code")); err != nil {
		trs.Rollback()
		resp.NewError("DeletePermissionModule", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = groupDefinition
	resp.Render(res, req)
}

// UpdatePermissionModuleInstance sends the request to model updating a users in group
func UpdatePermissionModuleInstance(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdatePermissionModuleInstance new transaction", err)
		resp.Render(res, req)
		return
	}

	columns, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdatePermissionModuleInstance request parse", err)
		resp.Render(res, req)
		return
	}

	groupDefinition := &group.Definitions{}
	if err := groupDefinition.UpdatePermissionModuleInstance(trs, chi.URLParam(req, "group_code"), chi.URLParam(req, "schema_code"), chi.URLParam(req, "module_code"), columns); err != nil {
		trs.Rollback()
		resp.NewError("UpdatePermissionModuleInstance", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = groupDefinition
	resp.Render(res, req)
}

// UpdatePermissionModuleField sends the request to model updating a users in group
func UpdatePermissionModuleField(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdatePermissionModuleField new transaction", err)
		resp.Render(res, req)
		return
	}

	columns, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdatePermissionModuleField request parse", err)
		resp.Render(res, req)
		return
	}

	groupDefinition := &group.Definitions{}
	if err := groupDefinition.UpdatePermissionModuleField(trs, chi.URLParam(req, "group_code"), chi.URLParam(req, "schema_code"), chi.URLParam(req, "module_code"), columns); err != nil {
		trs.Rollback()
		resp.NewError("UpdatePermissionModuleField", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = groupDefinition
	resp.Render(res, req)
}

// UpdatePermissionModuleFeature sends the request to model updating a users in group
func UpdatePermissionModuleFeature(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdatePermissionModuleFeature new transaction", err)
		resp.Render(res, req)
		return
	}

	columns, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdatePermissionModuleFeature request parse", err)
		resp.Render(res, req)
		return
	}

	groupDefinition := &group.Definitions{}
	if err := groupDefinition.UpdatePermissionModuleFeature(trs, chi.URLParam(req, "group_code"), chi.URLParam(req, "schema_code"), chi.URLParam(req, "module_code"), chi.URLParam(req, "feature_code"), columns); err != nil {
		trs.Rollback()
		resp.NewError("UpdatePermissionModuleFeature", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = groupDefinition
	resp.Render(res, req)
}
