package group

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/agile-work/srv-mdl-shared/models/customerror"
	"github.com/agile-work/srv-mdl-shared/models/security"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-mdl-shared/util"

	"github.com/agile-work/srv-shared/constants"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// Group defines the struct of this object
type Group struct {
	ID          string                    `json:"id" sql:"id" pk:"true"`
	Code        string                    `json:"code" sql:"code"`
	ContentCode string                    `json:"content_code" sql:"content_code"`
	Name        translation.Translation   `json:"name" sql:"name" field:"jsonb" validate:"required"`
	Description translation.Translation   `json:"description" sql:"description" field:"jsonb" validate:"required"`
	Permissions []security.Permission     `json:"permissions" sql:"permissions" field:"jsonb"`
	Users       []security.PermissionUser `json:"users" sql:"users" field:"jsonb"`
	Active      bool                      `json:"active" sql:"active"`
	CreatedBy   string                    `json:"created_by" sql:"created_by"`
	CreatedAt   time.Time                 `json:"created_at" sql:"created_at"`
	UpdatedBy   string                    `json:"updated_by" sql:"updated_by"`
	UpdatedAt   time.Time                 `json:"updated_at" sql:"updated_at"`
}

// Create persists the struct creating a new object in the database
func (g *Group) Create(trs *db.Transaction, columns ...string) error {
	if g.ContentCode != "" {
		prefix, err := util.GetContentPrefix(g.ContentCode)
		if err != nil {
			return customerror.New(http.StatusInternalServerError, "group create", err.Error())
		}
		g.Code = fmt.Sprintf("%s_%s", prefix, g.Code)
	} else {
		g.Code = fmt.Sprintf("%s_%s", "custom", g.Code)
	}

	if len(g.Code) > constants.DatabaseMaxLength {
		return customerror.New(http.StatusInternalServerError, "group create", "invalid code lenght")
	}

	id, err := db.InsertStructTx(trs.Tx, constants.TableCoreGroups, g, columns...)
	if err != nil {
		return customerror.New(http.StatusInternalServerError, "group create", err.Error())
	}
	g.ID = id

	return nil
}

// Load defines only one object from the database
func (g *Group) Load() error {
	if err := db.SelectStruct(constants.TableCoreGroups, g, &db.Options{
		Conditions: builder.Equal("code", g.Code),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "group load", err.Error())
	}
	return nil
}

// Update updates object data in the database
func (g *Group) Update(trs *db.Transaction, columns []string, translations map[string]string) error {
	opt := &db.Options{Conditions: builder.Equal("code", g.Code)}

	if g.ContentCode != "" {
		if err := util.ValidateContent(g.ContentCode); err != nil {
			return customerror.New(http.StatusInternalServerError, "group update", err.Error())
		}
	}

	if len(columns) > 0 {
		if err := db.UpdateStructTx(trs.Tx, constants.TableCoreGroups, g, opt, strings.Join(columns, ",")); err != nil {
			return customerror.New(http.StatusInternalServerError, "group update", err.Error())
		}
	}

	if len(translations) > 0 {
		statement := builder.Update(constants.TableCoreGroups)
		for col, val := range translations {
			statement.JSON(col, translation.FieldsRequestLanguageCode)
			jsonVal, _ := json.Marshal(val)
			statement.Values(jsonVal)
		}
		statement.Where(opt.Conditions)
		if _, err := trs.Query(statement); err != nil {
			return customerror.New(http.StatusInternalServerError, "group update", err.Error())
		}
	}

	return nil
}

// Delete deletes object from the database
func (g *Group) Delete(trs *db.Transaction) error {
	if err := db.DeleteStructTx(trs.Tx, constants.TableCoreGroups, &db.Options{
		Conditions: builder.Equal("code", g.Code),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "group delete", err.Error())
	}
	return nil
}

// Groups defines the array struct of this object
type Groups []Group

// LoadAll defines all instances from the object
func (g *Groups) LoadAll(opt *db.Options) error {
	if err := db.SelectStruct(constants.TableCoreGroups, g, opt); err != nil {
		return customerror.New(http.StatusInternalServerError, "groups load", err.Error())
	}
	return nil
}

// Validate check if group exists and is active
func Validate(codes []string) error {
	return nil
}

// // InsertUserInGroup persists the request creating a new object in the database
// func InsertUserInGroup(r *http.Request) *mdlShared.Response {
// 	groupID := chi.URLParam(r, "group_id")
// 	userID := chi.URLParam(r, "user_id")
// 	user := models.GroupUser{
// 		ID: userID,
// 	}

// 	response := db.GetResponse(r, &user, "InsertUserInGroup")
// 	if response.Code != http.StatusOK {
// 		return response
// 	}

// 	idColumn := fmt.Sprintf("%s.id", shared.TableCoreGroups)
// 	sql.InsertStructToJSON("users", shared.TableCoreGroups, &user, builder.Equal(idColumn, groupID))
// 	return response
// }

// // RemoveUserFromGroup deletes object from the database
// func RemoveUserFromGroup(r *http.Request) *mdlShared.Response {
// 	response := &mdlShared.Response{
// 		Code: http.StatusOK,
// 	}
// 	groupID := chi.URLParam(r, "group_id")
// 	userID := chi.URLParam(r, "user_id")

// 	err := sql.DeleteStructFromJSON(userID, groupID, "users", shared.TableCoreGroups)
// 	if err != nil {
// 		response.Code = http.StatusInternalServerError
// 		response.Errors = append(response.Errors, mdlShared.NewResponseError(shared.ErrorParsingRequest, "RemoveUserFromGroup", err.Error()))

// 		return response
// 	}

// 	return response
// }

// // LoadAllGroupPermissions return all instances from the object
// func LoadAllGroupPermissions(r *http.Request) *mdlShared.Response {
// 	permissions := []models.Permission{}

// 	groupID := chi.URLParam(r, "group_id")
// 	groupIDColumn := fmt.Sprintf("%s.parent_id", shared.ViewCoreStructurePermissions)
// 	languageCode := r.Header.Get("Content-Language")
// 	languageCodeColumn := fmt.Sprintf("%s.language_code", shared.ViewCoreStructurePermissions)
// 	condition := builder.And(
// 		builder.Equal(groupIDColumn, groupID),
// 		builder.Equal(languageCodeColumn, languageCode),
// 	)

// 	return db.Load(r, &permissions, "LoadAllGroupPermissions", shared.ViewCoreStructurePermissions, condition)
// }

// // InsertGroupPermission persists the request body creating a new object in the database
// func InsertGroupPermission(r *http.Request) *mdlShared.Response {
// 	groupID := chi.URLParam(r, "group_id")
// 	permission := models.Permission{}

// 	response := db.GetResponse(r, &permission, "InsertGroupPermission")
// 	if response.Code != http.StatusOK {
// 		return response
// 	}
// 	permission.ID = sql.UUID()

// 	idColumn := fmt.Sprintf("%s.id", shared.TableCoreGroups)
// 	sql.InsertStructToJSON("permissions", shared.TableCoreGroups, &permission, builder.Equal(idColumn, groupID))
// 	resp.Data = permission
// 	return response
// }

// // RemoveGroupPermission deletes object from the database
// func RemoveGroupPermission(r *http.Request) *mdlShared.Response {
// 	response := &mdlShared.Response{
// 		Code: http.StatusOK,
// 	}
// 	groupID := chi.URLParam(r, "group_id")
// 	permissionID := chi.URLParam(r, "permission_id")

// 	err := sql.DeleteStructFromJSON(permissionID, groupID, "permissions", shared.TableCoreGroups)
// 	if err != nil {
// 		response.Code = http.StatusInternalServerError
// 		response.Errors = append(response.Errors, mdlShared.NewResponseError(shared.ErrorParsingRequest, "RemoveGroupPermissionFromGroup", err.Error()))

// 		return response
// 	}

// 	return response
// }

// // LoadAllGroupsByUser return all instances from the object
// func LoadAllGroupsByUser(r *http.Request) *mdlShared.Response {
// 	viewUserGroups := []models.ViewUserGroup{}
// 	userID := chi.URLParam(r, "user_id")
// 	userIDColumn := fmt.Sprintf("%s.user_id", shared.ViewCoreUserGroups)
// 	languageCode := r.Header.Get("Content-Language")
// 	languageCodeColumn := fmt.Sprintf("%s.language_code", shared.ViewCoreUserGroups)
// 	condition := builder.And(
// 		builder.Equal(userIDColumn, userID),
// 		builder.Equal(languageCodeColumn, languageCode),
// 	)

// 	return db.Load(r, &viewUserGroups, "LoadAllGroupsByUser", shared.ViewCoreUserGroups, condition)
// }
