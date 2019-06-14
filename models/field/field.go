package field

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	sharedModels "github.com/agile-work/srv-mdl-shared/models"
	shared "github.com/agile-work/srv-shared"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

type Definition interface {
	load(payload json.RawMessage) error
}

// Field defines the struct of this object
type Field struct {
	ID           string                   `json:"id" sql:"id" pk:"true"`
	Code         string                   `json:"code" sql:"code" validate:"required"`
	SchemaCode   string                   `json:"schema_code" sql:"schema_code" updatable:"false"`
	Type         string                   `json:"field_type" sql:"field_type" updatable:"false" validate:"required"`
	Name         sharedModels.Translation `json:"name" sql:"name" field:"jsonb" validate:"required"`
	Description  sharedModels.Translation `json:"description" sql:"description" field:"jsonb"`
	DefaultValue json.RawMessage          `json:"default_value" sql:"default_value" field:"jsonb"`
	Definitions  json.RawMessage          `json:"definitions" sql:"definitions" field:"jsonb" updatable:"false" validate:"required"`
	Validations  json.RawMessage          `json:"validations" sql:"validations" field:"jsonb" updatable:"false"`
	Active       bool                     `json:"active" sql:"active"`
	CreatedBy    string                   `json:"created_by" sql:"created_by"`
	CreatedAt    time.Time                `json:"created_at" sql:"created_at"`
	UpdatedBy    string                   `json:"updated_by" sql:"updated_by"`
	UpdatedAt    time.Time                `json:"updated_at" sql:"updated_at"`
}

func (f *Field) Create(columns ...string) error {
	id, err := db.InsertStruct(shared.TableCoreSchemaFields, f, columns...)
	f.ID = id
	return err
}

func (f *Field) Update(columns []string, translations map[string]string) error {
	trs, err := db.NewTransaction()
	if err != nil {
		return err
	}
	conditions := builder.And(
		builder.Equal("code", f.Code),
		builder.Equal("schema_code", f.SchemaCode),
	)
	if len(columns) > 0 {
		trs.Add(db.StructUpdateStatement(shared.TableCoreSchemaFields, f, strings.Join(columns, ","), conditions))
	}

	if len(translations) > 0 {
		statement := builder.Update(shared.TableCoreSchemaFields)
		for col, val := range translations {
			statement.JSON(col, sharedModels.TranslationFieldsRequestLanguageCode)
			jsonVal, _ := json.Marshal(val)
			statement.Values(jsonVal)
		}
		statement.Where(conditions)
		trs.Add(statement)
	}

	return trs.Exec()
}

func (f *Field) Load() error {
	return db.SelectStruct(shared.TableCoreSchemaFields, f, builder.And(
		builder.Equal("code", f.Code),
		builder.Equal("schema_code", f.SchemaCode),
	))
}

func (f *Field) Delete() error {
	return db.DeleteStruct(shared.TableCoreSchemaFields, builder.And(
		builder.Equal("code", f.Code),
		builder.Equal("schema_code", f.SchemaCode),
	))
}

func (f *Field) SetDefinition(def Definition) {
	defBytes, _ := json.Marshal(def)
	json.Unmarshal(defBytes, &f.Definitions)
}

func (f *Field) GetDefinition() (Definition, error) {
	switch f.Type {
	case shared.FieldText:
		def := &TextDefinition{}
		err := def.load(f.Definitions)
		return def, err
	case shared.FieldNumber:
		def := &NumberDefinition{}
		err := def.load(f.Definitions)
		return def, err
	case shared.FieldDate:
		def := &DateDefinition{}
		err := def.load(f.Definitions)
		return def, err
	case shared.FieldLookup:
		def := &LookupDefinition{}
		err := def.load(f.Definitions)
		return def, err
	case shared.FieldAttachment:
		def := &AttachmentDefinition{}
		err := def.load(f.Definitions)
		return def, err
	default:
		return nil, errors.New("invalid field type")
	}
}

func LoadAll(schemaCode string) ([]Field, error) {
	fields := []Field{}
	err := db.SelectStruct(shared.TableCoreSchemaFields, &fields, builder.Equal("schema_code", schemaCode))
	return fields, err
}
