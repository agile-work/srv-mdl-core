package admin

import (
	"net/http"

	mdlShared "github.com/agile-work/srv-mdl-shared"
)

// PostContainerStructure sends the request to model creating a new containerStructure
func PostContainerStructure(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// GetAllContainerStructures return all containerStructure instances from the model
func GetAllContainerStructures(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// GetContainerStructure return only one containerStructure from the model
func GetContainerStructure(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// UpdateContainerStructure sends the request to model updating a containerStructure
func UpdateContainerStructure(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// DeleteContainerStructure sends the request to model deleting a containerStructure
func DeleteContainerStructure(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}
