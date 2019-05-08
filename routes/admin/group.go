package admin

import (
	controller "github.com/agile-work/srv-mdl-core/controllers/admin"
	"github.com/go-chi/chi"
)

// GroupRoutes creates the api methods
func GroupRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/admin/groups
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostGroup)
		r.Get("/", controller.GetAllGroups)
		r.Get("/{group_id}", controller.GetGroup)
		r.Patch("/{group_id}", controller.UpdateGroup)
		r.Delete("/{group_id}", controller.DeleteGroup)
		r.Post("/{group_id}/users/{user_id}", controller.AddUserInGroup)
		r.Get("/{group_id}/users", controller.GetAllUsersByGroup)
		r.Delete("/{group_id}/users/{user_id}", controller.DeleteGroupUser)
		r.Post("/{group_id}/permissions", controller.PostGroupPermission)
		r.Get("/{group_id}/permissions", controller.GetAllPermissionsByGroup)
		r.Delete("/{group_id}/permissions/{permission_id}", controller.DeleteGroupPermission)
	})

	return r
}
