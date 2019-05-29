package admin

import (
	controller "github.com/agile-work/srv-mdl-core/controllers/admin"
	"github.com/go-chi/chi"
)

// UserRoutes creates the api methods
func UserRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/admin/users
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostUser)
		r.Get("/", controller.GetAllUsers)
		r.Get("/{user_id}", controller.GetUser)
		r.Patch("/{user_id}", controller.UpdateUser)
		r.Delete("/{user_id}", controller.DeleteUser)
		r.Get("/{user_id}/groups", controller.GetAllGroupsByUser)
		r.Post("/{user_id}/groups/{group_id}", controller.AddGroupInUser)
		r.Delete("/{user_id}/groups/{group_id}", controller.RemoveGroupFromUser)
		r.Get("/{user_id}/permissions", controller.GetAllPermissionsByUser)
	})

	return r
}
