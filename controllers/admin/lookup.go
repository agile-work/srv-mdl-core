package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-core/models/lookup"
	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-mdl-shared/util"
	"github.com/agile-work/srv-shared/sql-builder/db"
	"github.com/go-chi/chi"
)

// PostLookup sends the request to service creating a new lookup
func PostLookup(res http.ResponseWriter, req *http.Request) {
	lkp := &lookup.Lookup{}
	resp := response.New()

	if err := resp.Parse(req, lkp); err != nil {
		resp.NewError("PostLookup response load", err)
		resp.Render(res, req)
		return
	}

	if err := lkp.ProcessDefinitions(req.Header.Get("Content-Language"), req.Method); err != nil {
		resp.NewError("PostLookup processing definitions", err)
		resp.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("PostLookup new transaction", err)
		resp.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"
	if err := lkp.Create(trs); err != nil {
		trs.Rollback()
		resp.NewError("PostLookup", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = lkp
	resp.Render(res, req)
}

// GetAllLookups return all lookup instances from the service
func GetAllLookups(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	metadata := response.Metadata{}
	if err := metadata.Load(req); err != nil {
		resp.NewError("GetLookupInstance metadata parse", err)
		resp.Render(res, req)
		return
	}
	opt := metadata.GenerateDBOptions()
	lkps := &lookup.Lookups{}
	if err := lkps.LoadAll(opt); err != nil {
		resp.NewError("GetAllLookups", err)
		resp.Render(res, req)
		return
	}
	resp.Data = lkps
	resp.Metadata = metadata
	resp.Render(res, req)
}

// GetLookup return only one lookup from the service
func GetLookup(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	lkp := &lookup.Lookup{Code: chi.URLParam(req, "lookup_code")}
	if err := lkp.Load(); err != nil {
		resp.NewError("GetLookup", err)
		resp.Render(res, req)
		return
	}
	resp.Data = lkp
	resp.Render(res, req)
}

// UpdateLookup sends the request to service updating a lookup
func UpdateLookup(res http.ResponseWriter, req *http.Request) {
	lkp := &lookup.Lookup{}
	resp := response.New()

	if err := resp.Parse(req, lkp); err != nil {
		resp.NewError("UpdateLookup lookup new transaction", err)
		resp.Render(res, req)
		return
	}

	lkp.Code = chi.URLParam(req, "lookup_code")

	body, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdateLookup lookup get body", err)
		resp.Render(res, req)
		return
	}

	columns, translations := util.GetColumnsFromBody(body, lkp)

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateLookup lookup new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := lkp.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		resp.NewError("UpdateLookup", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = lkp
	resp.Render(res, req)
}

// DeleteLookup sends the request to service deleting a lookup
func DeleteLookup(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeleteLookup new transaction", err)
		resp.Render(res, req)
		return
	}

	lkp := &lookup.Lookup{Code: chi.URLParam(req, "lookup_code")}
	if err := lkp.Delete(trs); err != nil {
		trs.Rollback()
		resp.NewError("DeleteLookup", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Render(res, req)
}

// AddLookupOption add a new option to a lookup
func AddLookupOption(res http.ResponseWriter, req *http.Request) {
	opt := &lookup.Option{}
	resp := response.New()

	if err := resp.Parse(req, opt); err != nil {
		resp.NewError("AddLookupOption response parse", err)
		resp.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("AddLookupOption new transaction", err)
		resp.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"
	if err := opt.Add(trs, chi.URLParam(req, "lookup_code")); err != nil {
		trs.Rollback()
		resp.NewError("AddLookupOption", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = opt
	resp.Render(res, req)
}

// UpdateLookupOption change lookup option data
func UpdateLookupOption(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// DeleteLookupOption delete an option
func DeleteLookupOption(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// UpdateLookupOrder delete an option
func UpdateLookupOrder(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// UpdateLookupQuery change dynamic lookup query
func UpdateLookupQuery(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// UpdateLookupDynamicField change dynamic lookup field
func UpdateLookupDynamicField(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// UpdateLookupDynamicParam change dynamic lookup param
func UpdateLookupDynamicParam(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// // AddLookupOption add a new option to a lookup
// func AddLookupOption(res http.ResponseWriter, req *http.Request) {
// 	response := services.AddLookupOption(r)

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }
// // PostLookup sends the request to service creating a new lookup
// func PostLookup(res http.ResponseWriter, req *http.Request) {
// 	response := services.CreateLookup(r)

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }

// // GetAllLookups return all lookup instances from the service
// func GetAllLookups(res http.ResponseWriter, req *http.Request) {
// 	response := services.LoadAllLookups(r)

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }

// // GetLookup return only one lookup from the service
// func GetLookup(res http.ResponseWriter, req *http.Request) {
// 	response := services.LoadLookup(r)

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }

// // UpdateLookup sends the request to service updating a lookup
// func UpdateLookup(res http.ResponseWriter, req *http.Request) {
// 	response := services.UpdateLookup(r)

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }

// // DeleteLookup sends the request to service deleting a lookup
// func DeleteLookup(res http.ResponseWriter, req *http.Request) {
// 	response := services.DeleteLookup(r)

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }

// // UpdateLookupOption change lookup option data
// func UpdateLookupOption(res http.ResponseWriter, req *http.Request) {
// 	response := services.UpdateLookupOption(r)

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }

// // DeleteLookupOption delete an option
// func DeleteLookupOption(res http.ResponseWriter, req *http.Request) {
// 	response := services.DeleteLookupOption(r)

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }

// // UpdateLookupOrder delete an option
// func UpdateLookupOrder(res http.ResponseWriter, req *http.Request) {
// 	response := services.UpdateLookupOrder(r)

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }

// // UpdateLookupQuery change dynamic lookup query
// func UpdateLookupQuery(res http.ResponseWriter, req *http.Request) {
// 	response := services.UpdateLookupQuery(r)

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }

// // UpdateLookupDynamicField change dynamic lookup field
// func UpdateLookupDynamicField(res http.ResponseWriter, req *http.Request) {
// 	response := services.UpdateLookupDynamicParam(r, "field")

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }

// // UpdateLookupDynamicParam change dynamic lookup param
// func UpdateLookupDynamicParam(res http.ResponseWriter, req *http.Request) {
// 	response := services.UpdateLookupDynamicParam(r, "param")

// 	render.Status(r, response.Code)
// 	render.JSON(w, r, response)
// }
