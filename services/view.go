package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/agile-work/srv-mdl-shared/db"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	sql "github.com/agile-work/srv-shared/sql-builder/db"
	"github.com/go-chi/chi"

	"github.com/agile-work/srv-mdl-core/models"
	moduleShared "github.com/agile-work/srv-mdl-shared"
	shared "github.com/agile-work/srv-shared"
)

// CreateView persists the request body creating a new object in the database
func CreateView(r *http.Request) *moduleShared.Response {
	schemaID := chi.URLParam(r, "schema_id")
	view := models.View{
		SchemaID: schemaID,
	}

	return db.Create(r, &view, "CreateView", shared.TableCoreSchViews)
}

// LoadAllViews return all instances from the object
func LoadAllViews(r *http.Request) *moduleShared.Response {
	views := []models.View{}
	schemaID := chi.URLParam(r, "schema_id")
	schemaIDColumn := fmt.Sprintf("%s.schema_id", shared.TableCoreSchViews)
	condition := builder.Equal(schemaIDColumn, schemaID)

	return db.Load(r, &views, "LoadAllViews", shared.TableCoreSchViews, condition)
}

// LoadView return only one object from the database
func LoadView(r *http.Request) *moduleShared.Response {
	view := models.View{}
	viewID := chi.URLParam(r, "view_id")
	viewIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchViews)
	condition := builder.Equal(viewIDColumn, viewID)

	return db.Load(r, &view, "LoadView", shared.TableCoreSchViews, condition)
}

// UpdateView updates object data in the database
func UpdateView(r *http.Request) *moduleShared.Response {
	viewID := chi.URLParam(r, "view_id")
	viewIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchViews)
	condition := builder.Equal(viewIDColumn, viewID)
	view := models.View{
		ID: viewID,
	}

	return db.Update(r, &view, "UpdateView", shared.TableCoreSchViews, condition)
}

// DeleteView deletes object from the database
func DeleteView(r *http.Request) *moduleShared.Response {
	viewID := chi.URLParam(r, "view_id")
	viewIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchViews)
	condition := builder.Equal(viewIDColumn, viewID)

	return db.Remove(r, "DeleteView", shared.TableCoreSchViews, condition)
}

// InsertPageInView persists the request creating a new object in the database
func InsertPageInView(r *http.Request) *moduleShared.Response {
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}

	viewID := chi.URLParam(r, "view_id")
	pageID := chi.URLParam(r, "page_id")

	userID := r.Header.Get("userID")
	now := time.Now()

	statemant := builder.Insert(
		shared.TableCoreViewsPages,
		"view_id",
		"page_id",
		"created_by",
		"created_at",
		"updated_by",
		"updated_at",
	).Values(
		viewID,
		pageID,
		userID,
		now,
		userID,
		now,
	)

	err := sql.Exec(statemant)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorInsertingRecord, "InsertPageInView", err.Error()))

		return response
	}

	return response
}

// LoadAllPagesByView return all instances from the object
func LoadAllPagesByView(r *http.Request) *moduleShared.Response {
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}

	page := []models.Page{}
	viewID := chi.URLParam(r, "view_id")
	tblTranslationName := fmt.Sprintf("%s AS %s_name", shared.TableCoreTranslations, shared.TableCoreTranslations)
	tblTranslationDescription := fmt.Sprintf("%s AS %s_description", shared.TableCoreTranslations, shared.TableCoreTranslations)
	languageCode := r.Header.Get("Content-Language")

	statemant := builder.Select(
		"core_sch_pages.id",
		"core_sch_pages.code",
		"core_translations_name.value AS name",
		"core_translations_description.value AS description",
		"core_sch_pages.schema_id",
		"core_sch_pages.type",
		"core_sch_pages.active",
		"core_sch_pages.created_by",
		"core_sch_pages.created_at",
		"core_sch_pages.updated_by",
		"core_sch_pages.updated_at",
	).From(shared.TableCoreSchPages).Join(
		tblTranslationName, "core_translations_name.structure_id = core_sch_pages.id and core_translations_name.structure_field = 'name'",
	).Join(
		tblTranslationDescription, "core_translations_description.structure_id = core_sch_pages.id and core_translations_description.structure_field = 'description'",
	).Join(
		shared.TableCoreViewsPages, "core_views_pages.page_id = core_sch_pages.id",
	).Where(
		builder.And(
			builder.Equal("core_views_pages.view_id", viewID),
			builder.Equal("core_translations_name.language_code", languageCode),
			builder.Equal("core_translations_description.language_code", languageCode),
		),
	)

	err := sql.QueryStruct(statemant, &page)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingData, "LoadAllPagesByView", err.Error()))

		return response
	}

	response.Data = page

	return response
}

// RemovePageFromView deletes object from the database
func RemovePageFromView(r *http.Request) *moduleShared.Response {
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}

	viewID := chi.URLParam(r, "view_id")
	pageID := chi.URLParam(r, "page_id")

	statemant := builder.Delete(shared.TableCoreViewsPages).Where(
		builder.And(
			builder.Equal("view_id", viewID),
			builder.Equal("page_id", pageID),
		),
	)

	err := sql.Exec(statemant)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorDeletingData, "RemovePageFromView", err.Error()))

		return response
	}

	return response
}
