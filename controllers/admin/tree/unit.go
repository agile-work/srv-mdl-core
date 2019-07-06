package tree

import (
	"net/http"

	"github.com/agile-work/srv-shared/sql-builder/builder"

	"github.com/agile-work/srv-mdl-core/models/tree"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-mdl-shared/util"

	"github.com/go-chi/chi"

	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// PostUnit sends the request to model creating a new unit
func PostUnit(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	unit := &tree.Unit{}
	resp := response.New()

	if err := resp.Parse(req, unit); err != nil {
		resp.NewError("PostUnit response load", err)
		resp.Render(res, req)
		return
	}

	unit.TreeCode = chi.URLParam(req, "tree_code")

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("PostUnit user new transaction", err)
		resp.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"
	if err := unit.Create(trs); err != nil {
		trs.Rollback()
		resp.NewError("PostUnit", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = unit
	resp.Render(res, req)
}

// GetAllUnits return all unit instances from the model
func GetAllUnits(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	metadata := response.Metadata{}
	if err := metadata.Load(req); err != nil {
		resp.NewError("GetLookupInstance metadata parse", err)
		resp.Render(res, req)
		return
	}
	opt := metadata.GenerateDBOptions()
	opt.AddCondition(builder.And(
		builder.Equal("tree_code", chi.URLParam(req, "tree_code")),
		builder.Equal("unit_code", chi.URLParam(req, "unit_code")),
	))
	units := &tree.Units{}
	if err := units.LoadAll(opt); err != nil {
		resp.NewError("GetAllUnits", err)
		resp.Render(res, req)
		return
	}
	resp.Data = units
	resp.Metadata = metadata
	resp.Render(res, req)
}

// GetUnit return only one unit from the model
func GetUnit(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	unit := &tree.Unit{
		TreeCode: chi.URLParam(req, "tree_code"),
		Code:     chi.URLParam(req, "unit_code"),
	}
	if err := unit.Load(); err != nil {
		resp.NewError("GetUnit", err)
		resp.Render(res, req)
		return
	}
	resp.Data = unit
	resp.Render(res, req)
}

// UpdateUnit sends the request to model updating a unit
func UpdateUnit(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	unit := &tree.Unit{}
	resp := response.New()

	if err := resp.Parse(req, unit); err != nil {
		resp.NewError("UpdateUnit unit new transaction", err)
		resp.Render(res, req)
		return
	}

	unit.TreeCode = chi.URLParam(req, "tree_code")
	unit.Code = chi.URLParam(req, "unit_code")

	body, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdateUnit unit get body", err)
		resp.Render(res, req)
		return
	}

	columns, translations := util.GetColumnsFromBody(body, unit)

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateUnit unit new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := unit.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		resp.NewError("UpdateUnit", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = unit
	resp.Render(res, req)
}

// DeleteUnit sends the request to model deleting a unit
func DeleteUnit(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeleteUnit unit new transaction", err)
		resp.Render(res, req)
		return
	}

	unit := &tree.Unit{
		TreeCode: chi.URLParam(req, "tree_code"),
		Code:     chi.URLParam(req, "unit_code"),
	}
	if err := unit.Delete(trs); err != nil {
		trs.Rollback()
		resp.NewError("DeleteUnit", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Render(res, req)
}
