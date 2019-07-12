package language

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

// Language defines the struct of this object
type Language struct {
	ID        string                  `json:"id" sql:"id" pk:"true"`
	Code      string                  `json:"code" sql:"code" updatable:"false" validate:"required"`
	Name      translation.Translation `json:"name" sql:"name" field:"jsonb" validate:"required"`
	Active    bool                    `json:"active" sql:"active"`
	CreatedBy string                  `json:"created_by" sql:"created_by"`
	CreatedAt time.Time               `json:"created_at" sql:"created_at"`
	UpdatedBy string                  `json:"updated_by" sql:"updated_by"`
	UpdatedAt time.Time               `json:"updated_at" sql:"updated_at"`
}

// Create persists the struct creating a new object in the database
func (l *Language) Create(trs *db.Transaction, columns ...string) error {
	id, err := db.InsertStructTx(trs.Tx, constants.TableCoreConfigLanguages, l, columns...)
	if err != nil {
		customerror.New(http.StatusInternalServerError, "language create", err.Error())
	}
	l.ID = id
	return nil
}

// Load defines only one object from the database
func (l *Language) Load() error {
	if err := db.SelectStruct(constants.TableCoreConfigLanguages, l, &db.Options{
		Conditions: builder.Equal("code", l.Code),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "language load", err.Error())
	}
	return nil
}

// Update updates object data in the database
func (l *Language) Update(trs *db.Transaction, columns []string, translations map[string]string) error {
	opt := &db.Options{Conditions: builder.Equal("code", l.Code)}

	if len(columns) > 0 {
		if err := db.UpdateStructTx(trs.Tx, constants.TableCoreConfigLanguages, l, opt, strings.Join(columns, ",")); err != nil {
			return customerror.New(http.StatusInternalServerError, "language update", err.Error())
		}
	}

	if len(translations) > 0 {
		statement := builder.Update(constants.TableCoreConfigLanguages)
		for col, val := range translations {
			statement.JSON(col, translation.FieldsRequestLanguageCode)
			jsonVal, _ := json.Marshal(val)
			statement.Values(jsonVal)
		}
		statement.Where(opt.Conditions)
		if _, err := trs.Query(statement); err != nil {
			return customerror.New(http.StatusInternalServerError, "language update", err.Error())
		}
	}

	return nil
}

// Delete deletes object from the database
func (l *Language) Delete(trs *db.Transaction) error {
	if err := db.DeleteStructTx(trs.Tx, constants.TableCoreConfigLanguages, &db.Options{
		Conditions: builder.And(
			builder.Equal("code", l.Code),
			builder.Equal("active", false),
		),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "language delete", err.Error())
	}
	return nil
}

// Languages defines the array struct of this object
type Languages []Language

// LoadAll defines all instances from the object
func (l *Languages) LoadAll(opt *db.Options) error {
	if err := db.SelectStruct(constants.TableCoreConfigLanguages, l, opt); err != nil {
		return customerror.New(http.StatusInternalServerError, "languages load", err.Error())
	}
	return nil
}
