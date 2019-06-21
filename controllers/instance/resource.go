package instance

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/response"
)

// GetAllResources return all schema instances from the service
func GetAllResources(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetResource return all schema instances from the service
func GetResource(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// UpdateResource return all schema instances from the service
func UpdateResource(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}
