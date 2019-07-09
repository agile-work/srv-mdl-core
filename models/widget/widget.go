package widget

import (
	"time"

	"github.com/agile-work/srv-mdl-shared/models/translation"
)

// Widget defines the struct of this object
type Widget struct {
	ID          string                  `json:"id" sql:"id" pk:"true"`
	Code        string                  `json:"code" sql:"code" updatable:"false" validate:"required"`
	ContentCode string                  `json:"content_code" sql:"content_code"`
	Name        translation.Translation `json:"name" sql:"name" field:"jsonb" validate:"required"`
	Description translation.Translation `json:"description" sql:"description" field:"jsonb" validate:"required"`
	WidgetType  string                  `json:"widget_type" sql:"widget_type" validate:"required"`
	Query       string                  `json:"query" sql:"query" validate:"required"`
	Active      bool                    `json:"active" sql:"active"`
	CreatedBy   string                  `json:"created_by" sql:"created_by"`
	CreatedAt   time.Time               `json:"created_at" sql:"created_at"`
	UpdatedBy   string                  `json:"updated_by" sql:"updated_by"`
	UpdatedAt   time.Time               `json:"updated_at" sql:"updated_at"`
}
