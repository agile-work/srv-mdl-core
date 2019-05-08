package auth

import (
	controller "github.com/agile-work/srv-mdl-core/controllers/auth"
	"github.com/go-chi/chi"
)

// Routes defines authentication endpoints
func Routes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/auth
	r.Route("/", func(r chi.Router) {
		r.Post("/login", controller.Login)
	})

	return r
}
