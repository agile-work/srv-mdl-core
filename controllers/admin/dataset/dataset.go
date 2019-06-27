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

// PostDataset sends the request to service creating a new dataset
func PostDataset(res http.ResponseWriter, req *http.Request) {
	ds := &dataset.Dataset{}
	resp := response.New()

	if err := resp.Parse(req, ds); err != nil {
		resp.NewError("PostDataset response load", err)
		resp.Render(res, req)
		return
	}

	if err := ds.ProcessDefinitions(req.Header.Get("Content-Language"), req.Method); err != nil {
		resp.NewError("PostDataset processing definitions", err)
		resp.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("PostDataset new transaction", err)
		resp.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"
	if err := ds.Create(trs); err != nil {
		trs.Rollback()
		resp.NewError("PostDataset", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = ds
	resp.Render(res, req)
}

// GetAllDatasets return all dataset instances from the service
func GetAllDatasets(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	metadata := response.Metadata{}
	if err := metadata.Load(req); err != nil {
		resp.NewError("GetAllDatasets metadata parse", err)
		resp.Render(res, req)
		return
	}
	opt := metadata.GenerateDBOptions()
	dses := &dataset.Datasets{}
	if err := dses.LoadAll(opt); err != nil {
		resp.NewError("GetAllDatasets", err)
		resp.Render(res, req)
		return
	}
	resp.Data = dses
	resp.Metadata = metadata
	resp.Render(res, req)
}

// GetDataset return only one dataset from the service
func GetDataset(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	ds := &dataset.Dataset{Code: chi.URLParam(req, "dataset_code")}
	if err := ds.Load(); err != nil {
		resp.NewError("GetDataset", err)
		resp.Render(res, req)
		return
	}
	resp.Data = ds
	resp.Render(res, req)
}

// UpdateDataset sends the request to service updating a dataset
func UpdateDataset(res http.ResponseWriter, req *http.Request) {
	ds := &dataset.Dataset{}
	resp := response.New()

	if err := resp.Parse(req, ds); err != nil {
		resp.NewError("UpdateDataset dataset new transaction", err)
		resp.Render(res, req)
		return
	}

	ds.Code = chi.URLParam(req, "dataset_code")

	body, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdateDataset dataset get body", err)
		resp.Render(res, req)
		return
	}

	columns, translations := util.GetColumnsFromBody(body, ds)

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateDataset dataset new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := ds.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		resp.NewError("UpdateDataset", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = ds
	resp.Render(res, req)
}

// DeleteDataset sends the request to service deleting a dataset
func DeleteDataset(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeleteDataset new transaction", err)
		resp.Render(res, req)
		return
	}

	ds := &dataset.Dataset{Code: chi.URLParam(req, "dataset_code")}
	if err := ds.Delete(trs); err != nil {
		trs.Rollback()
		resp.NewError("DeleteDataset", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Render(res, req)
}
