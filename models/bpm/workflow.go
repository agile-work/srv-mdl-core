package bpm

import (
	"time"

	"github.com/agile-work/srv-mdl-shared/models/translation"
)

// Workflow defines a new bussiness process on the database
type Workflow struct {
	ID          string                  `json:"id" sql:"id" pk:"true"`
	Code        string                  `json:"code" sql:"code"`
	ContentCode string                  `json:"content_code" sql:"content_code"`
	SchemaCode  string                  `json:"schema_code" sql:"schema_code"`
	Name        translation.Translation `json:"name" sql:"name" field:"jsonb" validate:"required"`
	Description translation.Translation `json:"description" sql:"description" field:"jsonb" validate:"required"`
	Active      bool                    `json:"active" sql:"active"`
	CreatedBy   string                  `json:"created_by" sql:"created_by"`
	CreatedAt   time.Time               `json:"created_at" sql:"created_at"`
	UpdatedBy   string                  `json:"updated_by" sql:"updated_by"`
	UpdatedAt   time.Time               `json:"updated_at" sql:"updated_at"`
}
