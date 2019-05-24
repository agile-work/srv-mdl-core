package services

import (
	"fmt"
	"net/http"

	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/agile-work/srv-mdl-core/models"
	moduleShared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-mdl-shared/db"
	shared "github.com/agile-work/srv-shared"
)

// CreateContainerStructure persists the request body creating a new object in the database
func CreateContainerStructure(r *http.Request) *moduleShared.Response {
	schemaID := chi.URLParam(r, "schema_id")
	pageID := chi.URLParam(r, "page_id")
	containerID := chi.URLParam(r, "container_id")
	containerType := chi.URLParam(r, "type")
	containerStructure := models.ContainerStructure{
		SchemaID:      schemaID,
		PageID:        pageID,
		ContainerID:   containerID,
		ContainerType: containerType,
	}

	return db.Create(r, &containerStructure, "CreateContainerStructure", shared.TableCoreSchemaPageContainerStructures)
}

// LoadAllContainerStructures return all instances from the object
func LoadAllContainerStructures(r *http.Request) *moduleShared.Response {
	containerStructures := []models.ContainerStructure{}
	containerID := chi.URLParam(r, "container_id")
	containerIDColumn := fmt.Sprintf("%s.container_id", shared.TableCoreSchemaPageContainerStructures)
	condition := builder.Equal(containerIDColumn, containerID)

	return db.Load(r, &containerStructures, "LoadAllContainerStructures", shared.TableCoreSchemaPageContainerStructures, condition)
}

// LoadContainerStructure return only one object from the database
func LoadContainerStructure(r *http.Request) *moduleShared.Response {
	containerStructure := models.ContainerStructure{}
	containerStructureID := chi.URLParam(r, "container_structure_id")
	containerStructureIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchemaPageContainerStructures)
	condition := builder.Equal(containerStructureIDColumn, containerStructureID)

	return db.Load(r, &containerStructure, "LoadContainerStructure", shared.TableCoreSchemaPageContainerStructures, condition)
}

// UpdateContainerStructure updates object data in the database
func UpdateContainerStructure(r *http.Request) *moduleShared.Response {
	containerStructureID := chi.URLParam(r, "container_structure_id")
	containerStructureIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchemaPageContainerStructures)
	condition := builder.Equal(containerStructureIDColumn, containerStructureID)
	containerStructure := models.ContainerStructure{
		ID: containerStructureID,
	}

	return db.Update(r, &containerStructure, "UpdateContainerStructure", shared.TableCoreSchemaPageContainerStructures, condition)
}

// DeleteContainerStructure deletes object from the database
func DeleteContainerStructure(r *http.Request) *moduleShared.Response {
	containerStructureID := chi.URLParam(r, "container_structure_id")
	containerStructureIDColumn := fmt.Sprintf("%s.id", shared.TableCoreSchemaPageContainerStructures)
	condition := builder.Equal(containerStructureIDColumn, containerStructureID)

	return db.Remove(r, "DeleteContainerStructure", shared.TableCoreSchemaPageContainerStructures, condition)
}
