package models

import "time"

// ContainerStructure defines the struct of this object
type ContainerStructure struct {
	ID             string    `json:"id" sql:"id" pk:"true"`
	SchemaID       string    `json:"schema_id" sql:"schema_id" fk:"true"`
	PageID         string    `json:"page_id" sql:"page_id" fk:"true"`
	ContainerID    string    `json:"container_id" sql:"container_id" fk:"true"`
	ContainerType  string    `json:"container_type" sql:"container_type"`
	StructureID    string    `json:"structure_id" sql:"structure_id" fk:"true"`
	StructureType  string    `json:"structure_type" sql:"structure_type"`
	PositionRow    int       `json:"position_row" sql:"position_row"`
	PositionColumn int       `json:"position_column" sql:"position_column"`
	Width          int       `json:"width" sql:"width"`
	Height         int       `json:"height" sql:"height"`
	CreatedBy      string    `json:"created_by" sql:"created_by"`
	CreatedByUser  *User     `json:"created_by_user" table:"core_users" alias:"created_by_user" on:"created_by_user.id = core_sch_pag_cnt_structures.created_by"`
	CreatedAt      time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy      string    `json:"updated_by" sql:"updated_by"`
	UpdatedByUser  *User     `json:"updated_by_user" table:"core_users" alias:"updated_by_user" on:"updated_by_user.id = core_sch_pag_cnt_structures.updated_by"`
	UpdatedAt      time.Time `json:"updated_at" sql:"updated_at"`
}
