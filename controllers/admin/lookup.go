package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-core/services"

	"github.com/go-chi/render"
)

// PostLookup sends the request to service creating a new lookup
func PostLookup(w http.ResponseWriter, r *http.Request) {
	response := services.CreateLookup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllLookups return all lookup instances from the service
func GetAllLookups(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllLookups(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetLookup return only one lookup from the service
func GetLookup(w http.ResponseWriter, r *http.Request) {
	response := services.LoadLookup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateLookup sends the request to service updating a lookup
func UpdateLookup(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateLookup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteLookup sends the request to service deleting a lookup
func DeleteLookup(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteLookup(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// PostLookupOption sends the request to service creating a new lookupOption
func PostLookupOption(w http.ResponseWriter, r *http.Request) {
	response := services.CreateLookupOption(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllLookupOptions return all lookupOption instances from the service
func GetAllLookupOptions(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllLookupOptions(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetLookupOption return only one lookupOption from the service
func GetLookupOption(w http.ResponseWriter, r *http.Request) {
	response := services.LoadLookupOption(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateLookupOption sends the request to service updating a lookupOption
func UpdateLookupOption(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateLookupOption(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteLookupOption sends the request to service deleting a lookupOption
func DeleteLookupOption(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteLookupOption(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
