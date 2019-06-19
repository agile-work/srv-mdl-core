package lookup

import (
	"encoding/json"
	"time"

	mdlShared "github.com/agile-work/srv-mdl-shared"
	mdlSharedModels "github.com/agile-work/srv-mdl-shared/models"
)

// StaticDefinition define specific fields for the lookup definition
type StaticDefinition struct {
	Options   map[string]Option `json:"options,omitempty"`
	OrderType string            `json:"order_type,omitempty"`
	Order     []string          `json:"order,omitempty"`
}

// Option defines the struct of a static option
type Option struct {
	Code      string                   `json:"code"`
	Label     mdlSharedModels.Translation `json:"label,omitempty"`
	Active    bool                     `json:"active"`
	CreatedBy string                   `json:"created_by"`
	CreatedAt time.Time                `json:"created_at"`
	UpdatedBy string                   `json:"updated_by"`
	UpdatedAt time.Time                `json:"updated_at"`
}

func (d *StaticDefinition) parse(payload json.RawMessage) error {
	if err := json.Unmarshal(payload, d); err != nil {
		return err
	}

	if err := mdlShared.Validate.Struct(d); err != nil {
		return err
	}
	return nil
}

func (d *StaticDefinition) GetValueAndLabel() (string, string) {
	return "code", "label"
}
