package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/response"
)

// PostPage sends the request to model creating a new page
func PostPage(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetAllPages return all page instances from the model
func GetAllPages(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetPage return only one page from the model
func GetPage(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// UpdatePage sends the request to model updating a page
func UpdatePage(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// DeletePage sends the request to model deleting a page
func DeletePage(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}
