package models

import (
	"time"
)

// View defines the struct of this object
type View struct {
	ID          string    `json:"id" sql:"id" pk:"true"`
	Code        string    `json:"code" sql:"code"`
	SchemaID    string    `json:"schema_id" sql:"schema_id" fk:"true"`
	Name        string    `json:"name" table:"core_translations" alias:"core_translations_name" sql:"value" on:"core_translations_name.structure_id = core_sch_views.id and core_translations_name.structure_field = 'name'"`
	Description string    `json:"description" table:"core_translations" alias:"core_translations_description" sql:"value" on:"core_translations_description.structure_id = core_sch_views.id and core_translations_description.structure_field = 'description'"`
	Active      bool      `json:"active" sql:"active"`
	CreatedBy   string    `json:"created_by" sql:"created_by"`
	CreatedAt   time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy   string    `json:"updated_by" sql:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at" sql:"updated_at"`
}
