package currency

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/agile-work/srv-mdl-shared/models/customerror"
	"github.com/agile-work/srv-mdl-shared/models/translation"
	"github.com/agile-work/srv-shared/constants"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// Currency defines the struct of this object
type Currency struct {
	ID        string                  `json:"id" sql:"id" pk:"true"`
	Code      string                  `json:"code" sql:"code" updatable:"false"`
	Name      translation.Translation `json:"name" sql:"name" field:"jsonb"`
	Rates     map[string][]Rate       `json:"rates" sql:"rates" field:"jsonb"`
	Active    bool                    `json:"active" sql:"active"`
	CreatedBy string                  `json:"created_by" sql:"created_by"`
	CreatedAt time.Time               `json:"created_at" sql:"created_at"`
	UpdatedBy string                  `json:"updated_by" sql:"updated_by"`
	UpdatedAt time.Time               `json:"updated_at" sql:"updated_at"`
}

// Rate defines the struct of this object
type Rate struct {
	Value float64    `json:"value"`
	Start *time.Time `json:"start_at,omitempty"`
	End   *time.Time `json:"end_at,omitempty"`
}

// Currencies defines the array struct of this object
type Currencies []Currency

// Create persists the struct creating a new object in the database
func (c *Currency) Create(trs *db.Transaction, columns ...string) error {
	id, err := db.InsertStructTx(trs.Tx, constants.TableCoreCurrencies, c, columns...)
	if err != nil {
		return customerror.New(http.StatusInternalServerError, "currency create", err.Error())
	}
	c.ID = id
	return nil
}

// LoadAll defines all instances from the object
func (c *Currencies) LoadAll(opt *db.Options) error {
	// TODO: Make a way of limit the columns for the get all. Passing fields to LoadStruct.
	if err := db.SelectStruct(constants.TableCoreCurrencies, c, opt); err != nil {
		return customerror.New(http.StatusInternalServerError, "currencies load", err.Error())
	}
	return nil
}

// Load defines only one object from the database
func (c *Currency) Load() error {
	// TODO: Limit the total of rates to +/- 100 records if has no filter
	if err := db.SelectStruct(constants.TableCoreCurrencies, c, &db.Options{
		Conditions: builder.Equal("code", c.Code),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "currency load", err.Error())
	}
	return nil
}

// Update updates object data in the database
func (c *Currency) Update(trs *db.Transaction, columns []string, translations map[string]string) error {
	opt := &db.Options{Conditions: builder.Equal("code", c.Code)}

	if len(columns) > 0 {
		if err := db.UpdateStructTx(trs.Tx, constants.TableCoreCurrencies, c, opt, strings.Join(columns, ",")); err != nil {
			return customerror.New(http.StatusInternalServerError, "currency update", err.Error())
		}
	}

	if len(translations) > 0 {
		statement := builder.Update(constants.TableCoreCurrencies)
		for col, val := range translations {
			statement.JSON(col, translation.FieldsRequestLanguageCode)
			jsonVal, _ := json.Marshal(val)
			statement.Values(jsonVal)
		}
		statement.Where(opt.Conditions)
		if _, err := trs.Query(statement); err != nil {
			return customerror.New(http.StatusInternalServerError, "currency update", err.Error())
		}
	}

	return nil
}

// Delete deletes object from the database
func (c *Currency) Delete(trs *db.Transaction) error {
	if err := db.DeleteStructTx(trs.Tx, constants.TableCoreCurrencies, &db.Options{
		Conditions: builder.And(
			builder.Equal("code", c.Code),
			builder.Equal("active", false),
		),
	}); err != nil {
		return customerror.New(http.StatusInternalServerError, "currency delete", err.Error())
	}
	return nil
}

// AddRate persists the request body creating a new object in the database
func (r *Rate) AddRate(trs *db.Transaction, fromCode, toCode string) error {
	querySQL := fmt.Sprintf(`update %s 
		set rates = jsonb_set(
			rates, 
			('{%s,-1}') ::text[], 
			rates::jsonb#>('{%s,-1}')::text[] || '{"end_at": "%s"}'
			,true
		) where code = '%s'`, constants.TableCoreCurrencies, toCode, toCode, time.Now().Format(time.RFC3339), fromCode)
	if _, err := trs.Query(builder.Raw(querySQL)); err != nil {
		return customerror.New(http.StatusInternalServerError, "add currency rate", err.Error())
	}

	t := time.Now()
	r.Start = &t
	rateBytes, _ := json.Marshal(r)
	querySQL = fmt.Sprintf(`update %s set rates = jsonb_insert(
		rates, '{%s,-1}', '%s', true) 
		where code = '%s'`, constants.TableCoreCurrencies, toCode, string(rateBytes), fromCode)
	if _, err := trs.Query(builder.Raw(querySQL)); err != nil {
		return customerror.New(http.StatusInternalServerError, "add currency rate", err.Error())
	}
	return nil
}
