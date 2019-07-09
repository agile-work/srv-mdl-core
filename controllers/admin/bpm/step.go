package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-mdl-shared/util"

	step "github.com/agile-work/srv-mdl-core/models/bpm"

	"github.com/go-chi/chi"

	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// PostStep sends the request to model creating a new step
func PostStep(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	step := &step.Step{}
	resp := response.New()

	if err := resp.Parse(req, step); err != nil {
		resp.NewError("PostStep response load", err)
		resp.Render(res, req)
		return
	}

	step.WorkflowCode = chi.URLParam(req, "workflow_code")

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("PostStep step new transaction", err)
		resp.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"
	if err := step.Create(trs); err != nil {
		trs.Rollback()
		resp.NewError("PostStep", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = step
	resp.Render(res, req)
}

// GetAllSteps return all step instances from the model
func GetAllSteps(res http.ResponseWriter, req *http.Request) {
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
	steps := &step.Steps{}
	if err := steps.LoadAll(opt); err != nil {
		resp.NewError("GetAllSteps", err)
		resp.Render(res, req)
		return
	}
	resp.Data = steps
	resp.Metadata = metadata
	resp.Render(res, req)
}

// GetStep return only one step from the model
func GetStep(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	step := &step.Step{
		WorkflowCode: chi.URLParam(req, "workflow_code"),
		Code:         chi.URLParam(req, "step_code"),
	}
	if err := step.Load(); err != nil {
		resp.NewError("GetStep", err)
		resp.Render(res, req)
		return
	}
	resp.Data = step
	resp.Render(res, req)
}

// UpdateStep sends the request to model updating a step
func UpdateStep(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	step := &step.Step{}
	resp := response.New()

	if err := resp.Parse(req, step); err != nil {
		resp.NewError("UpdateStep step new transaction", err)
		resp.Render(res, req)
		return
	}

	step.WorkflowCode = chi.URLParam(req, "workflow_code")
	step.Code = chi.URLParam(req, "step_code")

	body, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdateStep", err)
		resp.Render(res, req)
		return
	}

	columns, translations := util.GetColumnsFromBody(body, step)

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateStep step new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := step.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		resp.NewError("UpdateStep", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = step
	resp.Render(res, req)
}

// DeleteStep sends the request to model deleting a step
func DeleteStep(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeleteStep step new transaction", err)
		resp.Render(res, req)
		return
	}

	step := &step.Step{
		WorkflowCode: chi.URLParam(req, "workflow_code"),
		Code:         chi.URLParam(req, "step_code"),
	}
	if err := step.Delete(trs); err != nil {
		trs.Rollback()
		resp.NewError("DeleteStep", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Render(res, req)
}
