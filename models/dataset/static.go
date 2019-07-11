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
	total, err := db.Count("id", constants.TableCoreDatasets, &db.Options{
		Conditions: builder.Equal("code", datasetCode),
	})
	if err != nil || total == 0 {
		return customerror.New(http.StatusBadRequest, "validate dataset", "invalid dataset code")
	}

	total, err = db.Count(fmt.Sprintf("definitions->'options'->>'%s'", o.Code), constants.TableCoreDatasets, &db.Options{
		Conditions: builder.Equal("code", datasetCode),
	})
	if err != nil {
		return customerror.New(http.StatusBadRequest, "validate dataset", err.Error())
	}

	if total > 0 {
		return customerror.New(http.StatusNotFound, "validate dataset", "code already exists")
	}

	optionBytes, _ := json.Marshal(o)
	rows, err := trs.Query(builder.Update(
		constants.TableCoreDatasets,
	).InsertJSON(
		"definitions",
		fmt.Sprintf("'{options, %s}'", o.Code),
		fmt.Sprintf("'%s'", optionBytes),
		true,
	).Where(
		builder.Equal("code", datasetCode),
	))
	if err != nil {
		return customerror.New(http.StatusInternalServerError, "Add", err.Error())
	}
	rows.Close()

	rows, err = trs.Query(builder.Update(
		constants.TableCoreDatasets,
	).UpdateJSON(
		"definitions",
		"'{order}'",
		fmt.Sprintf(`(definitions->'order') || '"%s"'`, o.Code),
		true,
	).Where(
		builder.Equal("code", datasetCode),
	))
	if err != nil {
		return customerror.New(http.StatusInternalServerError, "Add", err.Error())
	}
	rows.Close()

	return nil
}

// Update updates a dataset option
func (o *Option) Update(trs *db.Transaction, datasetCode string, body map[string]interface{}) error {
	total, err := db.Count("id", constants.TableCoreDatasets, &db.Options{
		Conditions: builder.Equal("code", datasetCode),
	})
	if err != nil || total == 0 {
		return customerror.New(http.StatusBadRequest, "validate dataset", "invalid dataset code")
	}

	cols := util.GetBodyColumns(body)
	languageCode := translation.FieldsRequestLanguageCode
	if sharedUtil.Contains(cols, "label") && languageCode != "all" {
		_, err := trs.Query(builder.Update(
			constants.TableCoreDatasets,
		).UpdateJSON(
			"definitions",
			fmt.Sprintf("'{options,%s,label,%s}'", o.Code, languageCode),
			fmt.Sprintf(`'"%s"'`, o.Label.String(languageCode)),
			true,
		).Where(
			builder.Equal("code", datasetCode),
		))
		if err != nil {
			return customerror.New(http.StatusInternalServerError, "Update", err.Error())
		}
	} else if sharedUtil.Contains(cols, "label") && languageCode == "all" {
		jsonBytes, err := json.Marshal(o.Label)
		if err != nil {
			return customerror.New(http.StatusBadRequest, "validate dataset options", err.Error())
		}
		_, err = trs.Query(builder.Update(
			constants.TableCoreDatasets,
		).UpdateJSON(
			"definitions",
			fmt.Sprintf("'{options,%s,label,%s}'", o.Code, languageCode),
			fmt.Sprintf(`'%s'`, jsonBytes),
			true,
		).Where(
			builder.Equal("code", datasetCode),
		))
		if err != nil {
			return customerror.New(http.StatusInternalServerError, "Update", err.Error())
		}
	}

	// get fields from payload
	if sharedUtil.Contains(cols, "active") {
		_, err := trs.Query(builder.Update(
			constants.TableCoreDatasets,
		).UpdateJSON(
			"definitions",
			fmt.Sprintf("'{options,%s,active}'", o.Code),
			fmt.Sprintf(`'%v'`, o.Active),
			true,
		).Where(
			builder.Equal("code", datasetCode),
		))
		if err != nil {
			return customerror.New(http.StatusInternalServerError, "Update", err.Error())
		}
	}
	return nil
}

// Delete deletes a dataset option
func (o *Option) Delete(trs *db.Transaction, datasetCode string) error {
	total, err := db.Count("id", constants.TableCoreDatasets, &db.Options{
		Conditions: builder.Equal("code", datasetCode),
	})
	if err != nil || total == 0 {
		return customerror.New(http.StatusBadRequest, "validate dataset", "invalid dataset code")
	}

	total, err = db.Count(fmt.Sprintf("definitions->'options'->>'%s'", o.Code), constants.TableCoreDatasets, &db.Options{
		Conditions: builder.Equal("code", datasetCode),
	})
	if err != nil {
		return customerror.New(http.StatusBadRequest, "validate dataset", err.Error())
	}

	if total == 0 {
		return customerror.New(http.StatusNotFound, "validate dataset", "options not found")
	}

	rows, err := trs.Query(builder.Update(
		constants.TableCoreDatasets,
	).UpdateJSON(
		"definitions",
		"'{options}'",
		fmt.Sprintf("(definitions->'options') - '%s'", o.Code),
		true,
	).Where(
		builder.Equal("code", datasetCode),
	))
	if err != nil {
		return customerror.New(http.StatusInternalServerError, "Add", err.Error())
	}
	rows.Close()

	rows, err = trs.Query(builder.Update(
		constants.TableCoreDatasets,
	).UpdateJSON(
		"definitions",
		"'{order}'",
		fmt.Sprintf("(definitions->'order') - '%s'", o.Code),
		true,
	).Where(
		builder.Equal("code", datasetCode),
	))
	if err != nil {
		return customerror.New(http.StatusInternalServerError, "Add", err.Error())
	}
	rows.Close()

	if err != nil {
		return customerror.New(http.StatusInternalServerError, "Delete", err.Error())
	}

	return nil
}
