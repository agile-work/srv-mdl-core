package admin

import (
	"net/http"

	services "github.com/agile-work/srv-mdl-core/services/admin"

	"github.com/go-chi/render"
)

// PostContainerStructure sends the request to service creating a new sectionStructure
func PostContainerStructure(w http.ResponseWriter, r *http.Request) {
	response := services.CreateContainerStructure(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllContainerStructures return all sectionStructure instances from the service
func GetAllContainerStructures(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllContainerStructures(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetContainerStructure return only one sectionStructure from the service
func GetContainerStructure(w http.ResponseWriter, r *http.Request) {
	response := services.LoadContainerStructure(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateContainerStructure sends the request to service updating a sectionStructure
func UpdateContainerStructure(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateContainerStructure(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteContainerStructure sends the request to service deleting a sectionStructure
func DeleteContainerStructure(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteContainerStructure(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
