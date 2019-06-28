package instance

import (
	"net/http"

	"github.com/agile-work/srv-mdl-core/models/dataset"
	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/go-chi/chi"
)

// PostSchemaInstance sends the request to service creating a new schema
func PostSchemaInstance(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetAllSchemaInstances return all schema instances from the service
func GetAllSchemaInstances(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	metadata := response.Metadata{}
	if err := metadata.Load(req); err != nil {
		resp.NewError("GetDatasetInstance metadata parse", err)
		resp.Render(res, req)
		return
	}
	opt := metadata.GenerateDBOptions()

	ds := &dataset.Dataset{Code: chi.URLParam(req, "schema_code")}
	if err := ds.Load(); err != nil {
		resp.NewError("GetDatasetInstance dataset load", err)
		resp.Render(res, req)
		return
	}

	result, err := ds.GetUserInstances(req.Header.Get("username"), opt, nil)
	if err != nil {
		resp.NewError("GetDatasetInstance", err)
		resp.Render(res, req)
		return
	}

	resp.Data = result
	resp.Metadata = metadata
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
