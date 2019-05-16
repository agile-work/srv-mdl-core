package services

import (
	"fmt"
	"net/http"

	"github.com/agile-work/srv-mdl-shared/db"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/agile-work/srv-mdl-core/models"
	moduleShared "github.com/agile-work/srv-mdl-shared"
	shared "github.com/agile-work/srv-shared"
)

// CreateTree persists the request body creating a new object in the database
func CreateTree(r *http.Request) *moduleShared.Response {
	tree := models.Tree{}

	return db.Create(r, &tree, "CreateTree", shared.TableCoreTrees)
}

// LoadAllTrees return all instances from the object
func LoadAllTrees(r *http.Request) *moduleShared.Response {
	trees := []models.Tree{}

	return db.Load(r, &trees, "LoadAllTrees", shared.TableCoreTrees, nil)
}

// LoadTree return only one object from the database
func LoadTree(r *http.Request) *moduleShared.Response {
	tree := models.Tree{}
	treeID := chi.URLParam(r, "tree_id")
	treeIDColumn := fmt.Sprintf("%s.id", shared.TableCoreTrees)
	condition := builder.Equal(treeIDColumn, treeID)

	return db.Load(r, &tree, "LoadTree", shared.TableCoreTrees, condition)
}

// UpdateTree updates object data in the database
func UpdateTree(r *http.Request) *moduleShared.Response {
	treeID := chi.URLParam(r, "tree_id")
	treeIDColumn := fmt.Sprintf("%s.id", shared.TableCoreTrees)
	condition := builder.Equal(treeIDColumn, treeID)
	tree := models.Tree{
		ID: treeID,
	}

	return db.Update(r, &tree, "UpdateTree", shared.TableCoreTrees, condition)
}

// DeleteTree deletes object from the database
func DeleteTree(r *http.Request) *moduleShared.Response {
	treeID := chi.URLParam(r, "tree_id")
	treeIDColumn := fmt.Sprintf("%s.id", shared.TableCoreTrees)
	condition := builder.Equal(treeIDColumn, treeID)

	return db.Remove(r, "DeleteTree", shared.TableCoreTrees, condition)
}

// CreateTreeLevel persists the request body creating a new object in the database
func CreateTreeLevel(r *http.Request) *moduleShared.Response {
	treeID := chi.URLParam(r, "tree_id")
	treeLevel := models.TreeLevel{
		TreeID: treeID,
	}

	return db.Create(r, &treeLevel, "CreateTreeLevel", shared.TableCoreTreLevels)
}

// LoadAllTreeLevels return all instances from the object
func LoadAllTreeLevels(r *http.Request) *moduleShared.Response {
	treeLevels := []models.TreeLevel{}
	treeID := chi.URLParam(r, "tree_id")
	treeIDColumn := fmt.Sprintf("%s.tree_id", shared.TableCoreTreLevels)
	condition := builder.Equal(treeIDColumn, treeID)

	return db.Load(r, &treeLevels, "LoadAllTreeLevels", shared.TableCoreTreLevels, condition)
}

// LoadTreeLevel return only one object from the database
func LoadTreeLevel(r *http.Request) *moduleShared.Response {
	treeLevel := models.TreeLevel{}
	treeLevelID := chi.URLParam(r, "tree_level_id")
	treeLevelIDColumn := fmt.Sprintf("%s.id", shared.TableCoreTreLevels)
	condition := builder.Equal(treeLevelIDColumn, treeLevelID)

	return db.Load(r, &treeLevel, "LoadTreeLevel", shared.TableCoreTreLevels, condition)
}

// UpdateTreeLevel updates object data in the database
func UpdateTreeLevel(r *http.Request) *moduleShared.Response {
	treeLevelID := chi.URLParam(r, "tree_level_id")
	treeLevelIDColumn := fmt.Sprintf("%s.id", shared.TableCoreTreLevels)
	condition := builder.Equal(treeLevelIDColumn, treeLevelID)
	treeLevel := models.TreeLevel{
		ID: treeLevelID,
	}

	return db.Update(r, &treeLevel, "UpdateTreeLevel", shared.TableCoreTreLevels, condition)
}

// DeleteTreeLevel deletes object from the database
func DeleteTreeLevel(r *http.Request) *moduleShared.Response {
	treeLevelID := chi.URLParam(r, "tree_level_id")
	treeLevelIDColumn := fmt.Sprintf("%s.id", shared.TableCoreTreLevels)
	condition := builder.Equal(treeLevelIDColumn, treeLevelID)

	return db.Remove(r, "DeleteTreeLevel", shared.TableCoreTreLevels, condition)
}

// CreateTreeUnit persists the request body creating a new object in the database
func CreateTreeUnit(r *http.Request) *moduleShared.Response {
	treeID := chi.URLParam(r, "tree_id")
	treeUnit := models.TreeUnit{
		TreeID: treeID,
	}

	return db.Create(r, &treeUnit, "CreateTreeUnit", shared.TableCoreTreUnits)
}

// LoadAllTreeUnits return all instances from the object
func LoadAllTreeUnits(r *http.Request) *moduleShared.Response {
	treeUnits := []models.TreeUnit{}
	treeID := chi.URLParam(r, "tree_id")
	treeIDColumn := fmt.Sprintf("%s.tree_id", shared.TableCoreTreUnits)
	condition := builder.Equal(treeIDColumn, treeID)

	return db.Load(r, &treeUnits, "LoadAllTreeUnits", shared.TableCoreTreUnits, condition)
}

// LoadTreeUnit return only one object from the database
func LoadTreeUnit(r *http.Request) *moduleShared.Response {
	treeUnit := models.TreeUnit{}
	treeUnitID := chi.URLParam(r, "tree_unit_id")
	treeUnitIDColumn := fmt.Sprintf("%s.id", shared.TableCoreTreUnits)
	condition := builder.Equal(treeUnitIDColumn, treeUnitID)

	return db.Load(r, &treeUnit, "LoadTreeUnit", shared.TableCoreTreUnits, condition)
}

// UpdateTreeUnit updates object data in the database
func UpdateTreeUnit(r *http.Request) *moduleShared.Response {
	treeUnitID := chi.URLParam(r, "tree_unit_id")
	treeUnitIDColumn := fmt.Sprintf("%s.id", shared.TableCoreTreUnits)
	condition := builder.Equal(treeUnitIDColumn, treeUnitID)
	treeUnit := models.TreeUnit{
		ID: treeUnitID,
	}

	return db.Update(r, &treeUnit, "UpdateTreeUnit", shared.TableCoreTreUnits, condition)
}

// DeleteTreeUnit deletes object from the database
func DeleteTreeUnit(r *http.Request) *moduleShared.Response {
	treeUnitID := chi.URLParam(r, "tree_unit_id")
	treeUnitIDColumn := fmt.Sprintf("%s.id", shared.TableCoreTreUnits)
	condition := builder.Equal(treeUnitIDColumn, treeUnitID)

	return db.Remove(r, "DeleteTreeUnit", shared.TableCoreTreUnits, condition)
}
