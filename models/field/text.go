package field

import (
	"encoding/json"

	mdlShared "github.com/agile-work/srv-mdl-shared"
)

// TextDefinition defines custom attributes for the text type
type TextDefinition struct {
	Display string `json:"display" validate:"required"` // single_line, multi_line, readonly, enter_once, url
}

func (d *TextDefinition) load(payload json.RawMessage) error {
	if err := json.Unmarshal(payload, d); err != nil {
		return err
	}

	if err := mdlShared.Validate.Struct(d); err != nil {
		return err
	}
	return nil
}
