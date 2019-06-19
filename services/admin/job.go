package admin

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/agile-work/srv-mdl-shared/db"
// 	"github.com/agile-work/srv-shared/sql-builder/builder"
// 	sql "github.com/agile-work/srv-shared/sql-builder/db"
// 	"github.com/go-chi/chi"

// 	mdlShared "github.com/agile-work/srv-mdl-shared"
// 	"github.com/agile-work/srv-mdl-shared/models"
// 	shared "github.com/agile-work/srv-shared"
// )

// // CreateJob persists the request body creating a new object in the database
// func CreateJob(r *http.Request) *mdlShared.Response {
// 	job := models.Job{}

// 	return db.Create(r, &job, "CreateJob", shared.TableCoreJobs)
// }

// // LoadAllJobs return all instances from the object
// func LoadAllJobs(r *http.Request) *mdlShared.Response {
// 	jobs := []models.Job{}

// 	return db.Load(r, &jobs, "LoadAllJobs", shared.TableCoreJobs, nil)
// }

// // LoadJob return only one object from the database
// func LoadJob(r *http.Request) *mdlShared.Response {
// 	job := models.Job{}
// 	jobID := chi.URLParam(r, "job_id")
// 	jobIDColumn := fmt.Sprintf("%s.id", shared.TableCoreJobs)
// 	condition := builder.Equal(jobIDColumn, jobID)

// 	return db.Load(r, &job, "LoadJob", shared.TableCoreJobs, condition)
// }

// // UpdateJob updates object data in the database
// func UpdateJob(r *http.Request) *mdlShared.Response {
// 	jobID := chi.URLParam(r, "job_id")
// 	jobIDColumn := fmt.Sprintf("%s.id", shared.TableCoreJobs)
// 	condition := builder.Equal(jobIDColumn, jobID)
// 	job := models.Job{
// 		ID: jobID,
// 	}

// 	return db.Update(r, &job, "UpdateJob", shared.TableCoreJobs, condition)
// }

// // DeleteJob deletes object from the database
// func DeleteJob(r *http.Request) *mdlShared.Response {
// 	jobID := chi.URLParam(r, "job_id")
// 	jobIDColumn := fmt.Sprintf("%s.id", shared.TableCoreJobs)
// 	condition := builder.Equal(jobIDColumn, jobID)

// 	return db.Remove(r, "DeleteJob", shared.TableCoreJobs, condition)
// }

// // LoadAllJobInstances return all instances from the object
// func LoadAllJobInstances(r *http.Request) *mdlShared.Response {
// 	viewJobInstances := []models.ViewJobInstance{}
// 	languageCode := r.Header.Get("Content-Language")
// 	languageCodeColumn := fmt.Sprintf("%s.language_code", shared.ViewCoreJobInstances)
// 	condition := builder.Equal(languageCodeColumn, languageCode)

// 	return db.Load(r, &viewJobInstances, "LoadAllJobInstances", shared.ViewCoreJobInstances, condition)
// }

// // CreateJobTask persists the request body creating a new object in the database
// func CreateJobTask(r *http.Request) *mdlShared.Response {
// 	jobTaskID := chi.URLParam(r, "job_id")
// 	jobTask := models.JobTask{
// 		JobID: jobTaskID,
// 	}

// 	return db.Create(r, &jobTask, "CreateJobTask", shared.TableCoreJobTasks)
// }

// // LoadAllJobTasks return all instances from the object
// func LoadAllJobTasks(r *http.Request) *mdlShared.Response {
// 	jobTasks := []models.JobTask{}
// 	jobTaskID := chi.URLParam(r, "job_id")
// 	jobTaskIDColumn := fmt.Sprintf("%s.job_id", shared.TableCoreJobTasks)
// 	condition := builder.Equal(jobTaskIDColumn, jobTaskID)

// 	return db.Load(r, &jobTasks, "LoadAllJobTasks", shared.TableCoreJobTasks, condition)
// }

// // LoadJobTask return only one object from the database
// func LoadJobTask(r *http.Request) *mdlShared.Response {
// 	jobTask := models.JobTask{}
// 	jobTaskID := chi.URLParam(r, "job_task_id")
// 	jobTaskIDColumn := fmt.Sprintf("%s.id", shared.TableCoreJobTasks)
// 	condition := builder.Equal(jobTaskIDColumn, jobTaskID)

// 	return db.Load(r, &jobTask, "LoadJobTask", shared.TableCoreJobTasks, condition)
// }

// // UpdateJobTask updates object data in the database
// func UpdateJobTask(r *http.Request) *mdlShared.Response {
// 	jobTaskID := chi.URLParam(r, "job_task_id")
// 	jobTaskIDColumn := fmt.Sprintf("%s.id", shared.TableCoreJobTasks)
// 	condition := builder.Equal(jobTaskIDColumn, jobTaskID)
// 	jobTask := models.JobTask{
// 		ID: jobTaskID,
// 	}

// 	return db.Update(r, &jobTask, "UpdateJobTask", shared.TableCoreJobTasks, condition)
// }

// // DeleteJobTask deletes object from the database
// func DeleteJobTask(r *http.Request) *mdlShared.Response {
// 	jobTaskID := chi.URLParam(r, "job_task_id")
// 	jobTaskIDColumn := fmt.Sprintf("%s.id", shared.TableCoreJobTasks)
// 	condition := builder.Equal(jobTaskIDColumn, jobTaskID)

// 	return db.Remove(r, "DeleteJobTask", shared.TableCoreJobTasks, condition)
// }

// // LoadAllJobTaskInstances return all instances from the object
// func LoadAllJobTaskInstances(r *http.Request) *mdlShared.Response {
// 	viewJobTaskInstances := []models.ViewJobTaskInstance{}
// 	jobInstanceID := chi.URLParam(r, "job_instance_id")
// 	jobInstanceIDColumn := fmt.Sprintf("%s.job_instance_id", shared.ViewCoreJobTaskInstances)
// 	languageCode := r.Header.Get("Content-Language")
// 	languageCodeColumn := fmt.Sprintf("%s.language_code", shared.ViewCoreJobTaskInstances)
// 	condition := builder.And(
// 		builder.Equal(jobInstanceIDColumn, jobInstanceID),
// 		builder.Equal(languageCodeColumn, languageCode),
// 	)

// 	return db.Load(r, &viewJobTaskInstances, "LoadAllJobTaskInstances", shared.ViewCoreJobTaskInstances, condition)
// }

// // LoadAllJobFollowersAvaible return all instances from the object
// func LoadAllJobFollowersAvaible(r *http.Request) *mdlShared.Response {
// 	viewFollowersAvailable := []models.ViewFollowerAvailable{}
// 	activeColumn := fmt.Sprintf("%s.active", shared.ViewCoreUsersAndGroups)
// 	languageCode := r.Header.Get("Content-Language")
// 	languageCodeColumn := fmt.Sprintf("%s.language_code", shared.ViewCoreUsersAndGroups)
// 	condition := builder.And(
// 		builder.Equal(activeColumn, true),
// 		builder.Or(
// 			builder.Equal(languageCodeColumn, languageCode),
// 			builder.Equal(languageCodeColumn, nil),
// 		),
// 	)

// 	return db.Load(r, &viewFollowersAvailable, "LoadAllJobFollowersAvaible", shared.ViewCoreUsersAndGroups, condition)
// }

// // InsertFollowerInJob persists the request body creating a new object in the database
// func InsertFollowerInJob(r *http.Request) *mdlShared.Response {
// 	response := &mdlShared.Response{
// 		Code: http.StatusOK,
// 	}

// 	jobID := chi.URLParam(r, "job_id")
// 	followerID := chi.URLParam(r, "follower_id")
// 	languageCode := r.Header.Get("Content-Language")
// 	followerType := chi.URLParam(r, "follower_type")
// 	userID := r.Header.Get("userID")
// 	now := time.Now()

// 	follower := models.JobFollowers{
// 		ID:           sql.UUID(),
// 		JobID:        jobID,
// 		LanguageCode: languageCode,
// 		FollowerID:   followerID,
// 		FollowerType: followerType,
// 		CreatedBy:    userID,
// 		CreatedAt:    now,
// 	}

// 	jobIDColumn := fmt.Sprintf("%s.id", shared.TableCoreJobs)
// 	sql.InsertStructToJSON("followers", shared.TableCoreJobs, &follower, builder.Equal(jobIDColumn, jobID))
// 	response.Data = follower
// 	return response
// }

// // RemoveFollowerFromJob deletes object from the database
// func RemoveFollowerFromJob(r *http.Request) *mdlShared.Response {
// 	response := &mdlShared.Response{
// 		Code: http.StatusOK,
// 	}

// 	jobID := chi.URLParam(r, "job_id")
// 	followerID := chi.URLParam(r, "follower_id")

// 	err := sql.DeleteStructFromJSON(followerID, jobID, "followers", shared.TableCoreJobs)
// 	if err != nil {
// 		response.Code = http.StatusInternalServerError
// 		response.Errors = append(response.Errors, mdlShared.NewResponseError(shared.ErrorParsingRequest, "RemoveFollowerFromJob", err.Error()))

// 		return response
// 	}

// 	return response
// }

// // LoadAllFollowersByJob return all instances from the object
// func LoadAllFollowersByJob(r *http.Request) *mdlShared.Response {
// 	jobFollowers := []models.JobFollowers{}
// 	jobID := chi.URLParam(r, "job_id")
// 	jobIDColumn := fmt.Sprintf("%s.job_id", shared.ViewCoreJobFollowers)
// 	languageCode := r.Header.Get("Content-Language")
// 	languageCodeColumn := fmt.Sprintf("%s.language_code", shared.ViewCoreJobFollowers)
// 	followerTypeColumn := fmt.Sprintf("%s.follower_type", shared.ViewCoreJobFollowers)
// 	condition := builder.And(
// 		builder.Equal(jobIDColumn, jobID),
// 		builder.Or(
// 			builder.Equal(followerTypeColumn, "user"),
// 			builder.And(
// 				builder.Equal(followerTypeColumn, "group"),
// 				builder.Equal(languageCodeColumn, languageCode),
// 			),
// 		),
// 	)

// 	return db.Load(r, &jobFollowers, "LoadAllFollowersByJob", shared.ViewCoreJobFollowers, condition)
// }

// // CreateInstance persists the request body creating a new object in the database
// func CreateInstance(r *http.Request) *mdlShared.Response {
// 	response := &mdlShared.Response{
// 		Code: http.StatusOK,
// 	}

// 	jobID := chi.URLParam(r, "job_id")
// 	jobTable := shared.TableCoreJobs
// 	jobIDColumn := fmt.Sprintf("%s.id", jobTable)
// 	condition := builder.Equal(jobIDColumn, jobID)
// 	job := models.Job{}

// 	err := sql.SelectStruct(jobTable, &job, condition)
// 	if err != nil {
// 		response.Code = http.StatusInternalServerError
// 		response.Errors = append(response.Errors, mdlShared.NewResponseError(shared.ErrorParsingRequest, "CreateInstance load job", err.Error()))

// 		return response
// 	}

// 	params := map[string]interface{}{}
// 	err = json.NewDecoder(r.Body).Decode(&params)
// 	if err != nil {
// 		response.Code = http.StatusInternalServerError
// 		response.Errors = append(response.Errors, mdlShared.NewResponseError(shared.ErrorParsingRequest, "CreateInstance unmarshal body", err.Error()))

// 		return response
// 	}

// 	userID := r.Header.Get("userID")

// 	_, err = mdlShared.CreateInstance(userID, job.Code, params)
// 	if err != nil {
// 		response.Code = http.StatusInternalServerError
// 		response.Errors = append(response.Errors, mdlShared.NewResponseError(shared.ErrorJobExecution, "CreateInstance", err.Error()))

// 		return response
// 	}

// 	return response
// }
