package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-mdl-shared/util"

	"github.com/agile-work/srv-mdl-core/models/schema"

	"github.com/go-chi/chi"

	"github.com/agile-work/srv-shared/sql-builder/db"
)

// PostSchema sends the request to model creating a new schema
func PostSchema(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	schema := &schema.Schema{}
	resp := response.New()

	if err := resp.Parse(req, schema); err != nil {
		resp.NewError("PostSchema response load", err)
		resp.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("PostSchema schema new transaction", err)
		resp.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"
	if err := schema.Create(trs); err != nil {
		trs.Rollback()
		resp.NewError("PostSchema", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = schema
	resp.Render(res, req)
}

// GetAllSchemas return all schema instances from the model
func GetAllSchemas(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("GetAllSchemas schema new transaction", err)
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
	schemas := &schema.Schemas{}
	if err := schemas.LoadAll(trs, opt); err != nil {
		trs.Rollback()
		resp.NewError("GetAllSchemas", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = schemas
	resp.Metadata = metadata
	resp.Render(res, req)
}

// GetSchema return only one schema from the model
func GetSchema(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	schema := &schema.Schema{Code: chi.URLParam(req, "schema_code")}
	if err := schema.Load(); err != nil {
		resp.NewError("GetSchema", err)
		resp.Render(res, req)
		return
	}

	resp.Data = schema
	resp.Render(res, req)
}

// UpdateSchema sends the request to model updating a schema
func UpdateSchema(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	schema := &schema.Schema{}
	resp := response.New()

	if err := resp.Parse(req, schema); err != nil {
		resp.NewError("UpdateSchema schema new transaction", err)
		resp.Render(res, req)
		return
	}

	schema.Code = chi.URLParam(req, "schema_code")

	body, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdateSchema", err)
		resp.Render(res, req)
		return
	}

	columns, translations := util.GetColumnsFromBody(body, schema)

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateSchema schema new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := schema.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		resp.NewError("UpdateSchema", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = schema
	resp.Render(res, req)
}

// DeleteSchema sends the request to model deleting a schema
func DeleteSchema(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeleteSchema schema new transaction", err)
		resp.Render(res, req)
		return
	}

	schema := &schema.Schema{Code: chi.URLParam(req, "schema_code")}
	if err := schema.Delete(trs); err != nil {
		trs.Rollback()
		resp.NewError("DeleteSchema", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Render(res, req)
}

// CallDeleteSchema sends the request to service deleting a schema
func CallDeleteSchema(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("CallDeleteSchema schema new transaction", err)
		resp.Render(res, req)
		return
	}

	schema := &schema.Schema{Code: chi.URLParam(req, "schema_code")}
	if err := schema.CallDelete(trs); err != nil {
		trs.Rollback()
		resp.NewError("CallDeleteSchema", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Render(res, req)
}

// PostSchemaModule sends the request to service creating an association between group and user
func PostSchemaModule(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// GetAllModulesBySchema return all user instances by group from the service
func GetAllModulesBySchema(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}

// DeleteSchemaModule sends the request to service deleting a user from a group
func DeleteSchemaModule(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	resp.Render(res, req)
}
