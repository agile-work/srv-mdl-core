package models

import "time"

// Field defines the struct of this object
type Field struct {
	ID            string    `json:"id" sql:"id" pk:"true"`
	Code          string    `json:"code" sql:"code"`
	SchemaID      string    `json:"schema_id" sql:"schema_id" fk:"true"`
	Name          string    `json:"name" table:"core_translations" alias:"core_translations_name" sql:"value" on:"core_translations_name.structure_id = core_sch_fields.id and core_translations_name.structure_field = 'name'"`
	Description   string    `json:"description" table:"core_translations" alias:"core_translations_description" sql:"value" on:"core_translations_description.structure_id = core_sch_fields.id and core_translations_description.structure_field = 'description'"`
	FieldType     string    `json:"field_type" sql:"field_type"`
	Multivalue    bool      `json:"multivalue" sql:"multivalue"`
	LookupID      string    `json:"lookup_id" sql:"lookup_id" fk:"true"`
	Permission    int       `json:"permission"`
	Active        bool      `json:"active" sql:"active"`
	CreatedBy     string    `json:"created_by" sql:"created_by"`
	CreatedByUser *User     `json:"created_by_user" table:"core_users" alias:"created_by_user" on:"created_by_user.id = core_sch_fields.created_by"`
	CreatedAt     time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy     string    `json:"updated_by" sql:"updated_by"`
	UpdatedByUser *User     `json:"updated_by_user" table:"core_users" alias:"updated_by_user" on:"updated_by_user.id = core_sch_fields.updated_by"`
	UpdatedAt     time.Time `json:"updated_at" sql:"updated_at"`
}

// FieldValidation defines the struct of this object
type FieldValidation struct {
	ID            string    `json:"id" sql:"id" pk:"true"`
	SchemaID      string    `json:"schema_id" sql:"schema_id" fk:"true"`
	FieldID       string    `json:"field_id" sql:"field_id" fk:"true"`
	Validation    string    `json:"validation" sql:"validation"`
	ValidWhen     string    `json:"valid_when" sql:"valid_when"`
	CreatedBy     string    `json:"created_by" sql:"created_by"`
	CreatedByUser *User     `json:"created_by_user" table:"core_users" alias:"created_by_user" on:"created_by_user.id = core_sch_fld_validations.created_by"`
	CreatedAt     time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy     string    `json:"updated_by" sql:"updated_by"`
	UpdatedByUser *User     `json:"updated_by_user" table:"core_users" alias:"updated_by_user" on:"updated_by_user.id = core_sch_fld_validations.updated_by"`
	UpdatedAt     time.Time `json:"updated_at" sql:"updated_at"`
}
