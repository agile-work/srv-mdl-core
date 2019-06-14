package field

import (
	"encoding/json"

	mdlShared "github.com/agile-work/srv-mdl-shared"
)

// AttachmentDefinition defines custom attributes for the attachment type
// TODO: update to deal with module "document management"
type AttachmentDefinition struct {
	Display           string   `json:"display" validate:"required"` // single, multiple
	FilesMaxNum       int      `json:"files_max_number"`
	FileMaxSize       int      `json:"file_max_size"`
	FileTypeWhiteList []string `json:"file_type_white_list"`
}

func (d *AttachmentDefinition) load(payload json.RawMessage) error {
	if err := json.Unmarshal(payload, d); err != nil {
		return err
	}

	if err := mdlShared.Validate.Struct(d); err != nil {
		return err
	}
	return nil
}
