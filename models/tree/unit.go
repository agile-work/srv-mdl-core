package tree

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/agile-work/srv-mdl-shared/models/customerror"
	"github.com/agile-work/srv-mdl-shared/models/security"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-shared/constants"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// Unit defines the struct of this object
type Unit struct {
	ID              string                  `json:"id" sql:"id" pk:"true"`
	Code            string                  `json:"code" sql:"code" updatable:"false" validate:"required"`
	Name            translation.Translation `json:"name" sql:"name" field:"jsonb" validate:"required"`
	Description     translation.Translation `json:"description" sql:"description" field:"jsonb" validate:"required"`
	TreeCode        string                  `json:"tree_code" sql:"tree_code" updatable:"false" validate:"required" fk:"true"`
	Path            string                  `json:"path" sql:"path" updatable:"false" validate:"required"`
	PermissionScope string                  `json:"permission_scope" sql:"permission_scope"`
	Active          bool                    `json:"active" sql:"active"`
	CreatedBy       string                  `json:"created_by" sql:"created_by"`
	CreatedAt       time.Time               `json:"created_at" sql:"created_at"`
	UpdatedBy       string                  `json:"updated_by" sql:"updated_by"`
	UpdatedAt       time.Time               `json:"updated_at" sql:"updated_at"`
	Permissions     []security.Permission   `json:"permissions" sql:"permissions" field:"jsonb"`
}

// Create persists the struct creating a new object in the database
func (u *Unit) Create(trs *db.Transaction, columns ...string) error {
	id, err := db.InsertStructTx(trs.Tx, constants.TableCoreTreeUnits, u, columns...)
	if err != nil {
		customerror.New(http.StatusInternalServerError, "level create", err.Error())
	}
	u.ID = id
	return nil
}

// Load defines only one object from the database
func (u *Unit) Load() error {
	if err := db.SelectStruct(constants.TableCoreTreeUnits, u, &db.Options{
		Conditions: builder.And(
			builder.Equal("tree_code", u.TreeCode),
			builder.Equal("code", u.Code),
		),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "level load", err.Error())
	}
	return nil
}

// Update updates object data in the database
func (u *Unit) Update(trs *db.Transaction, columns []string, translations map[string]string) error {
	opt := &db.Options{Conditions: builder.And(
		builder.Equal("tree_code", u.TreeCode),
		builder.Equal("code", u.Code),
	)}

	if len(columns) > 0 {
		if err := db.UpdateStructTx(trs.Tx, constants.TableCoreTreeUnits, u, opt, strings.Join(columns, ",")); err != nil {
			return customerror.New(http.StatusInternalServerError, "level update", err.Error())
		}
	}

	if len(translations) > 0 {
		statement := builder.Update(constants.TableCoreTreeUnits)
		for col, val := range translations {
			statement.JSON(col, translation.FieldsRequestLanguageCode)
			jsonVal, _ := json.Marshal(val)
			statement.Values(jsonVal)
		}
		statement.Where(opt.Conditions)
		if _, err := trs.Query(statement); err != nil {
			return customerror.New(http.StatusInternalServerError, "level update", err.Error())
		}
	}

	return nil
}

// Delete deletes object from the database
func (u *Unit) Delete(trs *db.Transaction) error {
	if err := db.DeleteStructTx(trs.Tx, constants.TableCoreTreeUnits, &db.Options{
		Conditions: builder.And(
			builder.Equal("tree_code", u.TreeCode),
			builder.Equal("code", u.Code),
		),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "level delete", err.Error())
	}
	return nil
}

// Units defines the array struct of this object
type Units []Unit

// LoadAll defines all instances from the object
func (l *Units) LoadAll(opt *db.Options) error {
	if err := db.SelectStruct(constants.TableCoreTreeUnits, l, opt); err != nil {
		return customerror.New(http.StatusInternalServerError, "levels load", err.Error())
	}
	return nil
}
