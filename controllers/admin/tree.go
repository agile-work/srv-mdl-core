package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-core/services"

	"github.com/go-chi/render"
)

// PostTree sends the request to service creating a new schema
func PostTree(w http.ResponseWriter, r *http.Request) {
	response := services.CreateTree(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllTrees return all schema instances from the service
func GetAllTrees(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllTrees(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetTree return only one schema from the service
func GetTree(w http.ResponseWriter, r *http.Request) {
	response := services.LoadTree(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateTree sends the request to service updating a schema
func UpdateTree(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateTree(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteTree sends the request to service deleting a schema
func DeleteTree(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteTree(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// PostTreeLevel sends the request to service creating a new schema
func PostTreeLevel(w http.ResponseWriter, r *http.Request) {
	response := services.CreateTreeLevel(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllTreeLevels return all schema instances from the service
func GetAllTreeLevels(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllTreeLevels(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetTreeLevel return only one schema from the service
func GetTreeLevel(w http.ResponseWriter, r *http.Request) {
	response := services.LoadTreeLevel(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateTreeLevel sends the request to service updating a schema
func UpdateTreeLevel(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateTreeLevel(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteTreeLevel sends the request to service deleting a schema
func DeleteTreeLevel(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteTreeLevel(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// PostTreeUnit sends the request to service creating a new schema
func PostTreeUnit(w http.ResponseWriter, r *http.Request) {
	response := services.CreateTreeUnit(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllTreeUnits return all schema instances from the service
func GetAllTreeUnits(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllTreeUnits(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetTreeUnit return only one schema from the service
func GetTreeUnit(w http.ResponseWriter, r *http.Request) {
	response := services.LoadTreeUnit(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateTreeUnit sends the request to service updating a schema
func UpdateTreeUnit(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateTreeUnit(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteTreeUnit sends the request to service deleting a schema
func DeleteTreeUnit(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteTreeUnit(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
