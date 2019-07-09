package admin

import (
	controller "github.com/agile-work/srv-mdl-core/controllers/admin/content"
	"github.com/go-chi/chi"
)

// ContentRoutes creates the api methods
func ContentRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/admin/contents
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostContent)
		r.Get("/", controller.GetAllContents)
		r.Get("/{content_code}", controller.GetContent)
		r.Patch("/{content_code}", controller.UpdateContent)
		r.Delete("/{content_code}", controller.DeleteContent)
	})

	return r
}
