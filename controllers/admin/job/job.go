package job

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-mdl-shared/util"

	"github.com/agile-work/srv-mdl-shared/models/job"

	"github.com/go-chi/chi"

	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// PostJob sends the request to model creating a new job
func PostJob(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	job := &job.Job{}
	resp := response.New()

	if err := resp.Parse(req, job); err != nil {
		resp.NewError("PostJob response load", err)
		resp.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("PostJob user new transaction", err)
		resp.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"
	if err := job.Create(trs); err != nil {
		trs.Rollback()
		resp.NewError("PostJob", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = job
	resp.Render(res, req)
}

// GetAllJobs return all job instances from the model
func GetAllJobs(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	metadata := response.Metadata{}
	if err := metadata.Load(req); err != nil {
		resp.NewError("GetAllJobs metadata parse", err)
		resp.Render(res, req)
		return
	}
	opt := metadata.GenerateDBOptions()
	jobs := &job.Jobs{}
	if err := jobs.LoadAll(opt); err != nil {
		resp.NewError("GetAllJobs", err)
		resp.Render(res, req)
		return
	}
	resp.Data = jobs
	resp.Metadata = metadata
	resp.Render(res, req)
}

// GetJob return only one job from the model
func GetJob(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	job := &job.Job{Code: chi.URLParam(req, "job_code")}
	if err := job.Load(); err != nil {
		resp.NewError("GetJob", err)
		resp.Render(res, req)
		return
	}
	resp.Data = job
	resp.Render(res, req)
}

// UpdateJob sends the request to model updating a job
func UpdateJob(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	job := &job.Job{}
	resp := response.New()

	if err := resp.Parse(req, job); err != nil {
		resp.NewError("UpdateJob job new transaction", err)
		resp.Render(res, req)
		return
	}

	job.Code = chi.URLParam(req, "job_code")

	body, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdateJob job get body", err)
		resp.Render(res, req)
		return
	}

	columns, translations := util.GetColumnsFromBody(body, job)

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateJob job new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := job.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		resp.NewError("UpdateJob", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = job
	resp.Render(res, req)
}

// DeleteJob sends the request to model deleting a job
func DeleteJob(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeleteJob job new transaction", err)
		resp.Render(res, req)
		return
	}

	job := &job.Job{Code: chi.URLParam(req, "job_code")}
	if err := job.Delete(trs); err != nil {
		trs.Rollback()
		resp.NewError("DeleteJob", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Render(res, req)
}
