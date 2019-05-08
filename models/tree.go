package models

import "time"

// Tree defines the struct of this object
type Tree struct {
	ID            string    `json:"id" sql:"id" pk:"true"`
	Code          string    `json:"code" sql:"code"`
	Name          string    `json:"name" table:"core_translations" alias:"core_translations_name" sql:"value" on:"core_translations_name.structure_id = core_trees.id and core_translations_name.structure_field = 'name'"`
	Description   string    `json:"description" table:"core_translations" alias:"core_translations_description" sql:"value" on:"core_translations_description.structure_id = core_trees.id and core_translations_description.structure_field = 'description'"`
	Active        bool      `json:"active" sql:"active"`
	CreatedBy     string    `json:"created_by" sql:"created_by"`
	CreatedByUser *User     `json:"created_by_user" table:"core_users" alias:"created_by_user" on:"created_by_user.id = core_trees.created_by"`
	CreatedAt     time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy     string    `json:"updated_by" sql:"updated_by"`
	UpdatedByUser *User     `json:"updated_by_user" table:"core_users" alias:"updated_by_user" on:"updated_by_user.id = core_trees.updated_by"`
	UpdatedAt     time.Time `json:"updated_at" sql:"updated_at"`
}

// TreeLevel defines the struct of this object
type TreeLevel struct {
	ID            string    `json:"id" sql:"id" pk:"true"`
	Code          string    `json:"code" sql:"code"`
	TreeID        string    `json:"tree_id" sql:"tree_id" fk:"true"`
	Name          string    `json:"name" table:"core_translations" alias:"core_translations_name" sql:"value" on:"core_translations_name.structure_id = core_tre_levels.id and core_translations_name.structure_field = 'name'"`
	Description   string    `json:"description" table:"core_translations" alias:"core_translations_description" sql:"value" on:"core_translations_description.structure_id = core_tre_levels.id and core_translations_description.structure_field = 'description'"`
	CreatedBy     string    `json:"created_by" sql:"created_by"`
	CreatedByUser *User     `json:"created_by_user" table:"core_users" alias:"created_by_user" on:"created_by_user.id = core_tre_levels.created_by"`
	CreatedAt     time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy     string    `json:"updated_by" sql:"updated_by"`
	UpdatedByUser *User     `json:"updated_by_user" table:"core_users" alias:"updated_by_user" on:"updated_by_user.id = core_tre_levels.updated_by"`
	UpdatedAt     time.Time `json:"updated_at" sql:"updated_at"`
}

// TreeUnit defines the struct of this object
type TreeUnit struct {
	ID            string    `json:"id" sql:"id" pk:"true"`
	Code          string    `json:"code" sql:"code"`
	TreeID        string    `json:"tree_id" sql:"tree_id" fk:"true"`
	ParentID      string    `json:"parent_id" sql:"parent_id" fk:"true"`
	Name          string    `json:"name" table:"core_translations" alias:"core_translations_name" sql:"value" on:"core_translations_name.structure_id = core_tre_units.id and core_translations_name.structure_field = 'name'"`
	Description   string    `json:"description" table:"core_translations" alias:"core_translations_description" sql:"value" on:"core_translations_description.structure_id = core_tre_units.id and core_translations_description.structure_field = 'description'"`
	Active        bool      `json:"active" sql:"active"`
	CreatedBy     string    `json:"created_by" sql:"created_by"`
	CreatedByUser *User     `json:"created_by_user" table:"core_users" alias:"created_by_user" on:"created_by_user.id = core_tre_units.created_by"`
	CreatedAt     time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy     string    `json:"updated_by" sql:"updated_by"`
	UpdatedByUser *User     `json:"updated_by_user" table:"core_users" alias:"updated_by_user" on:"updated_by_user.id = core_tre_units.updated_by"`
	UpdatedAt     time.Time `json:"updated_at" sql:"updated_at"`
}
