package field

import (
	"encoding/json"

	mdlShared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-mdl-shared/models/translation"
)

// DateDefinition defines custom attributes for the date type
type DateDefinition struct {
	Display string                  `json:"display" validate:"required"` // date, date_time
	Format  translation.Translation `json:"format"`
}

func (d *DateDefinition) parse(payload json.RawMessage) error {
	if err := json.Unmarshal(payload, d); err != nil {
		return err
	}

	if len(d.Format.Language) <= 0 {
		d.Format.Language[translation.SystemDefaultLanguageCode] = "DD-MM-YYYY HH:MM:SS"
	}

	return nil
}

func (d *DateDefinition) prepare() error {
	if err := mdlShared.Validate.Struct(d); err != nil {
		return err
	}
	return nil
}
