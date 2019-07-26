package content

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/agile-work/srv-shared/sql-builder/builder"

	"github.com/agile-work/srv-mdl-shared/models/customerror"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-shared/constants"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// Content defines a package grouping structures
type Content struct {
	ID          string                  `json:"id" sql:"id" pk:"true"`
	Code        string                  `json:"code" sql:"code" updatable:"false" validate:"required"`
	Prefix      string                  `json:"prefix" sql:"prefix" validate:"required"`
	Name        translation.Translation `json:"name" sql:"name" field:"jsonb" validate:"required"`
	Description translation.Translation `json:"description" sql:"description" field:"jsonb"`
	IsModule    bool                    `json:"is_module" sql:"is_module"`
	IsSystem    bool                    `json:"is_system" sql:"is_system"`
	Active      bool                    `json:"active" sql:"active"`
	CreatedBy   string                  `json:"created_by" sql:"created_by"`
	CreatedAt   time.Time               `json:"created_at" sql:"created_at"`
	UpdatedBy   string                  `json:"updated_by" sql:"updated_by"`
	UpdatedAt   time.Time               `json:"updated_at" sql:"updated_at"`
}

// Create persists the struct creating a new object in the database
func (c *Content) Create(trs *db.Transaction) error {
	id, err := db.InsertStructTx(trs.Tx, constants.TableCoreContents, c)
	if err != nil {
		return customerror.New(http.StatusInternalServerError, "content create", err.Error())
	}
	c.ID = id
	return nil
}

// Update updates object data in the database
func (c *Content) Update(trs *db.Transaction, columns []string, translations map[string]string) error {
	opt := &db.Options{Conditions: builder.Equal("code", c.Code)}

	if len(columns) > 0 {
		if err := db.UpdateStructTx(trs.Tx, constants.TableCoreContents, c, opt, strings.Join(columns, ",")); err != nil {
			return customerror.New(http.StatusInternalServerError, "content update", err.Error())
		}
	}

	if len(translations) > 0 {
		statement := builder.Update(constants.TableCoreContents)
		for col, val := range translations {
			statement.JSON(col, translation.FieldsRequestLanguageCode)
			jsonVal, _ := json.Marshal(val)
			statement.Values(jsonVal)
		}
		statement.Where(opt.Conditions)
		if _, err := trs.Query(statement); err != nil {
			return customerror.New(http.StatusInternalServerError, "content update", err.Error())
		}
	}

	return nil
}

// Load defines only one object from the database
func (c *Content) Load() error {
	if err := db.SelectStruct(constants.TableCoreContents, c, &db.Options{
		Conditions: builder.Equal("code", c.Code),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "content load", err.Error())
	}
	return nil
}

// Delete the object from the database
func (c *Content) Delete(trs *db.Transaction) error {
	if err := db.DeleteStructTx(trs.Tx, constants.TableCoreContents, &db.Options{
		Conditions: builder.Equal("code", c.Code),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "content create", err.Error())
	}
	return nil
}

// Contents slice of content
type Contents []Content

// LoadAll defines all instances from the object
func (c *Contents) LoadAll(opt *db.Options) error {
	if err := db.SelectStruct(constants.TableCoreContents, c, opt); err != nil {
		return customerror.New(http.StatusInternalServerError, "contents load", err.Error())
	}
	return nil
}
