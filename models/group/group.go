package group

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/agile-work/srv-mdl-shared/models/customerror"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-mdl-shared/util"
	sharedUtil "github.com/agile-work/srv-shared/util"

	"github.com/agile-work/srv-shared/constants"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// Group struct
type Group struct {
	ID          string                  `json:"id" sql:"id" pk:"true"`
	Code        string                  `json:"code" sql:"code" updatable:"false" validate:"required"`
	Name        translation.Translation `json:"name" sql:"name" field:"jsonb" validate:"required"`
	Description translation.Translation `json:"description" sql:"description" field:"jsonb" validate:"required"`
	ContentCode string                  `json:"content_code,omitempty" sql:"content_code"`
	Type        string                  `json:"group_type" sql:"group_type" validate:"required"`
	Users       SecurityUser            `json:"users" sql:"users" field:"jsonb"`
	Permissions Permission              `json:"permissions" sql:"permissions" field:"jsonb"`
	Tree        SecurityTree            `json:"tree,omitempty" sql:"tree" field:"jsonb"`
	Active      bool                    `json:"active" sql:"active"`
	CreatedBy   string                  `json:"created_by" sql:"created_by"`
	CreatedAt   time.Time               `json:"created_at" sql:"created_at"`
	UpdatedBy   string                  `json:"updated_by" sql:"updated_by"`
	UpdatedAt   time.Time               `json:"updated_at" sql:"updated_at"`
}

// Create persists the struct creating a new object in the database
func (g *Group) Create(trs *db.Transaction) error {
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
		return customerror.New(http.StatusInternalServerError, "group create", "invalid code length")
	}

	id, err := db.InsertStructTx(trs.Tx, constants.TableCoreGroups, g)
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

// UpdateTree update tree in group
func (g *Group) UpdateTree(trs *db.Transaction) error {
	group := &Group{Code: g.Code}
	if err := group.Load(); err != nil {
		return err
	}

	if group.ID == "" {
		return customerror.New(http.StatusBadRequest, "load group", "invalid group code")
	}

	if group.Type != constants.GroupTypeTree {
		return customerror.New(http.StatusBadRequest, "load group", "invalid group type")
	}

	total, err := db.Count("id", constants.TableCoreTreeUnits, &db.Options{
		Conditions: builder.And(
			builder.Equal("tree_code", g.Tree.Code),
			builder.Equal("code", g.Tree.Unit),
		),
	})
	if err != nil {
		return customerror.New(http.StatusInternalServerError, "group tree update", err.Error())
	}
	if total <= 0 {
		return customerror.New(http.StatusInternalServerError, "group tree update", "invalid tree")
	}
	return g.Update(trs, []string{"tree"}, map[string]string{})
}

// UpdateUsers update users in group
func (g *Group) UpdateUsers(trs *db.Transaction, username string, columns []string) error {
	group := &Group{Code: g.Code}
	if err := group.Load(); err != nil {
		return err
	}

	if group.ID == "" {
		return customerror.New(http.StatusBadRequest, "load group", "invalid group code")
	}

	// TODO: Make a validation about usernames
	if len(g.Users.IncludeResources) > 0 {
		for username, security := range g.Users.IncludeResources {
			if security.Operator == "include" {
				if len(group.Users.IncludeResources) == 0 {
					group.Users.IncludeResources = map[string]SecurityAudit{}
				}
				securityAudit := SecurityAudit{}
				util.SetSchemaAudit(http.MethodPost, username, &securityAudit)
				if _, ok := group.Users.IncludeResources[username]; !ok {
					group.Users.IncludeResources[username] = securityAudit
				}
			} else {
				delete(group.Users.IncludeResources, username)
			}
		}
	} else if sharedUtil.Contains(columns, "include_resources") {
		group.Users.IncludeResources = g.Users.IncludeResources
	}
	if len(g.Users.ExcludeResources) > 0 {
		for username, security := range g.Users.ExcludeResources {
			if security.Operator == "include" {
				if len(group.Users.ExcludeResources) == 0 {
					group.Users.ExcludeResources = map[string]SecurityAudit{}
				}
				securityAudit := SecurityAudit{}
				util.SetSchemaAudit(http.MethodPost, username, &securityAudit)
				if _, ok := group.Users.IncludeResources[username]; !ok {
					group.Users.ExcludeResources[username] = securityAudit
				}
			} else {
				delete(group.Users.ExcludeResources, username)
			}
		}
	} else if sharedUtil.Contains(columns, "exclude_resources") {
		group.Users.ExcludeResources = g.Users.ExcludeResources
	}
	if sharedUtil.Contains(columns, "wildcard") {
		// TODO: Make a validation about wildcard value
		group.Users.Wildcard = g.Users.Wildcard
	}
	httpMethod := http.MethodPatch
	util.SetSchemaAudit(httpMethod, username, group)
	if group.Users.alreadyCreated() {
		httpMethod = http.MethodPost
	}
	util.SetSchemaAudit(httpMethod, username, &group.Users)
	return group.Update(trs, []string{"users"}, map[string]string{})
}

func (u SecurityUser) alreadyCreated() bool {
	if u.CreatedBy == "" {
		return false
	}
	return true
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
