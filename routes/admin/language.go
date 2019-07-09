package admin

import (
	controller "github.com/agile-work/srv-mdl-core/controllers/admin"
	"github.com/go-chi/chi"
)

// LanguageRoutes creates the api methods
func LanguageRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/admin/Languages
	r.Route("/languages", func(r chi.Router) {
		r.Post("/", controller.PostLanguage)
		r.Get("/", controller.GetAllLanguages)
		r.Get("/{language_id}", controller.GetLanguage)
		r.Patch("/{language_id}", controller.UpdateLanguage)
		r.Delete("/{language_id}", controller.DeleteLanguage)
	})

	return r
}
