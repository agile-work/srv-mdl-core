package admin

import (
	"fmt"
	"net/http"

	"github.com/agile-work/srv-mdl-shared/db"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/go-chi/chi"

	"github.com/agile-work/srv-mdl-core/models"
	moduleShared "github.com/agile-work/srv-mdl-shared"
	shared "github.com/agile-work/srv-shared"
)

// CreateLookup persists the request body creating a new object in the database
func CreateLookup(r *http.Request) *moduleShared.Response {
	lookup := models.Lookup{}

	return db.Create(r, &lookup, "CreateLookup", shared.TableCoreLookups)
}

// LoadAllLookups return all instances from the object
func LoadAllLookups(r *http.Request) *moduleShared.Response {
	lookups := []models.Lookup{}

	return db.Load(r, &lookups, "LoadAllLookups", shared.TableCoreLookups, nil)
}

// LoadLookup return only one object from the database
func LoadLookup(r *http.Request) *moduleShared.Response {
	lookup := models.Lookup{}
	lookupID := chi.URLParam(r, "lookup_id")
	lookupIDColumn := fmt.Sprintf("%s.id", shared.TableCoreLookups)
	condition := builder.Equal(lookupIDColumn, lookupID)

	return db.Load(r, &lookup, "LoadLookup", shared.TableCoreLookups, condition)
}

// UpdateLookup updates object data in the database
func UpdateLookup(r *http.Request) *moduleShared.Response {
	lookupID := chi.URLParam(r, "lookup_id")
	lookupIDColumn := fmt.Sprintf("%s.id", shared.TableCoreLookups)
	condition := builder.Equal(lookupIDColumn, lookupID)
	lookup := models.Lookup{
		ID: lookupID,
	}

	return db.Update(r, &lookup, "UpdateLookup", shared.TableCoreLookups, condition)
}

// DeleteLookup deletes object from the database
func DeleteLookup(r *http.Request) *moduleShared.Response {
	lookupID := chi.URLParam(r, "lookup_id")
	lookupIDColumn := fmt.Sprintf("%s.id", shared.TableCoreLookups)
	condition := builder.Equal(lookupIDColumn, lookupID)

	return db.Remove(r, "DeleteLookup", shared.TableCoreLookups, condition)
}

// CreateLookupOption persists the request body creating a new object in the database
func CreateLookupOption(r *http.Request) *moduleShared.Response {
	lookupID := chi.URLParam(r, "lookup_id")
	lookupOption := models.LookupOption{
		LookupID: lookupID,
	}

	return db.Create(r, &lookupOption, "CreateLookupOption", shared.TableCoreLkpOptions)
}

// LoadAllLookupOptions return all instances from the object
func LoadAllLookupOptions(r *http.Request) *moduleShared.Response {
	lookupOptions := []models.LookupOption{}
	lookupID := chi.URLParam(r, "lookup_id")
	lookupIDColumn := fmt.Sprintf("%s.lookup_id", shared.TableCoreLkpOptions)
	condition := builder.Equal(lookupIDColumn, lookupID)

	return db.Load(r, &lookupOptions, "LoadAllLookupOptions", shared.TableCoreLkpOptions, condition)
}

// LoadLookupOption return only one object from the database
func LoadLookupOption(r *http.Request) *moduleShared.Response {
	lookupOption := models.LookupOption{}
	lookupOptionID := chi.URLParam(r, "lookup_option_id")
	lookupOptionIDColumn := fmt.Sprintf("%s.id", shared.TableCoreLkpOptions)
	condition := builder.Equal(lookupOptionIDColumn, lookupOptionID)

	return db.Load(r, &lookupOption, "LoadLookupOption", shared.TableCoreLkpOptions, condition)
}

// UpdateLookupOption updates object data in the database
func UpdateLookupOption(r *http.Request) *moduleShared.Response {
	lookupOptionID := chi.URLParam(r, "lookup_option_id")
	lookupOptionIDColumn := fmt.Sprintf("%s.id", shared.TableCoreLkpOptions)
	condition := builder.Equal(lookupOptionIDColumn, lookupOptionID)
	lookupOption := models.LookupOption{
		ID: lookupOptionID,
	}

	return db.Update(r, &lookupOption, "UpdateLookupOption", shared.TableCoreLkpOptions, condition)
}

// DeleteLookupOption deletes object from the database
func DeleteLookupOption(r *http.Request) *moduleShared.Response {
	lookupOptionID := chi.URLParam(r, "lookup_option_id")
	lookupOptionIDColumn := fmt.Sprintf("%s.id", shared.TableCoreLkpOptions)
	condition := builder.Equal(lookupOptionIDColumn, lookupOptionID)

	return db.Remove(r, "DeleteLookupOption", shared.TableCoreLkpOptions, condition)
}
