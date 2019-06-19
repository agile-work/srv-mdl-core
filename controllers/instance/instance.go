package instance

import (
	"net/http"

	mdlShared "github.com/agile-work/srv-mdl-shared"
)

// PostSchemaInstance sends the request to service creating a new schema
func PostSchemaInstance(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// GetAllSchemaInstances return all schema instances from the service
func GetAllSchemaInstances(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// GetSchemaInstance return only one schema from the service
func GetSchemaInstance(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// UpdateSchemaInstance sends the request to service updating a schema
func UpdateSchemaInstance(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// DeleteSchemaInstance sends the request to service deleting a schema
func DeleteSchemaInstance(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// GetLookupInstance return all schema instances from the service
func GetLookupInstance(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}
