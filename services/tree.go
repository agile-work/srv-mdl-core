package services

import (
	"fmt"
	"net/http"

	"github.com/agile-work/srv-mdl-shared/db"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	sql "github.com/agile-work/srv-shared/sql-builder/db"
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
	treeCode := chi.URLParam(r, "tree_code")
	treeCodeColumn := fmt.Sprintf("%s.code", shared.TableCoreTrees)
	condition := builder.Equal(treeCodeColumn, treeCode)

	return db.Load(r, &tree, "LoadTree", shared.TableCoreTrees, condition)
}

// UpdateTree updates object data in the database
func UpdateTree(r *http.Request) *moduleShared.Response {
	treeCode := chi.URLParam(r, "tree_code")
	treeCodeColumn := fmt.Sprintf("%s.code", shared.TableCoreTrees)
	condition := builder.Equal(treeCodeColumn, treeCode)
	tree := models.Tree{
		ID: treeCode,
	}

	return db.Update(r, &tree, "UpdateTree", shared.TableCoreTrees, condition)
}

// DeleteTree deletes object from the database
func DeleteTree(r *http.Request) *moduleShared.Response {
	treeCode := chi.URLParam(r, "tree_code")
	treeCodeColumn := fmt.Sprintf("%s.code", shared.TableCoreTrees)
	condition := builder.Equal(treeCodeColumn, treeCode)

	return db.Remove(r, "DeleteTree", shared.TableCoreTrees, condition)
}

// CreateTreeLevel persists the request body creating a new object in the database
func CreateTreeLevel(r *http.Request) *moduleShared.Response {
	treeCode := chi.URLParam(r, "tree_code")
	treeLevel := models.TreeLevel{
		TreeCode: treeCode,
	}

	return db.Create(r, &treeLevel, "CreateTreeLevel", shared.TableCoreTreeLevels)
}

// LoadAllTreeLevels return all instances from the object
func LoadAllTreeLevels(r *http.Request) *moduleShared.Response {
	treeLevels := []models.TreeLevel{}
	treeCode := chi.URLParam(r, "tree_code")
	treeCodeColumn := fmt.Sprintf("%s.tree_code", shared.TableCoreTreeLevels)
	condition := builder.Equal(treeCodeColumn, treeCode)

	return db.Load(r, &treeLevels, "LoadAllTreeLevels", shared.TableCoreTreeLevels, condition)
}

// LoadTreeLevel return only one object from the database
func LoadTreeLevel(r *http.Request) *moduleShared.Response {
	treeLevel := models.TreeLevel{}
	treeLevelCode := chi.URLParam(r, "tree_level_code")
	treeLevelCodeColumn := fmt.Sprintf("%s.code", shared.TableCoreTreeLevels)
	condition := builder.Equal(treeLevelCodeColumn, treeLevelCode)

	return db.Load(r, &treeLevel, "LoadTreeLevel", shared.TableCoreTreeLevels, condition)
}

// UpdateTreeLevel updates object data in the database
func UpdateTreeLevel(r *http.Request) *moduleShared.Response {
	treeLevelCode := chi.URLParam(r, "tree_level_code")
	treeLevelCodeColumn := fmt.Sprintf("%s.code", shared.TableCoreTreeLevels)
	condition := builder.Equal(treeLevelCodeColumn, treeLevelCode)
	treeLevel := models.TreeLevel{
		ID: treeLevelCode,
	}

	return db.Update(r, &treeLevel, "UpdateTreeLevel", shared.TableCoreTreeLevels, condition)
}

// DeleteTreeLevel deletes object from the database
func DeleteTreeLevel(r *http.Request) *moduleShared.Response {
	treeLevelCode := chi.URLParam(r, "tree_level_code")
	treeLevelCodeColumn := fmt.Sprintf("%s.code", shared.TableCoreTreeLevels)
	condition := builder.Equal(treeLevelCodeColumn, treeLevelCode)

	return db.Remove(r, "DeleteTreeLevel", shared.TableCoreTreeLevels, condition)
}

// CreateTreeUnit persists the request body creating a new object in the database
func CreateTreeUnit(r *http.Request) *moduleShared.Response {
	treeCode := chi.URLParam(r, "tree_code")
	treeUnit := models.TreeUnit{
		TreeCode: treeCode,
	}

	return db.Create(r, &treeUnit, "CreateTreeUnit", shared.TableCoreTreeUnits)
}

// LoadAllTreeUnits return all instances from the object
func LoadAllTreeUnits(r *http.Request) *moduleShared.Response {
	treeUnits := []models.TreeUnit{}
	treeCode := chi.URLParam(r, "tree_code")
	treeCodeColumn := fmt.Sprintf("%s.tree_code", shared.TableCoreTreeUnits)
	condition := builder.Equal(treeCodeColumn, treeCode)

	return db.Load(r, &treeUnits, "LoadAllTreeUnits", shared.TableCoreTreeUnits, condition)
}

// LoadTreeUnit return only one object from the database
func LoadTreeUnit(r *http.Request) *moduleShared.Response {
	treeUnit := models.TreeUnit{}
	treeUnitCode := chi.URLParam(r, "tree_unit_code")
	treeUnitCodeColumn := fmt.Sprintf("%s.code", shared.TableCoreTreeUnits)
	condition := builder.Equal(treeUnitCodeColumn, treeUnitCode)

	return db.Load(r, &treeUnit, "LoadTreeUnit", shared.TableCoreTreeUnits, condition)
}

// UpdateTreeUnit updates object data in the database
func UpdateTreeUnit(r *http.Request) *moduleShared.Response {
	treeUnitCode := chi.URLParam(r, "tree_unit_code")
	treeUnitCodeColumn := fmt.Sprintf("%s.code", shared.TableCoreTreeUnits)
	condition := builder.Equal(treeUnitCodeColumn, treeUnitCode)
	treeUnit := models.TreeUnit{
		ID: treeUnitCode,
	}

	return db.Update(r, &treeUnit, "UpdateTreeUnit", shared.TableCoreTreeUnits, condition)
}

// DeleteTreeUnit deletes object from the database
func DeleteTreeUnit(r *http.Request) *moduleShared.Response {
	treeUnitCode := chi.URLParam(r, "tree_unit_code")
	treeUnitCodeColumn := fmt.Sprintf("%s.code", shared.TableCoreTreeUnits)
	condition := builder.Equal(treeUnitCodeColumn, treeUnitCode)

	return db.Remove(r, "DeleteTreeUnit", shared.TableCoreTreeUnits, condition)
}

// InsertTreeUnitPermission persists the request body creating a new object in the database
func InsertTreeUnitPermission(r *http.Request) *moduleShared.Response {
	treeUnitCode := chi.URLParam(r, "tree_unit_code")
	permission := models.Permission{}

	response := db.GetResponse(r, &permission, "InsertPermission")
	if response.Code != http.StatusOK {
		return response
	}
	permission.ID = sql.UUID()

	idColumn := fmt.Sprintf("%s.id", shared.TableCoreTreeUnits)
	sql.InsertStructToJSON("permissions", shared.TableCoreTreeUnits, &permission, builder.Equal(idColumn, treeUnitCode))
	response.Data = permission
	return response
}

// RemoveTreeUnitPermission deletes object from the database
func RemoveTreeUnitPermission(r *http.Request) *moduleShared.Response {
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}
	treeUnitCode := chi.URLParam(r, "tree_unit_code")
	permissionID := chi.URLParam(r, "permission_id")

	err := sql.DeleteStructFromJSON(permissionID, treeUnitCode, "permissions", shared.TableCoreTreeUnits)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorParsingRequest, "RemovePermissionFromGroup", err.Error()))

		return response
	}

	return response
}
