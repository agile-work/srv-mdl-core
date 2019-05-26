package admin

import (
	"fmt"
	"net/http"

	moduleShared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-mdl-shared/db"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/agile-work/srv-mdl-core/models"
	shared "github.com/agile-work/srv-shared"
)

// CreateWidget persists the request body creating a new object in the database
func CreateWidget(r *http.Request) *moduleShared.Response {
	widget := models.Widget{}

	return db.Create(r, &widget, "CreateWidget", shared.TableCoreWidgets)
}

// LoadAllWidgets return all instances from the object
func LoadAllWidgets(r *http.Request) *moduleShared.Response {
	widgets := []models.Widget{}

	return db.Load(r, &widgets, "LoadAllWidgets", shared.TableCoreWidgets, nil)
}

// LoadWidget return only one object from the database
func LoadWidget(r *http.Request) *moduleShared.Response {
	widget := models.Widget{}
	widgetID := chi.URLParam(r, "widget_id")
	widgetIDColumn := fmt.Sprintf("%s.id", shared.TableCoreWidgets)
	condition := builder.Equal(widgetIDColumn, widgetID)

	return db.Load(r, &widget, "LoadWidget", shared.TableCoreWidgets, condition)
}

// UpdateWidget updates object data in the database
func UpdateWidget(r *http.Request) *moduleShared.Response {
	widgetID := chi.URLParam(r, "widget_id")
	widgetIDColumn := fmt.Sprintf("%s.id", shared.TableCoreWidgets)
	condition := builder.Equal(widgetIDColumn, widgetID)
	widget := models.Widget{
		ID: widgetID,
	}

	return db.Update(r, &widget, "UpdateWidget", shared.TableCoreWidgets, condition)
}

// DeleteWidget deletes object from the database
func DeleteWidget(r *http.Request) *moduleShared.Response {
	widgetID := chi.URLParam(r, "widget_id")
	widgetIDColumn := fmt.Sprintf("%s.id", shared.TableCoreWidgets)
	condition := builder.Equal(widgetIDColumn, widgetID)

	return db.Remove(r, "DeleteWidget", shared.TableCoreWidgets, condition)
}
