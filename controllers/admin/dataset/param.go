package dataset

import (
	"net/http"

	"github.com/agile-work/srv-mdl-core/models/dataset"
	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-mdl-shared/util"
	"github.com/agile-work/srv-shared/sql-builder/db"
	"github.com/go-chi/chi"
)

// UpdateDatasetDynamicField change dynamic dataset field
func UpdateDatasetDynamicField(res http.ResponseWriter, req *http.Request) {
	param := &dataset.Param{}
	resp := response.New()

	if err := resp.Parse(req, param); err != nil {
		resp.NewError("UpdateDatasetDynamicField response parse", err)
		resp.Render(res, req)
		return
	}

	param.Code = chi.URLParam(req, "param_code")

	body, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdateDatasetDynamicField get body", err)
		resp.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateDatasetDynamicField new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := param.Update(trs, chi.URLParam(req, "dataset_code"), body, "fields"); err != nil {
		trs.Rollback()
		resp.NewError("UpdateDatasetDynamicField", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = param
	resp.Render(res, req)
}

// UpdateDatasetDynamicParam change dynamic dataset param
func UpdateDatasetDynamicParam(res http.ResponseWriter, req *http.Request) {
	param := &dataset.Param{}
	resp := response.New()

	if err := resp.Parse(req, param); err != nil {
		resp.NewError("UpdateDatasetDynamicParam response parse", err)
		resp.Render(res, req)
		return
	}

	param.Code = chi.URLParam(req, "param_code")

	body, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdateDatasetDynamicParam get body", err)
		resp.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateDatasetDynamicParam new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := param.Update(trs, chi.URLParam(req, "dataset_code"), body, "params"); err != nil {
		trs.Rollback()
		resp.NewError("UpdateDatasetDynamicParam", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = param
	resp.Render(res, req)
}
