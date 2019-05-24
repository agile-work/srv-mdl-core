package models

import (
	"time"
)

// Schema defines the struct of this object
type Schema struct {
	ID          string    `json:"id" sql:"id" pk:"true"`
	JobID       string    `json:"job_id" sql:"job_id" fk:"true"`
	Code        string    `json:"code" sql:"code"`
	Name        string    `json:"name" table:"core_translations" alias:"core_translations_name" sql:"value" on:"core_translations_name.structure_id = core_schemas.id and core_translations_name.structure_field = 'name'"`
	Description string    `json:"description" table:"core_translations" alias:"core_translations_description" sql:"value" on:"core_translations_description.structure_id = core_schemas.id and core_translations_description.structure_field = 'description'"`
	Parent      string    `json:"parent_id" sql:"parent_id"`
	IsExtension bool      `json:"is_extension" sql:"is_extension"`
	Module      bool      `json:"module" sql:"module"`
	Active      bool      `json:"active" sql:"active"`
	Status      string    `json:"status" sql:"status"`
	CreatedBy   string    `json:"created_by" sql:"created_by"`
	CreatedAt   time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy   string    `json:"updated_by" sql:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at" sql:"updated_at"`
}

// ViewSchModules defines the struct of this object
type ViewSchModules struct {
	ID           string    `json:"id" sql:"id" pk:"true"`
	SchemaID     string    `json:"schema_id" sql:"schema_id"`
	Code         string    `json:"code" sql:"code"`
	Name         string    `json:"name" sql:"name"`
	Description  string    `json:"description" sql:"description"`
	LanguageCode string    `json:"language_code" sql:"language_code"`
	Module       bool      `json:"module" sql:"module"`
	Active       bool      `json:"active" sql:"active"`
	CreatedBy    string    `json:"created_by" sql:"created_by"`
	CreatedAt    time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy    string    `json:"updated_by" sql:"updated_by"`
	UpdatedAt    time.Time `json:"updated_at" sql:"updated_at"`
}
