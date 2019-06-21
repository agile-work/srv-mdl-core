package instance

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/response"
)

// PostSchemaInstance sends the request to service creating a new schema
func PostSchemaInstance(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetAllSchemaInstances return all schema instances from the service
func GetAllSchemaInstances(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetSchemaInstance return only one schema from the service
func GetSchemaInstance(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// UpdateSchemaInstance sends the request to service updating a schema
func UpdateSchemaInstance(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// DeleteSchemaInstance sends the request to service deleting a schema
func DeleteSchemaInstance(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetLookupInstance return all schema instances from the service
func GetLookupInstance(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}
