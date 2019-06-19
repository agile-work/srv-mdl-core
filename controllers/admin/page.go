package admin

import (
	"net/http"

	mdlShared "github.com/agile-work/srv-mdl-shared"
)

// PostPage sends the request to model creating a new page
func PostPage(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// GetAllPages return all page instances from the model
func GetAllPages(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// GetPage return only one page from the model
func GetPage(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// UpdatePage sends the request to model updating a page
func UpdatePage(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// DeletePage sends the request to model deleting a page
func DeletePage(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}
