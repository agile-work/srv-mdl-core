package models

import (
	"time"
)

// Tree defines the struct of this object
type Tree struct {
	ID          string    `json:"id" sql:"id" pk:"true"`
	Code        string    `json:"code" sql:"code"`
	Name        string    `json:"name" table:"core_translations" alias:"core_translations_name" sql:"value" on:"core_translations_name.structure_id = core_trees.id and core_translations_name.structure_field = 'name'"`
	Description string    `json:"description" table:"core_translations" alias:"core_translations_description" sql:"value" on:"core_translations_description.structure_id = core_trees.id and core_translations_description.structure_field = 'description'"`
	Active      bool      `json:"active" sql:"active"`
	CreatedBy   string    `json:"created_by" sql:"created_by"`
	CreatedAt   time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy   string    `json:"updated_by" sql:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at" sql:"updated_at"`
}

// TreeLevel defines the struct of this object
type TreeLevel struct {
	ID          string    `json:"id" sql:"id" pk:"true"`
	Code        string    `json:"code" sql:"code"`
	TreeCode    string    `json:"tree_code" sql:"tree_code" fk:"true"`
	Name        string    `json:"name" table:"core_translations" alias:"core_translations_name" sql:"value" on:"core_translations_name.structure_id = core_tree_levels.id and core_translations_name.structure_field = 'name'"`
	Description string    `json:"description" table:"core_translations" alias:"core_translations_description" sql:"value" on:"core_translations_description.structure_id = core_tree_levels.id and core_translations_description.structure_field = 'description'"`
	CreatedBy   string    `json:"created_by" sql:"created_by"`
	CreatedAt   time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy   string    `json:"updated_by" sql:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at" sql:"updated_at"`
}

// TreeUnit defines the struct of this object
type TreeUnit struct {
	ID              string       `json:"id" sql:"id" pk:"true"`
	Code            string       `json:"code" sql:"code"`
	TreeCode        string       `json:"tree_code" sql:"tree_code" fk:"true"`
	Path            string       `json:"path" sql:"path"`
	Name            string       `json:"name" table:"core_translations" alias:"core_translations_name" sql:"value" on:"core_translations_name.structure_id = core_tree_units.id and core_translations_name.structure_field = 'name'"`
	Description     string       `json:"description" table:"core_translations" alias:"core_translations_description" sql:"value" on:"core_translations_description.structure_id = core_tree_units.id and core_translations_description.structure_field = 'description'"`
	PermissionScope string       `json:"permission_scope" sql:"permission_scope"`
	Permissions     []Permission `json:"permissions" sql:"permissions" field:"jsonb"`
	Active          bool         `json:"active" sql:"active"`
	CreatedBy       string       `json:"created_by" sql:"created_by"`
	CreatedAt       time.Time    `json:"created_at" sql:"created_at"`
	UpdatedBy       string       `json:"updated_by" sql:"updated_by"`
	UpdatedAt       time.Time    `json:"updated_at" sql:"updated_at"`
}
