package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/response"
)

// PostContainerStructure sends the request to model creating a new containerStructure
func PostContainerStructure(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetAllContainerStructures return all containerStructure instances from the model
func GetAllContainerStructures(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetContainerStructure return only one containerStructure from the model
func GetContainerStructure(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// UpdateContainerStructure sends the request to model updating a containerStructure
func UpdateContainerStructure(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// DeleteContainerStructure sends the request to model deleting a containerStructure
func DeleteContainerStructure(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}
