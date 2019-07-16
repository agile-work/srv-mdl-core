package group

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	mdlShared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-mdl-shared/models/customerror"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-mdl-shared/util"

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
	Definitions Definitions             `json:"definitions" sql:"definitions" field:"jsonb" updatable:"false"`
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

	if g.Type == constants.GroupTypeTree {
		if err := mdlShared.Validate.Struct(g.Definitions.Tree); err != nil {
			return customerror.New(http.StatusBadRequest, "tree invalid body", err.Error())
		}

		total, err := db.Count("id", constants.TableCoreTreeUnits, &db.Options{
			Conditions: builder.And(
				builder.Equal("tree_code", g.Definitions.Tree.Code),
				builder.Equal("code", g.Definitions.Tree.Unit),
			),
		})
		if err != nil {
			return customerror.New(http.StatusInternalServerError, "group tree create", err.Error())
		}
		if total <= 0 {
			return customerror.New(http.StatusInternalServerError, "group tree create", "invalid tree")
		}
	}

	translation.SetStructTranslationsLanguage(g, "all")
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
		if err := trs.Exec(statement); err != nil {
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

func (g *Group) handleUsers(def *Definitions, attribute, username string, columns map[string]interface{}) {
	var groupUsers, defUsers map[string]SecurityAudit

	groupUsers = g.Definitions.getUserIncludeResource()
	defUsers = def.getUserIncludeResource()

	if attribute == "exclude_resources" {
		groupUsers = g.Definitions.getUserExcludeResource()
		defUsers = def.getUserExcludeResource()
	}

	if len(defUsers) > 0 {
		for user, security := range defUsers {
			if security.Operator == "include" {
				if len(groupUsers) == 0 {
					groupUsers = map[string]SecurityAudit{}
				}
				securityAudit := SecurityAudit{}
				util.SetSchemaAudit(true, username, &securityAudit)
				if _, ok := groupUsers[user]; !ok {
					groupUsers[user] = securityAudit
				}
			} else {
				delete(groupUsers, user)
			}
		}
	} else if _, ok := columns[attribute]; ok {
		groupUsers = defUsers
	}

	if g.Definitions.Users == nil {
		g.Definitions.Users = &SecurityUser{}
	}

	if attribute == "exclude_resources" {
		g.Definitions.Users.ExcludeResources = groupUsers
	} else {
		g.Definitions.Users.IncludeResources = groupUsers
	}
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
