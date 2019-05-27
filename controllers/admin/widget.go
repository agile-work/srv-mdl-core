package admin

import (
	"net/http"

	services "github.com/agile-work/srv-mdl-core/services/admin"

	"github.com/go-chi/render"
)

// PostWidget sends the request to service creating a new schema
func PostWidget(w http.ResponseWriter, r *http.Request) {
	response := services.CreateWidget(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllWidgets return all schema instances from the service
func GetAllWidgets(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllWidgets(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetWidget return only one schema from the service
func GetWidget(w http.ResponseWriter, r *http.Request) {
	response := services.LoadWidget(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateWidget sends the request to service updating a schema
func UpdateWidget(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateWidget(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteWidget sends the request to service deleting a schema
func DeleteWidget(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteWidget(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
