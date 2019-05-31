package admin

import (
	"net/http"

	services "github.com/agile-work/srv-mdl-core/services/admin"

	"github.com/go-chi/render"
)

// PostCurrency sends the request to service creating a new schema
func PostCurrency(w http.ResponseWriter, r *http.Request) {
	response := services.CreateCurrency(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetAllCurrencies return all schema instances from the service
func GetAllCurrencies(w http.ResponseWriter, r *http.Request) {
	response := services.LoadAllCurrencies(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// GetCurrency return only one schema from the service
func GetCurrency(w http.ResponseWriter, r *http.Request) {
	response := services.LoadCurrency(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// UpdateCurrency sends the request to service updating a schema
func UpdateCurrency(w http.ResponseWriter, r *http.Request) {
	response := services.UpdateCurrency(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// DeleteCurrency sends the request to service deleting a schema
func DeleteCurrency(w http.ResponseWriter, r *http.Request) {
	response := services.DeleteCurrency(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}

// AddCurrencyRate sends the request to service creating a new schema
func AddCurrencyRate(w http.ResponseWriter, r *http.Request) {
	response := services.AddCurrencyRate(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
