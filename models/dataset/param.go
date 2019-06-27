package dataset

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/agile-work/srv-mdl-shared/util"

	"github.com/agile-work/srv-mdl-shared/models/customerror"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-shared/constants"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/agile-work/srv-shared/sql-builder/db"
	sharedUtil "github.com/agile-work/srv-shared/util"
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

// Update updates a dataset param
func (p *Param) Update(trs *db.Transaction, datasetCode string, body map[string]interface{}, typeList string) error {
	total, err := db.Count("id", constants.TableCoreDatasets, &db.Options{
		Conditions: builder.Equal("code", datasetCode),
	})
	if err != nil || total == 0 {
		return customerror.New(http.StatusBadRequest, "validate dataset", "invalid dataset code")
	}

	cols := util.GetBodyColumns(body)
	languageCode := translation.FieldsRequestLanguageCode
	if sharedUtil.Contains(cols, "label") && languageCode != "all" {
		translation.FieldsRequestLanguageCode = "all"
		if _, err := trs.Query(builder.Raw(fmt.Sprintf(`update %s set definitions = jsonb_set(
			definitions,
			('{%s,'|| data_object.obj_index ||',label}') ::text[],
			'{"%s": "%s"}',
			true
			) from (
				select index-1 as obj_index from core_datasets ,jsonb_array_elements(definitions->'%s') with ordinality arr(obj, index)
				where ((obj->>'code') = '%s') and (code = '%s')
			)data_object
			where (code = '%s')`,
			constants.TableCoreDatasets,
			typeList,
			languageCode,
			p.Label.String(languageCode),
			typeList,
			p.Code,
			datasetCode,
			datasetCode,
		))); err != nil {
			return customerror.New(http.StatusInternalServerError, "Update", err.Error())
		}
	} else if sharedUtil.Contains(cols, "label") && languageCode == "all" {
		jsonBytes, _ := json.Marshal(p.Label)
		sqlQuery := getQueryUpdateField("label", string(jsonBytes), p.Code, datasetCode, typeList)
		if _, err := trs.Query(builder.Raw(sqlQuery)); err != nil {
			return customerror.New(http.StatusInternalServerError, "Update", err.Error())
		}
	}

	// TODO: Verificar se era p/ alterar o tipo do field
	if sharedUtil.Contains(cols, "field_type") && typeList == "fields" {
		jsonBytes, err := json.Marshal(p.Type)
		if err != nil {
			return customerror.New(http.StatusInternalServerError, "field_type parse", err.Error())
		}
		sqlQuery := getQueryUpdateField("field_type", string(jsonBytes), p.Code, datasetCode, typeList)
		if _, err := trs.Query(builder.Raw(sqlQuery)); err != nil {
			return customerror.New(http.StatusInternalServerError, "Update", err.Error())
		}
	}

	if sharedUtil.Contains(cols, "security") && typeList == "fields" {
		jsonBytes, err := json.Marshal(p.Security)
		if err != nil {
			return customerror.New(http.StatusInternalServerError, "security parse", err.Error())
		}
		if _, err := trs.Query(builder.Raw(fmt.Sprintf(`update %s set definitions = jsonb_set(
			definitions,
			('{fields,'|| data_object.obj_index ||',security}') ::text[],
			'%s',
			true
			) from (
				select index-1 as obj_index from core_datasets ,jsonb_array_elements(definitions->'fields') with ordinality arr(obj, index)
				where ((obj->>'code') = '%s') and (code = '%s')
			)data_object
			where (code = '%s')`,
			constants.TableCoreDatasets,
			string(jsonBytes),
			p.Code,
			datasetCode,
			datasetCode,
		))); err != nil {
			return customerror.New(http.StatusInternalServerError, "Update", err.Error())
		}
	}
	return nil
}

func paramCodeExists(params []Param, code string) bool {
	for _, p := range params {
		if p.Code == code {
			return true
		}
	}
	return false
}

func getQueryUpdateField(field, value, paramCode, datasetCode, typeList string) string {
	qry := fmt.Sprintf(`update %s set definitions = jsonb_set(
		definitions,
		('{%s,'|| data_object.obj_index ||'}') ::text[],
		'{"%s": %s}',
		true
		) from (
			select index-1 as obj_index from core_datasets ,jsonb_array_elements(definitions->'%s') with ordinality arr(obj, index)
			where ((obj->>'code') = '%s') and (code = '%s')
		)data_object
		where (code = '%s')`, constants.TableCoreDatasets, typeList, field, value, typeList, paramCode, datasetCode, datasetCode)
	fmt.Println(qry)
	return qry
}
