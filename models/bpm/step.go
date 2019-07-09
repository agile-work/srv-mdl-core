package bpm

import (
	"encoding/json"
	"time"

	"github.com/agile-work/srv-mdl-shared/models/translation"
)

// Step defines a new business process on the database
type Step struct {
	ID           string                  `json:"id" sql:"id" pk:"true"`
	Code         string                  `json:"code" sql:"code" validate:"required"`
	WorkflowCode string                  `json:"bpm_code" sql:"bpm_code" validate:"required"`
	ContentCode  string                  `json:"content_code" sql:"content_code"`
	SchemaCode   string                  `json:"schema_code" sql:"schema_code"`
	Name         translation.Translation `json:"name" sql:"name" field:"jsonb" validate:"required"`
	Description  translation.Translation `json:"description" sql:"description" field:"jsonb" validate:"required"`
	Definitions  json.RawMessage         `json:"definitions" sql:"definitions" field:"jsonb" updatable:"false" validate:"required"`
	Active       bool                    `json:"active" sql:"active"`
	CreatedBy    string                  `json:"created_by" sql:"created_by"`
	CreatedAt    time.Time               `json:"created_at" sql:"created_at"`
	UpdatedBy    string                  `json:"updated_by" sql:"updated_by"`
	UpdatedAt    time.Time               `json:"updated_at" sql:"updated_at"`
}
