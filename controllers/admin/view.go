package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-core/services"

	"github.com/go-chi/render"
)

// PostView sends the request to service creating a new view
func PostView(w http.ResponseWriter, r *http.Request) {
	response := services.CreateView(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllViews return all view instances from the service
func GetAllViews(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllViews(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetView return only one view from the service
func GetView(w http.ResponseWriter, r *http.Request) {
	response := services.LoadView(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateView sends the request to service updating a view
func UpdateView(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateView(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteView sends the request to service deleting a view
func DeleteView(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteView(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// PostViewPage sends the request to service creating an association between group and user
func PostViewPage(w http.ResponseWriter, r *http.Request) {
	response := services.InsertPageInView(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllPagesByView return all user instances by group from the service
func GetAllPagesByView(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllPagesByView(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteViewPage sends the request to service deleting a user from a group
func DeleteViewPage(w http.ResponseWriter, r *http.Request) {
	response := services.RemovePageFromView(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
