package lookup

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	mdlShared "github.com/agile-work/srv-mdl-shared"
	"github.com/agile-work/srv-mdl-shared/models/customerror"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-shared/constants"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// Param defines the struct of a dynamic filter param
type Param struct {
	Code     string                  `json:"code"`
	DataType string                  `json:"data_type"`
	Label    translation.Translation `json:"label"`
	Type     string                  `json:"field_type,omitempty"`
	Pattern  string                  `json:"pattern,omitempty"`
	Security Security                `json:"security,omitempty"`
}

// Security defines the fields to set security to a field
type Security struct {
	SchemaCode string `json:"schema_code"`
	FieldCode  string `json:"field_code"`
}

// DynamicDefinition define specific fields for the lookup definition
type DynamicDefinition struct {
	Query     string    `json:"query"`
	Fields    []Param   `json:"fields"`
	Params    []Param   `json:"params,omitempty"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ParseQuery validate query and get fields and params from query
func (d *DynamicDefinition) ParseQuery(languageCode string) error {
	r := regexp.MustCompile("{{param:[^}}]*}}")
	params := r.FindAllString(d.Query, -1)
	params = unique(params)
	parsedQuery := d.Query

	for _, p := range params {
		param := Param{}
		fields := strings.Split(p[0:len(p)-2], ":")
		if len(fields) < 3 {
			return customerror.New(http.StatusBadRequest, "lookup parse query", "invalid query param")
		}

		if fields[1] == "security" {
			parsedQuery = strings.Replace(parsedQuery, p, "", -1)
			continue
		}

		param.Code = fields[1]
		if paramCodeExists(d.Params, param.Code) {
			return customerror.New(http.StatusBadRequest, "lookup parse query", "invalid query param, duplicated code "+param.Code)
		}
		param.Label.Language = make(map[string]string)
		param.Label.Language[languageCode] = param.Code
		param.DataType = fields[2]
		if len(fields) > 3 {
			param.Pattern = fields[3]
		}
		d.Params = append(d.Params, param)
		parsedQuery = strings.Replace(parsedQuery, p, "null", -1)
	}

	trs, err := db.NewTransaction()
	if err != nil {
		return customerror.New(http.StatusInternalServerError, "lookup parse query new transaction", err.Error())
	}

	// TODO: Tratar colunas jsonb
	query := fmt.Sprintf("CREATE TEMPORARY TABLE temp_table ON COMMIT DROP AS %s", parsedQuery)
	_, err = trs.Query(builder.Raw(query))
	if err != nil {
		trs.Rollback()
		return customerror.New(http.StatusInternalServerError, "lookup parse query create temporary table", err.Error())
	}
	query = "SELECT column_name as code, data_type FROM information_schema.columns WHERE table_name = 'temp_table'"
	rows, err := trs.Query(builder.Raw(query))
	if err != nil {
		trs.Rollback()
		return customerror.New(http.StatusInternalServerError, "lookup parse query select", err.Error())
	}

	err = db.StructScan(rows, &d.Fields)
	if err != nil {
		trs.Rollback()
		return customerror.New(http.StatusInternalServerError, "lookup parse query struct scan", err.Error())
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

// GetParamIndex returns the index of the param in the slice
func (d *DynamicDefinition) GetParamIndex(param Param) int {
	for i, p := range d.Params {
		if p.Code == param.Code {
			return i
		}
	}
	return -1
}

// ContainsField validate if dynamic definition fields contain a specific field
func (d *DynamicDefinition) ContainsField(field Param) bool {
	for _, f := range d.Fields {
		if f.Code == field.Code && f.DataType == field.DataType {
			return true
		}
	}
	return false
}

// ContainsParam validate if dynamic definition params contain a specific param and if the pattern has changed
func (d *DynamicDefinition) ContainsParam(param Param) int {
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

func (d *DynamicDefinition) parse(payload json.RawMessage) error {
	if err := json.Unmarshal(payload, d); err != nil {
		return customerror.New(http.StatusBadRequest, "lookup dynamic parse", err.Error())
	}

	if err := mdlShared.Validate.Struct(d); err != nil {
		return customerror.New(http.StatusBadRequest, "lookup dynamic invalid", err.Error())
	}
	return nil
}

// GetValueAndLabel returns the value and code columns og the lookup
func (d *DynamicDefinition) GetValueAndLabel() (string, string) {
	return "code", "label"
}

func (d *DynamicDefinition) getInstanceInformation(params map[string]interface{}) (string, string, []interface{}, error) {
	r := regexp.MustCompile("{{param:[^}}]*}}")
	paramsQuery := r.FindAllString(d.Query, -1)
	parsedQuery := d.Query
	values := []interface{}{}
	schema := ""

	for _, p := range paramsQuery {
		fields := strings.Split(p[0:len(p)-2], ":")
		if len(fields) < 3 {
			return "", "", nil, customerror.New(http.StatusBadRequest, "lookup parse query", "invalid query param")
		}

		if fields[1] == "security" {
			schema = fields[2]
			parsedQuery = strings.Replace(parsedQuery, p, "", -1)
			continue
		}

		code := fields[1]
		dataType := fields[2]
		pattern := ""
		if len(fields) > 3 {
			pattern = fields[3]
		}

		if v, ok := params[code]; ok {
			if pattern == "like" {
				v = strings.Replace(v.(string), "*", "%", -1)
			}
			values = append(values, v)
			parsedQuery = strings.Replace(parsedQuery, p, "?::"+dataType, 1)
		} else {
			parsedQuery = strings.Replace(parsedQuery, p, "NULL", 1)
		}
	}

	return schema, parsedQuery, values, nil
}

func (d *DynamicDefinition) getSecurityFields() map[string]map[string]string {
	result := map[string]map[string]string{}
	for _, f := range d.Fields {
		if f.Security.FieldCode != "" {
			column := map[string]string{}
			column[f.Code] = f.Security.FieldCode
			result[f.Security.SchemaCode] = column
		}
	}
	return result
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

func paramCodeExists(params []Param, code string) bool {
	for _, p := range params {
		if p.Code == code {
			return true
		}
	}
	return false
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
