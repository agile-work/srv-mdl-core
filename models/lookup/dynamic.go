package lookup

import (
	"encoding/json"
	"time"

	mdlShared "github.com/agile-work/srv-mdl-shared"
	mdlSharedModels "github.com/agile-work/srv-mdl-shared/models"
)

// DynamicDefinition define specific fields for the lookup definition
type DynamicDefinition struct {
	Query     string    `json:"query"`
	Fields    []Param   `json:"fields"`
	Params    []Param   `json:"params"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Param defines the struct of a dynamic filter param
type Param struct {
	Code     string                   `json:"code"`
	DataType string                   `json:"data_type"`
	Label    mdlSharedModels.Translation `json:"label"`
	Type     string                   `json:"field_type,omitempty"`
	Pattern  string                   `json:"pattern,omitempty"`
	Security Security                 `json:"security,omitempty"`
}

// Security defines the fields to set security to a field
type Security struct {
	SchemaCode string `json:"schema_code"`
	FieldCode  string `json:"field_code"`
}

func (d *DynamicDefinition) parse(payload json.RawMessage) error {
	if err := json.Unmarshal(payload, d); err != nil {
		return err
	}

	if err := mdlShared.Validate.Struct(d); err != nil {
		return err
	}
	return nil
}

func (d *DynamicDefinition) GetValueAndLabel() (string, string) {
	return "code", "label"
}
