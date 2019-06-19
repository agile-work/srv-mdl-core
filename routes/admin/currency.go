package admin

import (
	controller "github.com/agile-work/srv-mdl-core/controllers/admin"
	"github.com/go-chi/chi"
)

// CurrencyRoutes creates the api methods
func CurrencyRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/admin/currencies
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostCurrency)
		r.Get("/", controller.GetAllCurrencies)
		r.Get("/{currency_code}", controller.GetCurrency)
		r.Patch("/{currency_code}", controller.UpdateCurrency)
		r.Delete("/{currency_code}", controller.DeleteCurrency)
		r.Post("/{currency_code}/rates/{to_currency_code}", controller.AddRate)
	})

	return r
}
