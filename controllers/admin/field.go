package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-shared/sql-builder/builder"

	"github.com/agile-work/srv-mdl-shared/util"

	"github.com/agile-work/srv-mdl-core/models/field"

	"github.com/go-chi/chi"

	"github.com/agile-work/srv-shared/sql-builder/db"
)

// PostField sends the request to model creating a new field
func PostField(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	field := &field.Field{}
	resp := response.New()

	if err := resp.Parse(req, field); err != nil {
		resp.NewError("PostField response load", err)
		resp.Render(res, req)
		return
	}

	field.SchemaCode = chi.URLParam(req, "schema_code")

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("PostField field new transaction", err)
		resp.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"
	if err := field.Create(trs); err != nil {
		trs.Rollback()
		resp.NewError("PostField", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = field
	resp.Render(res, req)
}

// GetAllFields return all field instances from the model
func GetAllFields(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("GetAllFields field new transaction", err)
		resp.Render(res, req)
		return
	}

	metadata := response.Metadata{}
	if err := metadata.Load(req); err != nil {
		resp.NewError("GetLookupInstance metadata parse", err)
		resp.Render(res, req)
		return
	}
	opt := metadata.GenerateDBOptions()
	opt.AddCondition(builder.Equal("schema_code", chi.URLParam(req, "schema_code")))
	fields := &field.Fields{}
	if err := fields.LoadAll(trs, opt); err != nil {
		trs.Rollback()
		resp.NewError("GetAllFields", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = fields
	resp.Metadata = metadata
	resp.Render(res, req)
}

// GetField return only one field from the model
func GetField(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("GetField field new transaction", err)
		resp.Render(res, req)
		return
	}

	field := &field.Field{SchemaCode: chi.URLParam(req, "schema_code"), Code: chi.URLParam(req, "field_code")}
	if err := field.Load(trs); err != nil {
		trs.Rollback()
		resp.NewError("GetField", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = field
	resp.Render(res, req)
}

// UpdateField sends the request to model updating a field
func UpdateField(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	field := &field.Field{}
	resp := response.New()

	if err := resp.Parse(req, field); err != nil {
		resp.NewError("UpdateField field new transaction", err)
		resp.Render(res, req)
		return
	}

	field.SchemaCode = chi.URLParam(req, "schema_code")
	field.Code = chi.URLParam(req, "field_code")

	body, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdateField", err)
		resp.Render(res, req)
		return
	}

	columns, translations := util.GetColumnsFromBody(body, field)

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateField field new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := field.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		resp.NewError("UpdateField", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = field
	resp.Render(res, req)
}

// DeleteField sends the request to model deleting a field
func DeleteField(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeleteField field new transaction", err)
		resp.Render(res, req)
		return
	}

	field := &field.Field{SchemaCode: chi.URLParam(req, "schema_code"), Code: chi.URLParam(req, "field_code")}
	if err := field.Delete(trs); err != nil {
		trs.Rollback()
		resp.NewError("DeleteField", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Render(res, req)
}
