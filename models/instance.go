package models

import (
	"encoding/json"
	"time"
)

// Instance defines the struct of this object
type Instance struct {
	ID        string          `json:"id" sql:"id" pk:"true"`
	Data      json.RawMessage `json:"data" sql:"data" field:"jsonb"`
	CreatedBy string          `json:"created_by" sql:"created_by"`
	CreatedAt time.Time       `json:"created_at" sql:"created_at"`
	UpdatedBy string          `json:"updated_by" sql:"updated_by"`
	UpdatedAt time.Time       `json:"updated_at" sql:"updated_at"`
}
