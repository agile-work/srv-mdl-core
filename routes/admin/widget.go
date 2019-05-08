package admin

import (
	controller "github.com/agile-work/srv-mdl-core/controllers/admin"
	"github.com/go-chi/chi"
)

// WidgetRoutes creates the api methods
func WidgetRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/admin/widgets
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostWidget)
		r.Get("/", controller.GetAllWidgets)
		r.Get("/{widget_id}", controller.GetWidget)
		r.Patch("/{widget_id}", controller.UpdateWidget)
		r.Delete("/{widget_id}", controller.DeleteWidget)
	})

	return r
}
