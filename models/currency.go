package models

import (
	"time"

	sharedModels "github.com/agile-work/srv-mdl-shared/models"
)

// Currency defines the struct of this object
type Currency struct {
	ID        string                    `json:"id" sql:"id" pk:"true"`
	Code      string                    `json:"code" sql:"code"`
	Name      sharedModels.Translation  `json:"name" sql:"name" field:"jsonb"`
	Rates     map[string][]CurrencyRate `json:"rates" sql:"rates" field:"jsonb"`
	Active    bool                      `json:"active" sql:"active"`
	CreatedBy string                    `json:"created_by" sql:"created_by"`
	CreatedAt time.Time                 `json:"created_at" sql:"created_at"`
	UpdatedBy string                    `json:"updated_by" sql:"updated_by"`
	UpdatedAt time.Time                 `json:"updated_at" sql:"updated_at"`
}

// CurrencyRate defines the struct of this object
type CurrencyRate struct {
	Value float64    `json:"value"`
	Start *time.Time `json:"start_at,omitempty"`
	End   *time.Time `json:"end_at,omitempty"`
}
