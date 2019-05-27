package instance

import (
	"net/http"

	services "github.com/agile-work/srv-mdl-core/services/instance"

	"github.com/go-chi/render"
)

// GetAllResources return all schema instances from the service
func GetAllResources(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllResources(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetResource return all schema instances from the service
func GetResource(w http.ResponseWriter, r *http.Request) {
	response := services.LoadResource(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateResource return all schema instances from the service
func UpdateResource(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateResource(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
