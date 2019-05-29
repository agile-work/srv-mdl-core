package admin

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
	groupID := chi.URLParam(r, "group_id")
	userID := chi.URLParam(r, "user_id")
	user := models.GroupUser{
		ID: userID,
	}

	response := db.GetResponse(r, &user, "InsertUserInGroup")
	if response.Code != http.StatusOK {
		return response
	}

	idColumn := fmt.Sprintf("%s.id", shared.TableCoreGroups)
	sql.InsertStructToJSON("users", shared.TableCoreGroups, &user, builder.Equal(idColumn, groupID))
	return response
}

// RemoveUserFromGroup deletes object from the database
func RemoveUserFromGroup(r *http.Request) *moduleShared.Response {
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}
	groupID := chi.URLParam(r, "group_id")
	userID := chi.URLParam(r, "user_id")

	err := sql.DeleteStructFromJSON(userID, groupID, "users", shared.TableCoreGroups)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorParsingRequest, "RemoveUserFromGroup", err.Error()))

		return response
	}

	return response
}

// LoadAllGroupPermissions return all instances from the object
func LoadAllGroupPermissions(r *http.Request) *moduleShared.Response {
	permissions := []models.Permission{}

	groupID := chi.URLParam(r, "group_id")
	groupIDColumn := fmt.Sprintf("%s.parent_id", shared.ViewCoreStructurePermissions)
	languageCode := r.Header.Get("Content-Language")
	languageCodeColumn := fmt.Sprintf("%s.language_code", shared.ViewCoreStructurePermissions)
	condition := builder.And(
		builder.Equal(groupIDColumn, groupID),
		builder.Equal(languageCodeColumn, languageCode),
	)

	return db.Load(r, &permissions, "LoadAllGroupPermissions", shared.ViewCoreStructurePermissions, condition)
}

// InsertGroupPermission persists the request body creating a new object in the database
func InsertGroupPermission(r *http.Request) *moduleShared.Response {
	groupID := chi.URLParam(r, "group_id")
	permission := models.Permission{}

	response := db.GetResponse(r, &permission, "InsertGroupPermission")
	if response.Code != http.StatusOK {
		return response
	}
	permission.ID = sql.UUID()

	idColumn := fmt.Sprintf("%s.id", shared.TableCoreGroups)
	sql.InsertStructToJSON("permissions", shared.TableCoreGroups, &permission, builder.Equal(idColumn, groupID))
	response.Data = permission
	return response
}

// RemoveGroupPermission deletes object from the database
func RemoveGroupPermission(r *http.Request) *moduleShared.Response {
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}
	groupID := chi.URLParam(r, "group_id")
	permissionID := chi.URLParam(r, "permission_id")

	err := sql.DeleteStructFromJSON(permissionID, groupID, "permissions", shared.TableCoreGroups)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorParsingRequest, "RemoveGroupPermissionFromGroup", err.Error()))

		return response
	}

	return response
}

// LoadAllGroupsByUser return all instances from the object
func LoadAllGroupsByUser(r *http.Request) *moduleShared.Response {
	viewUserGroups := []models.ViewUserGroup{}
	userID := chi.URLParam(r, "user_id")
	userIDColumn := fmt.Sprintf("%s.user_id", shared.ViewCoreUserGroups)
	languageCode := r.Header.Get("Content-Language")
	languageCodeColumn := fmt.Sprintf("%s.language_code", shared.ViewCoreUserGroups)
	condition := builder.And(
		builder.Equal(userIDColumn, userID),
		builder.Equal(languageCodeColumn, languageCode),
	)

	return db.Load(r, &viewUserGroups, "LoadAllGroupsByUser", shared.ViewCoreUserGroups, condition)
}
