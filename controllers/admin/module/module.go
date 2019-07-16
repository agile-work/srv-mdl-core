package module

import (
	"net/http"

	"github.com/agile-work/srv-mdl-core/models/module"
	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-mdl-shared/util"
	"github.com/agile-work/srv-shared/sql-builder/db"
	"github.com/go-chi/chi"
)

// RegisterModule sends the request to register a new Module
func RegisterModule(res http.ResponseWriter, req *http.Request) {
	mdl := &module.Module{}
	resp := response.New()

	if err := resp.Parse(req, mdl); err != nil {
		resp.NewError("PostModule response load", err)
		resp.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("PostModule Module new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := mdl.Register(trs); err != nil {
		trs.Rollback()
		resp.NewError("PostModule", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = mdl
	resp.Render(res, req)
}

// GetAllModules return all Module instances from the model
func GetAllModules(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	metadata := response.Metadata{}
	if err := metadata.Load(req); err != nil {
		resp.NewError("GetLookupInstance metadata parse", err)
		resp.Render(res, req)
		return
	}
	opt := metadata.GenerateDBOptions()
	modules := &module.Modules{}
	if err := modules.LoadAll(opt); err != nil {
		resp.NewError("GetAllModules", err)
		resp.Render(res, req)
		return
	}
	translation.SetSliceTranslationsLanguage(modules, req.Header.Get("Content-Language"))
	resp.Data = modules
	resp.Metadata = metadata
	resp.Render(res, req)
}

// GetModule return only one Module from the model
func GetModule(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Module-Language")
	resp := response.New()

	mdl := &module.Module{Code: chi.URLParam(req, "module_code")}
	if err := mdl.Load(); err != nil {

		resp.NewError("GetModule", err)
		resp.Render(res, req)
		return
	}
	translation.SetStructTranslationsLanguage(mdl, req.Header.Get("Content-Language"))
	resp.Data = mdl
	resp.Render(res, req)
}

// UpdateModule sends the request to model updating a Module
func UpdateModule(res http.ResponseWriter, req *http.Request) {
	mdl := &module.Module{}
	resp := response.New()

	if err := resp.Parse(req, mdl); err != nil {
		resp.NewError("UpdateModule Module new transaction", err)
		resp.Render(res, req)
		return
	}

	mdl.Code = chi.URLParam(req, "module_code")

	body, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdateModule", err)
		resp.Render(res, req)
		return
	}

	columns, translations := util.GetColumnsFromBody(body, mdl)

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateModule Module new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := mdl.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		resp.NewError("UpdateModule", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = mdl
	resp.Render(res, req)
}

// DeleteModule sends the request to model deleting a Module
func DeleteModule(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeleteModule Module new transaction", err)
		resp.Render(res, req)
		return
	}

	mdl := &module.Module{Code: chi.URLParam(req, "module_code")}
	if err := mdl.Delete(trs); err != nil {
		trs.Rollback()
		resp.NewError("DeleteModule", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Render(res, req)
}
