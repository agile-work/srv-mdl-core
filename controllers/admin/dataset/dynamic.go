package dataset

import (
	"net/http"

	"github.com/agile-work/srv-mdl-core/models/dataset"
	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-shared/sql-builder/db"
	"github.com/go-chi/chi"
)

// UpdateDatasetQuery change dynamic dataset query
func UpdateDatasetQuery(res http.ResponseWriter, req *http.Request) {
	def := &dataset.DynamicDefinition{}
	resp := response.New()

	if err := resp.Parse(req, def); err != nil {
		resp.NewError("UpdateDatasetQuery response parse", err)
		resp.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateDatasetQuery new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := def.UpdateQuery(trs, chi.URLParam(req, "dataset_code")); err != nil {
		trs.Rollback()
		resp.NewError("UpdateDatasetQuery", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = def
	resp.Render(res, req)
}
