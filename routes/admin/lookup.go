package admin

import (
	controller "github.com/agile-work/srv-mdl-core/controllers/admin"
	"github.com/go-chi/chi"
)

// LookupRoutes creates the api methods
func LookupRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/admin/lookups
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostLookup)
		r.Get("/", controller.GetAllLookups)
		r.Get("/{lookup_id}", controller.GetLookup)
		r.Patch("/{lookup_id}", controller.UpdateLookup)
		r.Delete("/{lookup_id}", controller.DeleteLookup)
		r.Post("/{lookup_id}/options", controller.PostLookupOption)
		r.Get("/{lookup_id}/options", controller.GetAllLookupOptions)
		r.Get("/{lookup_id}/options/{option_id}", controller.GetLookupOption)
		r.Patch("/{lookup_id}/options/{option_id}", controller.PostLookupOption)
		r.Delete("/{lookup_id}/options/{option_id}", controller.GetLookupOption)
	})

	return r
}
