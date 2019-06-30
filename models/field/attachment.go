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

func (a *AttachmentDefinition) load(payload json.RawMessage) error {
	if err := json.Unmarshal(payload, a); err != nil {
		return err
	}

	return nil
}

func (a *AttachmentDefinition) validate() error {
	if err := mdlShared.Validate.Struct(a); err != nil {
		return err
	}
	return nil
}
