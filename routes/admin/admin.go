package admin

import (
	"github.com/go-chi/chi"
)

// Routes creates the api methods
func Routes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/admin/configs
	r.Route("/", func(r chi.Router) {
		r.Mount("/configs", ConfigRoutes())
		r.Mount("/users", UserRoutes())
		r.Mount("/trees", TreeRoutes())
		r.Mount("/schemas", SchemaRoutes())
		r.Mount("/widgets", WidgetRoutes())
		r.Mount("/lookups", LookupRoutes())
		r.Mount("/groups", GroupRoutes())
		r.Mount("/currencies", CurrencyRoutes())
		r.Mount("/jobs", JobRoutes())
	})

	return r
}
