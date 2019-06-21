package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/response"
)

// PostView sends the request to model creating a new view
func PostView(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetAllViews return all view instances from the model
func GetAllViews(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetView return only one view from the model
func GetView(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// UpdateView sends the request to model updating a view
func UpdateView(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// DeleteView sends the request to model deleting a view
func DeleteView(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}
