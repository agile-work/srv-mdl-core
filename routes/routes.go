package routes

import (
	"github.com/agile-work/srv-mdl-core/routes/admin"
	"github.com/agile-work/srv-mdl-core/routes/auth"
	"github.com/agile-work/srv-mdl-core/routes/instance"
	"github.com/go-chi/chi"
)

// Setup configure the API endpoints
func Setup() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/core", func(r chi.Router) {
		r.Mount("/admin", admin.Routes())
		//TODO: Ajustar o endpoint para remover o /admin da parte de autenticacao. Ajustar no client-admin-web
		r.Mount("/admin/auth", auth.Routes())
		r.Mount("/instance", instance.Routes())
	})

	return router
}
