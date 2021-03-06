package field

import (
	"encoding/json"
	"fmt"

	"github.com/agile-work/srv-mdl-core/models/dataset"
	"github.com/agile-work/srv-mdl-core/models/group"
	"github.com/agile-work/srv-mdl-shared/models/response"
	"github.com/agile-work/srv-shared/constants"

	"github.com/agile-work/srv-mdl-shared/models/translation"

	mdlShared "github.com/agile-work/srv-mdl-shared"
)

// LookupDefinition defines custom attributes for the lookup type
type LookupDefinition struct {
	Display        string            `json:"display" validate:"required"`     // select_single, select_multiple, radio_buttons, checkboxes
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
	ValueType string      `json:"value_type"` // column, constant
	Value     interface{} `json:"value"`
}

type lookupField struct {
	Code   string                  `json:"code"`
	Label  translation.Translation `json:"label"`
	Filter lookupFieldFilter       `json:"filter"`
}

type lookupFieldFilter struct {
	ValueType string      `json:"value_type"` // column, constant
	Value     interface{} `json:"value"`
	Operator  string      `json:"operator"`
	Readonly  bool        `json:"readonly"`
}

func (d *LookupDefinition) parse(payload json.RawMessage) error {
	if err := json.Unmarshal(payload, d); err != nil {
		return err
	}
	if d.Type == constants.FieldLookupStatic {
		d.LookupLabel = "label"
		d.LookupValue = "code"
	}
	return nil
}

func (d *LookupDefinition) prepare() error {
	if err := mdlShared.Validate.Struct(d); err != nil {
		return err
	}

	if d.Type != constants.FieldLookupStatic {
		if len(d.LookupFields) <= 0 {
			return fmt.Errorf("%s lookup without fields", d.Type)
		}

		columns := []string{}
		for _, f := range d.LookupFields {
			columns = append(columns, f.Code)
		}
		params := []string{}
		for _, p := range d.LookupParams {
			params = append(params, p.Code)
		}

		if err := dataset.Validate(d.DatasetCode, false, columns, params); err != nil {
			return err
		}

		if d.Type == constants.FieldLookupSecurity {
			if len(d.SecurityGroups) <= 0 {
				return fmt.Errorf("security lookup without security groups")
			}
			if err := group.Validate(d.SecurityGroups); err != nil {
				return err
			}
		}
	} else {
		if err := dataset.Validate(d.DatasetCode, true, nil, nil); err != nil {
			return err
		}
	}

	return nil
}

// UpdateLookupParam insert if not exists or change param
func (d *LookupDefinition) UpdateLookupParam() error { // TODO: Passar sempre todos os parâmetros
	return nil
}

// UpdateLookupField change field
func (d *LookupDefinition) UpdateLookupField() error { // TODO: passar sempre todos os fields
	return nil
}

// UpdateSecurityGroup insert if not exists or change security group
func (d *LookupDefinition) UpdateSecurityGroup() error {
	return nil
}
