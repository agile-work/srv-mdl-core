package models

import "time"

// Lookup defines the struct of this object
type Lookup struct {
	ID            string    `json:"id" sql:"id" pk:"true"`
	Code          string    `json:"code" sql:"code"`
	Name          string    `json:"name" table:"core_translations" alias:"core_translations_name" sql:"value" on:"core_translations_name.structure_id = core_lookups.id and core_translations_name.structure_field = 'name'"`
	Description   string    `json:"description" table:"core_translations" alias:"core_translations_description" sql:"value" on:"core_translations_description.structure_id = core_lookups.id and core_translations_description.structure_field = 'description'"`
	Type          string    `json:"type" sql:"type"`
	Query         string    `json:"query" sql:"query"`
	Value         string    `json:"value" sql:"value"`
	Label         string    `json:"label" sql:"label"`
	Autocomplete  string    `json:"autocomplete" sql:"autocomplete"`
	Active        bool      `json:"active" sql:"active"`
	CreatedBy     string    `json:"created_by" sql:"created_by"`
	CreatedByUser *User     `json:"created_by_user" table:"core_users" alias:"created_by_user" on:"created_by_user.id = core_lookups.created_by"`
	CreatedAt     time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy     string    `json:"updated_by" sql:"updated_by"`
	UpdatedByUser *User     `json:"updated_by_user" table:"core_users" alias:"updated_by_user" on:"updated_by_user.id = core_lookups.updated_by"`
	UpdatedAt     time.Time `json:"updated_at" sql:"updated_at"`
}

// LookupOption defines the struct of this object
type LookupOption struct {
	ID            string    `json:"id" sql:"id" pk:"true"`
	Code          string    `json:"code" sql:"code"`
	LookupID      string    `json:"lookup_id" sql:"lookup_id" fk:"true"`
	Value         string    `json:"value" sql:"value"`
	Label         string    `json:"label" table:"core_translations" alias:"core_translations_label" sql:"value" on:"core_translations_label.structure_id = core_lkp_options.id and core_translations_label.structure_field = 'label'"`
	Active        bool      `json:"active" sql:"active"`
	CreatedBy     string    `json:"created_by" sql:"created_by"`
	CreatedByUser *User     `json:"created_by_user" table:"core_users" alias:"created_by_user" on:"created_by_user.id = core_lkp_options.created_by"`
	CreatedAt     time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy     string    `json:"updated_by" sql:"updated_by"`
	UpdatedByUser *User     `json:"updated_by_user" table:"core_users" alias:"updated_by_user" on:"updated_by_user.id = core_lkp_options.updated_by"`
	UpdatedAt     time.Time `json:"updated_at" sql:"updated_at"`
}
