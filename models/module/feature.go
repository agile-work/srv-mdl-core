package module

// Feature defines each feature for this module
type Feature struct {
	Permissions []string `json:"permissions" validate:"required"`
	Type        string   `json:"type" validate:"required"`
}
