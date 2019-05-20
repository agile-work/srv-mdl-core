package routes

import (
	"github.com/agile-work/srv-mdl-core/routes/admin"
	"github.com/agile-work/srv-mdl-core/routes/auth"
	"github.com/go-chi/chi"
)

// Setup configure the API endpoints
func Setup() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/core/admin", func(r chi.Router) {
		r.Mount("/configs", admin.ConfigRoutes())
		r.Mount("/users", admin.UserRoutes())
		r.Mount("/trees", admin.TreeRoutes())
		r.Mount("/schemas", admin.SchemaRoutes())
		r.Mount("/widgets", admin.WidgetRoutes())
		r.Mount("/lookups", admin.LookupRoutes())
		r.Mount("/groups", admin.GroupRoutes())
		r.Mount("/currencies", admin.CurrencyRoutes())
		r.Mount("/jobs", admin.JobRoutes())
		r.Mount("/auth", auth.Routes())
	})

	return router
}
