package field

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

// Definition defines a interface to represent a definition by type
type Definition interface {
	load(payload json.RawMessage) error
	validate() error
}

// Field defines the struct of this object
type Field struct {
	ID           string                  `json:"id" sql:"id" pk:"true"`
	Code         string                  `json:"code" sql:"code" updatable:"false" validate:"required"`
	SchemaCode   string                  `json:"schema_code" sql:"schema_code" updatable:"false"`
	Type         string                  `json:"field_type" sql:"field_type" updatable:"false" validate:"required"`
	Name         translation.Translation `json:"name" sql:"name" field:"jsonb" validate:"required"`
	Description  translation.Translation `json:"description" sql:"description" field:"jsonb"`
	DefaultValue json.RawMessage         `json:"default_value" sql:"default_value" field:"jsonb"`
	Definitions  json.RawMessage         `json:"definitions" sql:"definitions" field:"jsonb" updatable:"false" validate:"required"`
	Validations  json.RawMessage         `json:"validations" sql:"validations" field:"jsonb" updatable:"false"`
	Active       bool                    `json:"active" sql:"active"`
	CreatedBy    string                  `json:"created_by" sql:"created_by"`
	CreatedAt    time.Time               `json:"created_at" sql:"created_at"`
	UpdatedBy    string                  `json:"updated_by" sql:"updated_by"`
	UpdatedAt    time.Time               `json:"updated_at" sql:"updated_at"`
}

// Fields defines the array struct of this object
type Fields []Field

// Create persists the struct creating a new object in the database
func (f *Field) Create(trs *db.Transaction, columns ...string) error {
	def := f.getDefinition()
	if err := def.load(f.Definitions); err != nil {
		return customerror.New(http.StatusBadRequest, "load definition", err.Error())
	}
	if err := def.validate(); err != nil {
		return customerror.New(http.StatusBadRequest, "validating definition", err.Error())
	}
	translation.FieldsRequestLanguageCode = "all"
	f.setDefinition(def)
	id, err := db.InsertStructTx(trs.Tx, constants.TableCoreSchemaFields, f, columns...)
	if err != nil {
		return customerror.New(http.StatusInternalServerError, "field create", err.Error())
	}
	f.ID = id
	return nil
}

// LoadAll defines all instances from the object
func (f *Fields) LoadAll(opt *db.Options) error {
	if err := db.SelectStruct(constants.TableCoreSchemaFields, f, opt); err != nil {
		return customerror.New(http.StatusInternalServerError, "fields load", err.Error())
	}
	return nil
}

// Load defines only one object from the database
func (f *Field) Load() error {
	if err := db.SelectStruct(constants.TableCoreSchemaFields, f, &db.Options{Conditions: builder.And(
		builder.Equal("code", f.Code),
		builder.Equal("schema_code", f.SchemaCode),
	)}); err != nil {
		return customerror.New(http.StatusInternalServerError, "field load", err.Error())
	}

	if f.Type == constants.FieldLookup {
		// TODO: process lookup options
	}

	return nil
}

// Update updates object data in the database
func (f *Field) Update(trs *db.Transaction, columns []string, translations map[string]string) error {
	opt := &db.Options{Conditions: builder.And(
		builder.Equal("code", f.Code),
		builder.Equal("schema_code", f.SchemaCode),
	)}

	if len(columns) > 0 {
		if err := db.UpdateStructTx(trs.Tx, constants.TableCoreSchemaFields, f, opt, strings.Join(columns, ",")); err != nil {
			return customerror.New(http.StatusInternalServerError, "field update", err.Error())
		}
	}

	if len(translations) > 0 {
		statement := builder.Update(constants.TableCoreSchemaFields)
		for col, val := range translations {
			statement.JSON(col, translation.FieldsRequestLanguageCode)
			jsonVal, _ := json.Marshal(val)
			statement.Values(jsonVal)
		}
		statement.Where(opt.Conditions)
		if _, err := trs.Query(statement); err != nil {
			return customerror.New(http.StatusInternalServerError, "field update", err.Error())
		}
	}

	return nil
}

// Delete deletes object from the database
func (f *Field) Delete(trs *db.Transaction) error {
	if err := db.DeleteStructTx(trs.Tx, constants.TableCoreSchemaFields, &db.Options{Conditions: builder.And(
		builder.Equal("code", f.Code),
		builder.Equal("schema_code", f.SchemaCode),
	)}); err != nil {
		return customerror.New(http.StatusInternalServerError, "field delete", err.Error())
	}
	return nil
}

// AddFieldValidation include a new validation to a field
func (f *Field) AddFieldValidation(trs *db.Transaction) error {
	return nil
}

// UpdateFieldValidation update the validation attributes
func (f *Field) UpdateFieldValidation(trs *db.Transaction) error {
	return nil
}

// DeleteFieldValidation delete a validation from the database
func (f *Field) DeleteFieldValidation(trs *db.Transaction) error {
	return nil
}

// setDefinition defines the definition in field struct
func (f *Field) setDefinition(def Definition) {
	defBytes, _ := json.Marshal(def)
	json.Unmarshal(defBytes, &f.Definitions)
}

// getDefinition get the definition in field struct by type
func (f *Field) getDefinition() Definition {
	switch f.Type {
	case constants.FieldText:
		return &TextDefinition{}
	case constants.FieldNumber:
		return &NumberDefinition{}
	case constants.FieldDate:
		return &DateDefinition{}
	case constants.FieldLookup:
		return &LookupDefinition{}
	case constants.FieldAttachment:
		return &AttachmentDefinition{}
	default:
		return nil
	}
}
