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

// CreateUser persists the request body creating a new object in the database
func CreateUser(r *http.Request) *moduleShared.Response {
	user := models.User{}

	return db.Create(r, &user, "CreateUser", shared.TableCoreUsers)
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
	groupIDColumn := fmt.Sprintf("%s.group_id", models.ViewCoreGroupUsers)
	condition := builder.Equal(groupIDColumn, groupID)

	return db.Load(r, &viewGroupUsers, "LoadAllUsersByGroup", models.ViewCoreGroupUsers, condition)
}
