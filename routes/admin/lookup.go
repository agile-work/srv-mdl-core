package admin

import (
	controller "github.com/agile-work/srv-mdl-core/controllers/admin"
	"github.com/go-chi/chi"
)

// LookupRoutes creates the api methods
func LookupRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/core/admin/lookups
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostLookup)
		r.Get("/", controller.GetAllLookups)
		r.Get("/{lookup_code}", controller.GetLookup)
		r.Patch("/{lookup_code}", controller.UpdateLookup)
		r.Delete("/{lookup_code}", controller.DeleteLookup)
		r.Post("/{lookup_code}/options", controller.AddLookupOption)
		r.Patch("/{lookup_code}/options/{option_code}", controller.UpdateLookupOption)
		r.Delete("/{lookup_code}/options/{option_code}", controller.DeleteLookupOption)
		r.Post("/{lookup_code}/order", controller.UpdateLookupOrder)
		r.Patch("/{lookup_code}/query", controller.UpdateLookupQuery)
		r.Patch("/{lookup_code}/field/{param_code}", controller.UpdateLookupDynamicField)
		r.Patch("/{lookup_code}/param/{param_code}", controller.UpdateLookupDynamicParam)
	})

	return r
}
