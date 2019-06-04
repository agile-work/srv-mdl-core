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
	})

	return r
}
