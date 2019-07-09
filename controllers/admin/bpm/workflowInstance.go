package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-mdl-shared/util"

	workflowInstance "github.com/agile-work/srv-mdl-core/models/bpm"

	"github.com/go-chi/chi"

	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// PostWorkflowInstance sends the request to model creating a new workflowInstance
func PostWorkflowInstance(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	workflowInstance := &workflowInstance.WorkflowInstance{}
	resp := response.New()

	if err := resp.Parse(req, workflowInstance); err != nil {
		resp.NewError("PostWorkflowInstance response load", err)
		resp.Render(res, req)
		return
	}

	workflowInstance.WorkflowCode = chi.URLParam(req, "workflow_code")

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("PostWorkflowInstance workflowInstance new transaction", err)
		resp.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"
	if err := workflowInstance.Create(trs); err != nil {
		trs.Rollback()
		resp.NewError("PostWorkflowInstance", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = workflowInstance
	resp.Render(res, req)
}

// GetAllWorkflowInstances return all workflowInstance instances from the model
func GetAllWorkflowInstances(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	metadata := response.Metadata{}
	if err := metadata.Load(req); err != nil {
		resp.NewError("GetLookupInstance metadata parse", err)
		resp.Render(res, req)
		return
	}
	opt := metadata.GenerateDBOptions()
	opt.AddCondition(
		builder.Equal("bpm_code", chi.URLParam(req, "workflow_code")),
	)
	workflowInstances := &workflowInstance.WorkflowInstances{}
	if err := workflowInstances.LoadAll(opt); err != nil {
		resp.NewError("GetAllWorkflowInstances", err)
		resp.Render(res, req)
		return
	}
	resp.Data = workflowInstances
	resp.Metadata = metadata
	resp.Render(res, req)
}

// GetWorkflowInstance return only one workflowInstance from the model
func GetWorkflowInstance(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	workflowInstance := &workflowInstance.WorkflowInstance{
		WorkflowCode: chi.URLParam(req, "workflow_code"),
		ID:           chi.URLParam(req, "instance_id"),
	}
	if err := workflowInstance.Load(); err != nil {
		resp.NewError("GetWorkflowInstance", err)
		resp.Render(res, req)
		return
	}
	resp.Data = workflowInstance
	resp.Render(res, req)
}

// UpdateWorkflowInstance sends the request to model updating a workflowInstance
func UpdateWorkflowInstance(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	workflowInstance := &workflowInstance.WorkflowInstance{}
	resp := response.New()

	if err := resp.Parse(req, workflowInstance); err != nil {
		resp.NewError("UpdateWorkflowInstance workflowInstance new transaction", err)
		resp.Render(res, req)
		return
	}

	workflowInstance.WorkflowCode = chi.URLParam(req, "workflow_code")
	workflowInstance.ID = chi.URLParam(req, "instance_id")

	body, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdateWorkflowInstance", err)
		resp.Render(res, req)
		return
	}

	columns, translations := util.GetColumnsFromBody(body, workflowInstance)

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateWorkflowInstance workflowInstance new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := workflowInstance.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		resp.NewError("UpdateWorkflowInstance", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = workflowInstance
	resp.Render(res, req)
}

// DeleteWorkflowInstance sends the request to model deleting a workflowInstance
func DeleteWorkflowInstance(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeleteWorkflowInstance workflowInstance new transaction", err)
		resp.Render(res, req)
		return
	}

	workflowInstance := &workflowInstance.WorkflowInstance{
		WorkflowCode: chi.URLParam(req, "workflow_code"),
		ID:           chi.URLParam(req, "instance_id"),
	}
	if err := workflowInstance.Delete(trs); err != nil {
		trs.Rollback()
		resp.NewError("DeleteWorkflowInstance", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Render(res, req)
}
