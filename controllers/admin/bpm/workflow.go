package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-mdl-shared/util"

	workflow "github.com/agile-work/srv-mdl-core/models/bpm"

	"github.com/go-chi/chi"

	"github.com/agile-work/srv-shared/sql-builder/db"
)

// PostWorkflow sends the request to model creating a new workflow
func PostWorkflow(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	workflow := &workflow.Workflow{}
	resp := response.New()

	if err := resp.Parse(req, workflow); err != nil {
		resp.NewError("PostWorkflow response load", err)
		resp.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("PostWorkflow workflow new transaction", err)
		resp.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"
	if err := workflow.Create(trs); err != nil {
		trs.Rollback()
		resp.NewError("PostWorkflow", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = workflow
	resp.Render(res, req)
}

// GetAllWorkflows return all workflow instances from the model
func GetAllWorkflows(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	metadata := response.Metadata{}
	if err := metadata.Load(req); err != nil {
		resp.NewError("GetLookupInstance metadata parse", err)
		resp.Render(res, req)
		return
	}
	opt := metadata.GenerateDBOptions()
	workflows := &workflow.Workflows{}
	if err := workflows.LoadAll(opt); err != nil {
		resp.NewError("GetAllWorkflows", err)
		resp.Render(res, req)
		return
	}
	resp.Data = workflows
	resp.Metadata = metadata
	resp.Render(res, req)
}

// GetWorkflow return only one workflow from the model
func GetWorkflow(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	workflow := &workflow.Workflow{Code: chi.URLParam(req, "workflow_code")}
	if err := workflow.Load(); err != nil {
		resp.NewError("GetWorkflow", err)
		resp.Render(res, req)
		return
	}
	resp.Data = workflow
	resp.Render(res, req)
}

// UpdateWorkflow sends the request to model updating a workflow
func UpdateWorkflow(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	workflow := &workflow.Workflow{}
	resp := response.New()

	if err := resp.Parse(req, workflow); err != nil {
		resp.NewError("UpdateWorkflow workflow new transaction", err)
		resp.Render(res, req)
		return
	}

	workflow.Code = chi.URLParam(req, "workflow_code")

	body, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdateWorkflow", err)
		resp.Render(res, req)
		return
	}

	columns, translations := util.GetColumnsFromBody(body, workflow)

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateWorkflow workflow new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := workflow.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		resp.NewError("UpdateWorkflow", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = workflow
	resp.Render(res, req)
}

// DeleteWorkflow sends the request to model deleting a workflow
func DeleteWorkflow(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeleteWorkflow workflow new transaction", err)
		resp.Render(res, req)
		return
	}

	workflow := &workflow.Workflow{Code: chi.URLParam(req, "workflow_code")}
	if err := workflow.Delete(trs); err != nil {
		trs.Rollback()
		resp.NewError("DeleteWorkflow", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Render(res, req)
}
