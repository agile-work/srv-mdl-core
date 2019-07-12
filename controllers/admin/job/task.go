package job

import (
	"net/http"

	"github.com/agile-work/srv-shared/sql-builder/builder"

	"github.com/agile-work/srv-mdl-shared/models/job"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-mdl-shared/util"

	"github.com/go-chi/chi"

	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// PostTask sends the request to model creating a new task
func PostTask(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	task := &job.Task{}
	resp := response.New()

	if err := resp.Parse(req, task); err != nil {
		resp.NewError("PostTask response load", err)
		resp.Render(res, req)
		return
	}

	task.JobCode = chi.URLParam(req, "job_code")

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("PostTask user new transaction", err)
		resp.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"
	if err := task.Create(trs); err != nil {
		trs.Rollback()
		resp.NewError("PostTask", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = task
	resp.Render(res, req)
}

// GetAllTasks return all task instances from the model
func GetAllTasks(res http.ResponseWriter, req *http.Request) {
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
		builder.Equal("job_code", chi.URLParam(req, "job_code")),
		builder.Equal("task_code", chi.URLParam(req, "task_code")),
	))
	tasks := &job.Tasks{}
	if err := tasks.LoadAll(opt); err != nil {
		resp.NewError("GetAllTasks", err)
		resp.Render(res, req)
		return
	}
	resp.Data = tasks
	resp.Metadata = metadata
	resp.Render(res, req)
}

// GetTask return only one task from the model
func GetTask(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	task := &job.Task{
		JobCode: chi.URLParam(req, "job_code"),
		Code:    chi.URLParam(req, "task_code"),
	}
	if err := task.Load(); err != nil {
		resp.NewError("GetTask", err)
		resp.Render(res, req)
		return
	}
	resp.Data = task
	resp.Render(res, req)
}

// UpdateTask sends the request to model updating a task
func UpdateTask(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	task := &job.Task{}
	resp := response.New()

	if err := resp.Parse(req, task); err != nil {
		resp.NewError("UpdateTask task new transaction", err)
		resp.Render(res, req)
		return
	}

	task.JobCode = chi.URLParam(req, "job_code")
	task.Code = chi.URLParam(req, "task_code")

	body, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdateTask task get body", err)
		resp.Render(res, req)
		return
	}

	columns, translations := util.GetColumnsFromBody(body, task)

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateTask task new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := task.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		resp.NewError("UpdateTask", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = task
	resp.Render(res, req)
}

// DeleteTask sends the request to model deleting a task
func DeleteTask(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeleteTask task new transaction", err)
		resp.Render(res, req)
		return
	}

	task := &job.Task{
		JobCode: chi.URLParam(req, "job_code"),
		Code:    chi.URLParam(req, "task_code"),
	}
	if err := task.Delete(trs); err != nil {
		trs.Rollback()
		resp.NewError("DeleteTask", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Render(res, req)
}
