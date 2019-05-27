package admin

import (
	"net/http"

	services "github.com/agile-work/srv-mdl-core/services/admin"

	"github.com/go-chi/render"
)

// PostUser sends the request to service creating a new user
func PostUser(w http.ResponseWriter, r *http.Request) {
	response := services.CreateUser(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllUsers return all user instances from the service
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllUsers(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetUser return only one user from the service
func GetUser(w http.ResponseWriter, r *http.Request) {
	response := services.LoadUser(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateUser sends the request to service updating an user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateUser(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteUser sends the request to service deleting an user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteUser(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// AddGroupInUser sends the request to service deleting an user
func AddGroupInUser(w http.ResponseWriter, r *http.Request) {
	response := services.InsertUserInGroup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// RemoveGroupFromUser sends the request to service deleting an user
func RemoveGroupFromUser(w http.ResponseWriter, r *http.Request) {
	response := services.RemoveUserFromGroup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllUsersByGroup return all user instances by group from the service
func GetAllUsersByGroup(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllUsersByGroup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
