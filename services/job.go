package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/agile-work/srv-mdl-shared/db"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	sql "github.com/agile-work/srv-shared/sql-builder/db"
	"github.com/go-chi/chi"

	moduleShared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-mdl-shared/models"
	shared "github.com/agile-work/srv-shared"
)

// CreateJob persists the request body creating a new object in the database
func CreateJob(r *http.Request) *moduleShared.Response {
	job := models.Job{}

	return db.Create(r, &job, "CreateJob", shared.TableCoreJobs)
}

// LoadAllJobs return all instances from the object
func LoadAllJobs(r *http.Request) *moduleShared.Response {
	jobs := []models.Job{}

	return db.Load(r, &jobs, "LoadAllJobs", shared.TableCoreJobs, nil)
}

// LoadJob return only one object from the database
func LoadJob(r *http.Request) *moduleShared.Response {
	job := models.Job{}
	jobID := chi.URLParam(r, "job_id")
	jobIDColumn := fmt.Sprintf("%s.id", shared.TableCoreJobs)
	condition := builder.Equal(jobIDColumn, jobID)

	return db.Load(r, &job, "LoadJob", shared.TableCoreJobs, condition)
}

// UpdateJob updates object data in the database
func UpdateJob(r *http.Request) *moduleShared.Response {
	jobID := chi.URLParam(r, "job_id")
	jobIDColumn := fmt.Sprintf("%s.id", shared.TableCoreJobs)
	condition := builder.Equal(jobIDColumn, jobID)
	job := models.Job{
		ID: jobID,
	}

	return db.Update(r, &job, "UpdateJob", shared.TableCoreJobs, condition)
}

// DeleteJob deletes object from the database
func DeleteJob(r *http.Request) *moduleShared.Response {
	jobID := chi.URLParam(r, "job_id")
	jobIDColumn := fmt.Sprintf("%s.id", shared.TableCoreJobs)
	condition := builder.Equal(jobIDColumn, jobID)

	return db.Remove(r, "DeleteJob", shared.TableCoreJobs, condition)
}

// CreateJobTask persists the request body creating a new object in the database
func CreateJobTask(r *http.Request) *moduleShared.Response {
	jobTaskID := chi.URLParam(r, "job_id")
	jobTask := models.JobTask{
		JobID: jobTaskID,
	}

	return db.Create(r, &jobTask, "CreateJobTask", shared.TableCoreJobTasks)
}

// LoadAllJobTasks return all instances from the object
func LoadAllJobTasks(r *http.Request) *moduleShared.Response {
	jobTasks := []models.JobTask{}
	jobTaskID := chi.URLParam(r, "job_id")
	jobTaskIDColumn := fmt.Sprintf("%s.job_id", shared.TableCoreJobTasks)
	condition := builder.Equal(jobTaskIDColumn, jobTaskID)

	return db.Load(r, &jobTasks, "LoadAllJobTasks", shared.TableCoreJobTasks, condition)
}

// LoadJobTask return only one object from the database
func LoadJobTask(r *http.Request) *moduleShared.Response {
	jobTask := models.JobTask{}
	jobTaskID := chi.URLParam(r, "job_task_id")
	jobTaskIDColumn := fmt.Sprintf("%s.id", shared.TableCoreJobTasks)
	condition := builder.Equal(jobTaskIDColumn, jobTaskID)

	return db.Load(r, &jobTask, "LoadJobTask", shared.TableCoreJobTasks, condition)
}

// UpdateJobTask updates object data in the database
func UpdateJobTask(r *http.Request) *moduleShared.Response {
	jobTaskID := chi.URLParam(r, "job_task_id")
	jobTaskIDColumn := fmt.Sprintf("%s.id", shared.TableCoreJobTasks)
	condition := builder.Equal(jobTaskIDColumn, jobTaskID)
	jobTask := models.JobTask{
		ID: jobTaskID,
	}

	return db.Update(r, &jobTask, "UpdateJobTask", shared.TableCoreJobTasks, condition)
}

// DeleteJobTask deletes object from the database
func DeleteJobTask(r *http.Request) *moduleShared.Response {
	jobTaskID := chi.URLParam(r, "job_task_id")
	jobTaskIDColumn := fmt.Sprintf("%s.id", shared.TableCoreJobTasks)
	condition := builder.Equal(jobTaskIDColumn, jobTaskID)

	return db.Remove(r, "DeleteJobTask", shared.TableCoreJobTasks, condition)
}

// LoadAllJobFollowersAvaible return all instances from the object
func LoadAllJobFollowersAvaible(r *http.Request) *moduleShared.Response {
	viewFollowersAvailable := []models.ViewFollowerAvailable{}
	activeColumn := fmt.Sprintf("%s.active", shared.ViewCoreUsersAndGroups)
	languageCode := r.Header.Get("Content-Language")
	languageCodeColumn := fmt.Sprintf("%s.language_code", shared.ViewCoreUsersAndGroups)
	condition := builder.And(
		builder.Equal(activeColumn, true),
		builder.Or(
			builder.Equal(languageCodeColumn, languageCode),
			builder.Equal(languageCodeColumn, nil),
		),
	)

	return db.Load(r, &viewFollowersAvailable, "LoadAllJobFollowersAvaible", shared.ViewCoreUsersAndGroups, condition)
}

// InsertFollowerInJob persists the request creating a new object in the database
func InsertFollowerInJob(r *http.Request) *moduleShared.Response {
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}

	jobID := chi.URLParam(r, "job_id")
	followerID := chi.URLParam(r, "follower_id")
	followerType := chi.URLParam(r, "follower_type")

	userID := r.Header.Get("userID")
	now := time.Now()

	statemant := builder.Insert(
		shared.TableCoreJobsFollowers,
		"job_id",
		"follower_id",
		"follower_type",
		"created_by",
		"created_at",
		"updated_by",
		"updated_at",
	).Values(
		jobID,
		followerID,
		followerType,
		userID,
		now,
		userID,
		now,
	)

	err := sql.Exec(statemant)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorInsertingRecord, "InsertFollowerInJob", err.Error()))

		return response
	}

	return response
}

// LoadAllFollowersByJob return all instances from the object
func LoadAllFollowersByJob(r *http.Request) *moduleShared.Response {
	jobFollowers := []models.JobFollowers{}
	jobID := chi.URLParam(r, "job_id")
	jobIDColumn := fmt.Sprintf("%s.job_id", shared.ViewCoreJobFollowers)
	languageCode := r.Header.Get("Content-Language")
	languageCodeColumn := fmt.Sprintf("%s.language_code", shared.ViewCoreJobFollowers)
	followerTypeColumn := fmt.Sprintf("%s.follower_type", shared.ViewCoreJobFollowers)
	condition := builder.And(
		builder.Equal(jobIDColumn, jobID),
		builder.Or(
			builder.Equal(followerTypeColumn, "user"),
			builder.And(
				builder.Equal(followerTypeColumn, "group"),
				builder.Equal(languageCodeColumn, languageCode),
			),
		),
	)

	return db.Load(r, &jobFollowers, "LoadAllFollowersByJob", shared.ViewCoreJobFollowers, condition)
}

// RemoveFollowerFromJob deletes object from the database
func RemoveFollowerFromJob(r *http.Request) *moduleShared.Response {
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}

	jobID := chi.URLParam(r, "job_id")
	followerID := chi.URLParam(r, "follower_id")

	statemant := builder.Delete(shared.TableCoreJobsFollowers).Where(
		builder.And(
			builder.Equal("job_id", jobID),
			builder.Equal("follower_id", followerID),
		),
	)

	err := sql.Exec(statemant)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorDeletingData, "RemoveFollowerFromJob", err.Error()))

		return response
	}

	return response
}

// CreateJobInstance persists the request body creating a new object in the database
func CreateJobInstance(r *http.Request) *moduleShared.Response {
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}

	jobID := chi.URLParam(r, "job_id")
	jobTable := shared.TableCoreJobs
	jobIDColumn := fmt.Sprintf("%s.id", jobTable)
	condition := builder.Equal(jobIDColumn, jobID)
	job := models.Job{}

	err := sql.LoadStruct(jobTable, &job, condition)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorParsingRequest, "CreateJobInstance load job", err.Error()))

		return response
	}

	params := map[string]interface{}{}
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorParsingRequest, "CreateJobInstance unmarshal body", err.Error()))

		return response
	}

	userID := r.Header.Get("userID")

	_, err = moduleShared.CreateJobInstance(userID, job.Code, params)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorJobExecution, "CreateJobInstance", err.Error()))

		return response
	}

	return response
}
