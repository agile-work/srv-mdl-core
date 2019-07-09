package bpm

import (
	"encoding/json"

	"github.com/agile-work/srv-mdl-shared/models/translation"
)

// StepDefinition define the specific data for a step
type StepDefinition struct {
	Actions        []Action
	PostConditions []Condition
}

// Condition define a condition to validate the next step
type Condition struct {
	Expression string `json:"expresion" validate:"required"`
	StepCode   string `json:"step_code" validate:"required"`
}

// Action define the action to be executed in the step
type Action struct {
	Code        string                  `json:"code" validate:"required"`
	Name        translation.Translation `json:"name" validate:"required"`
	Description translation.Translation `json:"description"`
	Type        string                  `json:"action_type" validate:"required"` // update_schema, set_param, notify, approval, execute_job
	Definitions json.RawMessage         `json:"action_definition" validate:"required"`
}

// UpdateSchemaAction specify the action type
type UpdateSchemaAction struct {
	Field string `json:"field" validate:"required"` // contract.cost, parent.cost, num_executions
	Value string `json:"value" validate:"required"` // get(workflow.param_code), get(step.step_code.param_code), const(1000), get(parent.cost), get(contract.cost), add(contract.cost, contract.cost_total)
}

// SetParam specify the action type
type SetParam struct {
	Key   string `json:"key" validate:"required"`
	Scope string `json:"scope" validate:"required"` // workflow, step
	Value string `json:"value" validate:"required"` // get(workflow.param_code), get(step.step_code.param_code), const(1000), get(parent.cost), get(contract.cost), add(contract.cost, contract.cost_total)
}

// NotifyAction specify the action type
type NotifyAction struct {
	Recipients []Recipient             `json:"recipients" validate:"required"`
	Subject    translation.Translation `json:"subject"`
	Message    translation.Translation `json:"description"`
}

// Recipient defines with structure should be notifyed
type Recipient struct {
	Code string `json:"code" validate:"required"`
	Type string `json:"recipient_type" validate:"required"` //user, group
}

// ApprovalAction specify the action type
type ApprovalAction struct {
	Recipients      []Recipient             `json:"recipients" validate:"required"`
	Subject         translation.Translation `json:"subject" validate:"required"`
	Message         translation.Translation `json:"description" validate:"required"`
	ApprovalOptions []string                `json:"approval_options" validate:"required"`
}

// ExecuteJobAction specify the action type
type ExecuteJobAction struct {
	Code         string     `json:"code" validate:"required"`
	Params       []JobParam `json:"params"`
	WaitComplete bool       `json:"wait_complete"`
}

// JobParam define the job initialization params
type JobParam struct {
	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"` // get(workflow.param_code), get(step.step_code.param_code), const(1000), get(parent.cost), get(contract.cost), add(contract.cost, contract.cost_total)
}
