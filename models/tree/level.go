package tree

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/agile-work/srv-mdl-shared/models/customerror"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-shared/constants"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// Level defines the struct of this object
type Level struct {
	ID          string                  `json:"id" sql:"id" pk:"true"`
	Code        string                  `json:"code" sql:"code" updatable:"false" validate:"required"`
	Name        translation.Translation `json:"name" sql:"name" field:"jsonb" validate:"required"`
	Description translation.Translation `json:"description" sql:"description" field:"jsonb" validate:"required"`
	TreeCode    string                  `json:"tree_code" sql:"tree_code" fk:"true" updatable:"false" validate:"required"`
	CreatedBy   string                  `json:"created_by" sql:"created_by"`
	CreatedAt   time.Time               `json:"created_at" sql:"created_at"`
	UpdatedBy   string                  `json:"updated_by" sql:"updated_by"`
	UpdatedAt   time.Time               `json:"updated_at" sql:"updated_at"`
}

// Create persists the struct creating a new object in the database
func (l *Level) Create(trs *db.Transaction) error {
	id, err := db.InsertStructTx(trs.Tx, constants.TableCoreTreeLevels, l)
	if err != nil {
		customerror.New(http.StatusInternalServerError, "level create", err.Error())
	}
	l.ID = id
	return nil
}

// Load defines only one object from the database
func (l *Level) Load() error {
	if err := db.SelectStruct(constants.TableCoreTreeLevels, l, &db.Options{
		Conditions: builder.And(
			builder.Equal("tree_code", l.TreeCode),
			builder.Equal("code", l.Code),
		),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "level load", err.Error())
	}
	return nil
}

// Update updates object data in the database
func (l *Level) Update(trs *db.Transaction, columns []string, translations map[string]string) error {
	opt := &db.Options{Conditions: builder.And(
		builder.Equal("tree_code", l.TreeCode),
		builder.Equal("code", l.Code),
	)}

	if len(columns) > 0 {
		if err := db.UpdateStructTx(trs.Tx, constants.TableCoreTreeLevels, l, opt, strings.Join(columns, ",")); err != nil {
			return customerror.New(http.StatusInternalServerError, "level update", err.Error())
		}
	}

	if len(translations) > 0 {
		statement := builder.Update(constants.TableCoreTreeLevels)
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
func (l *Level) Delete(trs *db.Transaction) error {
	if err := db.DeleteStructTx(trs.Tx, constants.TableCoreTreeLevels, &db.Options{
		Conditions: builder.And(
			builder.Equal("tree_code", l.TreeCode),
			builder.Equal("code", l.Code),
		),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "level delete", err.Error())
	}
	return nil
}

// Levels defines the array struct of this object
type Levels []Level

// LoadAll defines all instances from the object
func (l *Levels) LoadAll(opt *db.Options) error {
	if err := db.SelectStruct(constants.TableCoreTreeLevels, l, opt); err != nil {
		return customerror.New(http.StatusInternalServerError, "levels load", err.Error())
	}
	return nil
}
