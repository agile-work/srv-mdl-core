package dataset

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/agile-work/srv-mdl-shared/util"
	sharedUtil "github.com/agile-work/srv-shared/util"

	"github.com/agile-work/srv-mdl-shared/models/customerror"

	"github.com/agile-work/srv-shared/constants"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/agile-work/srv-shared/sql-builder/db"

	"github.com/agile-work/srv-mdl-shared/models/translation"
)

// StaticDefinition define specific fields for the dataset definition
type StaticDefinition struct {
	Options   map[string]Option `json:"options"`
	OrderType string            `json:"order_type,omitempty"`
	Order     []string          `json:"order,omitempty"`
}

// GetValueAndLabel returns the value and code columns og the dataset
func (d *StaticDefinition) GetValueAndLabel() (string, string) {
	return "code", "label"
}

func (d *StaticDefinition) getInstances() []map[string]interface{} {
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

// Option defines the struct of a static option
type Option struct {
	Code      string                  `json:"code"`
	Label     translation.Translation `json:"label"`
	Active    bool                    `json:"active"`
	CreatedBy string                  `json:"created_by"`
	CreatedAt time.Time               `json:"created_at"`
	UpdatedBy string                  `json:"updated_by"`
	UpdatedAt time.Time               `json:"updated_at"`
}

// Add inserts a new dataset option
func (o *Option) Add(trs *db.Transaction, datasetCode string) error {
	total, err := db.Count("id", constants.TableCoreLookups, &db.Options{
		Conditions: builder.Equal("code", datasetCode),
	})
	if err != nil || total == 0 {
		return customerror.New(http.StatusBadRequest, "validate dataset", "invalid dataset code")
	}

	total, err = db.Count(fmt.Sprintf("definitions->'options'->>'%s'", o.Code), constants.TableCoreLookups, &db.Options{
		Conditions: builder.Equal("code", datasetCode),
	})
	if err != nil {
		return customerror.New(http.StatusBadRequest, "validate dataset", err.Error())
	}

	if total > 0 {
		return customerror.New(http.StatusNotFound, "validate dataset", "code already exists")
	}

	optionBytes, _ := json.Marshal(o)
	sqlQuery := fmt.Sprintf(`update %s set definitions = jsonb_insert(
		definitions, '{options, %s}', '%s', true)
		where code = '%s';
		update %s set definitions = jsonb_set(
		definitions, '{order}', (definitions->'order') || '"%s"', true)
		where code = '%s';`,
		constants.TableCoreLookups,
		o.Code,
		string(optionBytes),
		datasetCode,
		constants.TableCoreLookups,
		o.Code,
		datasetCode)

	if _, err := trs.Query(builder.Raw(sqlQuery)); err != nil {
		return customerror.New(http.StatusInternalServerError, "Add", err.Error())
	}

	return nil
}

// Update updates a dataset option
func (o *Option) Update(trs *db.Transaction, datasetCode string, body map[string]interface{}) error {
	total, err := db.Count("id", constants.TableCoreLookups, &db.Options{
		Conditions: builder.Equal("code", datasetCode),
	})
	if err != nil || total == 0 {
		return customerror.New(http.StatusBadRequest, "validate dataset", "invalid dataset code")
	}

	cols := util.GetBodyColumns(body)
	languageCode := translation.FieldsRequestLanguageCode
	if sharedUtil.Contains(cols, "label") && languageCode != "all" {
		sqlQuery := fmt.Sprintf(`update %s set definitions = jsonb_set(
				definitions, '{options,%s,label,%s}', '"%s"', true
			)
			where code = '%s'`,
			constants.TableCoreLookups,
			o.Code,
			languageCode,
			o.Label.String(languageCode),
			datasetCode,
		)
		if _, err := trs.Query(builder.Raw(sqlQuery)); err != nil {
			return customerror.New(http.StatusInternalServerError, "Update", err.Error())
		}
	} else if sharedUtil.Contains(cols, "label") && languageCode == "all" {
		jsonBytes, err := json.Marshal(o.Label)
		if err != nil {
			return customerror.New(http.StatusBadRequest, "validate dataset options", err.Error())
		}
		sqlQuery := fmt.Sprintf(`update %s set definitions = jsonb_set(
				definitions, '{options,%s,label}', '%s', true
			)
			where code = '%s'`,
			constants.TableCoreLookups,
			o.Code,
			string(jsonBytes),
			datasetCode,
		)
		if _, err := trs.Query(builder.Raw(sqlQuery)); err != nil {
			return customerror.New(http.StatusInternalServerError, "Update", err.Error())
		}
	}

	// get fields from payload
	if sharedUtil.Contains(cols, "active") {
		sqlQuery := fmt.Sprintf(`update %s set definitions = jsonb_set(
			definitions, '{options,%s,active}', '%t', true
		)
		where code = '%s'`,
			constants.TableCoreLookups,
			o.Code,
			o.Active,
			datasetCode,
		)
		if _, err := trs.Query(builder.Raw(sqlQuery)); err != nil {
			return customerror.New(http.StatusInternalServerError, "Update", err.Error())
		}
	}
	return nil
}

// Delete delets a dataset option
func (o *Option) Delete(trs *db.Transaction, datasetCode string) error {
	total, err := db.Count("id", constants.TableCoreLookups, &db.Options{
		Conditions: builder.Equal("code", datasetCode),
	})
	if err != nil || total == 0 {
		return customerror.New(http.StatusBadRequest, "validate dataset", "invalid dataset code")
	}

	total, err = db.Count(fmt.Sprintf("definitions->'options'->>'%s'", o.Code), constants.TableCoreLookups, &db.Options{
		Conditions: builder.Equal("code", datasetCode),
	})
	if err != nil {
		return customerror.New(http.StatusBadRequest, "validate dataset", err.Error())
	}

	if total == 0 {
		return customerror.New(http.StatusNotFound, "validate dataset", "options not found")
	}

	sqlQuery := fmt.Sprintf(`update %s set definitions = jsonb_set(
		definitions, '{options}', (definitions->'options') - '%s', true)
		where code = '%s';
		update %s set definitions = jsonb_set(
		definitions, '{order}', (definitions->'order') - '%s', true)
		where code = '%s';`,
		constants.TableCoreLookups,
		o.Code,
		datasetCode,
		constants.TableCoreLookups,
		o.Code,
		datasetCode,
	)

	if _, err := trs.Query(builder.Raw(sqlQuery)); err != nil {
		return customerror.New(http.StatusInternalServerError, "Delete", err.Error())
	}

	return nil
}
