package admin

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/agile-work/srv-mdl-shared/models/translation"
// 	"github.com/agile-work/srv-shared/sql-builder/builder"
// 	"github.com/go-chi/chi"
// )

// // UpdateLookupOption update lookup option in the database
// func UpdateLookupOption(r *http.Request) *mdlShared.Response {
// 	option := &models.LookupOption{}
// 	lookupCode := chi.URLParam(r, "lookup_code")
// 	optionCode := chi.URLParam(r, "option_code")
// 	languageCode := r.Header.Get("Content-Language")
// 	translation.FieldsRequestLanguageCode = languageCode
// 	response := db.GetResponse(r, option, "AddLookupOption")
// 	if response.Code != http.StatusOK {
// 		return response
// 	}

// 	trs, err := sql.NewTransaction()
// 	if err != nil {
// 		response.Code = http.StatusInternalServerError
// 		response.Errors = append(response.Errors, mdlShared.NewResponseError(shared.ErrorInsertingRecord, "UpdateLookupOption new transaction", err.Error()))

// 		return response
// 	}

// 	cols := db.GetBodyColumns(r)

// 	if shared.Contains(cols, "label") && languageCode != "all" {
// 		label := option.Label.String(languageCode)
// 		sqlQuery := fmt.Sprintf(`update %s set definitions = jsonb_set(
// 			definitions,
// 			('{options,'|| data_object.obj_index ||',label}') ::text[],
// 			definitions::jsonb#>('{options,'|| data_object.obj_index ||',label}')::text[] || '{"%s": "%s"}',
// 			true
// 			) from (
// 				select index-1 as obj_index from core_lookups ,jsonb_array_elements(definitions->'options') with ordinality arr(obj, index)
// 				where ((obj->>'code') = '%s') and (code = '%s')
// 			)data_object
// 			where (code = '%s')`, shared.TableCoreLookups, languageCode, label, optionCode, lookupCode, lookupCode)
// 		trs.Add(builder.Raw(sqlQuery))
// 	} else if shared.Contains(cols, "label") && languageCode == "all" {
// 		jsonBytes, _ := json.Marshal(option.Label)
// 		sqlQuery := fmt.Sprintf(`update %s set definitions = jsonb_set(
// 			definitions,
// 			('{options,'|| data_object.obj_index ||'}') ::text[],
// 			definitions::jsonb#>('{options,'|| data_object.obj_index ||'}')::text[] || '{"label": %s}',
// 			true
// 			) from (
// 				select index-1 as obj_index from core_lookups ,jsonb_array_elements(definitions->'options') with ordinality arr(obj, index)
// 				where ((obj->>'code') = '%s') and (code = '%s')
// 			)data_object
// 			where (code = '%s')`, shared.TableCoreLookups, string(jsonBytes), optionCode, lookupCode, lookupCode)
// 		trs.Add(builder.Raw(sqlQuery))
// 	}

// 	// get fields from payload
// 	if shared.Contains(cols, "active") {
// 		sqlQuery := fmt.Sprintf(`update %s set definitions = jsonb_set(
// 			definitions,
// 			('{options,'|| data_object.obj_index ||'}') ::text[],
// 			definitions::jsonb#>('{options,'|| data_object.obj_index ||'}')::text[] || '{"active": %t}',
// 			true
// 			) from (
// 				select index-1 as obj_index from core_lookups ,jsonb_array_elements(definitions->'options') with ordinality arr(obj, index)
// 				where ((obj->>'code') = '%s') and (code = '%s')
// 			)data_object
// 			where (code = '%s')`, shared.TableCoreLookups, option.Active, optionCode, lookupCode, lookupCode)
// 		trs.Add(builder.Raw(sqlQuery))
// 	}

// 	err = trs.Exec()
// 	if err != nil {
// 		trs.Rollback()
// 		response.Code = http.StatusInternalServerError
// 		response.Errors = append(response.Errors, mdlShared.NewResponseError(shared.ErrorInsertingRecord, "UpdateLookupOption", err.Error()))

// 		return response
// 	}

// 	trs.Commit()
// 	resp.Data = option
// 	return response
// }

// // DeleteLookupOption deletes lookup option from the database
// func DeleteLookupOption(r *http.Request) *mdlShared.Response {
// 	response := &mdlShared.Response{
// 		Code: http.StatusOK,
// 	}
// 	lookupCode := chi.URLParam(r, "lookup_code")
// 	optionCode := chi.URLParam(r, "option_code")
// 	// TODO: check if this lookup is being used in any field
// 	sqlQuery := fmt.Sprintf(`update %s set definitions = jsonb_set(
// 		definitions,
// 		('{options}') ::text[],
// 		(definitions->'options') - data_object.obj_index::int,
// 		true
// 		)
// 		from (
// 			select index-1 as obj_index from %s ,jsonb_array_elements(definitions->'options') with ordinality arr(obj, index)
// 			where ((obj->>'code') = '%s') and (code = '%s')
// 		)data_object
// 		where (code = '%s')`, shared.TableCoreLookups, shared.TableCoreLookups, optionCode, lookupCode, lookupCode)
// 	err := sql.Exec(builder.Raw(sqlQuery))
// 	if err != nil {
// 		response.Code = http.StatusInternalServerError
// 		response.Errors = append(response.Errors, mdlShared.NewResponseError(shared.ErrorDeletingData, "DeleteLookupOption", err.Error()))

// 		return response
// 	}
// 	return response
// }

// // UpdateLookupOrder update lookup static order
// func UpdateLookupOrder(r *http.Request) *mdlShared.Response {
// 	staticLookup := &models.LookupStaticDefinition{}
// 	lookupCode := chi.URLParam(r, "lookup_code")
// 	response := db.GetResponse(r, staticLookup, "AddLookupOption")
// 	if response.Code != http.StatusOK {
// 		return response
// 	}

// 	cols := db.GetBodyColumns(r)

// 	if !shared.Contains(cols, "order_type") || !shared.Contains(cols, "order") || len(cols) != 2 {
// 		response.Code = http.StatusInternalServerError
// 		response.Errors = append(response.Errors, mdlShared.NewResponseError(shared.ErrorParsingRequest, "UpdateLookupOrder", "invalid request body"))

// 		return response
// 	}

// 	jsonBytes, _ := json.Marshal(staticLookup)
// 	sqlQuery := fmt.Sprintf(`update %s set
// 		definitions = definitions || '%s'
// 		where (code = '%s')`, shared.TableCoreLookups, string(jsonBytes), lookupCode)
// 	err := sql.Exec(builder.Raw(sqlQuery))
// 	if err != nil {
// 		response.Code = http.StatusInternalServerError
// 		response.Errors = append(response.Errors, mdlShared.NewResponseError(shared.ErrorDeletingData, "DeleteLookupOption", err.Error()))

// 		return response
// 	}

// 	fmt.Println(sqlQuery)

// 	resp.Data = staticLookup

// 	return response
// }

// // UpdateLookupQuery update dynamic lookup query
// // TODO: implement a feature to publish the new query to avoid break the system
// func UpdateLookupQuery(r *http.Request) *mdlShared.Response {
// 	dynamicLookup := &models.LookupDynamicDefinition{}
// 	languageCode := r.Header.Get("Content-Language")
// 	translation.FieldsRequestLanguageCode = "all"
// 	response := db.GetResponse(r, dynamicLookup, "AddLookupOption")
// 	if response.Code != http.StatusOK {
// 		return response
// 	}

// 	cols := db.GetBodyColumns(r)

// 	if !shared.Contains(cols, "query") || len(cols) != 1 {
// 		response.Code = http.StatusInternalServerError
// 		response.Errors = append(response.Errors, mdlShared.NewResponseError(shared.ErrorParsingRequest, "UpdateLookupQuery", "invalid request body"))

// 		return response
// 	}

// 	err := dynamicLookup.ParseQuery(languageCode)
// 	if err != nil {
// 		response.Code = http.StatusInternalServerError
// 		response.Errors = append(response.Errors, mdlShared.NewResponseError(shared.ErrorInsertingRecord, "UpdateLookupQuery invalid query", err.Error()))

// 		return response
// 	}

// 	lookup := &models.Lookup{}
// 	lookupCode := chi.URLParam(r, "lookup_code")
// 	lookupIDColumn := fmt.Sprintf("%s.code", shared.TableCoreLookups)
// 	condition := builder.Equal(lookupIDColumn, lookupCode)

// 	err = sql.SelectStruct(shared.TableCoreLookups, lookup, condition)
// 	if err != nil {
// 		response.Code = http.StatusInternalServerError
// 		response.Errors = append(response.Errors, mdlShared.NewResponseError(shared.ErrorLoadingData, "UpdateLookupQuery loading lookup", err.Error()))

// 		return response
// 	}

// 	if lookup.Type != shared.LookupDynamic {
// 		response.Code = http.StatusInternalServerError
// 		response.Errors = append(response.Errors, mdlShared.NewResponseError(shared.ErrorParsingRequest, "UpdateLookupQuery validation", "only dynamic lookups can update query"))

// 		return response
// 	}

// 	currentDynamicLookup := &models.LookupDynamicDefinition{}
// 	jsonBytes, _ := json.Marshal(lookup.Definitions)
// 	json.Unmarshal(jsonBytes, currentDynamicLookup)

// 	/*
// 		// TODO: Activate this validation when implementing publish lookup
// 		if len(currentDynamicLookup.Fields) > len(dynamicLookup.Fields) || len(currentDynamicLookup.Params) > len(dynamicLookup.Params) {
// 			response.Code = http.StatusInternalServerError
// 			response.Errors = append(response.Errors, mdlShared.NewResponseError(shared.ErrorParsingRequest, "UpdateLookupQuery validation", "can't change query structure"))

// 			return response
// 		}

// 		errors := []string{}

// 		for _, f := range currentDynamicLookup.Fields {
// 			if !dynamicLookup.ContainsField(f) {
// 				errors = append(errors, f.Code)
// 			}
// 		}

// 		for _, p := range currentDynamicLookup.Params {
// 			if dynamicLookup.ContainsParam(p) == -1 {
// 				errors = append(errors, p.Code)
// 			}
// 		}

// 		if len(errors) > 0 {
// 			msg := fmt.Sprintf("can't change query structure, invalid: %s", strings.Join(errors, ", "))
// 			response.Code = http.StatusInternalServerError
// 			response.Errors = append(response.Errors, mdlShared.NewResponseError(shared.ErrorInsertingRecord, "UpdateLookupQuery validation", msg))

// 			return response
// 		}
// 	*/

// 	for _, f := range dynamicLookup.Fields {
// 		if !currentDynamicLookup.ContainsField(f) {
// 			currentDynamicLookup.Fields = append(currentDynamicLookup.Fields, f)
// 		}
// 	}

// 	for _, p := range dynamicLookup.Params {
// 		val := currentDynamicLookup.ContainsParam(p)
// 		if val == -1 {
// 			currentDynamicLookup.Params = append(currentDynamicLookup.Params, p)
// 		} else if val == 1 {
// 			index := currentDynamicLookup.GetParamIndex(p)
// 			currentDynamicLookup.Params[index].Pattern = p.Pattern
// 		}
// 	}

// 	currentDynamicLookup.UpdatedAt = dynamicLookup.UpdatedAt
// 	currentDynamicLookup.UpdatedBy = dynamicLookup.UpdatedBy
// 	currentDynamicLookup.Query = dynamicLookup.Query

// 	jsonBytes, _ = json.Marshal(currentDynamicLookup)
// 	json.Unmarshal(jsonBytes, &lookup.Definitions)

// 	sqlQuery := fmt.Sprintf(`update %s set
// 		definitions = '%s',
// 		updated_by = '%s',
// 		updated_at = current_date
// 		where (code = '%s')`, shared.TableCoreLookups, string(jsonBytes), dynamicLookup.UpdatedBy, lookupCode)
// 	err = sql.Exec(builder.Raw(sqlQuery))
// 	if err != nil {
// 		response.Code = http.StatusInternalServerError
// 		response.Errors = append(response.Errors, mdlShared.NewResponseError(shared.ErrorDeletingData, "UpdateLookupQuery", err.Error()))

// 		return response
// 	}

// 	translation.FieldsRequestLanguageCode = languageCode
// 	resp.Data = lookup

// 	return response
// }

// // UpdateLookupDynamicParam update dynamic lookup generic param (fields or filter params)
// func UpdateLookupDynamicParam(r *http.Request, paramType string) *mdlShared.Response {
// 	param := &models.LookupParam{}
// 	lookupCode := chi.URLParam(r, "lookup_code")
// 	paramCode := chi.URLParam(r, "param_code")
// 	languageCode := r.Header.Get("Content-Language")
// 	translation.FieldsRequestLanguageCode = languageCode
// 	response := db.GetResponse(r, param, "AddLookupOption")
// 	if response.Code != http.StatusOK {
// 		return response
// 	}

// 	trs, err := sql.NewTransaction()
// 	if err != nil {
// 		response.Code = http.StatusInternalServerError
// 		response.Errors = append(response.Errors, mdlShared.NewResponseError(shared.ErrorInsertingRecord, "UpdateLookupDynamicField new transaction", err.Error()))

// 		return response
// 	}

// 	cols := db.GetBodyColumns(r)

// 	if shared.Contains(cols, "label") && languageCode != "all" {
// 		label := param.Label.String(languageCode)
// 		sqlQuery := fmt.Sprintf(`update %s set definitions = jsonb_set(
// 			definitions,
// 			('{fields,'|| data_object.obj_index ||',label}') ::text[],
// 			definitions::jsonb#>('{fields,'|| data_object.obj_index ||',label}')::text[] || '{"%s": "%s"}',
// 			true
// 			) from (
// 				select index-1 as obj_index from core_lookups ,jsonb_array_elements(definitions->'fields') with ordinality arr(obj, index)
// 				where ((obj->>'code') = '%s') and (code = '%s')
// 			)data_object
// 			where (code = '%s')`, shared.TableCoreLookups, languageCode, label, paramCode, lookupCode, lookupCode)
// 		trs.Add(builder.Raw(sqlQuery))
// 	} else if shared.Contains(cols, "label") && languageCode == "all" {
// 		jsonBytes, _ := json.Marshal(param.Label)
// 		sqlQuery := getQueryUpdateField("label", string(jsonBytes), paramCode, lookupCode)
// 		trs.Add(builder.Raw(sqlQuery))
// 	}

// 	if shared.Contains(cols, "field_type") && paramType == "field" {
// 		jsonBytes, _ := json.Marshal(param.Type)
// 		sqlQuery := getQueryUpdateField("field_type", string(jsonBytes), paramCode, lookupCode)
// 		trs.Add(builder.Raw(sqlQuery))
// 	}

// 	if shared.Contains(cols, "security") && paramType == "field" {
// 		jsonBytes, _ := json.Marshal(param.Security)
// 		sqlQuery := getQueryUpdateField("security", string(jsonBytes), paramCode, lookupCode)
// 		trs.Add(builder.Raw(sqlQuery))
// 	}

// 	err = trs.Exec()
// 	if err != nil {
// 		trs.Rollback()
// 		response.Code = http.StatusInternalServerError
// 		response.Errors = append(response.Errors, mdlShared.NewResponseError(shared.ErrorInsertingRecord, "UpdateLookupOption", err.Error()))

// 		return response
// 	}

// 	trs.Commit()
// 	resp.Data = param
// 	return response
// }

// func getQueryUpdateField(field, value, paramCode, lookupCode string) string {
// 	return fmt.Sprintf(`update %s set definitions = jsonb_set(
// 		definitions,
// 		('{fields,'|| data_object.obj_index ||'}') ::text[],
// 		definitions::jsonb#>('{fields,'|| data_object.obj_index ||'}')::text[] || '{"%s": %s}',
// 		true
// 		) from (
// 			select index-1 as obj_index from core_lookups ,jsonb_array_elements(definitions->'fields') with ordinality arr(obj, index)
// 			where ((obj->>'code') = '%s') and (code = '%s')
// 		)data_object
// 		where (code = '%s')`, shared.TableCoreLookups, field, value, paramCode, lookupCode, lookupCode)
// }
