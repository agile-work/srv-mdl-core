package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/agile-work/srv-mdl-shared/db"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	sql "github.com/agile-work/srv-shared/sql-builder/db"
	"github.com/go-chi/chi"

	"github.com/agile-work/srv-mdl-core/models"
	moduleShared "github.com/agile-work/srv-mdl-shared"
	shared "github.com/agile-work/srv-shared"
)

// CreateGroup persists the request body creating a new object in the database
func CreateGroup(r *http.Request) *moduleShared.Response {
	group := models.Group{}

	return db.Create(r, &group, "CreateGroup", shared.TableCoreGroups)
}

// LoadAllGroups return all instances from the object
func LoadAllGroups(r *http.Request) *moduleShared.Response {
	groups := []models.Group{}

	return db.Load(r, &groups, "LoadAllGroups", shared.TableCoreGroups, nil)
}

// LoadGroup return only one object from the database
func LoadGroup(r *http.Request) *moduleShared.Response {
	group := models.Group{}
	groupID := chi.URLParam(r, "group_id")
	groupIDColumn := fmt.Sprintf("%s.id", shared.TableCoreGroups)
	condition := builder.Equal(groupIDColumn, groupID)

	return db.Load(r, &group, "LoadGroup", shared.TableCoreGroups, condition)
}

// UpdateGroup updates object data in the database
func UpdateGroup(r *http.Request) *moduleShared.Response {
	groupID := chi.URLParam(r, "group_id")
	groupIDColumn := fmt.Sprintf("%s.id", shared.TableCoreGroups)
	condition := builder.Equal(groupIDColumn, groupID)
	group := models.Group{
		ID: groupID,
	}

	return db.Update(r, &group, "UpdateGroup", shared.TableCoreGroups, condition)
}

// DeleteGroup deletes object from the database
func DeleteGroup(r *http.Request) *moduleShared.Response {
	groupID := chi.URLParam(r, "group_id")
	groupIDColumn := fmt.Sprintf("%s.id", shared.TableCoreGroups)
	condition := builder.Equal(groupIDColumn, groupID)

	return db.Remove(r, "DeleteGroup", shared.TableCoreGroups, condition)
}

// InsertUserInGroup persists the request creating a new object in the database
func InsertUserInGroup(r *http.Request) *moduleShared.Response {
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}

	permissionGroupID := chi.URLParam(r, "group_id")
	permissionUserID := chi.URLParam(r, "user_id")

	userID := r.Header.Get("userID")
	now := time.Now()

	statemant := builder.Insert(
		shared.TableCoreGroupsUsers,
		"group_id",
		"user_id",
		"created_by",
		"created_at",
		"updated_by",
		"updated_at",
	).Values(
		permissionGroupID,
		permissionUserID,
		userID,
		now,
		userID,
		now,
	)

	err := sql.Exec(statemant)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(moduleShared.ErrorInsertingRecord, "InsertUserInGroup", err.Error()))

		return response
	}

	return response
}

// RemoveUserFromGroup deletes object from the database
func RemoveUserFromGroup(r *http.Request) *moduleShared.Response {
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}

	groupID := chi.URLParam(r, "group_id")
	userID := chi.URLParam(r, "user_id")

	statemant := builder.Delete(shared.TableCoreGroupsUsers).Where(
		builder.And(
			builder.Equal("group_id", groupID),
			builder.Equal("user_id", userID),
		),
	)

	err := sql.Exec(statemant)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(moduleShared.ErrorDeletingData, "RemoveUserFromGroup", err.Error()))

		return response
	}

	return response
}

// InsertPermission persists the request body creating a new object in the database
func InsertPermission(r *http.Request) *moduleShared.Response {
	groupID := chi.URLParam(r, "group_id")
	permission := models.Permission{
		GroupID: groupID,
	}

	return db.Create(r, &permission, "InsertPermission", shared.TableCoreGrpPermissions)
}

// LoadAllPermissionsByGroup return all instances from the object
func LoadAllPermissionsByGroup(r *http.Request) *moduleShared.Response {
	permissions := []models.Permission{}
	groupID := chi.URLParam(r, "group_id")
	groupIDColumn := fmt.Sprintf("%s.group_id", shared.TableCoreGrpPermissions)
	condition := builder.Equal(groupIDColumn, groupID)

	return db.Load(r, &permissions, "LoadAllPermissionsByGroup", shared.TableCoreGrpPermissions, condition)
}

// RemovePermission deletes object from the database
func RemovePermission(r *http.Request) *moduleShared.Response {
	permissionID := chi.URLParam(r, "permission_id")
	permissionIDColumn := fmt.Sprintf("%s.id", shared.TableCoreGrpPermissions)
	condition := builder.Equal(permissionIDColumn, permissionID)

	return db.Remove(r, "RemovePermission", shared.TableCoreGrpPermissions, condition)
}

// LoadAllGroupsByUser return all instances from the object
func LoadAllGroupsByUser(r *http.Request) *moduleShared.Response {
	viewUserGroups := []models.ViewUserGroup{}
	userID := chi.URLParam(r, "user_id")
	userIDColumn := fmt.Sprintf("%s.user_id", models.ViewCoreUserGroups)
	languageCode := r.Header.Get("Content-Language")
	languageCodeColumn := fmt.Sprintf("%s.language_code", models.ViewCoreUserGroups)
	condition := builder.And(
		builder.Equal(userIDColumn, userID),
		builder.Equal(languageCodeColumn, languageCode),
	)

	return db.Load(r, &viewUserGroups, "LoadAllGroupsByUser", models.ViewCoreUserGroups, condition)
}
