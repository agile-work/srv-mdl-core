package field

import (
	"encoding/json"

	mdlShared "github.com/agile-work/srv-mdl-shared"
)

// NumberDefinition defines custom attributes for the number type
type NumberDefinition struct {
	Display  string      `json:"display" validate:"required"` // percentage, number, money
	Decimals int         `json:"decimals"`
	Scale    NumberScale `json:"scale,omitempty"`
}

// NumberScale defines a lookup to define a custom scale to a number field
type NumberScale struct {
	LookupCode       string                 `json:"lookup_code" validate:"required"`
	LookupLabel      string                 `json:"lookup_label"`
	LookupValue      string                 `json:"lookup_value"`
	LookupParams     []lookupParam          `json:"lookup_params,omitempty"`
	AggregationRates map[string]interface{} `json:"aggr_rates"`
}

func (n *NumberDefinition) load(payload json.RawMessage) error {
	if err := json.Unmarshal(payload, n); err != nil {
		return err
	}

	return nil
}

func (n *NumberDefinition) validate() error {
	if err := mdlShared.Validate.Struct(n); err != nil {
		return err
	}
	return nil
}
