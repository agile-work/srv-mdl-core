package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-core/services"

	"github.com/go-chi/render"
)

// PostJob sends the request to service creating a new schema
func PostJob(w http.ResponseWriter, r *http.Request) {
	response := services.CreateJob(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllJobs return all schema instances from the service
func GetAllJobs(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllJobs(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetJob return only one schema from the service
func GetJob(w http.ResponseWriter, r *http.Request) {
	response := services.LoadJob(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateJob sends the request to service updating a schema
func UpdateJob(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateJob(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteJob sends the request to service deleting a schema
func DeleteJob(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteJob(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// PostJobTask sends the request to service creating a new schema
func PostJobTask(w http.ResponseWriter, r *http.Request) {
	response := services.CreateJobTask(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllJobTasks return all schema instances from the service
func GetAllJobTasks(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllJobTasks(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetJobTask return only one schema from the service
func GetJobTask(w http.ResponseWriter, r *http.Request) {
	response := services.LoadJobTask(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateJobTask sends the request to service updating a schema
func UpdateJobTask(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateJobTask(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteJobTask sends the request to service deleting a schema
func DeleteJobTask(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteJobTask(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// LoadAllJobFollowersAvaible sends the request to service deleting a schema
func LoadAllJobFollowersAvaible(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllJobFollowersAvaible(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// InsertFollowerInJob sends the request to service deleting a schema
func InsertFollowerInJob(w http.ResponseWriter, r *http.Request) {
	response := services.InsertFollowerInJob(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// LoadAllFollowersByJob sends the request to service deleting a schema
func LoadAllFollowersByJob(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllFollowersByJob(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// RemoveFollowerFromJob sends the request to service deleting a schema
func RemoveFollowerFromJob(w http.ResponseWriter, r *http.Request) {
	response := services.RemoveFollowerFromJob(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// PostJobInstance sends the request to service deleting a schema
func PostJobInstance(w http.ResponseWriter, r *http.Request) {
	response := services.CreateJobInstance(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
