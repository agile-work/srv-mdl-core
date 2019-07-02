package field

import (
	"encoding/json"

	"github.com/agile-work/srv-mdl-core/models/dataset"
	mdlShared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-shared/constants"
)

// NumberDefinition defines custom attributes for the number type
type NumberDefinition struct {
	Display  string       `json:"display" validate:"required"` // percentage, number, money
	Decimals int          `json:"decimals"`
	Scale    *NumberScale `json:"scale,omitempty"`
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

	if n.Display == constants.FieldNumberDisplayMoney {
		n.Scale.DatasetCode = "ds_currencies"
		n.Scale.AggregationRates = nil
	}

	return nil
}

func (n *NumberDefinition) prepare() error {
	if err := mdlShared.Validate.Struct(n); err != nil {
		return err
	}
	if n.Scale == nil {
		return nil
	}
	if err := dataset.Validate(n.Scale.DatasetCode, true, nil, nil); err != nil {
		return err
	}
	return nil
}
