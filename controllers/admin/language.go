package admin

import (
	"net/http"

	services "github.com/agile-work/srv-mdl-core/services/admin"

	"github.com/go-chi/render"
)

// PostLanguage sends the request to service creating a new schema
func PostLanguage(w http.ResponseWriter, r *http.Request) {
	response := services.CreateLanguage(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllLanguages return all schema instances from the service
func GetAllLanguages(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllLanguages(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetLanguage return only one schema from the service
func GetLanguage(w http.ResponseWriter, r *http.Request) {
	response := services.LoadLanguage(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateLanguage sends the request to service updating a schema
func UpdateLanguage(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateLanguage(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteLanguage sends the request to service deleting a schema
func DeleteLanguage(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteLanguage(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
