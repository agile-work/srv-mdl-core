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

// CreateSection persists the request body creating a new object in the database
func CreateSection(r *http.Request) *moduleShared.Response {
	schemaID := chi.URLParam(r, "schema_id")
	pageID := chi.URLParam(r, "page_id")
	section := models.Section{
		SchemaID: schemaID,
		PageID:   pageID,
	}

	return db.Create(r, &section, "CreateSection", shared.TableCoreSchemaPagSections)
}

// LoadAllSections return all instances from the object
func LoadAllSections(r *http.Request) *moduleShared.Response {
	sections := []models.Section{}
	pageID := chi.URLParam(r, "page_id")
	pageIDColumn := fmt.Sprintf("%s.page_id", shared.TableCoreSchemaPagSections)
	condition := builder.Equal(pageIDColumn, pageID)

	return db.Load(r, &sections, "LoadAllSections", shared.TableCoreSchemaPagSections, condition)
}

// LoadSection return only one object from the database
func LoadSection(r *http.Request) *moduleShared.Response {
	section := models.Section{}
	sectionID := chi.URLParam(r, "section_id")
	sectionIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchemaPagSections)
	condition := builder.Equal(sectionIDColumn, sectionID)

	return db.Load(r, &section, "LoadSection", shared.TableCoreSchemaPagSections, condition)
}

// UpdateSection updates object data in the database
func UpdateSection(r *http.Request) *moduleShared.Response {
	sectionID := chi.URLParam(r, "section_id")
	sectionIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchemaPagSections)
	condition := builder.Equal(sectionIDColumn, sectionID)
	section := models.Section{
		ID: sectionID,
	}

	return db.Update(r, &section, "UpdateSection", shared.TableCoreSchemaPagSections, condition)
}

// DeleteSection deletes object from the database
func DeleteSection(r *http.Request) *moduleShared.Response {
	sectionID := chi.URLParam(r, "section_id")
	sectionIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchemaPagSections)
	condition := builder.Equal(sectionIDColumn, sectionID)

	return db.Remove(r, "DeleteSection", shared.TableCoreSchemaPagSections, condition)
}
