package models

import (
	"time"
)

// Group defines the struct of this object
type Group struct {
	ID                      string       `json:"id" sql:"id" pk:"true"`
	Code                    string       `json:"code" sql:"code"`
	Name                    string       `json:"name" table:"core_translations" alias:"core_translations_name" sql:"value" on:"core_translations_name.structure_id = core_groups.id and core_translations_name.structure_field = 'name'"`
	Description             string       `json:"description" table:"core_translations" alias:"core_translations_description" sql:"value" on:"core_translations_description.structure_id = core_groups.id and core_translations_description.structure_field = 'description'"`
	TreeUnitID              *string      `json:"tree_unit_id" sql:"tree_unit_id"`
	TreeUnitPermissionScope *string      `json:"tree_unit_permission_scope" sql:"tree_unit_permission_scope"`
	Permissions             []Permission `json:"permissions" sql:"permissions" field:"jsonb"`
	Users                   []GroupUser  `json:"users" sql:"users" field:"jsonb"`
	Active                  bool         `json:"active" sql:"active"`
	CreatedBy               string       `json:"created_by" sql:"created_by"`
	CreatedAt               time.Time    `json:"created_at" sql:"created_at"`
	UpdatedBy               string       `json:"updated_by" sql:"updated_by"`
	UpdatedAt               time.Time    `json:"updated_at" sql:"updated_at"`
}

// GroupUser defines the struct to the jsonb field in group
type GroupUser struct {
	ID        string    `json:"id"`
	CreatedBy string    `json:"created_by,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Permission defines the struct of this object
type Permission struct {
	ID             string    `json:"id,omitempty" sql:"id"`
	ParentID       string    `json:"parent_id,omitempty" sql:"parent_id"`
	StructureID    string    `json:"structure_id,omitempty" sql:"structure_id"`
	StructureType  string    `json:"structure_type,omitempty" sql:"structure_type"`
	StructureName  string    `json:"structure_name,omitempty" sql:"structure_name"`
	PermissionType int       `json:"permission_type,omitempty" sql:"permission_type"`
	ConditionQuery string    `json:"condition_query,omitempty"`
	LanguageCode   string    `json:"language_code,omitempty" sql:"language_code"`
	CreatedBy      string    `json:"created_by,omitempty" sql:"created_by"`
	CreatedAt      time.Time `json:"created_at,omitempty" sql:"created_at"`
}

// ViewUserGroup defines the struct of this object
type ViewUserGroup struct {
	ID           string    `json:"id" sql:"id" pk:"true"`
	UserID       string    `json:"user_id" sql:"user_id" fk:"true"`
	Code         string    `json:"code" sql:"code"`
	Name         string    `json:"name" sql:"name"`
	Description  string    `json:"description" sql:"description"`
	LanguageCode string    `json:"language_code" sql:"language_code"`
	Active       bool      `json:"active" sql:"active"`
	CreatedBy    string    `json:"created_by" sql:"created_by"`
	CreatedAt    time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy    string    `json:"updated_by" sql:"updated_by"`
	UpdatedAt    time.Time `json:"updated_at" sql:"updated_at"`
}
