package field

import (
	"encoding/json"
	"fmt"

	"github.com/agile-work/srv-mdl-shared/models/response"

	"github.com/agile-work/srv-mdl-shared/models/translation"

	mdlShared "github.com/agile-work/srv-mdl-shared"
)

// LookupDefinition defines custom attributes for the lookup type
type LookupDefinition struct {
	Display        string            `json:"display" validate:"required"`     // select_single, select_multiple, checkbox, radio_buttons
	Type           string            `json:"lookup_type" validate:"required"` // static, dynamic, tree, security
	LookupLabel    string            `json:"lookup_label" validate:"required"`
	LookupValue    string            `json:"lookup_value" validate:"required"`
	DatasetCode    string            `json:"dataset_code" validate:"required"`
	LookupFields   []lookupField     `json:"lookup_fields,omitempty"`
	LookupParams   []lookupParam     `json:"lookup_params,omitempty"`
	SecurityGroups []string          `json:"security_groups,omitempty"`
	Options        response.Metadata `json:"options,omitempty"`
}

type lookupParam struct {
	Code      string      `json:"code"`
	DataType  string      `json:"data_type"`
	ValueType string      `json:"value_type"` // column, constant
	Value     interface{} `json:"value"`
}

type lookupField struct {
	Code     string                  `json:"code"`
	DataType string                  `json:"data_type"`
	Label    translation.Translation `json:"label"`
	Filter   lookupFieldFilter       `json:"filter"`
}

type lookupFieldFilter struct {
	ValueType string      `json:"value_type"` // column, constant
	Value     interface{} `json:"value"`
	Readonly  bool        `json:"readonly"`
}

func (d *LookupDefinition) load(payload json.RawMessage) error {
	if err := json.Unmarshal(payload, d); err != nil {
		return err
	}
	return nil
}

func (d *LookupDefinition) validate() error {
	if err := mdlShared.Validate.Struct(d); err != nil {
		return err
	}
	if len(d.LookupFields) <= 0 {
		return fmt.Errorf("lookup without fields")
	}

	return nil
}

// UpdateLookupParam inset if not exists or change param
func (d *LookupDefinition) UpdateLookupParam() error {
	return nil
}

// UpdateLookupField inset if not exists or change field
func (d *LookupDefinition) UpdateLookupField() error {
	return nil
}

// UpdateSecurityGroup inset if not exists or change security group
func (d *LookupDefinition) UpdateSecurityGroup() error {
	return nil
}
