package group

import (
	"net/http"

	"github.com/agile-work/srv-mdl-shared/models/customerror"
	"github.com/agile-work/srv-mdl-shared/util"
	"github.com/agile-work/srv-shared/constants"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// UpdatePermissionModule update module in permission group
func (d *Definitions) UpdatePermissionModule(trs *db.Transaction, groupCode, schemaCode, moduleCode string, columns map[string]interface{}) error {
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

	securityModule := SecurityModule{}
	if err := util.DataToStruct(columns, &securityModule); err != nil {
		return customerror.New(http.StatusBadRequest, "users parse", err.Error())
	}
	if len(permissions.Schemas) == 0 {
		permissions.Schemas = make(map[string]SecuritySchema)
	}
	if _, ok := permissions.Schemas[schemaCode]; !ok {
		permissions.Schemas[schemaCode] = SecuritySchema{}
	}
	permissionSchema := permissions.Schemas[schemaCode]
	if len(permissions.Schemas[schemaCode].Modules) == 0 {
		permissionSchema.Modules = make(map[string]SecurityModule)
	}
	permissionSchema.Modules[moduleCode] = securityModule
	permissions.Schemas[schemaCode] = permissionSchema
	group.Definitions.Permissions = permissions

	if err := db.UpdateJSONAttributeTx(trs.Tx, constants.TableCoreGroups, "definitions", "{permissions}", group.Definitions.Permissions, builder.Equal("code", groupCode)); err != nil {
		return customerror.New(http.StatusInternalServerError, "group permission module update", err.Error())
	}
	return nil
}

// DeletePermissionModule delete module in permission group
func (d *Definitions) DeletePermissionModule(trs *db.Transaction, groupCode, schemaCode, moduleCode string) error {
	group := &Group{Code: groupCode}
	if err := group.Load(); err != nil {
		return err
	}

	if group.ID == "" {
		return customerror.New(http.StatusBadRequest, "load group", "invalid group code")
	}

	permissions := group.Definitions.Permissions

	if permissions != nil {
		if securitySchema, ok := permissions.Schemas[schemaCode]; ok {
			if _, ok := securitySchema.Modules[moduleCode]; ok {
				delete(securitySchema.Modules, moduleCode)
				permissions.Schemas[schemaCode] = securitySchema
				if err := db.UpdateJSONAttributeTx(trs.Tx, constants.TableCoreGroups, "definitions", "{permissions}", group.Definitions.Permissions, builder.Equal("code", groupCode)); err != nil {
					return customerror.New(http.StatusInternalServerError, "group permission module delete", err.Error())
				}
			}
		}
	}
	return nil
}

// UpdatePermissionModuleInstance update module in permission group
func (d *Definitions) UpdatePermissionModuleInstance(trs *db.Transaction, groupCode, schemaCode, moduleCode string, columns map[string]interface{}) error {
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

	securityInstance := SecurityInstance{}
	if err := util.DataToStruct(columns, &securityInstance); err != nil {
		return customerror.New(http.StatusBadRequest, "users parse", err.Error())
	}
	if len(permissions.Schemas) == 0 {
		permissions.Schemas = make(map[string]SecuritySchema)
	}
	if _, ok := permissions.Schemas[schemaCode]; !ok {
		permissions.Schemas[schemaCode] = SecuritySchema{}
	}
	permissionSchema := permissions.Schemas[schemaCode]
	if len(permissions.Schemas[schemaCode].Modules) == 0 {
		permissionSchema.Modules = make(map[string]SecurityModule)
	}
	securityModule := permissionSchema.Modules[moduleCode]
	securityModule.Instances = securityInstance
	permissionSchema.Modules[moduleCode] = securityModule
	permissions.Schemas[schemaCode] = permissionSchema
	group.Definitions.Permissions = permissions

	if err := db.UpdateJSONAttributeTx(trs.Tx, constants.TableCoreGroups, "definitions", "{permissions}", group.Definitions.Permissions, builder.Equal("code", groupCode)); err != nil {
		return customerror.New(http.StatusInternalServerError, "group permission schema update", err.Error())
	}
	return nil
}

// UpdatePermissionModuleField update module in permission group
func (d *Definitions) UpdatePermissionModuleField(trs *db.Transaction, groupCode, schemaCode, moduleCode string, columns map[string]interface{}) error {
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

	securityField := SecurityField{}
	if err := util.DataToStruct(columns, &securityField); err != nil {
		return customerror.New(http.StatusBadRequest, "users parse", err.Error())
	}
	if len(permissions.Schemas) == 0 {
		permissions.Schemas = make(map[string]SecuritySchema)
	}
	if _, ok := permissions.Schemas[schemaCode]; !ok {
		permissions.Schemas[schemaCode] = SecuritySchema{}
	}
	permissionSchema := permissions.Schemas[schemaCode]
	if len(permissions.Schemas[schemaCode].Modules) == 0 {
		permissionSchema.Modules = make(map[string]SecurityModule)
	}
	securityModule := permissionSchema.Modules[moduleCode]
	securityModule.Fields = securityField
	permissionSchema.Modules[moduleCode] = securityModule
	permissions.Schemas[schemaCode] = permissionSchema
	group.Definitions.Permissions = permissions

	if err := db.UpdateJSONAttributeTx(trs.Tx, constants.TableCoreGroups, "definitions", "{permissions}", group.Definitions.Permissions, builder.Equal("code", groupCode)); err != nil {
		return customerror.New(http.StatusInternalServerError, "group permission schema update", err.Error())
	}
	return nil
}

// UpdatePermissionModuleFeature update module in permission group
func (d *Definitions) UpdatePermissionModuleFeature(trs *db.Transaction, groupCode, schemaCode, moduleCode, featureCode string, columns map[string]interface{}) error {
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

	if len(permissions.Schemas) == 0 {
		permissions.Schemas = make(map[string]SecuritySchema)
	}
	if _, ok := permissions.Schemas[schemaCode]; !ok {
		permissions.Schemas[schemaCode] = SecuritySchema{}
	}
	permissionSchema := permissions.Schemas[schemaCode]
	if len(permissionSchema.Modules) == 0 {
		permissionSchema.Modules = make(map[string]SecurityModule)
	}
	securityModule := permissionSchema.Modules[moduleCode]
	if len(securityModule.Features) == 0 {
		securityModule.Features = make(map[string]map[string]interface{})
	}
	securityModule.Features[featureCode] = columns
	permissionSchema.Modules[moduleCode] = securityModule
	permissions.Schemas[schemaCode] = permissionSchema
	group.Definitions.Permissions = permissions

	if err := db.UpdateJSONAttributeTx(trs.Tx, constants.TableCoreGroups, "definitions", "{permissions}", group.Definitions.Permissions, builder.Equal("code", groupCode)); err != nil {
		return customerror.New(http.StatusInternalServerError, "group permission feature update", err.Error())
	}
	return nil
}
