package module

import (
	"net/http"

	"github.com/agile-work/srv-mdl-core/models/module"
	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-shared/sql-builder/db"
	"github.com/go-chi/chi"
)

// AddModuleInstance sends the request add a new instance to the Module
func AddModuleInstance(res http.ResponseWriter, req *http.Request) {
	instance := &module.Instance{}
	resp := response.New()

	if err := resp.Parse(req, instance); err != nil {
		resp.NewError("AddModuleInstance response load", err)
		resp.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("AddModuleInstance new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := instance.Add(trs, chi.URLParam(req, "module_code")); err != nil {
		trs.Rollback()
		resp.NewError("AddModuleInstance", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = instance
	resp.Render(res, req)
}

// DeleteModuleInstance delete an existing instance from the Module
func DeleteModuleInstance(res http.ResponseWriter, req *http.Request) {
	instance := &module.Instance{}
	resp := response.New()

	if err := resp.Parse(req, instance); err != nil {
		resp.NewError("DeleteModuleInstance response load", err)
		resp.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeleteModuleInstance new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := instance.Delete(trs, chi.URLParam(req, "module_code")); err != nil {
		trs.Rollback()
		resp.NewError("DeleteModuleInstance", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Render(res, req)
}
