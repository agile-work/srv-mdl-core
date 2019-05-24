package models

import (
	"encoding/json"
	"time"
)

// Instance defines the struct of this object
type Instance struct {
	ID        string          `json:"id" sql:"id" pk:"true"`
	Data      json.RawMessage `json:"data" sql:"data" field:"jsonb"`
	CreatedBy string          `json:"created_by" sql:"created_by"`
	CreatedAt time.Time       `json:"created_at" sql:"created_at"`
	UpdatedBy string          `json:"updated_by" sql:"updated_by"`
	UpdatedAt time.Time       `json:"updated_at" sql:"updated_at"`
}

// EntityInstancePermission defines the struct of this object
type EntityInstancePermission struct {
	ID           string               `json:"id" sql:"id" pk:"true"`
	UserID       string               `json:"user_id" sql:"user_id"`
	SourceType   string               `json:"source_type" sql:"source_type"`
	SourceID     string               `json:"source_id" sql:"source_id"`
	InstanceID   string               `json:"instance_id" sql:"instance_id"`
	InstanceType string               `json:"instance_type" sql:"instance_type"`
	Permissions  []InstancePermission `json:"permissions" sql:"permissions" field:"jsonb"`
	CreatedBy    string               `json:"created_by" sql:"created_by"`
	CreatedAt    time.Time            `json:"created_at" sql:"created_at"`
	UpdatedBy    string               `json:"updated_by" sql:"updated_by"`
	UpdatedAt    time.Time            `json:"updated_at" sql:"updated_at"`
}

// InstancePermission defines the struct of this object
type InstancePermission struct {
	ID             string `json:"id" sql:"id" pk:"true"`
	StructureType  string `json:"structure_type" sql:"structure_type"`
	StructureID    string `json:"structure_id" sql:"structure_id"`
	PermissionType int    `json:"permission_type" sql:"permission_type"`
	ConditionQuery string `json:"condition_query" sql:"condition_query"`
}
