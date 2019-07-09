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

// Step defines a new business process on the database
type Step struct {
	ID           string                  `json:"id" sql:"id" pk:"true"`
	Code         string                  `json:"code" sql:"code" validate:"required"`
	WorkflowCode string                  `json:"bpm_code" sql:"bpm_code" validate:"required"`
	ContentCode  string                  `json:"content_code" sql:"content_code"`
	SchemaCode   string                  `json:"schema_code" sql:"schema_code"`
	Name         translation.Translation `json:"name" sql:"name" field:"jsonb" validate:"required"`
	Description  translation.Translation `json:"description" sql:"description" field:"jsonb" validate:"required"`
	Definitions  json.RawMessage         `json:"definitions" sql:"definitions" field:"jsonb" updatable:"false" validate:"required"`
	Active       bool                    `json:"active" sql:"active"`
	CreatedBy    string                  `json:"created_by" sql:"created_by"`
	CreatedAt    time.Time               `json:"created_at" sql:"created_at"`
	UpdatedBy    string                  `json:"updated_by" sql:"updated_by"`
	UpdatedAt    time.Time               `json:"updated_at" sql:"updated_at"`
}

// Create persists the struct creating a new object in the database
func (s *Step) Create(trs *db.Transaction, columns ...string) error {
	if s.ContentCode != "" {
		prefix, err := util.GetContentPrefix(s.ContentCode)
		if err != nil {
			return customerror.New(http.StatusInternalServerError, "step create", err.Error())
		}
		s.Code = fmt.Sprintf("%s_%s", prefix, s.Code)
	} else {
		s.Code = fmt.Sprintf("%s_%s", "custom", s.Code)
	}

	if len(s.Code) > constants.DatabaseMaxLength {
		return customerror.New(http.StatusInternalServerError, "step create", "invalid code lenght")
	}

	id, err := db.InsertStructTx(trs.Tx, constants.TableCoreBPMSteps, s, columns...)
	if err != nil {
		return customerror.New(http.StatusInternalServerError, "step create", err.Error())
	}
	s.ID = id

	return nil
}

// Load defines only one object from the database
func (s *Step) Load() error {
	if err := db.SelectStruct(constants.TableCoreBPMSteps, s, &db.Options{
		Conditions: builder.Equal("code", s.Code),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "step load", err.Error())
	}
	return nil
}

// Update updates object data in the database
func (s *Step) Update(trs *db.Transaction, columns []string, translations map[string]string) error {
	opt := &db.Options{Conditions: builder.Equal("code", s.Code)}

	if s.ContentCode != "" {
		if err := util.ValidateContent(s.ContentCode); err != nil {
			return customerror.New(http.StatusInternalServerError, "step update", err.Error())
		}
	}

	if len(columns) > 0 {
		if err := db.UpdateStructTx(trs.Tx, constants.TableCoreBPMSteps, s, opt, strings.Join(columns, ",")); err != nil {
			return customerror.New(http.StatusInternalServerError, "step update", err.Error())
		}
	}

	if len(translations) > 0 {
		statement := builder.Update(constants.TableCoreBPMSteps)
		for col, val := range translations {
			statement.JSON(col, translation.FieldsRequestLanguageCode)
			jsonVal, _ := json.Marshal(val)
			statement.Values(jsonVal)
		}
		statement.Where(opt.Conditions)
		if _, err := trs.Query(statement); err != nil {
			return customerror.New(http.StatusInternalServerError, "step update", err.Error())
		}
	}

	return nil
}

// Delete deletes object from the database
func (s *Step) Delete(trs *db.Transaction) error {
	if err := db.DeleteStructTx(trs.Tx, constants.TableCoreBPMSteps, &db.Options{
		Conditions: builder.Equal("code", s.Code),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "step delete", err.Error())
	}
	return nil
}

// Steps defines the array struct of this object
type Steps []Step

// LoadAll defines all instances from the object
func (s *Steps) LoadAll(opt *db.Options) error {
	if err := db.SelectStruct(constants.TableCoreBPMSteps, s, opt); err != nil {
		return customerror.New(http.StatusInternalServerError, "steps load", err.Error())
	}
	return nil
}
