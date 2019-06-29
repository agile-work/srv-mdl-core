package field

import (
	"encoding/json"

	"github.com/agile-work/srv-mdl-shared/models/response"

	mdlShared "github.com/agile-work/srv-mdl-shared"
)

// LookupDefinition defines custom attributes for the lookup type
type LookupDefinition struct {
	Display        string            `json:"display" validate:"required"`     // select_single, select_multiple, checkbox, radio_buttons
	Type           string            `json:"lookup_type" validate:"required"` // static, dynamic, tree, security
	LookupCode     string            `json:"lookup_code" validate:"required"`
	LookupLabel    string            `json:"lookup_label" validate:"required"`
	LookupValue    string            `json:"lookup_value" validate:"required"`
	LookupParams   []LookupParam     `json:"lookup_params,omitempty"`
	Options        response.Metadata `json:"options,omitempty"`
	SecurityGroups []string          `json:"security_groups,omitempty"`
}

// LookupParam defines the values for a lookup param in the field
type LookupParam struct {
	Code       string           `json:"code"`
	DataType   string           `json:"data_type"`
	Definition LookupParamValue `json:"definition"`
}

// LookupParamValue defines the values for a lookup param in the field
type LookupParamValue struct {
	ValueType string      `json:"value_type"` // column, constant
	Value     interface{} `json:"value"`
}

func (d *LookupDefinition) load(payload json.RawMessage) error {
	if err := json.Unmarshal(payload, d); err != nil {
		return err
	}

	if err := mdlShared.Validate.Struct(d); err != nil {
		return err
	}
	return nil
}
