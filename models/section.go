package models

import "time"

// Section defines the struct of this object
type Section struct {
	ID            string    `json:"id" sql:"id" pk:"true"`
	Code          string    `json:"code" sql:"code"`
	Name          string    `json:"name" table:"core_translations" alias:"core_translations_name" sql:"value" on:"core_translations_name.structure_id = core_sch_pag_sections.id and core_translations_name.structure_field = 'name'"`
	Description   string    `json:"description" table:"core_translations" alias:"core_translations_description" sql:"value" on:"core_translations_description.structure_id = core_sch_pag_sections.id and core_translations_description.structure_field = 'description'"`
	SchemaID      string    `json:"schema_id" sql:"schema_id" fk:"true"`
	PageID        string    `json:"page_id" sql:"page_id" fk:"true"`
	CreatedBy     string    `json:"created_by" sql:"created_by"`
	CreatedByUser *User     `json:"created_by_user" table:"core_users" alias:"created_by_user" on:"created_by_user.id = core_sch_pag_sections.created_by"`
	CreatedAt     time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy     string    `json:"updated_by" sql:"updated_by"`
	UpdatedByUser *User     `json:"updated_by_user" table:"core_users" alias:"updated_by_user" on:"updated_by_user.id = core_sch_pag_sections.updated_by"`
	UpdatedAt     time.Time `json:"updated_at" sql:"updated_at"`
}
