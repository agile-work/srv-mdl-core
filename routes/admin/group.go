package admin

import (
	controller "github.com/agile-work/srv-mdl-core/controllers/admin/group"
	"github.com/go-chi/chi"
)

// GroupRoutes creates the api methods
func GroupRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/admin/groups
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostGroup)
		r.Get("/", controller.GetAllGroups)
		r.Get("/{group_code}", controller.GetGroup)
		r.Patch("/{group_code}", controller.UpdateGroup)
		r.Delete("/{group_code}", controller.DeleteGroup)

		r.Patch("/{group_code}/trees", controller.UpdateTree)
		r.Patch("/{group_code}/users", controller.UpdateUsers)
	})

	return r
}
