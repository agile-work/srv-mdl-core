package admin

import (
	"net/http"

	services "github.com/agile-work/srv-mdl-core/services/admin"

	"github.com/go-chi/render"
)

// PostGroup sends the request to service creating a new group
func PostGroup(w http.ResponseWriter, r *http.Request) {
	response := services.CreateGroup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllGroups return all group instances from the service
func GetAllGroups(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllGroups(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetGroup return only one group from the service
func GetGroup(w http.ResponseWriter, r *http.Request) {
	response := services.LoadGroup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateGroup sends the request to service updating a group
func UpdateGroup(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateGroup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteGroup sends the request to service deleting a group
func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteGroup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// AddUserInGroup sends the request to service deleting an user
func AddUserInGroup(w http.ResponseWriter, r *http.Request) {
	response := services.InsertUserInGroup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteGroupUser sends the request to service deleting a user from a group
func DeleteGroupUser(w http.ResponseWriter, r *http.Request) {
	response := services.RemoveUserFromGroup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// PostGroupPermission sends the request to service creating a permission in a group
func PostGroupPermission(w http.ResponseWriter, r *http.Request) {
	response := services.InsertPermission(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteGroupPermission sends the request to service deleting a permission from a group
func DeleteGroupPermission(w http.ResponseWriter, r *http.Request) {
	response := services.RemovePermission(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllGroupsByUser sends the request to service deleting a permission from a group
func GetAllGroupsByUser(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllGroupsByUser(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
