package field

import (
	"encoding/json"

	mdlShared "github.com/agile-work/srv-mdl-shared"
)

// TextDefinition defines custom attributes for the text type
type TextDefinition struct {
	Display string `json:"display" validate:"required"` // single_line, multi_line, readonly, enter_once, url
}

func (t *TextDefinition) load(payload json.RawMessage) error {
	if err := json.Unmarshal(payload, t); err != nil {
		return err
	}

	return nil
}

func (t *TextDefinition) validate() error {
	if err := mdlShared.Validate.Struct(t); err != nil {
		return err
	}
	return nil
}
