package group

import (
	"encoding/json"
	"strings"
	"time"

	mdlShared "github.com/agile-work/srv-mdl-shared"
	mdlSharedModels "github.com/agile-work/srv-mdl-shared/models"
	"github.com/agile-work/srv-shared/constants"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// Group defines the struct of this object
type Group struct {
	ID                      string                      `json:"id" sql:"id" pk:"true"`
	Code                    string                      `json:"code" sql:"code"`
	Name                    mdlSharedModels.Translation `json:"name" sql:"name" field:"jsonb" validate:"required"`
	Description             mdlSharedModels.Translation `json:"description" sql:"description" field:"jsonb" validate:"required"`
	TreeUnitID              *string                     `json:"tree_unit_id" sql:"tree_unit_id"`
	TreeUnitPermissionScope *string                     `json:"tree_unit_permission_scope" sql:"tree_unit_permission_scope"`
	Permissions             []Permission                `json:"permissions" sql:"permissions" field:"jsonb"`
	Users                   []GroupUser                 `json:"users" sql:"users" field:"jsonb"`
	Active                  bool                        `json:"active" sql:"active"`
	CreatedBy               string                      `json:"created_by" sql:"created_by"`
	CreatedAt               time.Time                   `json:"created_at" sql:"created_at"`
	UpdatedBy               string                      `json:"updated_by" sql:"updated_by"`
	UpdatedAt               time.Time                   `json:"updated_at" sql:"updated_at"`
}

// GroupUser defines the struct to the jsonb field in group
type GroupUser struct {
	ID        string    `json:"id"`
	CreatedBy string    `json:"created_by,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Permission defines the struct of this object
type Permission struct {
	ID             string    `json:"id,omitempty" sql:"id"`
	ParentID       string    `json:"parent_id,omitempty" sql:"parent_id"`
	StructureID    string    `json:"structure_id,omitempty" sql:"structure_id"`
	StructureType  string    `json:"structure_type,omitempty" sql:"structure_type"`
	StructureName  string    `json:"structure_name,omitempty" sql:"structure_name"`
	PermissionType int       `json:"permission_type,omitempty" sql:"permission_type"`
	ConditionQuery string    `json:"condition_query,omitempty"`
	LanguageCode   string    `json:"language_code,omitempty" sql:"language_code"`
	CreatedBy      string    `json:"created_by,omitempty" sql:"created_by"`
	CreatedAt      time.Time `json:"created_at,omitempty" sql:"created_at"`
}

// ViewUserGroup defines the struct of this object
type ViewUserGroup struct {
	ID           string    `json:"id" sql:"id" pk:"true"`
	UserID       string    `json:"user_id" sql:"user_id" fk:"true"`
	Code         string    `json:"code" sql:"code"`
	Name         string    `json:"name" sql:"name"`
	Description  string    `json:"description" sql:"description"`
	LanguageCode string    `json:"language_code" sql:"language_code"`
	Active       bool      `json:"active" sql:"active"`
	CreatedBy    string    `json:"created_by" sql:"created_by"`
	CreatedAt    time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy    string    `json:"updated_by" sql:"updated_by"`
	UpdatedAt    time.Time `json:"updated_at" sql:"updated_at"`
}

// Groups defines the array struct of this object
type Groups []Group

// Create persists the struct creating a new object in the database
func (g *Group) Create(trs *db.Transaction, columns ...string) error {
	id, err := db.InsertStructTx(trs.Tx, constants.TableCoreGroups, g, columns...)
	if err != nil {
		return mdlShared.NewError("group create", err.Error())
	}
	g.ID = id

	return nil
}

// LoadAll defines all instances from the object
func (g *Groups) LoadAll(trs *db.Transaction, opt *db.Options) error {
	if err := db.SelectStructTx(trs.Tx, constants.TableCoreGroups, g, opt); err != nil {
		return mdlShared.NewError("groups load", err.Error())
	}
	return nil
}

// Load defines only one object from the database
func (g *Group) Load(trs *db.Transaction) error {
	if err := db.SelectStructTx(trs.Tx, constants.TableCoreGroups, g, &db.Options{
		Conditions: builder.Equal("code", g.Code),
	}); err != nil {
		return mdlShared.NewError("group load", err.Error())
	}
	return nil
}

// Update updates object data in the database
func (g *Group) Update(trs *db.Transaction, columns []string, translations map[string]string) error {
	opt := &db.Options{Conditions: builder.Equal("code", g.Code)}

	if len(columns) > 0 {
		if err := db.UpdateStructTx(trs.Tx, constants.TableCoreGroups, g, opt, strings.Join(columns, ",")); err != nil {
			return mdlShared.NewError("group update", err.Error())
		}
	}

	if len(translations) > 0 {
		statement := builder.Update(constants.TableCoreGroups)
		for col, val := range translations {
			statement.JSON(col, mdlSharedModels.TranslationFieldsRequestLanguageCode)
			jsonVal, _ := json.Marshal(val)
			statement.Values(jsonVal)
		}
		statement.Where(opt.Conditions)
		if _, err := trs.Query(statement); err != nil {
			return mdlShared.NewError("group update", err.Error())
		}
	}

	return nil
}

// Delete deletes object from the database
func (g *Group) Delete(trs *db.Transaction) error {
	if err := db.DeleteStructTx(trs.Tx, constants.TableCoreGroups, &db.Options{
		Conditions: builder.Equal("code", g.Code),
	}); err != nil {
		return mdlShared.NewError("group delete", err.Error())
	}
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
// 	response.Data = permission
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
