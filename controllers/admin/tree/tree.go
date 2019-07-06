package tree

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-mdl-shared/util"

	"github.com/agile-work/srv-mdl-core/models/tree"

	"github.com/go-chi/chi"

	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// PostTree sends the request to model creating a new tree
func PostTree(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	tree := &tree.Tree{}
	resp := response.New()

	if err := resp.Parse(req, tree); err != nil {
		resp.NewError("PostTree response load", err)
		resp.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("PostTree user new transaction", err)
		resp.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"
	if err := tree.Create(trs); err != nil {
		trs.Rollback()
		resp.NewError("PostTree", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = tree
	resp.Render(res, req)
}

// GetAllTrees return all tree instances from the model
func GetAllTrees(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	metadata := response.Metadata{}
	if err := metadata.Load(req); err != nil {
		resp.NewError("GetLookupInstance metadata parse", err)
		resp.Render(res, req)
		return
	}
	opt := metadata.GenerateDBOptions()
	trees := &tree.Trees{}
	if err := trees.LoadAll(opt); err != nil {
		resp.NewError("GetAllTrees", err)
		resp.Render(res, req)
		return
	}
	resp.Data = trees
	resp.Metadata = metadata
	resp.Render(res, req)
}

// GetTree return only one tree from the model
func GetTree(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	tree := &tree.Tree{Code: chi.URLParam(req, "tree_code")}
	if err := tree.Load(); err != nil {
		resp.NewError("GetTree", err)
		resp.Render(res, req)
		return
	}
	resp.Data = tree
	resp.Render(res, req)
}

// UpdateTree sends the request to model updating a tree
func UpdateTree(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	tree := &tree.Tree{}
	resp := response.New()

	if err := resp.Parse(req, tree); err != nil {
		resp.NewError("UpdateTree tree new transaction", err)
		resp.Render(res, req)
		return
	}

	tree.Code = chi.URLParam(req, "tree_code")

	body, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdateTree tree get body", err)
		resp.Render(res, req)
		return
	}

	columns, translations := util.GetColumnsFromBody(body, tree)

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateTree tree new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := tree.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		resp.NewError("UpdateTree", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = tree
	resp.Render(res, req)
}

// DeleteTree sends the request to model deleting a tree
func DeleteTree(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeleteTree tree new transaction", err)
		resp.Render(res, req)
		return
	}

	tree := &tree.Tree{Code: chi.URLParam(req, "tree_code")}
	if err := tree.Delete(trs); err != nil {
		trs.Rollback()
		resp.NewError("DeleteTree", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Render(res, req)
}
