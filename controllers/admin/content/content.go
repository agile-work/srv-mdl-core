package content

import (
	"net/http"

	"github.com/agile-work/srv-mdl-core/models/content"
	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-mdl-shared/util"
	"github.com/agile-work/srv-shared/sql-builder/db"
	"github.com/go-chi/chi"
)

// PostContent sends the request to model creating a new Content
func PostContent(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	cnt := &content.Content{}
	resp := response.New()

	if err := resp.Parse(req, cnt); err != nil {
		resp.NewError("PostContent response load", err)
		resp.Render(res, req)
		return
	}

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("PostContent Content new transaction", err)
		resp.Render(res, req)
		return
	}

	translation.FieldsRequestLanguageCode = "all"
	if err := cnt.Create(trs); err != nil {
		trs.Rollback()
		resp.NewError("PostContent", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()

	resp.Data = cnt
	resp.Render(res, req)
}

// GetAllContents return all Content instances from the model
func GetAllContents(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	metadata := response.Metadata{}
	if err := metadata.Load(req); err != nil {
		resp.NewError("GetLookupInstance metadata parse", err)
		resp.Render(res, req)
		return
	}
	opt := metadata.GenerateDBOptions()
	cnt := &content.Contents{}
	if err := cnt.LoadAll(opt); err != nil {
		resp.NewError("GetAllContents", err)
		resp.Render(res, req)
		return
	}

	resp.Data = cnt
	resp.Metadata = metadata
	resp.Render(res, req)
}

// GetContent return only one Content from the model
func GetContent(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	resp := response.New()

	cnt := &content.Content{Code: chi.URLParam(req, "content_code")}
	if err := cnt.Load(); err != nil {

		resp.NewError("GetContent", err)
		resp.Render(res, req)
		return
	}
	resp.Data = cnt
	resp.Render(res, req)
}

// UpdateContent sends the request to model updating a Content
func UpdateContent(res http.ResponseWriter, req *http.Request) {
	translation.FieldsRequestLanguageCode = req.Header.Get("Content-Language")
	cnt := &content.Content{}
	resp := response.New()

	if err := resp.Parse(req, cnt); err != nil {
		resp.NewError("UpdateContent Content new transaction", err)
		resp.Render(res, req)
		return
	}

	cnt.Code = chi.URLParam(req, "content_code")

	body, err := util.GetBodyMap(req)
	if err != nil {
		resp.NewError("UpdateContent", err)
		resp.Render(res, req)
		return
	}

	columns, translations := util.GetColumnsFromBody(body, cnt)

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("UpdateContent Content new transaction", err)
		resp.Render(res, req)
		return
	}

	if err := cnt.Update(trs, columns, translations); err != nil {
		trs.Rollback()
		resp.NewError("UpdateContent", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Data = cnt
	resp.Render(res, req)
}

// DeleteContent sends the request to model deleting a Content
func DeleteContent(res http.ResponseWriter, req *http.Request) {
	resp := response.New()

	trs, err := db.NewTransaction()
	if err != nil {
		resp.NewError("DeleteContent Content new transaction", err)
		resp.Render(res, req)
		return
	}

	cnt := &content.Content{Code: chi.URLParam(req, "content_code")}
	if err := cnt.Delete(trs); err != nil {
		trs.Rollback()
		resp.NewError("DeleteContent", err)
		resp.Render(res, req)
		return
	}
	trs.Commit()
	resp.Render(res, req)
}
