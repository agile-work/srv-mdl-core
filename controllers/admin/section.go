package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/response"
)

// PostSection sends the request to model creating a new section
func PostSection(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetAllSections return all section instances from the model
func GetAllSections(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetSection return only one section from the model
func GetSection(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// UpdateSection sends the request to model updating a section
func UpdateSection(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// DeleteSection sends the request to model deleting a section
func DeleteSection(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}
