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

// NumberScale defines a static dataset to define a custom scale to a number field
type NumberScale struct {
	DatasetCode      string                        `json:"dataset_code" validate:"required"`
	AggregationRates map[string]map[string]float32 `json:"aggr_rates,omitempty"`
}

func (n *NumberDefinition) parse(payload json.RawMessage) error {
	if err := json.Unmarshal(payload, n); err != nil {
		return err
	}

	if n.Display == FieldNumberDisplayMoney {
		n.Scale.DatasetCode = "ds_currencies"
		n.Scale.AggregationRates = nil
	}

	return nil
}

func (n *NumberDefinition) prepare() error {
	if err := mdlShared.Validate.Struct(n); err != nil {
		return err
	}
	if err := dataset.Validate(d.DatasetCode, true, nil, nil); err!= nil {
		return err
	}
	return nil
}
