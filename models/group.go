package models

import "time"

// Group defines the struct of this object
type Group struct {
	ID            string    `json:"id" sql:"id" pk:"true"`
	Code          string    `json:"code" sql:"code"`
	Name          string    `json:"name" table:"core_translations" alias:"core_translations_name" sql:"value" on:"core_translations_name.structure_id = core_groups.id and core_translations_name.structure_field = 'name'"`
	Description   string    `json:"description" table:"core_translations" alias:"core_translations_description" sql:"value" on:"core_translations_description.structure_id = core_groups.id and core_translations_description.structure_field = 'description'"`
	Active        bool      `json:"active" sql:"active"`
	CreatedBy     string    `json:"created_by" sql:"created_by"`
	CreatedByUser *User     `json:"created_by_user" table:"core_users" alias:"created_by_user" on:"created_by_user.id = core_groups.created_by"`
	CreatedAt     time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy     string    `json:"updated_by" sql:"updated_by"`
	UpdatedByUser *User     `json:"updated_by_user" table:"core_users" alias:"updated_by_user" on:"updated_by_user.id = core_groups.updated_by"`
	UpdatedAt     time.Time `json:"updated_at" sql:"updated_at"`
}

// Permission defines the struct of this object
type Permission struct {
	ID             string    `json:"id" sql:"id" pk:"true"`
	GroupID        string    `json:"group_id" sql:"group_id" fk:"true"`
	StructureType  string    `json:"structure_type" sql:"structure_type"`
	StructureID    string    `json:"structure_id" sql:"structure_id" fk:"true"`
	PermissionType int       `json:"permission_type" sql:"permission_type"`
	ConditionQuery string    `json:"condition_query" sql:"condition_query"`
	CreatedBy      string    `json:"created_by" sql:"created_by"`
	CreatedByUser  *User     `json:"created_by_user" table:"core_users" alias:"created_by_user" on:"created_by_user.id = core_grp_permissions.created_by"`
	CreatedAt      time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy      string    `json:"updated_by" sql:"updated_by"`
	UpdatedByUser  *User     `json:"updated_by_user" table:"core_users" alias:"updated_by_user" on:"updated_by_user.id = core_grp_permissions.updated_by"`
	UpdatedAt      time.Time `json:"updated_at" sql:"updated_at"`
}

// ViewUserGroup defines the struct of this object
type ViewUserGroup struct {
	ID            string    `json:"id" sql:"id" pk:"true"`
	UserID        string    `json:"user_id" sql:"user_id" fk:"true"`
	Code          string    `json:"code" sql:"code"`
	Name          string    `json:"name" sql:"name"`
	Description   string    `json:"description" sql:"description"`
	LanguageCode  string    `json:"language_code" sql:"language_code"`
	Active        bool      `json:"active" sql:"active"`
	CreatedBy     string    `json:"created_by" sql:"created_by"`
	CreatedByUser *User     `json:"created_by_user" table:"core_users" alias:"created_by_user" on:"created_by_user.id = core_v_user_groups.created_by"`
	CreatedAt     time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy     string    `json:"updated_by" sql:"updated_by"`
	UpdatedByUser *User     `json:"updated_by_user" table:"core_users" alias:"updated_by_user" on:"updated_by_user.id = core_v_user_groups.updated_by"`
	UpdatedAt     time.Time `json:"updated_at" sql:"updated_at"`
}
