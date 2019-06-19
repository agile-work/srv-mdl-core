package field

import (
	"encoding/json"

	mdlShared "github.com/agile-work/srv-mdl-shared"
	mdlSharedModels "github.com/agile-work/srv-mdl-shared/models"
)

// DateDefinition defines custom attributes for the date type
type DateDefinition struct {
	Display string                   `json:"display" validate:"required"` // date, date_time
	Format  mdlSharedModels.Translation `json:"format"`
}

func (d *DateDefinition) load(payload json.RawMessage) error {
	if err := json.Unmarshal(payload, d); err != nil {
		return err
	}

	if err := mdlShared.Validate.Struct(d); err != nil {
		return err
	}
	return nil
}
