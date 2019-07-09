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

// WorkflowInstance defines a new running instance for a workflow
type WorkflowInstance struct {
	ID           string    `json:"id" sql:"id" pk:"true"`
	WorkflowCode string    `json:"bpm_code" sql:"bpm_code"`
	Status       string    `json:"status" sql:"status" validate:"required"`
	Active       bool      `json:"active" sql:"active"`
	CreatedBy    string    `json:"created_by" sql:"created_by"`
	CreatedAt    time.Time `json:"created_at" sql:"created_at"`
	StartedAt    time.Time `json:"started_at" sql:"started_at"`
	EndAt        time.Time `json:"end_at" sql:"end_at"`
}

// Create persists the struct creating a new object in the database
func (w *WorkflowInstance) Create(trs *db.Transaction) error {
	id, err := db.InsertStructTx(trs.Tx, constants.TableCoreBPMInstances, w)
	if err != nil {
		return customerror.New(http.StatusInternalServerError, "workflow instance create", err.Error())
	}
	w.ID = id
	return nil
}

// Load defines only one object from the database
func (w *WorkflowInstance) Load() error {
	if err := db.SelectStruct(constants.TableCoreBPMInstances, w, &db.Options{
		Conditions: builder.Equal("id", w.ID),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "workflow instance load", err.Error())
	}
	return nil
}

// Update updates object data in the database
func (w *WorkflowInstance) Update(trs *db.Transaction, columns []string, translations map[string]string) error {
	opt := &db.Options{Conditions: builder.Equal("id", w.ID)}

	if len(columns) > 0 {
		if err := db.UpdateStructTx(trs.Tx, constants.TableCoreGroups, w, opt, strings.Join(columns, ",")); err != nil {
			return customerror.New(http.StatusInternalServerError, "workflow instance update", err.Error())
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
			return customerror.New(http.StatusInternalServerError, "workflow instance update", err.Error())
		}
	}

	return nil
}

// Delete deletes object from the database
func (w *WorkflowInstance) Delete(trs *db.Transaction) error {
	if err := db.DeleteStructTx(trs.Tx, constants.TableCoreGroups, &db.Options{
		Conditions: builder.Equal("code", w.ID),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "workflow instance delete", err.Error())
	}
	return nil
}

// WorkflowInstances defines a list of workflow instances
type WorkflowInstances []WorkflowInstance

// LoadAll defines all instances from the object
func (w *WorkflowInstances) LoadAll(opt *db.Options) error {
	if err := db.SelectStruct(constants.TableCoreBPMInstances, w, opt); err != nil {
		return customerror.New(http.StatusInternalServerError, "workflow instances load", err.Error())
	}
	return nil
}
