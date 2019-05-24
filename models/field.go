package models

import (
	"time"
)

// Field defines the struct of this object
type Field struct {
	ID          string       `json:"id" sql:"id" pk:"true"`
	Code        string       `json:"code" sql:"code"`
	SchemaID    string       `json:"schema_id" sql:"schema_id" fk:"true"`
	Name        string       `json:"name" table:"core_translations" alias:"core_translations_name" sql:"value" on:"core_translations_name.structure_id = core_sch_fields.id and core_translations_name.structure_field = 'name'"`
	Description string       `json:"description" table:"core_translations" alias:"core_translations_description" sql:"value" on:"core_translations_description.structure_id = core_sch_fields.id and core_translations_description.structure_field = 'description'"`
	FieldType   string       `json:"field_type" sql:"field_type"`
	Multivalue  bool         `json:"multivalue" sql:"multivalue"`
	LookupID    string       `json:"lookup_id" sql:"lookup_id" fk:"true"`
	Groups      []FieldGroup `json:"groups" sql:"groups" field:"jsonb"`
	Active      bool         `json:"active" sql:"active"`
	CreatedBy   string       `json:"created_by" sql:"created_by"`
	CreatedAt   time.Time    `json:"created_at" sql:"created_at"`
	UpdatedBy   string       `json:"updated_by" sql:"updated_by"`
	UpdatedAt   time.Time    `json:"updated_at" sql:"updated_at"`
}

// FieldGroup defines the struct of this object
type FieldGroup struct {
	ID        string    `json:"id"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

// FieldValidation defines the struct of this object
type FieldValidation struct {
	ID         string    `json:"id" sql:"id" pk:"true"`
	SchemaID   string    `json:"schema_id" sql:"schema_id" fk:"true"`
	FieldID    string    `json:"field_id" sql:"field_id" fk:"true"`
	Validation string    `json:"validation" sql:"validation"`
	ValidWhen  string    `json:"valid_when" sql:"valid_when"`
	CreatedBy  string    `json:"created_by" sql:"created_by"`
	CreatedAt  time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy  string    `json:"updated_by" sql:"updated_by"`
	UpdatedAt  time.Time `json:"updated_at" sql:"updated_at"`
}
