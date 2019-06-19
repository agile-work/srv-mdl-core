package schema

import (
	"encoding/json"
	"strings"
	"time"

	mdlShared "github.com/agile-work/srv-mdl-shared"
	mdlSharedModels "github.com/agile-work/srv-mdl-shared/models"
	"github.com/agile-work/srv-mdl-shared/models/job"
	"github.com/agile-work/srv-shared/constants"
	"github.com/agile-work/srv-shared/sql-builder/builder"
	"github.com/agile-work/srv-shared/sql-builder/db"
)

// Schema defines the struct of this object
type Schema struct {
	ID          string                      `json:"id" sql:"id" pk:"true"`
	JobID       string                      `json:"job_id" sql:"job_id" fk:"true"`
	Code        string                      `json:"code" sql:"code" updatable:"false" validate:"required"`
	Name        mdlSharedModels.Translation `json:"name" sql:"name" field:"jsonb" validate:"required"`
	Description mdlSharedModels.Translation `json:"description" sql:"description" field:"jsonb" validate:"required"`
	Parent      string                      `json:"parent_id" sql:"parent_id"`
	IsExtension bool                        `json:"is_extension" sql:"is_extension"`
	Module      bool                        `json:"module" sql:"module"`
	Active      bool                        `json:"active" sql:"active"`
	Status      string                      `json:"status" sql:"status"`
	CreatedBy   string                      `json:"created_by" sql:"created_by"`
	CreatedAt   time.Time                   `json:"created_at" sql:"created_at"`
	UpdatedBy   string                      `json:"updated_by" sql:"updated_by"`
	UpdatedAt   time.Time                   `json:"updated_at" sql:"updated_at"`
}

// ViewSchModules defines the struct of this object
type ViewSchModules struct {
	ID           string    `json:"id" sql:"id" pk:"true"`
	SchemaID     string    `json:"schema_id" sql:"schema_id"`
	Code         string    `json:"code" sql:"code"`
	Name         string    `json:"name" sql:"name"`
	Description  string    `json:"description" sql:"description"`
	LanguageCode string    `json:"language_code" sql:"language_code"`
	Module       bool      `json:"module" sql:"module"`
	Active       bool      `json:"active" sql:"active"`
	CreatedBy    string    `json:"created_by" sql:"created_by"`
	CreatedAt    time.Time `json:"created_at" sql:"created_at"`
	UpdatedBy    string    `json:"updated_by" sql:"updated_by"`
	UpdatedAt    time.Time `json:"updated_at" sql:"updated_at"`
}

// Schemas defines the array struct of this object
type Schemas []Schema

// Create persists the struct creating a new object in the database
func (s *Schema) Create(trs *db.Transaction, columns ...string) error {
	s.Status = constants.SchemaStatusProcessing
	id, err := db.InsertStructTx(trs.Tx, constants.TableCoreSchemas, s, columns...)
	if err != nil {
		return mdlShared.NewError("schema create", err.Error())
	}
	s.ID = id
	params := map[string]interface{}{
		"schema_code": s.Code,
	}

	id, err = job.CreateInstance(s.CreatedBy, constants.JobSystemCreateSchema, params)
	if err != nil {
		return mdlShared.NewError("schema create job execution", err.Error())
	}

	s.JobID = id
	if err := s.Update(trs, []string{"job_id"}, nil); err != nil {
		return mdlShared.NewError("schema create update job id", err.Error())
	}

	return nil
}

// LoadAll defines all instances from the object
func (s *Schemas) LoadAll(trs *db.Transaction, opt *db.Options) error {
	if err := db.SelectStructTx(trs.Tx, constants.TableCoreSchemas, s, opt); err != nil {
		return mdlShared.NewError("schemas load", err.Error())
	}
	return nil
}

// Load defines only one object from the database
func (s *Schema) Load(trs *db.Transaction) error {
	if err := db.SelectStructTx(trs.Tx, constants.TableCoreSchemas, s, &db.Options{
		Conditions: builder.Equal("code", s.Code),
	}); err != nil {
		return mdlShared.NewError("schema load", err.Error())
	}
	return nil
}

// Update updates object data in the database
func (s *Schema) Update(trs *db.Transaction, columns []string, translations map[string]string) error {
	opt := &db.Options{Conditions: builder.Equal("code", s.Code)}

	if len(columns) > 0 {
		if err := db.UpdateStructTx(trs.Tx, constants.TableCoreSchemas, s, opt, strings.Join(columns, ",")); err != nil {
			return mdlShared.NewError("schema update", err.Error())
		}
	}

	if len(translations) > 0 {
		statement := builder.Update(constants.TableCoreSchemas)
		for col, val := range translations {
			statement.JSON(col, mdlSharedModels.TranslationFieldsRequestLanguageCode)
			jsonVal, _ := json.Marshal(val)
			statement.Values(jsonVal)
		}
		statement.Where(opt.Conditions)
		if _, err := trs.Query(statement); err != nil {
			return mdlShared.NewError("schema update", err.Error())
		}
	}

	return nil
}

// Delete deletes object from the database
func (s *Schema) Delete(trs *db.Transaction) error {
	if err := db.DeleteStructTx(trs.Tx, constants.TableCoreSchemas, &db.Options{
		Conditions: builder.Equal("code", s.Code),
	}); err != nil {
		return mdlShared.NewError("schema delete", err.Error())
	}
	return nil
}

// CallDelete call job to delete object from the database
func (s *Schema) CallDelete(trs *db.Transaction) error {
	s.Status = constants.SchemaStatusDeleting
	s.Active = false

	if err := s.Update(trs, []string{"status", "active"}, nil); err != nil {
		return mdlShared.NewError("schema delete update status", err.Error())
	}

	params := map[string]interface{}{
		"schema_code": s.Code,
	}

	if _, err := job.CreateInstance(s.UpdatedBy, constants.JobSystemDeleteSchema, params); err != nil {
		return mdlShared.NewError("schema delete job execution", err.Error())
	}

	return nil
}

// // InsertModuleInSchema persists the request creating a new object in the database
// func InsertModuleInSchema(r *http.Request) *mdlShared.Response {
// 	response := &mdlShared.Response{
// 		Code: http.StatusOK,
// 	}

// 	schemaID := chi.URLParam(r, "schema_id")
// 	moduleID := chi.URLParam(r, "module_id")

// 	userID := r.Header.Get("userID")
// 	now := time.Now()

// 	statemant := builder.Insert(
// 		constants.TableCoreSchemasModels,
// 		"schema_id",
// 		"module_id",
// 		"created_by",
// 		"created_at",
// 		"updated_by",
// 		"updated_at",
// 	).Values(
// 		schemaID,
// 		moduleID,
// 		userID,
// 		now,
// 		userID,
// 		now,
// 	)

// 	err := sql.Exec(statemant)
// 	if err != nil {
// 		response.Code = http.StatusInternalServerError
// 		response.Errors = append(response.Errors, mdlShared.NewResponseError(constants.ErrorInsertingRecord, "InsertModuleSchema", err.Error()))

// 		return response
// 	}

// 	return response
// }

// // LoadAllModulesBySchema return all instances from the object
// func LoadAllModulesBySchema(r *http.Request) *mdlShared.Response {
// 	modules := []models.Schema{}
// 	schemaID := chi.URLParam(r, "schema_id")
// 	schemaIDColumn := fmt.Sprintf("%s.schema_id", constants.ViewCoreSchemaModules)
// 	languageCode := r.Header.Get("Content-Language")
// 	languageCodeColumn := fmt.Sprintf("%s.language_code", constants.ViewCoreSchemaModules)
// 	condition := builder.And(
// 		builder.Equal(schemaIDColumn, schemaID),
// 		builder.Equal(languageCodeColumn, languageCode),
// 	)

// 	return db.Load(r, &modules, "LoadAllModulesBySchema", constants.ViewCoreSchemaModules, condition)
// }

// // RemoveModuleFromSchema deletes object from the database
// func RemoveModuleFromSchema(r *http.Request) *mdlShared.Response {
// 	response := &mdlShared.Response{
// 		Code: http.StatusOK,
// 	}

// 	schemaID := chi.URLParam(r, "schema_id")
// 	moduleID := chi.URLParam(r, "module_id")

// 	statemant := builder.Delete(constants.TableCoreSchemasModels).Where(
// 		builder.And(
// 			builder.Equal("schema_id", schemaID),
// 			builder.Equal("module_id", moduleID),
// 		),
// 	)

// 	err := sql.Exec(statemant)
// 	if err != nil {
// 		response.Code = http.StatusInternalServerError
// 		response.Errors = append(response.Errors, mdlShared.NewResponseError(constants.ErrorDeletingData, "RemoveModuleFromSchema", err.Error()))

// 		return response
// 	}

// 	return response
// }
