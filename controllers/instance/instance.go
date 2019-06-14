package instance

import (
	"net/http"

	services "github.com/agile-work/srv-mdl-core/services/instance"

	"github.com/go-chi/render"
)

// PostSchemaInstance sends the request to service creating a new schema
func PostSchemaInstance(w http.ResponseWriter, r *http.Request) {
	response := services.CreateSchemaInstance(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllSchemaInstances return all schema instances from the service
func GetAllSchemaInstances(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllSchemaInstances(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetSchemaInstance return only one schema from the service
func GetSchemaInstance(w http.ResponseWriter, r *http.Request) {
	response := services.LoadSchemaInstance(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateSchemaInstance sends the request to service updating a schema
func UpdateSchemaInstance(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateSchemaInstance(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteSchemaInstance sends the request to service deleting a schema
func DeleteSchemaInstance(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteSchemaInstance(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetLookupInstance return all schema instances from the service
func GetLookupInstance(w http.ResponseWriter, r *http.Request) {
	response := services.LoadLookupInstance(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
