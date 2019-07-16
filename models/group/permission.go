package group

import (
	"net/http"
	"time"

	mdlShared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-mdl-shared/models/customerror"
	"github.com/agile-work/srv-mdl-shared/util"
	"github.com/agile-work/srv-shared/constants"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// SecurityTree struct
type SecurityTree struct {
	Code  string `json:"code,omitempty" sql:"code" validate:"required"`
	Unit  string `json:"unit,omitempty" sql:"unit" validate:"required"`
	Scope string `json:"scope,omitempty" sql:"scope" validate:"required"`
}

// SecurityAudit struct
type SecurityAudit struct {
	Operator  string    `json:"operator,omitempty" sql:"operator"`
	CreatedBy string    `json:"created_by" sql:"created_by"`
	CreatedAt time.Time `json:"created_at" sql:"created_at"`
}

// SecurityUser struct
type SecurityUser struct {
	IncludeResources map[string]SecurityAudit `json:"include_resources,omitempty" sql:"include_resources"`
	ExcludeResources map[string]SecurityAudit `json:"exclude_resources,omitempty" sql:"exclude_resources"`
	Wildcard         string                   `json:"wildcard,omitempty" sql:"wildcard"`
	UpdatedBy        string                   `json:"updated_by" sql:"updated_by"`
	UpdatedAt        *time.Time               `json:"updated_at" sql:"updated_at"`
}

// SecurityField struct
type SecurityField struct {
	EditAll        bool     `json:"edit_all,omitempty" sql:"edit_all"`
	EditExceptions []string `json:"edit_exceptions,omitempty" sql:"edit_exceptions"`
	ViewAll        bool     `json:"view_all,omitempty" sql:"view_all"`
	ViewExceptions []string `json:"view_exceptions,omitempty" sql:"view_exceptions"`
}

// SecurityInstance struct
type SecurityInstance struct {
	ViewAll bool `json:"view_all" sql:"view_all"`
	Create  bool `json:"create" sql:"create"`
	Delete  bool `json:"delete" sql:"delete"`
}

// SecuritySchema struct
type SecuritySchema struct {
	Instances SecurityInstance          `json:"instances,omitempty" sql:"instances"`
	Fields    SecurityField             `json:"fields,omitempty" sql:"fields"`
	Modules   map[string]SecurityModule `json:"modules,omitempty" sql:"modules"`
}

// SecurityModule struct
type SecurityModule struct {
	Instances SecurityInstance                  `json:"instances,omitempty" sql:"instances"`
	Fields    SecurityField                     `json:"fields,omitempty" sql:"fields"`
	Features  map[string]map[string]interface{} `json:"features,omitempty" sql:"features"`
}

// SecurityProcess struct
type SecurityProcess struct {
	View   bool `json:"view" sql:"view"`
	Start  bool `json:"start" sql:"start"`
	Cancel bool `json:"cancel" sql:"cancel"`
	Delete bool `json:"delete" sql:"delete"`
}

// SecurityWidget struct
type SecurityWidget struct {
	View      bool `json:"view" sql:"view"`
	Customize bool `json:"customize" sql:"customize"`
}

// Permission struct
type Permission struct {
	Widgets   map[string]SecurityWidget  `json:"widgets,omitempty" sql:"widgets"`
	Processes map[string]SecurityProcess `json:"processes,omitempty" sql:"processes"`
	Schemas   map[string]SecuritySchema  `json:"schemas,omitempty" sql:"schemas"`
}

// Definitions struct
type Definitions struct {
	Users       *SecurityUser `json:"users,omitempty" sql:"users"`
	Permissions *Permission   `json:"permissions,omitempty" sql:"permissions"`
	Tree        *SecurityTree `json:"tree,omitempty" sql:"tree"`
}

// UpdateTree update tree in group
func (d *Definitions) UpdateTree(trs *db.Transaction, groupCode string, columns map[string]interface{}) error {
	group := &Group{Code: groupCode}
	if err := group.Load(); err != nil {
		return err
	}

	if group.ID == "" {
		return customerror.New(http.StatusBadRequest, "load group", "invalid group code")
	}

	if group.Type != constants.GroupTypeTree {
		return customerror.New(http.StatusBadRequest, "load group", "invalid group type")
	}

	tree := &SecurityTree{}
	if err := util.DataToStruct(columns, tree); err != nil {
		return customerror.New(http.StatusBadRequest, "tree parse", err.Error())
	}

	if err := mdlShared.Validate.Struct(tree); err != nil {
		return customerror.New(http.StatusBadRequest, "tree invalid body", err.Error())
	}

	total, err := db.Count("id", constants.TableCoreTreeUnits, &db.Options{
		Conditions: builder.And(
			builder.Equal("tree_code", tree.Code),
			builder.Equal("code", tree.Unit),
		),
	})
	if err != nil {
		return customerror.New(http.StatusInternalServerError, "group tree update", err.Error())
	}
	if total <= 0 {
		return customerror.New(http.StatusInternalServerError, "group tree update", "invalid tree")
	}

	if err := db.UpdateJSONAttributeTx(trs.Tx, constants.TableCoreGroups, "definitions", "{tree}", columns, builder.Equal("code", groupCode)); err != nil {
		return customerror.New(http.StatusInternalServerError, "group tree update", err.Error())
	}
	return nil
}

// UpdateUsers update users in group
func (d *Definitions) UpdateUsers(trs *db.Transaction, groupCode, username string, columns map[string]interface{}) error {
	group := &Group{Code: groupCode}
	if err := group.Load(); err != nil {
		return err
	}

	if group.ID == "" {
		return customerror.New(http.StatusBadRequest, "load group", "invalid group code")
	}

	if err := util.DataToStruct(columns, &d.Users); err != nil {
		return customerror.New(http.StatusBadRequest, "users parse", err.Error())
	}

	if _, ok := columns["include_resources"]; ok {
		group.handleUsers(d, "include_resources", username, columns)
	}

	if _, ok := columns["exclude_resources"]; ok {
		group.handleUsers(d, "exclude_resources", username, columns)
	}

	if val, ok := columns["wildcard"]; ok {
		group.Definitions.Users.Wildcard = val.(string)
	}

	util.SetSchemaAudit(false, username, group.Definitions.Users)

	if err := db.UpdateJSONAttributeTx(trs.Tx, constants.TableCoreGroups, "definitions", "{users}", group.Definitions.Users, builder.Equal("code", groupCode)); err != nil {
		return customerror.New(http.StatusInternalServerError, "group users update", err.Error())
	}
	return nil
}

// UpdatePermissionWidgets update widget in permission group
func (d *Definitions) UpdatePermissionWidgets(trs *db.Transaction, groupCode string, columns map[string]interface{}) error {
	group := &Group{Code: groupCode}
	if err := group.Load(); err != nil {
		return err
	}

	if group.ID == "" {
		return customerror.New(http.StatusBadRequest, "load group", "invalid group code")
	}

	permissions := group.Definitions.Permissions
	if permissions == nil {
		permissions = &Permission{}
	}

	for widgetCode, value := range columns {
		securityWidget := SecurityWidget{}
		if err := util.DataToStruct(value, &securityWidget); err != nil {
			return customerror.New(http.StatusBadRequest, "users parse", err.Error())
		}
		if len(permissions.Widgets) == 0 {
			permissions.Widgets = make(map[string]SecurityWidget)
		}
		permissions.Widgets[widgetCode] = securityWidget
	}

	group.Definitions.Permissions = permissions

	if err := db.UpdateJSONAttributeTx(trs.Tx, constants.TableCoreGroups, "definitions", "{permissions}", group.Definitions.Permissions, builder.Equal("code", groupCode)); err != nil {
		return customerror.New(http.StatusInternalServerError, "group permission widget update", err.Error())
	}
	return nil
}

// DeletePermissionWidgets delete widget in permission group
func (d *Definitions) DeletePermissionWidgets(trs *db.Transaction, groupCode, widgetCode string) error {
	group := &Group{Code: groupCode}
	if err := group.Load(); err != nil {
		return err
	}

	if group.ID == "" {
		return customerror.New(http.StatusBadRequest, "load group", "invalid group code")
	}

	permissions := group.Definitions.Permissions

	if permissions != nil {
		delete(permissions.Widgets, widgetCode)
		if err := db.UpdateJSONAttributeTx(trs.Tx, constants.TableCoreGroups, "definitions", "{permissions}", group.Definitions.Permissions, builder.Equal("code", groupCode)); err != nil {
			return customerror.New(http.StatusInternalServerError, "group permission widget delete", err.Error())
		}
	}
	return nil
}

// UpdatePermissionProcesses update processe in permission group
func (d *Definitions) UpdatePermissionProcesses(trs *db.Transaction, groupCode string, columns map[string]interface{}) error {
	group := &Group{Code: groupCode}
	if err := group.Load(); err != nil {
		return err
	}

	if group.ID == "" {
		return customerror.New(http.StatusBadRequest, "load group", "invalid group code")
	}

	permissions := group.Definitions.Permissions
	if permissions == nil {
		permissions = &Permission{}
	}

	for processCode, value := range columns {
		securityProcess := SecurityProcess{}
		if err := util.DataToStruct(value, &securityProcess); err != nil {
			return customerror.New(http.StatusBadRequest, "users parse", err.Error())
		}
		if len(permissions.Processes) == 0 {
			permissions.Processes = make(map[string]SecurityProcess)
		}
		permissions.Processes[processCode] = securityProcess
	}

	group.Definitions.Permissions = permissions

	if err := db.UpdateJSONAttributeTx(trs.Tx, constants.TableCoreGroups, "definitions", "{permissions}", group.Definitions.Permissions, builder.Equal("code", groupCode)); err != nil {
		return customerror.New(http.StatusInternalServerError, "group permission process update", err.Error())
	}
	return nil
}

// DeletePermissionProcesses delete process in permission group
func (d *Definitions) DeletePermissionProcesses(trs *db.Transaction, groupCode, processCode string) error {
	group := &Group{Code: groupCode}
	if err := group.Load(); err != nil {
		return err
	}

	if group.ID == "" {
		return customerror.New(http.StatusBadRequest, "load group", "invalid group code")
	}

	permissions := group.Definitions.Permissions

	if permissions != nil {
		delete(permissions.Processes, processCode)
		if err := db.UpdateJSONAttributeTx(trs.Tx, constants.TableCoreGroups, "definitions", "{permissions}", group.Definitions.Permissions, builder.Equal("code", groupCode)); err != nil {
			return customerror.New(http.StatusInternalServerError, "group permission process delete", err.Error())
		}
	}
	return nil
}

func (d *Definitions) getUserIncludeResource() map[string]SecurityAudit {
	result := make(map[string]SecurityAudit)
	if d.Users != nil {
		result = d.Users.IncludeResources
	}
	return result
}

func (d *Definitions) getUserExcludeResource() map[string]SecurityAudit {
	result := make(map[string]SecurityAudit)
	if d.Users != nil {
		result = d.Users.ExcludeResources
	}
	return result
}
