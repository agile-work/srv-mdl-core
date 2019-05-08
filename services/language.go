package services

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

// CreateLanguage persists the request body creating a new object in the database
func CreateLanguage(r *http.Request) *moduleShared.Response {
	language := models.Language{}

	return db.Create(r, &language, "CreateLanguage", shared.TableCoreConfigLanguages)
}

// LoadAllLanguages return all instances from the object
func LoadAllLanguages(r *http.Request) *moduleShared.Response {
	languages := []models.Language{}

	return db.Load(r, &languages, "LoadAllLanguages", shared.TableCoreConfigLanguages, nil)
}

// LoadLanguage return only one object from the database
func LoadLanguage(r *http.Request) *moduleShared.Response {
	language := models.Language{}
	languageID := chi.URLParam(r, "language_id")
	languageIDColumn := fmt.Sprintf("%s.id", shared.TableCoreConfigLanguages)
	condition := builder.Equal(languageIDColumn, languageID)

	return db.Load(r, &language, "LoadLanguage", shared.TableCoreConfigLanguages, condition)
}

// UpdateLanguage updates object data in the database
func UpdateLanguage(r *http.Request) *moduleShared.Response {
	languageID := chi.URLParam(r, "language_id")
	languageIDColumn := fmt.Sprintf("%s.id", shared.TableCoreConfigLanguages)
	condition := builder.Equal(languageIDColumn, languageID)
	language := models.Language{
		ID: languageID,
	}

	return db.Update(r, &language, "UpdateLanguage", shared.TableCoreConfigLanguages, condition)
}

// DeleteLanguage deletes object from the database
func DeleteLanguage(r *http.Request) *moduleShared.Response {
	languageID := chi.URLParam(r, "language_id")
	languageIDColumn := fmt.Sprintf("%s.id", shared.TableCoreConfigLanguages)
	condition := builder.Equal(languageIDColumn, languageID)

	return db.Remove(r, "DeleteLanguage", shared.TableCoreConfigLanguages, condition)
}
