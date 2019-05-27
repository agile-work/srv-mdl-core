package admin

import (
	"net/http"

	services "github.com/agile-work/srv-mdl-core/services/admin"

	"github.com/go-chi/render"
)

// PostTab sends the request to service creating a new field
func PostTab(w http.ResponseWriter, r *http.Request) {
	response := services.CreateTab(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllTabs return all field instances from the service
func GetAllTabs(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllTabs(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetTab return only one field from the service
func GetTab(w http.ResponseWriter, r *http.Request) {
	response := services.LoadTab(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateTab sends the request to service updating a field
func UpdateTab(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateTab(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteTab sends the request to service deleting a field
func DeleteTab(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteTab(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
