package dataset

import (
	"net/http"

	"github.com/agile-work/srv-mdl-core/models/dataset"
	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-mdl-shared/util"
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

	if err := def.UpdateQuery(trs, chi.URLParam(req, "dataset_code"), req.Header.Get("username")); err != nil {
		trs.Rollback()
		resp.NewError("UpdateDatasetQuery", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = def
	resp.Render(res, req)
}

// UpdateDatasetDynamicParam change dynamic dataset field
func UpdateDatasetDynamicParam(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateDatasetDynamicParam new transaction", err)
		resp.Render(res, req)
		return
	}

	columns, err := util.GetBodyUpdatableJSONColumns(req, &dataset.Param{}, req.Header.Get("username"), req.Header.Get("Content-Language"))
	if err != nil {
		resp.NewError("UpdateDatasetDynamicParam request parse", err)
		resp.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"
	dynamicDefinition := &dataset.DynamicDefinition{}

	if err := dynamicDefinition.UpdateParam(trs, chi.URLParam(req, "param_type"), chi.URLParam(req, "param_code"), chi.URLParam(req, "dataset_code"), columns); err != nil {
		trs.Rollback()
		resp.NewError("UpdateDatasetDynamicParam", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Render(res, req)
}
