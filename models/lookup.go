package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-shared/constants"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	sql "github.com/agile-work/srv-shared/sql-builder/db"
)

// Lookup defines the struct of this object
type Lookup struct {
	ID          string                  `json:"id" sql:"id" pk:"true"`
	Code        string                  `json:"code" sql:"code" updatable:"false"`
	Type        string                  `json:"type" sql:"type" updatable:"false"`
	Name        translation.Translation `json:"name" sql:"name" field:"jsonb"`
	Description translation.Translation `json:"description" sql:"description" field:"jsonb"`
	Definitions json.RawMessage         `json:"definitions" sql:"definitions" field:"jsonb" updatable:"false"`
	Active      bool                    `json:"active" sql:"active"`
	CreatedBy   string                  `json:"created_by" sql:"created_by"`
	CreatedAt   time.Time               `json:"created_at" sql:"created_at"`
	UpdatedBy   string                  `json:"updated_by" sql:"updated_by"`
	UpdatedAt   time.Time               `json:"updated_at" sql:"updated_at"`
}

// GetDynamicDefinition returns the specific definition for a dynamic lookup
func (l *Lookup) GetDynamicDefinition() (*LookupDynamicDefinition, error) {
	if l.Type != constants.LookupDynamic {
		return nil, errors.New("invalid lookup type")
	}
	def := &LookupDynamicDefinition{}
	err := json.Unmarshal(l.Definitions, def)
	if err != nil {
		return nil, err
	}
	return def, nil
}

// GetStaticDefinition returns the specific definition for a static lookup
func (l *Lookup) GetStaticDefinition() (*LookupStaticDefinition, error) {
	if l.Type != constants.LookupStatic {
		return nil, errors.New("invalid lookup type")
	}
	def := &LookupStaticDefinition{}
	err := json.Unmarshal(l.Definitions, def)
	if err != nil {
		return nil, err
	}
	return def, nil
}

// ParseDefinition parse to a specific definition type for the lookup
func (l *Lookup) ParseDefinition() error {
	var def interface{}
	invalidType := false

	switch l.Type {
	case constants.LookupDynamic:
		def = &LookupDynamicDefinition{}
		err := json.Unmarshal(l.Definitions, def)
		if err != nil {
			return err
		}
	case constants.LookupStatic:
		def = &LookupStaticDefinition{}
		err := json.Unmarshal(l.Definitions, def)
		if err != nil {
			return err
		}
	default:
		invalidType = true
	}

	if invalidType {
		return errors.New("invalid lookup type")
	}

	jsonBytes, err := json.Marshal(def)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonBytes, &l.Definitions)
}

// ProcessDefinitions parse generic definition to a specific type processing necessary fields
func (l *Lookup) ProcessDefinitions(languageCode, method string) error {
	switch l.Type {
	case constants.LookupDynamic:
		dynamicDef := &LookupDynamicDefinition{}
		if method == http.MethodPost {
			dynamicDef.CreatedBy = l.CreatedBy
			dynamicDef.CreatedAt = l.CreatedAt
		}
		dynamicDef.UpdatedBy = l.UpdatedBy
		dynamicDef.UpdatedAt = l.UpdatedAt
		err := json.Unmarshal(l.Definitions, dynamicDef)
		if err != nil {
			return err
		}
		err = dynamicDef.ParseQuery(languageCode)
		if err != nil {
			return err
		}
		jsonBytes, err := json.Marshal(dynamicDef)
		if err != nil {
			return err
		}
		return json.Unmarshal(jsonBytes, &l.Definitions)
	case constants.LookupStatic:
		translation.FieldsRequestLanguageCode = languageCode
		staticDef := LookupStaticDefinition{}
		err := json.Unmarshal(l.Definitions, &staticDef)
		if err != nil {
			return err
		}
		for code, option := range staticDef.Options {
			if method == http.MethodPost {
				option.CreatedBy = l.CreatedBy
				option.CreatedAt = l.CreatedAt
				staticDef.Options[code] = option
			}
			option.UpdatedBy = l.UpdatedBy
			option.UpdatedAt = l.UpdatedAt
			staticDef.Options[code] = option
		}
		translation.FieldsRequestLanguageCode = "all"
		jsonBytes, err := json.Marshal(staticDef)
		if err != nil {
			return err
		}
		return json.Unmarshal(jsonBytes, &l.Definitions)
	}
	return errors.New("invalid lookup type")
}

// LookupDynamicDefinition define specific fields for the lookup definition
type LookupDynamicDefinition struct {
	Query     string        `json:"query"`
	Fields    []LookupParam `json:"fields"`
	Params    []LookupParam `json:"params"`
	CreatedBy string        `json:"created_by"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedBy string        `json:"updated_by"`
	UpdatedAt time.Time     `json:"updated_at"`
}

// GetValueAndLabel returns the code of the value and label fields in the slice
func (d *LookupDynamicDefinition) GetValueAndLabel() (string, string) {
	label := ""
	value := ""
	for _, f := range d.Fields {
		if f.Type == "label" {
			label = f.Code
		}
		if f.Type == "value" {
			value = f.Code
		}
	}
	return value, label
}

// GetParamIndex returns the index of the param in the slice
func (d *LookupDynamicDefinition) GetParamIndex(param LookupParam) int {
	for i, p := range d.Params {
		if p.Code == param.Code {
			return i
		}
	}
	return -1
}

// ContainsField validate if dynamic definition fields contain a specific field
func (d *LookupDynamicDefinition) ContainsField(field LookupParam) bool {
	for _, f := range d.Fields {
		if f.Code == field.Code && f.DataType == field.DataType {
			return true
		}
	}
	return false
}

// ContainsParam validate if dynamic definition params contain a specific param and if the pattern has changed
func (d *LookupDynamicDefinition) ContainsParam(param LookupParam) int {
	for _, f := range d.Params {
		if f.Code == param.Code && f.DataType == param.DataType {
			if f.Pattern == param.Pattern {
				return 0
			}
			return 1
		}
	}
	return -1
}

// ParseQuery validate query and get fields and params from query
func (d *LookupDynamicDefinition) ParseQuery(languageCode string) error {
	r := regexp.MustCompile("{{param:[^}}]*}}")
	params := r.FindAllString(d.Query, -1)
	params = unique(params)
	parsedQuery := d.Query

	for _, p := range params {
		param := LookupParam{}
		fields := strings.Split(p[0:len(p)-2], ":")
		if len(fields) < 3 {
			return errors.New("invalid query param")
		}

		if fields[1] == "security" {
			parsedQuery = strings.Replace(parsedQuery, p, "1 = 1", -1)
			continue
		}

		param.Code = fields[1]
		if paramCodeExists(d.Params, param.Code) {
			return errors.New("invalid query param, duplicated code " + param.Code)
		}
		param.Label.Language = make(map[string]string)
		param.Label.Language[languageCode] = param.Code
		param.DataType = fields[2]
		if len(fields) > 3 {
			param.Pattern = fields[3]
		}

		d.Params = append(d.Params, param)

		replaceParam := ""
		switch param.DataType {
		case constants.SQLDataTypeText:
			replaceParam = "''"
		case constants.SQLDataTypeDate:
			replaceParam = "current_date"
		case constants.SQLDataTypeNumber:
			replaceParam = "0"
		case constants.SQLDataTypeBool:
			replaceParam = "true"
		}
		parsedQuery = strings.Replace(parsedQuery, p, replaceParam, -1)
	}

	trs, err := sql.NewTransaction()
	if err != nil {
		return err
	}

	query := fmt.Sprintf("CREATE TEMPORARY TABLE temp_table ON COMMIT DROP AS %s", parsedQuery)
	_, err = trs.Query(builder.Raw(query))
	if err != nil {
		trs.Rollback()
		return err
	}
	query = "SELECT column_name as code, data_type FROM information_schema.columns WHERE table_name = 'temp_table'"
	rows, err := trs.Query(builder.Raw(query))
	if err != nil {
		trs.Rollback()
		return err
	}

	err = sql.StructScan(rows, &d.Fields)
	if err != nil {
		trs.Rollback()
		return err
	}
	for i, f := range d.Fields {
		d.Fields[i].DataType = parseSQLType(f.DataType)
		d.Fields[i].Label.Language = make(map[string]string)
		d.Fields[i].Label.Language[languageCode] = f.Code
		d.Fields[i].Type = "field"
		if f.Code == "id" {
			d.Fields[i].Type = "id"
		}
		if f.Code == "label" {
			d.Fields[i].Type = "label"
		}
	}

	trs.Commit()
	return nil
}

// LookupParam defines the struct of a dynamic filter param
type LookupParam struct {
	Code     string                  `json:"code"`
	DataType string                  `json:"data_type"`
	Label    translation.Translation `json:"label"`
	Type     string                  `json:"field_type,omitempty"`
	Pattern  string                  `json:"pattern,omitempty"`
	Security LookupSecurity          `json:"security,omitempty"`
}

// LookupSecurity defines the fields to set security to a field
type LookupSecurity struct {
	SchemaCode string `json:"schema_code"`
	FieldCode  string `json:"field_code"`
}

// LookupStaticDefinition define specific fields for the lookup definition
type LookupStaticDefinition struct {
	Options   map[string]LookupOption `json:"options,omitempty"`
	OrderType string                  `json:"order_type,omitempty"`
	Order     []string                `json:"order,omitempty"`
}

// LookupOption defines the struct of a static option
type LookupOption struct {
	Code      string                  `json:"code"`
	Label     translation.Translation `json:"label,omitempty"`
	Active    bool                    `json:"active"`
	CreatedBy string                  `json:"created_by"`
	CreatedAt time.Time               `json:"created_at"`
	UpdatedBy string                  `json:"updated_by"`
	UpdatedAt time.Time               `json:"updated_at"`
}

func parseSQLType(sqlType string) string {
	if strings.Contains(sqlType, "timestamp") {
		return constants.SQLDataTypeDate
	}
	switch sqlType {
	case "character varying":
		return constants.SQLDataTypeText
	case "integer", "numeric":
		return constants.SQLDataTypeNumber
	case "boolean":
		return constants.SQLDataTypeBool
	default:
		return constants.SQLDataTypeText
	}
}

func unique(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func paramCodeExists(params []LookupParam, code string) bool {
	for _, p := range params {
		if p.Code == code {
			return true
		}
	}
	return false
}

// GetInstances returns lookup instances according to type
func (l *Lookup) GetInstances() ([]map[string]interface{}, error) {
	results := []map[string]interface{}{}
	switch l.Type {
	case constants.LookupStatic:
		staticDefinition, err := l.GetStaticDefinition()
		if err != nil {
			return nil, err
		}
		results = staticDefinition.getInstances()
	case constants.LookupDynamic:
		dynamicDefinition, err := l.GetDynamicDefinition()
		if err != nil {
			return nil, err
		}
		results, err = dynamicDefinition.getInstances()
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}

func (d *LookupStaticDefinition) getInstances() []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, code := range d.Order {
		item := map[string]interface{}{}
		item["code"] = code
		if option, ok := d.Options[code]; ok && option.Active {
			item["label"] = option.Label
		}
		result = append(result, item)
	}
	return result
}

func (d *LookupDynamicDefinition) getInstances() ([]map[string]interface{}, error) {
	return nil, nil
}
