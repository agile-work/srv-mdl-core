package admin

import (
	"net/http"

	services "github.com/agile-work/srv-mdl-core/services/admin"

	"github.com/go-chi/render"
)

// PostSchema sends the request to service creating a new schema
func PostSchema(w http.ResponseWriter, r *http.Request) {
	response := services.CreateSchema(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllSchemas return all schema instances from the service
func GetAllSchemas(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllSchemas(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetSchema return only one schema from the service
func GetSchema(w http.ResponseWriter, r *http.Request) {
	response := services.LoadSchema(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateSchema sends the request to service updating a schema
func UpdateSchema(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateSchema(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteSchema sends the request to service deleting a schema
func DeleteSchema(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteSchema(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// CallDeleteSchema sends the request to service deleting a schema
func CallDeleteSchema(w http.ResponseWriter, r *http.Request) {
	response := services.CallDeleteSchema(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// PostSchemaModule sends the request to service creating an association between group and user
func PostSchemaModule(w http.ResponseWriter, r *http.Request) {
	response := services.InsertModuleInSchema(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllModulesBySchema return all user instances by group from the service
func GetAllModulesBySchema(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllModulesBySchema(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteSchemaModule sends the request to service deleting a user from a group
func DeleteSchemaModule(w http.ResponseWriter, r *http.Request) {
	response := services.RemoveModuleFromSchema(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
