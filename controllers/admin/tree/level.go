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

// PostLevel sends the request to model creating a new level
func PostLevel(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	level := &tree.Level{}
	resp := response.New()

	if err := resp.Parse(req, level); err != nil {
		resp.NewError("PostLevel response load", err)
		resp.Render(res, req)
		return
	}

	level.TreeCode = chi.URLParam(req, "tree_code")

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("PostLevel user new transaction", err)
		resp.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"
	if err := level.Create(trs); err != nil {
		trs.Rollback()
		resp.NewError("PostLevel", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = level
	resp.Render(res, req)
}

// GetAllLevels return all level instances from the model
func GetAllLevels(res http.ResponseWriter, req *http.Request) {
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
		builder.Equal("level_code", chi.URLParam(req, "level_code")),
	))
	levels := &tree.Levels{}
	if err := levels.LoadAll(opt); err != nil {
		resp.NewError("GetAllLevels", err)
		resp.Render(res, req)
		return
	}
	resp.Data = levels
	resp.Metadata = metadata
	resp.Render(res, req)
}

// GetLevel return only one level from the model
func GetLevel(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	level := &tree.Level{
		TreeCode: chi.URLParam(req, "tree_code"),
		Code:     chi.URLParam(req, "level_code"),
	}
	if err := level.Load(); err != nil {
		resp.NewError("GetLevel", err)
		resp.Render(res, req)
		return
	}
	resp.Data = level
	resp.Render(res, req)
}

// UpdateLevel sends the request to model updating a level
func UpdateLevel(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	level := &tree.Level{}
	resp := response.New()

	if err := resp.Parse(req, level); err != nil {
		resp.NewError("UpdateLevel level new transaction", err)
		resp.Render(res, req)
		return
	}

	level.TreeCode = chi.URLParam(req, "tree_code")
	level.Code = chi.URLParam(req, "level_code")

	body, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdateLevel level get body", err)
		resp.Render(res, req)
		return
	}

	columns, translations := util.GetColumnsFromBody(body, level)

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateLevel level new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := level.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		resp.NewError("UpdateLevel", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = level
	resp.Render(res, req)
}

// DeleteLevel sends the request to model deleting a level
func DeleteLevel(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeleteLevel level new transaction", err)
		resp.Render(res, req)
		return
	}

	level := &tree.Level{
		TreeCode: chi.URLParam(req, "tree_code"),
		Code:     chi.URLParam(req, "level_code"),
	}
	if err := level.Delete(trs); err != nil {
		trs.Rollback()
		resp.NewError("DeleteLevel", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Render(res, req)
}
