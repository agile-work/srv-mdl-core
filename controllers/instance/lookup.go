package instance

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/util"

	"github.com/agile-work/srv-mdl-core/models/lookup"
	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/go-chi/chi"
)

// GetLookupInstance return all schema instances from the service
func GetLookupInstance(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	params, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("GetLookupInstance body parse", err)
		resp.Render(res, req)
		return
	}

	lkp := &lookup.Lookup{Code: chi.URLParam(req, "lookup_code")}
	if err := lkp.Load(); err != nil {
		resp.NewError("GetLookupInstance lookup load", err)
		resp.Render(res, req)
		return
	}

	result, err := lkp.GetInstances(params)
	if err != nil {
		resp.NewError("GetLookupInstance", err)
		resp.Render(res, req)
		return
	}

	resp.Data = result
	resp.Render(res, req)
}
