package bpm

import (
	"encoding/json"
	"time"
)

// StepInstance defines a new instance for a workflow step
type StepInstance struct {
	ID              string          `json:"id" sql:"id" pk:"true"`
	BPMStepCode     string          `json:"bpm_step_code" sql:"bpm_step_code" validate:"required"`
	BPMInstanceCode string          `json:"bpm_code" sql:"bpm_code" validate:"required"`
	Definitions     json.RawMessage `json:"definitions" sql:"definitions" field:"jsonb" updatable:"false" validate:"required"`
	Status          string          `json:"status" sql:"status" validate:"required"`
	StartedAt       time.Time       `json:"started_at" sql:"started_at"`
	EndAt           time.Time       `json:"end_at" sql:"end_at"`
}
