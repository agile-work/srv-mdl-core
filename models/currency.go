package models

import "time"

// Currency defines the struct of this object
type Currency struct {
	ID            string    `json:"id" sql:"id" pk:"true"`
	Code          string    `json:"code" sql:"code"`
	Name          string    `json:"name" table:"core_translations" alias:"core_translations_name" sql:"value" on:"core_translations_name.structure_id = core_currencies.id and core_translations_name.structure_field = 'name'"`
	Active        bool      `json:"active" sql:"active"`
	CreatedBy     string    `json:"created_by" sql:"created_by"`
	CreatedByUser *User     `json:"created_by_user" table:"core_users" alias:"created_by_user" on:"created_by_user.id = core_currencies.created_by"`
	CreatedAt     time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy     string    `json:"updated_by" sql:"updated_by"`
	UpdatedByUser *User     `json:"updated_by_user" table:"core_users" alias:"updated_by_user" on:"updated_by_user.id = core_currencies.updated_by"`
	UpdatedAt     time.Time `json:"updated_at" sql:"updated_at"`
}

// CurrencyRate defines the struct of this object
type CurrencyRate struct {
	ID               string     `json:"id" sql:"id" pk:"true"`
	FromCurrencyID   string     `json:"from_currency_id" sql:"from_currency_id" fk:"true"`
	ToCurrencyID     string     `json:"to_currency_id" sql:"to_currency_id" fk:"true"`
	FromCurrencyCode string     `json:"from_currency_code" sql:"from_currency_code" fk:"true"`
	ToCurrencyCode   string     `json:"to_currency_code" sql:"to_currency_code" fk:"true"`
	Value            int        `json:"value" sql:"value"`
	StartAt          *time.Time `json:"start_at" sql:"start_at"`
	EndAt            *time.Time `json:"end_at" sql:"end_at"`
	CreatedBy        string     `json:"created_by" sql:"created_by"`
	CreatedByUser    *User      `json:"created_by_user" table:"core_users" alias:"created_by_user" on:"created_by_user.id = core_cry_rates.created_by"`
	CreatedAt        time.Time  `json:"created_at" sql:"created_at"`
	UpdatedBy        string     `json:"updated_by" sql:"updated_by"`
	UpdatedByUser    *User      `json:"updated_by_user" table:"core_users" alias:"updated_by_user" on:"updated_by_user.id = core_cry_rates.updated_by"`
	UpdatedAt        time.Time  `json:"updated_at" sql:"updated_at"`
}
