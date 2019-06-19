package instance

import (
	"net/http"

	mdlShared "github.com/agile-work/srv-mdl-shared"
)

// GetAllResources return all schema instances from the service
func GetAllResources(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// GetResource return all schema instances from the service
func GetResource(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// UpdateResource return all schema instances from the service
func UpdateResource(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}
