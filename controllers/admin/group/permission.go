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

// UpdatePermissionWidget sends the request to model updating a users in group
func UpdatePermissionWidget(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdatePermissionWidget new transaction", err)
		resp.Render(res, req)
		return
	}

	columns, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdatePermissionWidget request parse", err)
		resp.Render(res, req)
		return
	}

	groupDefinition := &group.Definitions{}
	if err := groupDefinition.UpdatePermissionWidget(trs, chi.URLParam(req, "group_code"), chi.URLParam(req, "widget_code"), columns); err != nil {
		trs.Rollback()
		resp.NewError("UpdatePermissionWidget", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = groupDefinition
	resp.Render(res, req)
}

// DeletePermissionWidget sends the request to model updating a users in group
func DeletePermissionWidget(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeletePermissionWidget new transaction", err)
		resp.Render(res, req)
		return
	}

	groupDefinition := &group.Definitions{}
	if err := groupDefinition.DeletePermissionWidget(trs, chi.URLParam(req, "group_code"), chi.URLParam(req, "widget_code")); err != nil {
		trs.Rollback()
		resp.NewError("DeletePermissionWidget", err)
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

// UpdatePermissionProcess sends the request to model updating a users in group
func UpdatePermissionProcess(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdatePermissionProcess new transaction", err)
		resp.Render(res, req)
		return
	}

	columns, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdatePermissionProcess request parse", err)
		resp.Render(res, req)
		return
	}

	groupDefinition := &group.Definitions{}
	if err := groupDefinition.UpdatePermissionProcess(trs, chi.URLParam(req, "group_code"), chi.URLParam(req, "process_code"), columns); err != nil {
		trs.Rollback()
		resp.NewError("UpdatePermissionProcess", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = groupDefinition
	resp.Render(res, req)
}

// DeletePermissionProcess sends the request to model updating a users in group
func DeletePermissionProcess(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeletePermissionProcess new transaction", err)
		resp.Render(res, req)
		return
	}

	groupDefinition := &group.Definitions{}
	if err := groupDefinition.DeletePermissionProcess(trs, chi.URLParam(req, "group_code"), chi.URLParam(req, "process_code")); err != nil {
		trs.Rollback()
		resp.NewError("DeletePermissionProcess", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = groupDefinition
	resp.Render(res, req)
}

// UpdatePermissionSchema sends the request to model updating a users in group
func UpdatePermissionSchema(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdatePermissionSchema new transaction", err)
		resp.Render(res, req)
		return
	}

	columns, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdatePermissionSchema request parse", err)
		resp.Render(res, req)
		return
	}

	groupDefinition := &group.Definitions{}
	if err := groupDefinition.UpdatePermissionSchema(trs, chi.URLParam(req, "group_code"), chi.URLParam(req, "schema_code"), columns); err != nil {
		trs.Rollback()
		resp.NewError("UpdatePermissionSchema", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = groupDefinition
	resp.Render(res, req)
}

// DeletePermissionSchema sends the request to model updating a users in group
func DeletePermissionSchema(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeletePermissionSchema new transaction", err)
		resp.Render(res, req)
		return
	}

	groupDefinition := &group.Definitions{}
	if err := groupDefinition.DeletePermissionSchema(trs, chi.URLParam(req, "group_code"), chi.URLParam(req, "schema_code")); err != nil {
		trs.Rollback()
		resp.NewError("DeletePermissionSchema", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = groupDefinition
	resp.Render(res, req)
}

// UpdatePermissionSchemaInstance sends the request to model updating a users in group
func UpdatePermissionSchemaInstance(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdatePermissionSchemaInstance new transaction", err)
		resp.Render(res, req)
		return
	}

	columns, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdatePermissionSchemaInstance request parse", err)
		resp.Render(res, req)
		return
	}

	groupDefinition := &group.Definitions{}
	if err := groupDefinition.UpdatePermissionSchemaInstance(trs, chi.URLParam(req, "group_code"), chi.URLParam(req, "schema_code"), columns); err != nil {
		trs.Rollback()
		resp.NewError("UpdatePermissionSchemaInstance", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = groupDefinition
	resp.Render(res, req)
}

// UpdatePermissionSchemaField sends the request to model updating a users in group
func UpdatePermissionSchemaField(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdatePermissionSchemaField new transaction", err)
		resp.Render(res, req)
		return
	}

	columns, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdatePermissionSchemaField request parse", err)
		resp.Render(res, req)
		return
	}

	groupDefinition := &group.Definitions{}
	if err := groupDefinition.UpdatePermissionSchemaField(trs, chi.URLParam(req, "group_code"), chi.URLParam(req, "schema_code"), columns); err != nil {
		trs.Rollback()
		resp.NewError("UpdatePermissionSchemaField", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = groupDefinition
	resp.Render(res, req)
}
