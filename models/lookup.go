package models

import (
	"encoding/json"
	"errors"
	"fmt"
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
	Type        string                   `json:"type" sql:"type"`
	Name        sharedModels.Translation `json:"name" sql:"name" field:"jsonb"`
	Description sharedModels.Translation `json:"description" sql:"description" field:"jsonb"`
	Definitions json.RawMessage          `json:"definitions" sql:"definitions" field:"jsonb"`
	Active      bool                     `json:"active" sql:"active"`
	CreatedBy   string                   `json:"created_by" sql:"created_by"`
	CreatedAt   time.Time                `json:"created_at" sql:"created_at"`
	UpdatedBy   string                   `json:"updated_by" sql:"updated_by"`
	UpdatedAt   time.Time                `json:"updated_at" sql:"updated_at"`
}

// ProcessDefinitions parse generic definition to a specific type processing necessary fields
func (l *Lookup) ProcessDefinitions(languageCode string) error {
	switch l.Type {
	case shared.LookupDynamic:
		dynamicDef := &LookupDynamicDefinition{}
		err := json.Unmarshal(l.Definitions, dynamicDef)
		if err != nil {
			return err
		}
		err = dynamicDef.parseQuery(languageCode)
		if err != nil {
			return err
		}
		jsonBytes, err := json.Marshal(dynamicDef)
		if err != nil {
			return err
		}
		return json.Unmarshal(jsonBytes, &l.Definitions)
	case shared.LookupStatic:
		staticDef := LookupStaticDefinition{}
		err := json.Unmarshal(l.Definitions, &staticDef)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("invalid lookup type")
}

// LookupDynamicDefinition define specific fields for the lookup definition
type LookupDynamicDefinition struct {
	Query     string        `json:"query"`
	Fields    []LookupParam `json:"field"`
	Params    []LookupParam `json:"params"`
	CreatedBy string        `json:"created_by"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedBy string        `json:"updated_by"`
	UpdatedAt time.Time     `json:"updated_at"`
}

func (d *LookupDynamicDefinition) parseQuery(languageCode string) error {
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
		param.Code = fields[1]
		if paramCodeExists(d.Params, param.Code) {
			return errors.New("invalid query param, duplicated code " + param.Code)
		}
		param.Label = make(map[string]string)
		param.Label[languageCode] = param.Code
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
		d.Fields[i].Label = make(map[string]string)
		d.Fields[i].Label[languageCode] = f.Code
	}

	trs.Commit()
	return nil
}

// LookupParam defines the struct of a dynamic filter param
type LookupParam struct {
	Code     string            `json:"code"`
	DataType string            `json:"data_type"`
	Label    map[string]string `json:"label"`
	Pattern  string            `json:"pattern,omitempty"`
}

// LookupStaticDefinition define specific fields for the lookup definition
type LookupStaticDefinition struct {
	Options []LookupOption `json:"options"`
}

// LookupOption defines the struct of a static option
type LookupOption struct {
	Code      string            `json:"code"`
	Label     map[string]string `json:"label"`
	Active    bool              `json:"active"`
	CreatedBy string            `json:"created_by"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedBy string            `json:"updated_by"`
	UpdatedAt time.Time         `json:"updated_at"`
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
