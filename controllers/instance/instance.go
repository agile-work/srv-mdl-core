package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-core/services"

	"github.com/go-chi/render"
)

// PostInstance sends the request to service creating a new schema
func PostInstance(w http.ResponseWriter, r *http.Request) {
	response := services.CreateInstance(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllInstances return all schema instances from the service
func GetAllInstances(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllInstances(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetInstance return only one schema from the service
func GetInstance(w http.ResponseWriter, r *http.Request) {
	response := services.LoadInstance(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateInstance sends the request to service updating a schema
func UpdateInstance(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateInstance(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteInstance sends the request to service deleting a schema
func DeleteInstance(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteInstance(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
