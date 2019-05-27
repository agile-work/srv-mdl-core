package services

import (
	"fmt"
	"net/http"

	"github.com/agile-work/srv-mdl-shared/db"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/go-chi/chi"

	localModels "github.com/agile-work/srv-mdl-core/models"
	moduleShared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-mdl-shared/models"
	shared "github.com/agile-work/srv-shared"
	sql "github.com/agile-work/srv-shared/sql-builder/db"
)

// CreateUser persists the request body creating a new object in the database
func CreateUser(r *http.Request) *moduleShared.Response {
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}

	trs, err := sql.NewTransaction()
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorInsertingRecord, "CreateUser initiating transaction", err.Error()))
		return response
	}

	user := models.User{}
	db.LoadBodyToStruct(r, &user)
	user.ID = sql.UUID()
	db.SetSchemaAudit(r, &user)
	trs.Add(sql.StructInsertQuery(shared.TableCoreUsers, &user, "", true))

	resource := localModels.Instance{}
	resource.ID = sql.UUID()
	resource.ParentID = user.ID
	db.SetSchemaAudit(r, &resource)
	trs.Add(sql.StructInsertQuery(shared.TableCustomResources, &resource, "", true))

	err = trs.Exec()
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorInsertingRecord, "CreateUser", err.Error()))
		return response
	}
	return response
}

// LoadAllUsers return all instances from the object
func LoadAllUsers(r *http.Request) *moduleShared.Response {
	users := []models.User{}

	return db.Load(r, &users, "LoadAllUsers", shared.TableCoreUsers, nil)
}

// LoadUser return only one object from the database
func LoadUser(r *http.Request) *moduleShared.Response {
	user := models.User{}
	userID := chi.URLParam(r, "user_id")
	userIDColumn := fmt.Sprintf("%s.id", shared.TableCoreUsers)
	condition := builder.Equal(userIDColumn, userID)

	return db.Load(r, &user, "LoadUser", shared.TableCoreUsers, condition)
}

// UpdateUser updates object data in the database
func UpdateUser(r *http.Request) *moduleShared.Response {
	userID := chi.URLParam(r, "user_id")
	userIDColumn := fmt.Sprintf("%s.id", shared.TableCoreUsers)
	condition := builder.Equal(userIDColumn, userID)
	user := models.User{
		ID: userID,
	}

	return db.Update(r, &user, "UpdateUser", shared.TableCoreUsers, condition)
}

// DeleteUser deletes object from the database
func DeleteUser(r *http.Request) *moduleShared.Response {
	userID := chi.URLParam(r, "user_id")
	userIDColumn := fmt.Sprintf("%s.id", shared.TableCoreUsers)
	condition := builder.Equal(userIDColumn, userID)

	return db.Remove(r, "DeleteUser", shared.TableCoreUsers, condition)
}

// LoadAllUsersByGroup return all instances from the object
func LoadAllUsersByGroup(r *http.Request) *moduleShared.Response {
	viewGroupUsers := []models.ViewGroupUser{}
	groupID := chi.URLParam(r, "group_id")
	groupIDColumn := fmt.Sprintf("%s.group_id", shared.ViewCoreGroupUsers)
	condition := builder.Equal(groupIDColumn, groupID)

	return db.Load(r, &viewGroupUsers, "LoadAllUsersByGroup", shared.ViewCoreGroupUsers, condition)
}
