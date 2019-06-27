package instance

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/util"

	"github.com/agile-work/srv-mdl-core/models/dataset"
	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-mdl-shared/models/user"
	"github.com/go-chi/chi"
)

// GetDatasetInstance return all schema instances from the service
func GetDatasetInstance(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	params, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("GetDatasetInstance body parse", err)
		resp.Render(res, req)
		return
	}

	metadata := response.Metadata{}
	if err := metadata.Load(req); err != nil {
		resp.NewError("GetDatasetInstance metadata parse", err)
		resp.Render(res, req)
		return
	}
	opt := metadata.GenerateDBOptions()

	usr := &user.User{Username: req.Header.Get("username")}
	if err := usr.Load(); err != nil {
		resp.NewError("GetDatasetInstance user load", err)
		resp.Render(res, req)
		return
	}

	ds := &dataset.Dataset{Code: chi.URLParam(req, "dataset_code")}
	if err := ds.Load(); err != nil {
		resp.NewError("GetDatasetInstance dataset load", err)
		resp.Render(res, req)
		return
	}

	result, err := ds.GetInstances(params, usr, opt)
	if err != nil {
		resp.NewError("GetDatasetInstance", err)
		resp.Render(res, req)
		return
	}

	resp.Data = result
	resp.Metadata = metadata
	resp.Render(res, req)
}
