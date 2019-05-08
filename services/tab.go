package services

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

// CreateTab persists the request body creating a new object in the database
func CreateTab(r *http.Request) *moduleShared.Response {
	tab := models.Tab{}

	return db.Create(r, &tab, "CreateTab", shared.TableCoreSchPagSecTabs)
}

// LoadAllTabs return all instances from the object
func LoadAllTabs(r *http.Request) *moduleShared.Response {
	tabs := []models.Tab{}
	sectionID := chi.URLParam(r, "section_id")
	sectionIDColumn := fmt.Sprintf("%s.section_id", shared.TableCoreSchPagSecTabs)
	condition := builder.Equal(sectionIDColumn, sectionID)

	return db.Load(r, &tabs, "LoadAllTabs", shared.TableCoreSchPagSecTabs, condition)
}

// LoadTab return only one object from the database
func LoadTab(r *http.Request) *moduleShared.Response {
	tab := models.Tab{}
	tabID := chi.URLParam(r, "tab_id")
	tabIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchPagSecTabs)
	condition := builder.Equal(tabIDColumn, tabID)

	return db.Load(r, &tab, "LoadTab", shared.TableCoreSchPagSecTabs, condition)
}

// UpdateTab updates object data in the database
func UpdateTab(r *http.Request) *moduleShared.Response {
	tabID := chi.URLParam(r, "tab_id")
	tabIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchPagSecTabs)
	condition := builder.Equal(tabIDColumn, tabID)
	tab := models.Tab{
		ID: tabID,
	}

	return db.Update(r, &tab, "UpdateTab", shared.TableCoreSchPagSecTabs, condition)
}

// DeleteTab deletes object from the database
func DeleteTab(r *http.Request) *moduleShared.Response {
	tabID := chi.URLParam(r, "tab_id")
	tabIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchPagSecTabs)
	condition := builder.Equal(tabIDColumn, tabID)

	return db.Remove(r, "DeleteTab", shared.TableCoreSchPagSecTabs, condition)
}
