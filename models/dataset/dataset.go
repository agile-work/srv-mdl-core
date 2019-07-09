package dataset

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	mdlShared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-mdl-shared/models/customerror"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-mdl-shared/models/user"
	"github.com/agile-work/srv-mdl-shared/util"
	"github.com/agile-work/srv-shared/constants"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// Definition interface to represent dynamic and static dataset definition
type Definition interface{}

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
	ContentCode string                  `json:"content_code" sql:"content_code"`
	Type        string                  `json:"type" sql:"type" updatable:"false" validate:"required"`
	Name        translation.Translation `json:"name" sql:"name" field:"jsonb" validate:"required"`
	Description translation.Translation `json:"description" sql:"description" field:"jsonb" validate:"required"`
	Definitions json.RawMessage         `json:"definitions" sql:"definitions" field:"jsonb" updatable:"false"` // TODO: estudar um jeito de validar levando em consideração o valor de outros campos
	Active      bool                    `json:"active" sql:"active"`
	CreatedBy   string                  `json:"created_by" sql:"created_by"`
	CreatedAt   time.Time               `json:"created_at" sql:"created_at"`
	UpdatedBy   string                  `json:"updated_by" sql:"updated_by"`
	UpdatedAt   time.Time               `json:"updated_at" sql:"updated_at"`
}

// Create persists the struct creating a new object in the database
func (ds *Dataset) Create(trs *db.Transaction, columns ...string) error {
	if ds.ContentCode != "" {
		prefix, err := util.GetContentPrefix(ds.ContentCode)
		if err != nil {
			return customerror.New(http.StatusInternalServerError, "dataset create", err.Error())
		}
		ds.Code = fmt.Sprintf("%s_%s", prefix, ds.Code)
	} else {
		ds.Code = fmt.Sprintf("%s_%s", "custom", ds.Code)
	}

	if len(ds.Code) > constants.DatabaseMaxLength {
		return customerror.New(http.StatusInternalServerError, "dataset create", "invalid code lenght")
	}

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

	if ds.ID == "" {
		return customerror.New(http.StatusNotFound, "dataset load", "not found")
	}

	def, err := ds.getDefinition()
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

	if ds.ContentCode != "" {
		if err := util.ValidateContent(ds.ContentCode); err != nil {
			return customerror.New(http.StatusInternalServerError, "dataset update", err.Error())
		}
	}

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

// GetUserInstances returns dataset instances according to type
func (ds *Dataset) GetUserInstances(username string, opt *db.Options, params map[string]interface{}) ([]map[string]interface{}, error) {
	def, err := ds.getDefinition()
	if err != nil {
		return nil, customerror.New(http.StatusBadRequest, "GetBody read", err.Error())
	}

	results := []map[string]interface{}{}

	switch ds.Type {
	case constants.DatasetDynamic:
		usr := &user.User{Username: username}
		if err := usr.Load(); err != nil {
			return nil, err
		}

		dsDynDef := def.(*DynamicDefinition)
		statement, err := dsDynDef.getQueryStatement(usr.LanguageCode, params)
		if err != nil {
			return nil, customerror.New(http.StatusBadRequest, "dataset dynamic get query", err.Error())
		}

		securitySchema := dsDynDef.getSecuritySchema()

		if securitySchema != "" {
			results, err = usr.GetSecurityInstances(securitySchema, opt, statement, dsDynDef.getSecurityFields())
			if err != nil {
				return nil, customerror.New(http.StatusInternalServerError, "dataset dynamic get intances", err.Error())
			}
		}

		rows, err := db.Query(statement)
		if err != nil {
			return nil, customerror.New(http.StatusInternalServerError, "dataset dynamic query exec", err.Error())
		}

		results, err = db.MapScan(rows)
		if err != nil {
			return nil, customerror.New(http.StatusInternalServerError, "dataset dynamic query map scan", err.Error())
		}
	case constants.DatasetStatic:
		dsStaDef := def.(*StaticDefinition)
		results = dsStaDef.getInstances()
	case constants.DatasetSchema:
		usr := &user.User{Username: username}
		if err := usr.Load(); err != nil {
			return nil, err
		}

		results, err = usr.GetSecurityInstances(ds.Code, opt, nil, nil)
		if err != nil {
			return nil, customerror.New(http.StatusInternalServerError, "dataset schema get intances", err.Error())
		}
	}

	return results, nil
}

// getDefinition returns the definition of the dataset according to the type
func (ds *Dataset) getDefinition() (Definition, error) {
	switch ds.Type {
	case constants.DatasetStatic:
		def := &StaticDefinition{}
		if err := parse(ds.Definitions, def); err != nil {
			return nil, customerror.New(http.StatusBadRequest, "dataset dynamic get definition", err.Error())
		}
		return def, nil
	case constants.DatasetDynamic:
		def := &DynamicDefinition{}
		if err := parse(ds.Definitions, def); err != nil {
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
	def, err := ds.getDefinition()
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
		if err := dsDynDef.parseQuery(languageCode); err != nil {
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

// Validate checks if dataset exists, if has the correct type
// and if has all the columns and params defined
func Validate(code string, isStaticDataset bool, columns, params []string) error {
	ds := Dataset{Code: code}
	if err := ds.Load(); err != nil {
		return err
	}

	if !ds.Active {
		return customerror.New(http.StatusBadRequest, "dataset validate", "deactivated dataset")
	}

	if isStaticDataset && ds.Type != constants.DatasetStatic {
		return customerror.New(http.StatusBadRequest, "dataset validate", "invalid type")
	}

	if isStaticDataset {
		return nil
	}

	invalidColumns := []string{}
	msgParams := ""

	switch ds.Type {
	case constants.DatasetDynamic:
		def := &DynamicDefinition{}
		if err := json.Unmarshal(ds.Definitions, def); err != nil {
			return err
		}

	LoopColumns:
		for _, col := range columns {
			for _, f := range def.Fields {
				if f.Code == col {
					continue LoopColumns
				}
			}
			invalidColumns = append(invalidColumns, col)
		}

		if params != nil && len(params) > 0 {
			invalidParams := []string{}
		LoopParams:
			for _, param := range params {
				for _, f := range def.Fields {
					if f.Code == param {
						continue LoopParams
					}
				}
				invalidParams = append(invalidParams, param)
			}
			if len(invalidParams) > 0 {
				msgParams = fmt.Sprintf("invalid params (%s)", strings.Join(invalidParams, ","))
			}
		}
	case constants.DatasetSchema:
		fields, err := getFieldsCodes(code)
		if err != nil {
			return err
		}
		for _, col := range columns {
			for _, fieldCode := range fields {
				if fieldCode == col {
					break
				}
			}
			invalidColumns = append(invalidColumns, col)
		}
	}

	if len(invalidColumns) > 0 {
		msg := fmt.Sprintf("invalid fields (%s) %s", strings.Join(invalidColumns, ","), msgParams)
		return customerror.New(http.StatusBadRequest, "dataset validate", msg)
	} else if msgParams != "" {
		return customerror.New(http.StatusBadRequest, "dataset validate", msgParams)
	}

	return nil
}

func getFieldsCodes(code string) ([]string, error) {
	codes := []string{}
	statement := builder.Select("code").From(constants.TableCoreSchemaFields).Where(builder.Equal("schema_code", code))
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	fieldCodes, _ := db.MapScan(rows)
	for _, row := range fieldCodes {
		for _, v := range row {
			codes = append(codes, v.(string))
		}
	}
	return codes, nil
}

func parse(payload json.RawMessage, def Definition) error {
	if err := json.Unmarshal(payload, def); err != nil {
		return err
	}

	if err := mdlShared.Validate.Struct(def); err != nil {
		return err
	}
	return nil
}
