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

// Tree defines the struct of this object
type Tree struct {
	ID          string                  `json:"id" sql:"id" pk:"true"`
	Code        string                  `json:"code" sql:"code" updatable:"false" validate:"required"`
	Name        translation.Translation `json:"name" sql:"name" field:"jsonb" validate:"required"`
	Description translation.Translation `json:"description" sql:"description" field:"jsonb" validate:"required"`
	Active      bool                    `json:"active" sql:"active"`
	CreatedBy   string                  `json:"created_by" sql:"created_by"`
	CreatedAt   time.Time               `json:"created_at" sql:"created_at"`
	UpdatedBy   string                  `json:"updated_by" sql:"updated_by"`
	UpdatedAt   time.Time               `json:"updated_at" sql:"updated_at"`
}

// Trees defines the array struct of this object
type Trees []Tree

// Create persists the struct creating a new object in the database
func (t *Tree) Create(trs *db.Transaction, columns ...string) error {
	id, err := db.InsertStructTx(trs.Tx, constants.TableCoreTrees, t, columns...)
	if err != nil {
		customerror.New(http.StatusInternalServerError, "tree create", err.Error())
	}
	t.ID = id
	return nil
}

// LoadAll defines all instances from the object
func (t *Trees) LoadAll(opt *db.Options) error {
	if err := db.SelectStruct(constants.TableCoreTrees, t, opt); err != nil {
		return customerror.New(http.StatusInternalServerError, "trees load", err.Error())
	}
	return nil
}

// Load defines only one object from the database
func (t *Tree) Load() error {
	if err := db.SelectStruct(constants.TableCoreTrees, t, &db.Options{
		Conditions: builder.Equal("code", t.Code),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "tree load", err.Error())
	}
	return nil
}

// Update updates object data in the database
func (t *Tree) Update(trs *db.Transaction, columns []string, translations map[string]string) error {
	opt := &db.Options{Conditions: builder.Equal("code", t.Code)}

	if len(columns) > 0 {
		if err := db.UpdateStructTx(trs.Tx, constants.TableCoreTrees, t, opt, strings.Join(columns, ",")); err != nil {
			return customerror.New(http.StatusInternalServerError, "tree update", err.Error())
		}
	}

	if len(translations) > 0 {
		statement := builder.Update(constants.TableCoreTrees)
		for col, val := range translations {
			statement.JSON(col, translation.FieldsRequestLanguageCode)
			jsonVal, _ := json.Marshal(val)
			statement.Values(jsonVal)
		}
		statement.Where(opt.Conditions)
		if _, err := trs.Query(statement); err != nil {
			return customerror.New(http.StatusInternalServerError, "tree update", err.Error())
		}
	}

	return nil
}

// Delete deletes object from the database
func (t *Tree) Delete(trs *db.Transaction) error {
	if err := db.DeleteStructTx(trs.Tx, constants.TableCoreTrees, &db.Options{
		Conditions: builder.Equal("code", t.Code),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "tree delete", err.Error())
	}
	return nil
}
