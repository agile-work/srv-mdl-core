package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/response"
)

// PostTree sends the request to service creating a new schema
func PostTree(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetAllTrees return all schema instances from the service
func GetAllTrees(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetTree return only one schema from the service
func GetTree(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// UpdateTree sends the request to service updating a schema
func UpdateTree(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// DeleteTree sends the request to service deleting a schema
func DeleteTree(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// PostTreeLevel sends the request to service creating a new schema
func PostTreeLevel(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetAllTreeLevels return all schema instances from the service
func GetAllTreeLevels(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetTreeLevel return only one schema from the service
func GetTreeLevel(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// UpdateTreeLevel sends the request to service updating a schema
func UpdateTreeLevel(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// DeleteTreeLevel sends the request to service deleting a schema
func DeleteTreeLevel(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// PostTreeUnit sends the request to service creating a new schema
func PostTreeUnit(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetAllTreeUnits return all schema instances from the service
func GetAllTreeUnits(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetTreeUnit return only one schema from the service
func GetTreeUnit(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// UpdateTreeUnit sends the request to service updating a schema
func UpdateTreeUnit(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// DeleteTreeUnit sends the request to service deleting a schema
func DeleteTreeUnit(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetAllTreeUnitPermissions return all schema instances from the service
func GetAllTreeUnitPermissions(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// PostTreeUnitPermission sends the request to service deleting a schema
func PostTreeUnitPermission(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// DeleteTreeUnitPermission sends the request to service deleting a schema
func DeleteTreeUnitPermission(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}
