package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-mdl-shared/util"

	stepInstance "github.com/agile-work/srv-mdl-core/models/bpm"

	"github.com/go-chi/chi"

	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// PostStepInstance sends the request to model creating a new stepInstance
func PostStepInstance(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	stepInstance := &stepInstance.StepInstance{}
	resp := response.New()

	if err := resp.Parse(req, stepInstance); err != nil {
		resp.NewError("PostStepInstance response load", err)
		resp.Render(res, req)
		return
	}

	stepInstance.WorkflowCode = chi.URLParam(req, "workflow_code")
	stepInstance.InstanceID = chi.URLParam(req, "instance_id")

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("PostStepInstance stepInstance new transaction", err)
		resp.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"
	if err := stepInstance.Create(trs); err != nil {
		trs.Rollback()
		resp.NewError("PostStepInstance", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = stepInstance
	resp.Render(res, req)
}

// GetAllStepInstances return all stepInstance instances from the model
func GetAllStepInstances(res http.ResponseWriter, req *http.Request) {
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
		builder.Equal("bpm_code", chi.URLParam(req, "workflow_code")),
		builder.Equal("bpm_instance_id", chi.URLParam(req, "instance_id")),
	))
	stepInstances := &stepInstance.StepInstances{}
	if err := stepInstances.LoadAll(opt); err != nil {
		resp.NewError("GetAllStepInstances", err)
		resp.Render(res, req)
		return
	}
	resp.Data = stepInstances
	resp.Metadata = metadata
	resp.Render(res, req)
}

// GetStepInstance return only one stepInstance from the model
func GetStepInstance(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	stepInstance := &stepInstance.StepInstance{
		WorkflowCode: chi.URLParam(req, "workflow_code"),
		InstanceID:   chi.URLParam(req, "instance_id"),
		ID:           chi.URLParam(req, "step_id"),
	}
	if err := stepInstance.Load(); err != nil {
		resp.NewError("GetStepInstance", err)
		resp.Render(res, req)
		return
	}
	resp.Data = stepInstance
	resp.Render(res, req)
}

// UpdateStepInstance sends the request to model updating a stepInstance
func UpdateStepInstance(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	stepInstance := &stepInstance.StepInstance{}
	resp := response.New()

	if err := resp.Parse(req, stepInstance); err != nil {
		resp.NewError("UpdateStepInstance stepInstance new transaction", err)
		resp.Render(res, req)
		return
	}

	stepInstance.WorkflowCode = chi.URLParam(req, "workflow_code")
	stepInstance.InstanceID = chi.URLParam(req, "instance_id")
	stepInstance.ID = chi.URLParam(req, "step_id")

	body, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdateStepInstance", err)
		resp.Render(res, req)
		return
	}

	columns, translations := util.GetColumnsFromBody(body, stepInstance)

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateStepInstance stepInstance new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := stepInstance.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		resp.NewError("UpdateStepInstance", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = stepInstance
	resp.Render(res, req)
}

// DeleteStepInstance sends the request to model deleting a stepInstance
func DeleteStepInstance(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeleteStepInstance stepInstance new transaction", err)
		resp.Render(res, req)
		return
	}

	stepInstance := &stepInstance.StepInstance{
		WorkflowCode: chi.URLParam(req, "workflow_code"),
		InstanceID:   chi.URLParam(req, "instance_id"),
		ID:           chi.URLParam(req, "step_id"),
	}
	if err := stepInstance.Delete(trs); err != nil {
		trs.Rollback()
		resp.NewError("DeleteStepInstance", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Render(res, req)
}
