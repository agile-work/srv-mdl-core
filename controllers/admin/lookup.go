package admin

import (
	"net/http"

	services "github.com/agile-work/srv-mdl-core/services/admin"

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

// AddLookupOption add a new option to a lookup
func AddLookupOption(w http.ResponseWriter, r *http.Request) {
	response := services.AddLookupOption(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateLookupOption change lookup option data
func UpdateLookupOption(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateLookupOption(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteLookupOption delete an option
func DeleteLookupOption(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteLookupOption(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateLookupOrder delete an option
func UpdateLookupOrder(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateLookupOrder(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
