package instance

import (
	"net/http"

	services "github.com/agile-work/srv-mdl-core/services/instance"

	"github.com/go-chi/render"
)

// PostPermission sends the request to service creating a new schema
func PostPermission(w http.ResponseWriter, r *http.Request) {
	response := services.CreatePermission(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllPermissions return all schema instances from the service
func GetAllPermissions(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllPermissions(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdatePermission return all schema instances from the service
func UpdatePermission(w http.ResponseWriter, r *http.Request) {
	response := services.UpdatePermission(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeletePermission sends the request to service deleting a schema
func DeletePermission(w http.ResponseWriter, r *http.Request) {
	response := services.DeletePermission(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
