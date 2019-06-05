package admin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/agile-work/srv-mdl-shared/db"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/agile-work/srv-mdl-core/models"
	moduleShared "github.com/agile-work/srv-mdl-shared"
	sharedModels "github.com/agile-work/srv-mdl-shared/models"
	shared "github.com/agile-work/srv-shared"
	sql "github.com/agile-work/srv-shared/sql-builder/db"
)

// CreateLookup persists the request body creating a new object in the database
func CreateLookup(r *http.Request) *moduleShared.Response {
	lookup := &models.Lookup{}
	languageCode := r.Header.Get("Content-Language")
	sharedModels.TranslationFieldsRequestLanguageCode = languageCode
	response := db.GetResponse(r, lookup, "CreateLookup")
	if response.Code != http.StatusOK {
		return response
	}
	// TODO: validate required fields on struct
	sharedModels.TranslationFieldsRequestLanguageCode = "all"
	err := lookup.ProcessDefinitions(languageCode, r.Method)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorInsertingRecord, "CreateLookup processing definitions", err.Error()))

		return response
	}

	id, err := sql.InsertStruct(shared.TableCoreLookups, lookup)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorInsertingRecord, "CreateLookup", err.Error()))

		return response
	}
	lookup.ID = id
	response.Data = lookup
	return response
}

// LoadAllLookups return all instances from the object
func LoadAllLookups(r *http.Request) *moduleShared.Response {
	lookups := []models.Lookup{}
	response := db.Load(r, &lookups, "LoadAllLookups", shared.TableCoreLookups, nil)
	for i := range lookups {
		err := lookups[i].ParseDefinition()
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingData, "LoadLookup ParseDefinition", err.Error()))

			return response
		}
	}
	response.Data = lookups
	return response
}

// LoadLookup return only one object from the database
func LoadLookup(r *http.Request) *moduleShared.Response {
	sharedModels.TranslationFieldsRequestLanguageCode = r.Header.Get("Content-Language")
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}

	lookup := &models.Lookup{}
	lookupCode := chi.URLParam(r, "lookup_code")
	lookupIDColumn := fmt.Sprintf("%s.code", shared.TableCoreLookups)
	condition := builder.Equal(lookupIDColumn, lookupCode)

	err := sql.LoadStruct(shared.TableCoreLookups, lookup, condition)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingData, "LoadLookup", err.Error()))

		return response
	}

	err = lookup.ParseDefinition()
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorLoadingData, "LoadLookup ParseDefinition", err.Error()))

		return response
	}

	response.Data = lookup
	return response
}

// UpdateLookup updates object data in the database
func UpdateLookup(r *http.Request) *moduleShared.Response {
	lookupCode := chi.URLParam(r, "lookup_code")
	lookupIDColumn := fmt.Sprintf("%s.code", shared.TableCoreLookups)
	condition := builder.Equal(lookupIDColumn, lookupCode)
	lookup := models.Lookup{}

	return db.Update(r, &lookup, "UpdateLookup", shared.TableCoreLookups, condition)
}

// DeleteLookup deletes object from the database
func DeleteLookup(r *http.Request) *moduleShared.Response {
	lookupID := chi.URLParam(r, "lookup_id")
	lookupIDColumn := fmt.Sprintf("%s.code", shared.TableCoreLookups)
	condition := builder.Equal(lookupIDColumn, lookupID)

	// TODO: validate lookup is not used in any field

	return db.Remove(r, "DeleteLookup", shared.TableCoreLookups, condition)
}

// AddLookupOption include new options to a lookup
func AddLookupOption(r *http.Request) *moduleShared.Response {
	option := &models.LookupOption{}
	lookupCode := chi.URLParam(r, "lookup_code")
	languageCode := r.Header.Get("Content-Language")
	sharedModels.TranslationFieldsRequestLanguageCode = languageCode
	response := db.GetResponse(r, option, "AddLookupOption")
	if response.Code != http.StatusOK {
		return response
	}

	// TODO: validate option required fields
	querySQL := fmt.Sprintf(`select count(id) total from %s, jsonb_array_elements(definitions->'options') opt 
		where code = '%s' and opt->>'code' = '%s'`, shared.TableCoreLookups, lookupCode, option.Code)
	rows, err := sql.Query(builder.Raw(querySQL))
	total := 0
	for rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorInsertingRecord, "AddLookupOption checking if code already exists", err.Error()))

			return response
		}
	}

	if total > 0 {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorInsertingRecord, "AddLookupOption", "code already exists"))

		return response
	}

	sharedModels.TranslationFieldsRequestLanguageCode = "all"
	optionBytes, _ := json.Marshal(option)
	querySQL = fmt.Sprintf(`update %s set definitions = jsonb_insert(
		definitions, '{options,-1}', '%s', true) 
		where code = '%s'`, shared.TableCoreLookups, string(optionBytes), lookupCode)

	err = sql.Exec(builder.Raw(querySQL))
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorInsertingRecord, "LoadLookup", err.Error()))

		return response
	}

	response.Data = option
	return response
}

// UpdateLookupOption update lookup option in the database
func UpdateLookupOption(r *http.Request) *moduleShared.Response {
	option := &models.LookupOption{}
	lookupCode := chi.URLParam(r, "lookup_code")
	optionCode := chi.URLParam(r, "option_code")
	languageCode := r.Header.Get("Content-Language")
	sharedModels.TranslationFieldsRequestLanguageCode = languageCode
	response := db.GetResponse(r, option, "AddLookupOption")
	if response.Code != http.StatusOK {
		return response
	}

	trs, err := sql.NewTransaction()
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorInsertingRecord, "UpdateLookupOption new transaction", err.Error()))

		return response
	}

	cols := db.GetBodyColumns(r)

	if shared.Contains(cols, "label") && languageCode != "all" {
		label := option.Label.String(languageCode)
		sqlQuery := fmt.Sprintf(`update %s set definitions = jsonb_set(
			definitions,
			('{options,'|| data_object.obj_index ||',label}') ::text[],
			definitions::jsonb#>('{options,'|| data_object.obj_index ||',label}')::text[] || '{"%s": "%s"}',
			true
			) from ( 
				select index-1 as obj_index from core_lookups ,jsonb_array_elements(definitions->'options') with ordinality arr(obj, index)
				where ((obj->>'code') = '%s') and (code = '%s')
			)data_object
			where (code = '%s')`, shared.TableCoreLookups, languageCode, label, optionCode, lookupCode, lookupCode)
		trs.Add(builder.Raw(sqlQuery))
	} else if shared.Contains(cols, "label") && languageCode == "all" {
		jsonBytes, _ := json.Marshal(option.Label)
		sqlQuery := fmt.Sprintf(`update %s set definitions = jsonb_set(
			definitions,
			('{options,'|| data_object.obj_index ||'}') ::text[],
			definitions::jsonb#>('{options,'|| data_object.obj_index ||'}')::text[] || '{"label": %s}',
			true
			) from ( 
				select index-1 as obj_index from core_lookups ,jsonb_array_elements(definitions->'options') with ordinality arr(obj, index)
				where ((obj->>'code') = '%s') and (code = '%s')
			)data_object
			where (code = '%s')`, shared.TableCoreLookups, string(jsonBytes), optionCode, lookupCode, lookupCode)
		trs.Add(builder.Raw(sqlQuery))
	}

	// get fields from payload
	if shared.Contains(cols, "active") {
		sqlQuery := fmt.Sprintf(`update %s set definitions = jsonb_set(
			definitions,
			('{options,'|| data_object.obj_index ||'}') ::text[],
			definitions::jsonb#>('{options,'|| data_object.obj_index ||'}')::text[] || '{"active": %t}',
			true
			) from ( 
				select index-1 as obj_index from core_lookups ,jsonb_array_elements(definitions->'options') with ordinality arr(obj, index)
				where ((obj->>'code') = '%s') and (code = '%s')
			)data_object
			where (code = '%s')`, shared.TableCoreLookups, option.Active, optionCode, lookupCode, lookupCode)
		trs.Add(builder.Raw(sqlQuery))
	}

	err = trs.Exec()
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorInsertingRecord, "UpdateLookupOption", err.Error()))

		return response
	}

	response.Data = option
	return response
}

// DeleteLookupOption deletes lookup option from the database
func DeleteLookupOption(r *http.Request) *moduleShared.Response {
	response := &moduleShared.Response{
		Code: http.StatusOK,
	}
	lookupCode := chi.URLParam(r, "lookup_code")
	optionCode := chi.URLParam(r, "option_code")
	// TODO: check if this lookup is being used in any field
	sqlQuery := fmt.Sprintf(`update %s set definitions = jsonb_set(
		definitions,
		('{options}') ::text[],
		(definitions->'options') - data_object.obj_index::int,
		true
		) 
		from ( 
			select index-1 as obj_index from %s ,jsonb_array_elements(definitions->'options') with ordinality arr(obj, index)
			where ((obj->>'code') = '%s') and (code = '%s')
		)data_object
		where (code = '%s')`, shared.TableCoreLookups, shared.TableCoreLookups, optionCode, lookupCode, lookupCode)
	err := sql.Exec(builder.Raw(sqlQuery))
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorDeletingData, "DeleteLookupOption", err.Error()))

		return response
	}
	return response
}

// UpdateLookupOrder update lookup static order
func UpdateLookupOrder(r *http.Request) *moduleShared.Response {
	staticLookup := &models.LookupStaticDefinition{}
	lookupCode := chi.URLParam(r, "lookup_code")
	response := db.GetResponse(r, staticLookup, "AddLookupOption")
	if response.Code != http.StatusOK {
		return response
	}

	cols := db.GetBodyColumns(r)

	if !shared.Contains(cols, "order_type") || !shared.Contains(cols, "order") || len(cols) != 2 {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorParsingRequest, "UpdateLookupOrder", "invalid request body"))

		return response
	}

	jsonBytes, _ := json.Marshal(staticLookup)
	sqlQuery := fmt.Sprintf(`update %s set
		definitions = definitions || '%s'
		where (code = '%s')`, shared.TableCoreLookups, string(jsonBytes), lookupCode)
	err := sql.Exec(builder.Raw(sqlQuery))
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorDeletingData, "DeleteLookupOption", err.Error()))

		return response
	}

	fmt.Println(sqlQuery)

	response.Data = staticLookup

	return response
}
