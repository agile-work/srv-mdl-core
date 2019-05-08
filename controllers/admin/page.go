package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-core/services"

	"github.com/go-chi/render"
)

// PostPage sends the request to service creating a new page
func PostPage(w http.ResponseWriter, r *http.Request) {
	response := services.CreatePage(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllPages return all page instances from the service
func GetAllPages(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllPages(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetPage return only one page from the service
func GetPage(w http.ResponseWriter, r *http.Request) {
	response := services.LoadPage(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdatePage sends the request to service updating a page
func UpdatePage(w http.ResponseWriter, r *http.Request) {
	response := services.UpdatePage(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeletePage sends the request to service deleting a page
func DeletePage(w http.ResponseWriter, r *http.Request) {
	response := services.DeletePage(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
