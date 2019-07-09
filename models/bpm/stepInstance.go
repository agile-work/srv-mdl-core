package bpm

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

// StepInstance defines a new instance for a workflow step
type StepInstance struct {
	ID              string          `json:"id" sql:"id" pk:"true"`
	BPMStepCode     string          `json:"bpm_step_code" sql:"bpm_step_code" validate:"required"`
	BPMInstanceCode string          `json:"bpm_code" sql:"bpm_code" validate:"required"`
	Definitions     json.RawMessage `json:"definitions" sql:"definitions" field:"jsonb" updatable:"false" validate:"required"`
	Status          string          `json:"status" sql:"status" validate:"required"`
	StartedAt       time.Time       `json:"started_at" sql:"started_at"`
	EndAt           time.Time       `json:"end_at" sql:"end_at"`
}

// Create persists the struct creating a new object in the database
func (s *StepInstance) Create(trs *db.Transaction) error {
	id, err := db.InsertStructTx(trs.Tx, constants.TableCoreBPMStepInstances, s)
	if err != nil {
		return customerror.New(http.StatusInternalServerError, "step instance create", err.Error())
	}
	s.ID = id
	return nil
}

// Load defines only one object from the database
func (s *StepInstance) Load() error {
	if err := db.SelectStruct(constants.TableCoreBPMStepInstances, s, &db.Options{
		Conditions: builder.Equal("id", s.ID),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "step instance load", err.Error())
	}
	return nil
}

// Update updates object data in the database
func (s *StepInstance) Update(trs *db.Transaction, columns []string, translations map[string]string) error {
	opt := &db.Options{Conditions: builder.Equal("id", s.ID)}

	if len(columns) > 0 {
		if err := db.UpdateStructTx(trs.Tx, constants.TableCoreGroups, s, opt, strings.Join(columns, ",")); err != nil {
			return customerror.New(http.StatusInternalServerError, "step instance update", err.Error())
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
			return customerror.New(http.StatusInternalServerError, "step instance update", err.Error())
		}
	}

	return nil
}

// Delete deletes object from the database
func (s *StepInstance) Delete(trs *db.Transaction) error {
	if err := db.DeleteStructTx(trs.Tx, constants.TableCoreGroups, &db.Options{
		Conditions: builder.Equal("code", s.ID),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "step instance delete", err.Error())
	}
	return nil
}

// StepInstances defines a list of step instances
type StepInstances []StepInstance

// LoadAll defines all instances from the object
func (s *StepInstances) LoadAll(opt *db.Options) error {
	if err := db.SelectStruct(constants.TableCoreBPMStepInstances, s, opt); err != nil {
		return customerror.New(http.StatusInternalServerError, "step instances load", err.Error())
	}
	return nil
}
