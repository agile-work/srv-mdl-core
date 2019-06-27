package dataset

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/agile-work/srv-mdl-shared/models/customerror"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-mdl-shared/models/user"
	"github.com/agile-work/srv-shared/constants"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// Definition interface to represent dynamic and static dataset definition
type Definition interface {
	parse(payload json.RawMessage) error
	GetValueAndLabel() (string, string)
}

// Datasets defines the array struct of this object
type Datasets []Dataset

// LoadAll defines all instances from the object
func (ds *Datasets) LoadAll(opt *db.Options) error {
	if err := db.SelectStruct(constants.TableCoreDatasets, ds, opt); err != nil {
		return customerror.New(http.StatusInternalServerError, "datasets load", err.Error())
	}
	return nil
}

// Dataset defines the struct of this object
type Dataset struct {
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
func (ds *Dataset) Create(trs *db.Transaction, columns ...string) error {
	id, err := db.InsertStructTx(trs.Tx, constants.TableCoreDatasets, ds, columns...)
	if err != nil {
		return customerror.New(http.StatusInternalServerError, "dataset create", err.Error())
	}
	ds.ID = id
	return nil
}

// Load defines only one object from the database
func (ds *Dataset) Load() error {
	if err := db.SelectStruct(constants.TableCoreDatasets, ds, &db.Options{
		Conditions: builder.Equal("code", ds.Code),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "dataset load", err.Error())
	}

	def, err := ds.GetDefinition()
	if err != nil {
		return customerror.New(http.StatusInternalServerError, "dataset get definition", err.Error())
	}

	if err := ds.setDefinition(def); err != nil {
		return customerror.New(http.StatusInternalServerError, "dataset set definition", err.Error())
	}

	return nil
}

// Update updates object data in the database
func (ds *Dataset) Update(trs *db.Transaction, columns []string, translations map[string]string) error {
	opt := &db.Options{Conditions: builder.Equal("code", ds.Code)}

	if len(columns) > 0 {
		if err := db.UpdateStructTx(trs.Tx, constants.TableCoreDatasets, ds, opt, strings.Join(columns, ",")); err != nil {
			return customerror.New(http.StatusInternalServerError, "dataset update", err.Error())
		}
	}

	if len(translations) > 0 {
		statement := builder.Update(constants.TableCoreDatasets)
		for col, val := range translations {
			statement.JSON(col, translation.FieldsRequestLanguageCode)
			jsonVal, _ := json.Marshal(val)
			statement.Values(jsonVal)
		}
		statement.Where(opt.Conditions)
		if _, err := trs.Query(statement); err != nil {
			return customerror.New(http.StatusInternalServerError, "dataset update", err.Error())
		}
	}

	return nil
}

// Delete deletes object from the database
func (ds *Dataset) Delete(trs *db.Transaction) error {
	if err := db.DeleteStructTx(trs.Tx, constants.TableCoreDatasets, &db.Options{
		Conditions: builder.Equal("code", ds.Code),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "dataset delete", err.Error())
	}
	return nil
}

// GetInstances returns dataset instances according to type
func (ds *Dataset) GetInstances(params map[string]interface{}, usr *user.User, opt *db.Options) ([]map[string]interface{}, error) {
	def, err := ds.GetDefinition()
	if err != nil {
		return nil, customerror.New(http.StatusBadRequest, "GetBody read", err.Error())
	}

	results := []map[string]interface{}{}

	switch ds.Type {
	case constants.DatasetDynamic:
		dsDynDef := def.(*DynamicDefinition)
		schema, dsQuery, values, err := dsDynDef.getInstanceInformation(params)
		if err != nil {
			return nil, customerror.New(http.StatusBadRequest, "dataset dynamic get instances", err.Error())
		}

		statement, err := usr.GetSecurityQueryWithSub(schema, dsQuery, opt)
		if err != nil {
			return nil, customerror.New(http.StatusInternalServerError, "dataset dynamic get security query", err.Error())
		}

		query, _ := statement.Query()

		rows, err := db.Query(builder.Raw(query, values...))
		if err != nil {
			return nil, customerror.New(http.StatusInternalServerError, "dataset dynamic exec query", err.Error())
		}

		results, err = usr.SecurityMapScanWithFields(schema, rows, opt, dsDynDef.getSecurityFields())
		if err != nil {
			return nil, customerror.New(http.StatusInternalServerError, "dataset dynamic scan", err.Error())
		}
	case constants.DatasetStatic:
		dsStaDef := def.(*StaticDefinition)
		results = dsStaDef.getInstances()
	}

	return results, nil
}

// GetDefinition returns the definition of the dataset according to the type
func (ds *Dataset) GetDefinition() (Definition, error) {
	switch ds.Type {
	case constants.DatasetStatic:
		def := &StaticDefinition{}
		if err := def.parse(ds.Definitions); err != nil {
			return nil, customerror.New(http.StatusBadRequest, "dataset dynamic get definition", err.Error())
		}
		return def, nil
	case constants.DatasetDynamic:
		def := &DynamicDefinition{}
		if err := def.parse(ds.Definitions); err != nil {
			return nil, customerror.New(http.StatusBadRequest, "dataset static get definition", err.Error())
		}
		return def, nil
	}
	return nil, nil
}

// setDefinition defines the definition in dataset struct
func (ds *Dataset) setDefinition(def Definition) error {
	defBytes, err := json.Marshal(def)
	if err != nil {
		return customerror.New(http.StatusBadRequest, "dataset set definition to byte", err.Error())
	}
	err = json.Unmarshal(defBytes, &ds.Definitions)
	if err != nil {
		return customerror.New(http.StatusBadRequest, "dataset set definition parse", err.Error())
	}
	return nil
}

// ProcessDefinitions parse generic definition to a specific type processing necessary fields
func (ds *Dataset) ProcessDefinitions(languageCode, method string) error {
	def, err := ds.GetDefinition()
	if err != nil {
		return customerror.New(http.StatusBadRequest, "dataset process definition get definition", err.Error())
	}

	switch ds.Type {
	case constants.DatasetDynamic:
		dsDynDef := def.(*DynamicDefinition)
		if method == http.MethodPost {
			dsDynDef.CreatedBy = ds.CreatedBy
			dsDynDef.CreatedAt = ds.CreatedAt
		}
		dsDynDef.UpdatedBy = ds.UpdatedBy
		dsDynDef.UpdatedAt = ds.UpdatedAt
		if err := dsDynDef.ParseQuery(languageCode); err != nil {
			return customerror.New(http.StatusBadRequest, "dataset parse query", err.Error())
		}
	case constants.DatasetStatic:
		dsStaDef := def.(*StaticDefinition)
		for code, option := range dsStaDef.Options {
			if method == http.MethodPost {
				option.CreatedBy = ds.CreatedBy
				option.CreatedAt = ds.CreatedAt
			}
			option.UpdatedBy = ds.UpdatedBy
			option.UpdatedAt = ds.UpdatedAt
			dsStaDef.Options[code] = option
		}
	}
	translation.FieldsRequestLanguageCode = "all"
	if err := ds.setDefinition(def); err != nil {
		return customerror.New(http.StatusBadRequest, "dataset process definition set definition", err.Error())
	}
	return nil
}
