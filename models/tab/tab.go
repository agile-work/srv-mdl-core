package tab

import (
	"time"
)

// Tab defines the struct of this object
type Tab struct {
	ID          string    `json:"id" sql:"id" pk:"true"`
	Code        string    `json:"code" sql:"code"`
	Name        string    `json:"name" table:"core_translations" alias:"core_translations_name" sql:"value" on:"core_translations_name.structure_id = core_sch_pag_sec_tabs.id and core_translations_name.structure_field = 'name'"`
	Description string    `json:"description" table:"core_translations" alias:"core_translations_description" sql:"value" on:"core_translations_description.structure_id = core_sch_pag_sec_tabs.id and core_translations_description.structure_field = 'description'"`
	SchemaID    string    `json:"schema_id" sql:"schema_id" fk:"true"`
	PageID      string    `json:"page_id" sql:"page_id" fk:"true"`
	SectionID   string    `json:"section_id" sql:"section_id" fk:"true"`
	TabOrder    int       `json:"tab_order" sql:"tab_order"`
	CreatedBy   string    `json:"created_by" sql:"created_by"`
	CreatedAt   time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy   string    `json:"updated_by" sql:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at" sql:"updated_at"`
}
