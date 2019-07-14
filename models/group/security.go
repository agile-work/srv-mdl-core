package group

import "time"

// SecurityTree struct
type SecurityTree struct {
	Code      string     `json:"code,omitempty"`
	Unit      string     `json:"unit,omitempty"`
	Scope     string     `json:"scope,omitempty"`
	CreatedBy string     `json:"created_by" sql:"created_by"`
	CreatedAt *time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy string     `json:"updated_by" sql:"updated_by"`
	UpdatedAt *time.Time `json:"updated_at" sql:"updated_at"`
}

// SecurityAudit struct
type SecurityAudit struct {
	Operator  string    `json:"operator,omitempty"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

// SecurityUser struct
type SecurityUser struct {
	IncludeResources map[string]SecurityAudit `json:"include_resources,omitempty"`
	ExcludeResources map[string]SecurityAudit `json:"exclude_resources,omitempty"`
	Wildcard         string                   `json:"wildcard,omitempty"`
	CreatedBy        string                   `json:"created_by"`
	CreatedAt        *time.Time               `json:"created_at"`
	UpdatedBy        string                   `json:"updated_by" sql:"updated_by"`
	UpdatedAt        *time.Time               `json:"updated_at" sql:"updated_at"`
}

// SecurityField struct
type SecurityField struct {
	EditAll        bool     `json:"edit_all,omitempty"`
	EditExceptions []string `json:"edit_exceptions,omitempty"`
	ViewAll        bool     `json:"view_all,omitempty"`
	ViewExceptions []string `json:"view_exceptions,omitempty"`
}

// SecurityInstance struct
type SecurityInstance struct {
	ViewAll bool `json:"view_all"`
	Create  bool `json:"create"`
	Delete  bool `json:"delete"`
}

// SecuritySchema struct
type SecuritySchema struct {
	Instances SecurityInstance          `json:"instances,omitempty"`
	Fields    SecurityField             `json:"fields,omitempty"`
	Modules   map[string]SecurityModule `json:"modules,omitempty"`
}

// SecurityModule struct
type SecurityModule struct {
	Instances SecurityInstance                  `json:"instances,omitempty"`
	Fields    SecurityField                     `json:"fields,omitempty"`
	Features  map[string]map[string]interface{} `json:"features,omitempty"`
}

// SecurityProcesse struct
type SecurityProcesse struct {
	View   bool `json:"view"`
	Start  bool `json:"start"`
	Cancel bool `json:"cancel"`
	Delete bool `json:"delete"`
}

// SecurityWidget struct
type SecurityWidget struct {
	View      bool `json:"view"`
	Customize bool `json:"customize"`
}

// Permission struct
type Permission struct {
	Widgets   map[string]SecurityWidget   `json:"widgets,omitempty"`
	Processes map[string]SecurityProcesse `json:"processes,omitempty"`
	Schemas   map[string]SecuritySchema   `json:"schemas,omitempty"`
}
