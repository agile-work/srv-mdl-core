package admin

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-shared/util"

	"github.com/agile-work/srv-mdl-core/models/schema"

	"github.com/go-chi/chi"

	mdlShared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// PostSchema sends the request to model creating a new schema
func PostSchema(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	schema := &schema.Schema{}
	response := mdlShared.NewResponse()

	if err := response.Load(req, schema); err != nil {
		response.NewError(http.StatusInternalServerError, "PostSchema response load", err.Error())
		response.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "PostSchema schema new transaction", err.Error())
		response.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"
	if err := schema.Create(trs); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "PostSchema "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()

	response.Data = schema
	response.Render(res, req)
}

// GetAllSchemas return all schema instances from the model
func GetAllSchemas(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	response := mdlShared.NewResponse()

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "GetAllSchemas schema new transaction", err.Error())
		response.Render(res, req)
		return
	}

	metaData := mdlShared.Metadata{}
	metaData.Load(req)
	opt := metaData.GenerateDBOptions()
	schemas := &schema.Schemas{}
	if err := schemas.LoadAll(trs, opt); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "GetAllSchemas "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Data = schemas
	response.Metadata = metaData
	response.Render(res, req)
}

// GetSchema return only one schema from the model
func GetSchema(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	response := mdlShared.NewResponse()

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "GetSchema schema new transaction", err.Error())
		response.Render(res, req)
		return
	}

	schema := &schema.Schema{Code: chi.URLParam(req, "schema_code")}
	if err := schema.Load(trs); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "GetSchema "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Data = schema
	response.Render(res, req)
}

// UpdateSchema sends the request to model updating a schema
func UpdateSchema(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	schema := &schema.Schema{}
	response := mdlShared.NewResponse()

	if err := response.Load(req, schema); err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateSchema schema new transaction", err.Error())
		response.Render(res, req)
		return
	}

	schema.Code = chi.URLParam(req, "schema_code")

	body, err := util.GetBodyMap(req)
	if err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateSchema "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}

	columns, translations, err := util.GetColumnsFromBody(body, schema)
	if err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateSchema "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "UpdateSchema schema new transaction", err.Error())
		response.Render(res, req)
		return
	}

	if err := schema.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "UpdateSchema "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Data = schema
	response.Render(res, req)
}

// DeleteSchema sends the request to model deleting a schema
func DeleteSchema(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "DeleteSchema schema new transaction", err.Error())
		response.Render(res, req)
		return
	}

	schema := &schema.Schema{Code: chi.URLParam(req, "schema_code")}
	if err := schema.Delete(trs); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "DeleteSchema "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Render(res, req)
}

// CallDeleteSchema sends the request to service deleting a schema
func CallDeleteSchema(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	trs, err := db.NewTransaction()
	if err != nil {
		response.NewError(http.StatusInternalServerError, "CallDeleteSchema schema new transaction", err.Error())
		response.Render(res, req)
		return
	}

	schema := &schema.Schema{Code: chi.URLParam(req, "schema_code")}
	if err := schema.CallDelete(trs); err != nil {
		trs.Rollback()
		response.NewError(http.StatusInternalServerError, "CallDeleteSchema "+mdlShared.GetErrorStruct(err).Scope, mdlShared.GetErrorStruct(err).ErrorMessage)
		response.Render(res, req)
		return
	}
	trs.Commit()
	response.Render(res, req)
}

// PostSchemaModule sends the request to service creating an association between group and user
func PostSchemaModule(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// GetAllModulesBySchema return all user instances by group from the service
func GetAllModulesBySchema(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}

// DeleteSchemaModule sends the request to service deleting a user from a group
func DeleteSchemaModule(res http.ResponseWriter, req *http.Request) {
	response := mdlShared.NewResponse()

	response.Render(res, req)
}
