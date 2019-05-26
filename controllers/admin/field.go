package admin

import (
	"net/http"

	services "github.com/agile-work/srv-mdl-core/services/admin"

	"github.com/go-chi/render"
)

// PostField sends the request to service creating a new field
func PostField(w http.ResponseWriter, r *http.Request) {
	response := services.CreateField(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllFields return all field instances from the service
func GetAllFields(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllFields(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetField return only one field from the service
func GetField(w http.ResponseWriter, r *http.Request) {
	response := services.LoadField(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateField sends the request to service updating a field
func UpdateField(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateField(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteField sends the request to service deleting a field
func DeleteField(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteField(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// PostFieldGroup sends the request to service deleting a schema
func PostFieldGroup(w http.ResponseWriter, r *http.Request) {
	response := services.InsertGroupToField(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteFieldGroup sends the request to service deleting a schema
func DeleteFieldGroup(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteGroupFromField(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// PostFieldValidation sends the request to service creating a new field
func PostFieldValidation(w http.ResponseWriter, r *http.Request) {
	response := services.CreateFieldValidation(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllFieldValidations return all field instances from the service
func GetAllFieldValidations(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllFieldValidations(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetFieldValidation return only one field from the service
func GetFieldValidation(w http.ResponseWriter, r *http.Request) {
	response := services.LoadFieldValidation(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateFieldValidation sends the request to service updating a field
func UpdateFieldValidation(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateFieldValidation(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteFieldValidation sends the request to service deleting a field
func DeleteFieldValidation(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteFieldValidation(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
