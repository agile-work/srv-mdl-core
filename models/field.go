package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/agile-work/srv-shared"

	moduleShared "github.com/agile-work/srv-mdl-shared"
	sharedModels "github.com/agile-work/srv-mdl-shared/models"
)

// Field defines the struct of this object
type Field struct {
	ID           string                   `json:"id" sql:"id" pk:"true"`
	Code         string                   `json:"code" sql:"code" validate:"required"`
	SchemaCode   string                   `json:"schema_code" sql:"schema_code" validate:"required"`
	Type         string                   `json:"field_type" sql:"field_type" validate:"required"`
	Name         sharedModels.Translation `json:"name" sql:"name" field:"jsonb" validate:"required"`
	Description  sharedModels.Translation `json:"description" sql:"description" field:"jsonb"`
	Definitions  json.RawMessage          `json:"definitions" sql:"definitions" field:"jsonb" updatable:"false" validate:"required"`
	Validations  json.RawMessage          `json:"validations" sql:"validations" field:"jsonb"`
	DefaultValue json.RawMessage          `json:"default_value" sql:"default_value" field:"jsonb"`
	Active       bool                     `json:"active" sql:"active"`
	CreatedBy    string                   `json:"created_by" sql:"created_by"`
	CreatedAt    time.Time                `json:"created_at" sql:"created_at"`
	UpdatedBy    string                   `json:"updated_by" sql:"updated_by"`
	UpdatedAt    time.Time                `json:"updated_at" sql:"updated_at"`
}

// ProcessDefinitions parse generic definition to a specific type
func (f *Field) ProcessDefinitions(languageCode, method string) error {
	definitionJSON, err := json.Marshal(f.Definitions)
	if err != nil {
		return err
	}

	def, err := getDefinitionByType(f.Type)
	if err != nil {
		return err
	}

	switch ds := def.(type) {
	case FieldTextDefinition:
		err := parseAndValidate(definitionJSON, &ds)
		if err != nil {
			return err
		}
		definitionJSON, _ = json.Marshal(ds)
	case FieldNumberDefinition:
		err := parseAndValidate(definitionJSON, &ds)
		if err != nil {
			return err
		}
		definitionJSON, _ = json.Marshal(ds)
	case FieldDateDefinition:
		err := parseAndValidate(definitionJSON, &ds)
		if err != nil {
			return err
		}
		definitionJSON, _ = json.Marshal(ds)
	case FieldLookupDefinition:
		err := parseAndValidate(definitionJSON, &ds)
		if err != nil {
			return err
		}
		definitionJSON, _ = json.Marshal(ds)
	case FieldAttachmentDefinition:
		err := parseAndValidate(definitionJSON, &ds)
		if err != nil {
			return err
		}
		definitionJSON, _ = json.Marshal(ds)
	default:
		return errors.New("invalid definition type")
	}

	return json.Unmarshal(definitionJSON, &f.Definitions)
}

// FieldTextDefinition defines custom attributes for the text type
type FieldTextDefinition struct {
	Display string `json:"display" validate:"required"` // single_line, multi_line, readonly, enter_once, url
}

// FieldNumberDefinition defines custom attributes for the number type
type FieldNumberDefinition struct {
	Display  string           `json:"display" validate:"required"` // percentage, number, money
	Decimals int              `json:"decimals"`
	Scale    FieldNumberScale `json:"scale,omitempty"`
}

// FieldNumberScale defines a lookup to define a custom scale to a number field
type FieldNumberScale struct {
	LookupCode       string                 `json:"lookup_code" validate:"required"`
	LookupLabel      string                 `json:"lookup_label" validate:"required"`
	LookupValue      string                 `json:"lookup_value" validate:"required"`
	LookupParams     []FieldLookupParam     `json:"lookup_params"`
	AggregationRates map[string]interface{} `json:"aggr_rates"`
}

// FieldLookupParam defines the values for a lookup param in the field
type FieldLookupParam struct {
	Code     string      `json:"code"`
	DataType string      `json:"data_type"`
	Value    interface{} `json:"value"`
}

// FieldDateDefinition defines custom attributes for the date type
type FieldDateDefinition struct {
	Display string                   `json:"display" validate:"required"` // date, date_time
	Format  sharedModels.Translation `json:"format"`
}

// FieldLookupDefinition defines custom attributes for the lookup type
type FieldLookupDefinition struct {
	Type           string                 `json:"type" validate:"required"`    // static, dynamic, tree, security
	Display        string                 `json:"display" validate:"required"` // select_single, select_multiple, checkbox, radio_buttons
	LookupCode     string                 `json:"lookup_code" validate:"required"`
	LookupLabel    string                 `json:"lookup_label" validate:"required"`
	LookupValue    string                 `json:"lookup_value" validate:"required"`
	Params         map[string]interface{} `json:"lookup_params,omitempty"`
	SecurityGroups []string               `json:"security_groups,omitempty"`
	OrderType      string                 `json:"order_type,omitempty"`
	Order          []string               `json:"order,omitempty"`
}

// FieldAttachmentDefinition defines custom attributes for the attachment type
// TODO: update to deal with module "document management"
type FieldAttachmentDefinition struct {
	Display           string   `json:"display" validate:"required"` // single, multiple
	FilesMaxNum       int      `json:"files_max_number"`
	FileMaxSize       int      `json:"file_max_size"`
	FileTypeWhiteList []string `json:"file_type_white_list"`
}

func getDefinitionByType(fieldType string) (interface{}, error) {
	switch fieldType {
	case shared.FieldText:
		return FieldTextDefinition{}, nil
	case shared.FieldNumber:
		return FieldNumberDefinition{}, nil
	case shared.FieldDate:
		return FieldDateDefinition{}, nil
	case shared.FieldLookup:
		return FieldLookupDefinition{}, nil
	case shared.FieldAttachment:
		return FieldAttachmentDefinition{}, nil
	default:
		return nil, errors.New("invalid field type")
	}
}

func parseAndValidate(jsonBytes []byte, obj interface{}) error {
	err := json.Unmarshal(jsonBytes, obj)
	if err != nil {
		return err
	}
	err = moduleShared.Validate.Struct(obj)
	if err != nil {
		return err
	}
	return nil
}
