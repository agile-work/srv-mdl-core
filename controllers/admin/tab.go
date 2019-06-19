package admin

import (
	"net/http"

	mdlShared "github.com/agile-work/srv-mdl-shared"
)

// PostTab sends the request to model creating a new tab
func PostTab(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// GetAllTabs return all tab instances from the model
func GetAllTabs(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// GetTab return only one tab from the model
func GetTab(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// UpdateTab sends the request to model updating a tab
func UpdateTab(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// DeleteTab sends the request to model deleting a tab
func DeleteTab(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}
