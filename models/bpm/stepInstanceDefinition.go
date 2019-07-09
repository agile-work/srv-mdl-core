package bpm

import (
	"encoding/json"

	"github.com/agile-work/srv-mdl-shared/models/user"
)

// StepInstanceDefinition define the specific data for a step
type StepInstanceDefinition struct {
	Params  map[string]interface{} `json:"params"`
	Actions []InstanceAction       `json:"actions"`
}

// InstanceAction define the action to be executed in the step
type InstanceAction struct {
	ID          string          `json:"id" validate:"required"`
	Type        string          `json:"action_type" validate:"required"` // update_schema, set_param, notify, approval, execute_job
	Definitions json.RawMessage `json:"action_definition" validate:"required"`
}

// UpdateSchemaActionInstance specify the action type
type UpdateSchemaActionInstance struct {
	Field  string      `json:"field" validate:"required"` // contract.cost, parent.cost, num_executions
	Value  interface{} `json:"value" validate:"required"` // processed value from definition
	Status string      `json:"status" validate:"required"`
}

// SetParamActionInstance specify the action type
type SetParamActionInstance struct {
	Status string `json:"status" validate:"required"`
}

// NotifyActionInstance specify the action type
type NotifyActionInstance struct {
	Users  []user.User `json:"users" validate:"required"`
	Status string      `json:"status" validate:"required"`
}

// ApprovalActionInstance specify the action type
type ApprovalActionInstance struct {
	Users           []user.User       `json:"users" validate:"required"`
	ApprovalResults map[string]string `json:"approval_options" validate:"required"` // map[username]option_code
	Status          string            `json:"status" validate:"required"`
}

// ExecuteJobAcionInstance specify the action type
type ExecuteJobAcionInstance struct {
	Status        string `json:"status" validate:"required"`
	JobInstanceID string `json:"job_instance_id"`
}
