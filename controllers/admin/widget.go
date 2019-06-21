package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/response"
)

// PostWidget sends the request to model creating a new widget
func PostWidget(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetAllWidgets return all widget instances from the model
func GetAllWidgets(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetWidget return only one widget from the model
func GetWidget(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// UpdateWidget sends the request to model updating a widget
func UpdateWidget(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// DeleteWidget sends the request to model deleting a widget
func DeleteWidget(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}
