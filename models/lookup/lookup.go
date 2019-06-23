package lookup

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

// Definition interface to represent dynamic and static lookup definition
type Definition interface {
	parse(payload json.RawMessage) error
	GetValueAndLabel() (string, string)
}

// Lookups defines the array struct of this object
type Lookups []Lookup

// LoadAll defines all instances from the object
func (l *Lookups) LoadAll(opt *db.Options) error {
	if err := db.SelectStruct(constants.TableCoreLookups, l, opt); err != nil {
		return customerror.New(http.StatusInternalServerError, "lookups load", err.Error())
	}
	return nil
}

// Lookup defines the struct of this object
type Lookup struct {
	ID          string                  `json:"id" sql:"id" pk:"true"`
	Code        string                  `json:"code" sql:"code" updatable:"false" validate:"required"`
	Type        string                  `json:"type" sql:"type" updatable:"false" validate:"required"`
	Name        translation.Translation `json:"name" sql:"name" field:"jsonb" validate:"required"`
	Description translation.Translation `json:"description" sql:"description" field:"jsonb" validate:"required"`
	Definitions json.RawMessage         `json:"definitions" sql:"definitions" field:"jsonb" updatable:"false" validate:"required"`
	Active      bool                    `json:"active" sql:"active"`
	CreatedBy   string                  `json:"created_by" sql:"created_by"`
	CreatedAt   time.Time               `json:"created_at" sql:"created_at"`
	UpdatedBy   string                  `json:"updated_by" sql:"updated_by"`
	UpdatedAt   time.Time               `json:"updated_at" sql:"updated_at"`
}

// Create persists the struct creating a new object in the database
func (l *Lookup) Create(trs *db.Transaction, columns ...string) error {
	id, err := db.InsertStructTx(trs.Tx, constants.TableCoreLookups, l, columns...)
	if err != nil {
		return customerror.New(http.StatusInternalServerError, "lookup create", err.Error())
	}
	l.ID = id
	return nil
}

// Load defines only one object from the database
func (l *Lookup) Load() error {
	if err := db.SelectStruct(constants.TableCoreLookups, l, &db.Options{
		Conditions: builder.Equal("code", l.Code),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "lookup load", err.Error())
	}

	def, err := l.GetDefinition()
	if err != nil {
		return customerror.New(http.StatusInternalServerError, "lookup get definition", err.Error())
	}

	if err := l.setDefinition(def); err != nil {
		return customerror.New(http.StatusInternalServerError, "lookup set definition", err.Error())
	}

	return nil
}

// Update updates object data in the database
func (l *Lookup) Update(trs *db.Transaction, columns []string, translations map[string]string) error {
	opt := &db.Options{Conditions: builder.Equal("code", l.Code)}

	if len(columns) > 0 {
		if err := db.UpdateStructTx(trs.Tx, constants.TableCoreLookups, l, opt, strings.Join(columns, ",")); err != nil {
			return customerror.New(http.StatusInternalServerError, "lookup update", err.Error())
		}
	}

	if len(translations) > 0 {
		statement := builder.Update(constants.TableCoreLookups)
		for col, val := range translations {
			statement.JSON(col, translation.FieldsRequestLanguageCode)
			jsonVal, _ := json.Marshal(val)
			statement.Values(jsonVal)
		}
		statement.Where(opt.Conditions)
		if _, err := trs.Query(statement); err != nil {
			return customerror.New(http.StatusInternalServerError, "lookup update", err.Error())
		}
	}

	return nil
}

// Delete deletes object from the database
func (l *Lookup) Delete(trs *db.Transaction) error {
	if err := db.DeleteStructTx(trs.Tx, constants.TableCoreLookups, &db.Options{
		Conditions: builder.Equal("code", l.Code),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "lookup delete", err.Error())
	}
	return nil
}

// GetInstances returns lookup instances according to type
func (l *Lookup) GetInstances(params map[string]interface{}) ([]map[string]interface{}, error) {
	def, err := l.GetDefinition()
	if err != nil {
		return nil, customerror.New(http.StatusBadRequest, "GetBody read", err.Error())
	}

	results := []map[string]interface{}{}

	switch l.Type {
	case constants.LookupDynamic:
		lkpDynDef := def.(*DynamicDefinition)
		results, err = lkpDynDef.getInstances(params)
		if err != nil {
			return nil, customerror.New(http.StatusBadRequest, "lookup dynamic get instances", err.Error())
		}
	case constants.LookupStatic:
		lkpStaDef := def.(*StaticDefinition)
		results = lkpStaDef.getInstances()
	}

	return results, nil
}

// GetDefinition returns the definition of the lookup according to the type
func (l *Lookup) GetDefinition() (Definition, error) {
	switch l.Type {
	case constants.LookupStatic:
		def := &StaticDefinition{}
		err := def.parse(l.Definitions)
		return def, customerror.New(http.StatusBadRequest, "lookup dynamic get definition", err.Error())
	case constants.LookupDynamic:
		def := &DynamicDefinition{}
		err := def.parse(l.Definitions)
		return def, customerror.New(http.StatusBadRequest, "lookup static get definition", err.Error())
	}
	return nil, nil
}

// setDefinition defines the definition in lookup struct
func (l *Lookup) setDefinition(def Definition) error {
	defBytes, err := json.Marshal(def)
	if err != nil {
		return customerror.New(http.StatusBadRequest, "lookup set definition to byte", err.Error())
	}
	err = json.Unmarshal(defBytes, &l.Definitions)
	if err != nil {
		return customerror.New(http.StatusBadRequest, "lookup set definition parse", err.Error())
	}
	return nil
}

// ProcessDefinitions parse generic definition to a specific type processing necessary fields
func (l *Lookup) ProcessDefinitions(languageCode, method string) error {
	def, err := l.GetDefinition()
	if err != nil {
		return customerror.New(http.StatusBadRequest, "lookup process definition get definition", err.Error())
	}

	switch l.Type {
	case constants.LookupDynamic:
		lkpDynDef := def.(*DynamicDefinition)
		if method == http.MethodPost {
			lkpDynDef.CreatedBy = l.CreatedBy
			lkpDynDef.CreatedAt = l.CreatedAt
		}
		lkpDynDef.UpdatedBy = l.UpdatedBy
		lkpDynDef.UpdatedAt = l.UpdatedAt
		if err := lkpDynDef.ParseQuery(languageCode); err != nil {
			return customerror.New(http.StatusBadRequest, "lookup parse query", err.Error())
		}
	case constants.LookupStatic:
		lkpStaDef := def.(*StaticDefinition)
		for code, option := range lkpStaDef.Options {
			if method == http.MethodPost {
				option.CreatedBy = l.CreatedBy
				option.CreatedAt = l.CreatedAt
			}
			option.UpdatedBy = l.UpdatedBy
			option.UpdatedAt = l.UpdatedAt
			lkpStaDef.Options[code] = option
		}
	}
	translation.FieldsRequestLanguageCode = "all"
	if err := l.setDefinition(def); err != nil {
		return customerror.New(http.StatusBadRequest, "lookup process definition set definition", err.Error())
	}
	return nil
}
