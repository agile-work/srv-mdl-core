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
		r.Get("/{currency_id}", controller.GetCurrency)
		r.Patch("/{currency_id}", controller.UpdateCurrency)
		r.Delete("/{currency_id}", controller.DeleteCurrency)
		r.Post("/{currency_id}/rates", controller.PostCurrencyRate)
		r.Get("/{currency_id}/rates", controller.GetAllCurrencyRates)
		r.Get("/{currency_id}/rates/{currency_rate_id}", controller.GetCurrencyRate)
		r.Patch("/{currency_id}/rates/{currency_rate_id}", controller.UpdateCurrencyRate)
		r.Delete("/{currency_id}/rates/{currency_rate_id}", controller.DeleteCurrencyRate)
	})

	return r
}
