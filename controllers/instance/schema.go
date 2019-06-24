package instance

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-mdl-shared/models/user"
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
		resp.NewError("GetAllSchemaInstances metadata parse", err)
		resp.Render(res, req)
		return
	}

	opt := metadata.GenerateDBOptions()
	usr := &user.User{Username: req.Header.Get("username")}
	if err := usr.Load(); err != nil {
		resp.NewError("GetAllSchemaInstances user load", err)
		resp.Render(res, req)
		return
	}

	results, err := usr.GetSecurityInstances(chi.URLParam(req, "schema_code"), opt)
	if err != nil {
		resp.NewError("GetAllSchemaInstances", err)
		resp.Render(res, req)
		return
	}

	resp.Data = results
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
