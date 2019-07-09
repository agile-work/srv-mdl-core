package bpm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/agile-work/srv-mdl-shared/models/customerror"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-mdl-shared/util"
	"github.com/agile-work/srv-shared/constants"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// Workflow defines a new bussiness process on the database
type Workflow struct {
	ID          string                  `json:"id" sql:"id" pk:"true"`
	Code        string                  `json:"code" sql:"code"`
	ContentCode string                  `json:"content_code" sql:"content_code"`
	SchemaCode  string                  `json:"schema_code" sql:"schema_code"`
	Name        translation.Translation `json:"name" sql:"name" field:"jsonb" validate:"required"`
	Description translation.Translation `json:"description" sql:"description" field:"jsonb" validate:"required"`
	Params      map[string]interface{}  `json:"params" sql:"params" field:"jsonb"`
	Active      bool                    `json:"active" sql:"active"`
	CreatedBy   string                  `json:"created_by" sql:"created_by"`
	CreatedAt   time.Time               `json:"created_at" sql:"created_at"`
	UpdatedBy   string                  `json:"updated_by" sql:"updated_by"`
	UpdatedAt   time.Time               `json:"updated_at" sql:"updated_at"`
}

// Create persists the struct creating a new object in the database
func (w *Workflow) Create(trs *db.Transaction) error {
	if w.ContentCode != "" {
		prefix, err := util.GetContentPrefix(w.ContentCode)
		if err != nil {
			return customerror.New(http.StatusInternalServerError, "workflow create", err.Error())
		}
		w.Code = fmt.Sprintf("%s_%s", prefix, w.Code)
	} else {
		w.Code = fmt.Sprintf("%s_%s", "custom", w.Code)
	}

	if len(w.Code) > constants.DatabaseMaxLength {
		return customerror.New(http.StatusInternalServerError, "workflow create", "invalid code lenght")
	}

	id, err := db.InsertStructTx(trs.Tx, constants.TableCoreBPM, w)
	if err != nil {
		return customerror.New(http.StatusInternalServerError, "workflow create", err.Error())
	}
	w.ID = id

	return nil
}

// Load defines only one object from the database
func (w *Workflow) Load() error {
	if err := db.SelectStruct(constants.TableCoreBPM, w, &db.Options{
		Conditions: builder.Equal("code", w.Code),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "workflow load", err.Error())
	}
	return nil
}

// Update updates object data in the database
func (w *Workflow) Update(trs *db.Transaction, columns []string, translations map[string]string) error {
	opt := &db.Options{Conditions: builder.Equal("code", w.Code)}

	if w.ContentCode != "" {
		if err := util.ValidateContent(w.ContentCode); err != nil {
			return customerror.New(http.StatusInternalServerError, "workflow update", err.Error())
		}
	}

	if len(columns) > 0 {
		if err := db.UpdateStructTx(trs.Tx, constants.TableCoreBPM, w, opt, strings.Join(columns, ",")); err != nil {
			return customerror.New(http.StatusInternalServerError, "workflow update", err.Error())
		}
	}

	if len(translations) > 0 {
		statement := builder.Update(constants.TableCoreBPM)
		for col, val := range translations {
			statement.JSON(col, translation.FieldsRequestLanguageCode)
			jsonVal, _ := json.Marshal(val)
			statement.Values(jsonVal)
		}
		statement.Where(opt.Conditions)
		if _, err := trs.Query(statement); err != nil {
			return customerror.New(http.StatusInternalServerError, "workflow update", err.Error())
		}
	}

	return nil
}

// Delete deletes object from the database
func (w *Workflow) Delete(trs *db.Transaction) error {
	if err := db.DeleteStructTx(trs.Tx, constants.TableCoreBPM, &db.Options{
		Conditions: builder.Equal("code", w.Code),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "workflow delete", err.Error())
	}
	return nil
}

// Workflows defines the array struct of this object
type Workflows []Workflow

// LoadAll defines all instances from the object
func (w *Workflows) LoadAll(opt *db.Options) error {
	if err := db.SelectStruct(constants.TableCoreBPM, w, opt); err != nil {
		return customerror.New(http.StatusInternalServerError, "workflows load", err.Error())
	}
	return nil
}
