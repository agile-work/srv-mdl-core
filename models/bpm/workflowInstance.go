package bpm

import (
	"time"
)

// WorkflowInstance defines a new running instance for a workflow
type WorkflowInstance struct {
	ID        string    `json:"id" sql:"id" pk:"true"`
	BPMCode   string    `json:"bpm_code" sql:"bpm_code"`
	Status    string    `json:"status" sql:"status" validate:"required"`
	Active    bool      `json:"active" sql:"active"`
	CreatedBy string    `json:"created_by" sql:"created_by"`
	CreatedAt time.Time `json:"created_at" sql:"created_at"`
	StartedAt time.Time `json:"started_at" sql:"started_at"`
	EndAt     time.Time `json:"end_at" sql:"end_at"`
}
