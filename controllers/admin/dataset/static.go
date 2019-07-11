package dataset

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/util"

	"github.com/agile-work/srv-mdl-core/models/dataset"
	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-shared/sql-builder/db"
	"github.com/go-chi/chi"
)

// AddDatasetOption add a new option to a dataset
func AddDatasetOption(res http.ResponseWriter, req *http.Request) {
	opt := &dataset.Option{}
	resp := response.New()

	if err := resp.Parse(req, opt); err != nil {
		resp.NewError("AddDatasetOption response parse", err)
		resp.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("AddDatasetOption new transaction", err)
		resp.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"

	staticDefinition := &dataset.StaticDefinition{}
	if err := staticDefinition.AddOption(trs, opt, chi.URLParam(req, "dataset_code")); err != nil {
		trs.Rollback()
		resp.NewError("AddDatasetOption", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = opt
	resp.Render(res, req)
}

// UpdateDatasetOption change dataset option data
func UpdateDatasetOption(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateDatasetOption new transaction", err)
		resp.Render(res, req)
		return
	}

	columns, err := util.GetBodyUpdatableJSONColumns(req, &dataset.Option{}, req.Header.Get("username"), req.Header.Get("Content-Language"))
	if err != nil {
		resp.NewError("AddDatasetOption request parse", err)
		resp.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"
	staticDefinition := &dataset.StaticDefinition{}

	if err := staticDefinition.UpdateOption(trs, chi.URLParam(req, "option_code"), chi.URLParam(req, "dataset_code"), columns); err != nil {
		trs.Rollback()
		resp.NewError("UpdateDatasetOption", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Render(res, req)
}

// DeleteDatasetOption delete an option
func DeleteDatasetOption(res http.ResponseWriter, req *http.Request) {
	resp := response.New()
	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeleteDatasetOption new transaction", err)
		resp.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"
	staticDefinition := &dataset.StaticDefinition{}
	if err := staticDefinition.DeleteOption(trs, chi.URLParam(req, "option_code"), chi.URLParam(req, "dataset_code")); err != nil {
		trs.Rollback()
		resp.NewError("DeleteDatasetOption", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Render(res, req)
}

// UpdateDatasetOrder delete an option
func UpdateDatasetOrder(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}
