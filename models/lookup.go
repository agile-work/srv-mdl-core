package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	sharedModels "github.com/agile-work/srv-mdl-shared/models"
	shared "github.com/agile-work/srv-shared"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	sql "github.com/agile-work/srv-shared/sql-builder/db"
)

// Lookup defines the struct of this object
type Lookup struct {
	ID          string                   `json:"id" sql:"id" pk:"true"`
	Code        string                   `json:"code" sql:"code" updatable:"false"`
	Type        string                   `json:"type" sql:"type" updatable:"false"`
	Name        sharedModels.Translation `json:"name" sql:"name" field:"jsonb"`
	Description sharedModels.Translation `json:"description" sql:"description" field:"jsonb"`
	Definitions json.RawMessage          `json:"definitions" sql:"definitions" field:"jsonb" updatable:"false"`
	Active      bool                     `json:"active" sql:"active"`
	CreatedBy   string                   `json:"created_by" sql:"created_by"`
	CreatedAt   time.Time                `json:"created_at" sql:"created_at"`
	UpdatedBy   string                   `json:"updated_by" sql:"updated_by"`
	UpdatedAt   time.Time                `json:"updated_at" sql:"updated_at"`
}

// ParseDefinition parse to a specific definition type for the lookup
func (l *Lookup) ParseDefinition() error {
	var def interface{}
	invalidType := false

	switch l.Type {
	case shared.LookupDynamic:
		def = &LookupDynamicDefinition{}
		err := json.Unmarshal(l.Definitions, def)
		if err != nil {
			return err
		}
	case shared.LookupStatic:
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
	case shared.LookupDynamic:
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
	case shared.LookupStatic:
		sharedModels.TranslationFieldsRequestLanguageCode = languageCode
		staticDef := LookupStaticDefinition{}
		err := json.Unmarshal(l.Definitions, &staticDef)
		if err != nil {
			return err
		}
		for i := range staticDef.Options {
			if method == http.MethodPost {
				staticDef.Options[i].CreatedBy = l.CreatedBy
				staticDef.Options[i].CreatedAt = l.CreatedAt
			}
			staticDef.Options[i].UpdatedBy = l.UpdatedBy
			staticDef.Options[i].UpdatedAt = l.UpdatedAt
		}
		sharedModels.TranslationFieldsRequestLanguageCode = "all"
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
		case shared.SQLDataTypeText:
			replaceParam = "''"
		case shared.SQLDataTypeDate:
			replaceParam = "current_date"
		case shared.SQLDataTypeNumber:
			replaceParam = "0"
		case shared.SQLDataTypeBool:
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
	}

	trs.Commit()
	return nil
}

// LookupParam defines the struct of a dynamic filter param
type LookupParam struct {
	Code     string                   `json:"code"`
	DataType string                   `json:"data_type"`
	Label    sharedModels.Translation `json:"label"`
	Pattern  string                   `json:"pattern,omitempty"`
	Security LookupSecurity           `json:"security,omitempty"`
}

// LookupSecurity defines the fields to set security to a field
type LookupSecurity struct {
	SchemaCode string `json:"schema_code"`
	FieldCode  string `json:"field_code"`
}

// LookupStaticDefinition define specific fields for the lookup definition
type LookupStaticDefinition struct {
	Options   []LookupOption `json:"options,omitempty"`
	OrderType string         `json:"order_type,omitempty"`
	Order     []string       `json:"order,omitempty"`
}

// LookupOption defines the struct of a static option
type LookupOption struct {
	Code      string                   `json:"code"`
	Label     sharedModels.Translation `json:"label,omitempty"`
	Active    bool                     `json:"active"`
	CreatedBy string                   `json:"created_by"`
	CreatedAt time.Time                `json:"created_at"`
	UpdatedBy string                   `json:"updated_by"`
	UpdatedAt time.Time                `json:"updated_at"`
}

func parseSQLType(sqlType string) string {
	if strings.Contains(sqlType, "timestamp") {
		return shared.SQLDataTypeDate
	}
	switch sqlType {
	case "character varying":
		return shared.SQLDataTypeText
	case "integer", "numeric":
		return shared.SQLDataTypeNumber
	case "boolean":
		return shared.SQLDataTypeBool
	default:
		return shared.SQLDataTypeText
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
