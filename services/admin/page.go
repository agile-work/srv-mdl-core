package admin

import (
	"fmt"
	"net/http"

	"github.com/agile-work/srv-mdl-shared/db"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/agile-work/srv-mdl-core/models"
	moduleShared "github.com/agile-work/srv-mdl-shared"
	shared "github.com/agile-work/srv-shared"
)

// CreatePage persists the request body creating a new object in the database
func CreatePage(r *http.Request) *moduleShared.Response {
	schemaID := chi.URLParam(r, "schema_id")
	page := models.Page{
		SchemaID: schemaID,
	}

	return db.Create(r, &page, "CreatePage", shared.TableCoreSchemaPages)
}

// LoadAllPages return all instances from the object
func LoadAllPages(r *http.Request) *moduleShared.Response {
	pages := []models.Page{}
	schemaID := chi.URLParam(r, "schema_id")
	schemaIDColumn := fmt.Sprintf("%s.schema_id", shared.TableCoreSchemaPages)
	condition := builder.Equal(schemaIDColumn, schemaID)

	return db.Load(r, &pages, "LoadAllPages", shared.TableCoreSchemaPages, condition)
}

// LoadPage return only one object from the database
func LoadPage(r *http.Request) *moduleShared.Response {
	page := models.Page{}
	pageID := chi.URLParam(r, "page_id")
	pageIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchemaPages)
	condition := builder.Equal(pageIDColumn, pageID)

	return db.Load(r, &page, "LoadPage", shared.TableCoreSchemaPages, condition)
}

// UpdatePage updates object data in the database
func UpdatePage(r *http.Request) *moduleShared.Response {
	pageID := chi.URLParam(r, "page_id")
	pageIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchemaPages)
	condition := builder.Equal(pageIDColumn, pageID)
	page := models.Page{
		ID: pageID,
	}

	return db.Update(r, &page, "UpdatePage", shared.TableCoreSchemaPages, condition)
}

// DeletePage deletes object from the database
func DeletePage(r *http.Request) *moduleShared.Response {
	pageID := chi.URLParam(r, "page_id")
	pageIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchemaPages)
	condition := builder.Equal(pageIDColumn, pageID)

	return db.Remove(r, "DeletePage", shared.TableCoreSchemaPages, condition)
}
