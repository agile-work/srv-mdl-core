package models

import (
	"time"

	"github.com/agile-work/srv-mdl-shared/models"
)

// Widget defines the struct of this object
type Widget struct {
	ID            string       `json:"id" sql:"id" pk:"true"`
	Code          string       `json:"code" sql:"code"`
	Name          string       `json:"name" table:"core_translations" alias:"core_translations_name" sql:"value" on:"core_translations_name.structure_id = core_widgets.id and core_translations_name.structure_field = 'name'"`
	Description   string       `json:"description" table:"core_translations" alias:"core_translations_description" sql:"value" on:"core_translations_description.structure_id = core_widgets.id and core_translations_description.structure_field = 'description'"`
	WidgetType    string       `json:"widget_type" sql:"widget_type"`
	Query         string       `json:"query" sql:"query"`
	Active        bool         `json:"active" sql:"active"`
	CreatedBy     string       `json:"created_by" sql:"created_by"`
	CreatedByUser *models.User `json:"created_by_user" table:"core_users" alias:"created_by_user" on:"created_by_user.id = core_widgets.created_by"`
	CreatedAt     time.Time    `json:"created_at" sql:"created_at"`
	UpdatedBy     string       `json:"updated_by" sql:"updated_by"`
	UpdatedByUser *models.User `json:"updated_by_user" table:"core_users" alias:"updated_by_user" on:"updated_by_user.id = core_widgets.updated_by"`
	UpdatedAt     time.Time    `json:"updated_at" sql:"updated_at"`
}
