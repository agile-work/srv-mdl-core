package admin

import (
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
	err := lookup.ProcessDefinitions(languageCode)
	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Errors = append(response.Errors, moduleShared.NewResponseError(shared.ErrorInsertingRecord, "CreateLookup processing definitions", err.Error()))

		return response
	}

	sharedModels.TranslationFieldsRequestLanguageCode = "all"
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
	return db.Load(r, &lookups, "LoadAllLookups", shared.TableCoreLookups, nil)
}

// LoadLookup return only one object from the database
func LoadLookup(r *http.Request) *moduleShared.Response {
	lookup := models.Lookup{}
	lookupCode := chi.URLParam(r, "lookup_code")
	lookupIDColumn := fmt.Sprintf("%s.code", shared.TableCoreLookups)
	condition := builder.Equal(lookupIDColumn, lookupCode)

	return db.Load(r, &lookup, "LoadLookup", shared.TableCoreLookups, condition)
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
