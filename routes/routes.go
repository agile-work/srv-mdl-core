package routes

import (
	core "github.com/agile-work/srv-mdl-core/middlewares"
	"github.com/agile-work/srv-mdl-core/routes/admin"
	"github.com/agile-work/srv-mdl-core/routes/auth"
	"github.com/go-chi/chi"
)

// Setup configure the API endpoints
func Setup() *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		core.Cors().Handler,
		core.Authorization,
	)
	//TODO: Retirar router.use
	// router.Use(
	// 	render.SetContentType(render.ContentTypeJSON),
	// 	middleware.Logger,
	// 	middleware.DefaultCompress,
	// 	middleware.RedirectSlashes,
	// 	middleware.Recoverer,
	// 	cryo.Authorization,
	// 	cryo.Cors().Handler,
	// )

	router.Route("/core", func(r chi.Router) {
		r.Mount("/admin/configs", admin.ConfigRoutes())
		r.Mount("/admin/users", admin.UserRoutes())
		r.Mount("/admin/trees", admin.TreeRoutes())
		r.Mount("/admin/schemas", admin.SchemaRoutes())
		r.Mount("/admin/widgets", admin.WidgetRoutes())
		r.Mount("/admin/lookups", admin.LookupRoutes())
		r.Mount("/admin/groups", admin.GroupRoutes())
		r.Mount("/admin/currencies", admin.CurrencyRoutes())
		r.Mount("/admin/jobs", admin.JobRoutes())
		r.Mount("/admin/auth", auth.Routes())
	})

	return router
}
