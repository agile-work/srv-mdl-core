package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/response"
)

// PostJob sends the request to service creating a new schema
func PostJob(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetAllJobs return all schema instances from the service
func GetAllJobs(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetJob return only one schema from the service
func GetJob(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// UpdateJob sends the request to service updating a schema
func UpdateJob(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// DeleteJob sends the request to service deleting a schema
func DeleteJob(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetAllJobInstances return all schema instances from the service
func GetAllJobInstances(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// PostJobTask sends the request to service creating a new schema
func PostJobTask(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetAllJobTasks return all schema instances from the service
func GetAllJobTasks(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetJobTask return only one schema from the service
func GetJobTask(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// UpdateJobTask sends the request to service updating a schema
func UpdateJobTask(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// DeleteJobTask sends the request to service deleting a schema
func DeleteJobTask(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetAllJobTaskInstances return all schema instances from the service
func GetAllJobTaskInstances(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// LoadAllJobFollowersAvaible sends the request to service deleting a schema
func LoadAllJobFollowersAvaible(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// InsertFollowerInJob sends the request to service deleting a schema
func InsertFollowerInJob(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// LoadAllFollowersByJob sends the request to service deleting a schema
func LoadAllFollowersByJob(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// RemoveFollowerFromJob sends the request to service deleting a schema
func RemoveFollowerFromJob(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// PostJobInstance sends the request to service deleting a schema
func PostJobInstance(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}
