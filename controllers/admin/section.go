package admin

import (
	"net/http"

	services "github.com/agile-work/srv-mdl-core/services/admin"

	"github.com/go-chi/render"
)

// PostSection sends the request to service creating a new section
func PostSection(w http.ResponseWriter, r *http.Request) {
	response := services.CreateSection(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllSections return all section instances from the service
func GetAllSections(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllSections(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetSection return only one section from the service
func GetSection(w http.ResponseWriter, r *http.Request) {
	response := services.LoadSection(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateSection sends the request to service updating a section
func UpdateSection(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateSection(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteSection sends the request to service deleting a section
func DeleteSection(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteSection(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
