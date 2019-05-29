package admin

import (
	controller "github.com/agile-work/srv-mdl-core/controllers/admin"
	"github.com/go-chi/chi"
)

// TreeRoutes creates the api methods
func TreeRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/admin/trees
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostTree)
		r.Get("/", controller.GetAllTrees)
		r.Get("/{tree_id}", controller.GetTree)
		r.Patch("/{tree_id}", controller.UpdateTree)
		r.Delete("/{tree_id}", controller.DeleteTree)
		// Units
		r.Post("/{tree_id}/units", controller.PostTreeUnit)
		r.Get("/{tree_id}/units", controller.GetAllTreeUnits)
		r.Get("/{tree_id}/units/{tree_unit_id}", controller.GetTreeUnit)
		r.Patch("/{tree_id}/units/{tree_unit_id}", controller.UpdateTreeUnit)
		r.Delete("/{tree_id}/units/{tree_unit_id}", controller.DeleteTreeUnit)
		// Permissions
		r.Get("/{tree_id}/units/{tree_unit_id}/permissions", controller.GetAllTreeUnitPermissions)
		r.Post("/{tree_id}/units/{tree_unit_id}/permissions", controller.PostTreeUnitPermission)
		r.Delete("/{tree_id}/units/{tree_unit_id}/permissions/{permission_id}", controller.DeleteTreeUnitPermission)
		// Levels
		r.Post("/{tree_id}/levels", controller.PostTreeLevel)
		r.Get("/{tree_id}/levels", controller.GetAllTreeLevels)
		r.Get("/{tree_id}/levels/{tree_level_id}", controller.GetTreeLevel)
		r.Patch("/{tree_id}/levels/{tree_level_id}", controller.UpdateTreeLevel)
		r.Delete("/{tree_id}/levels/{tree_level_id}", controller.DeleteTreeLevel)
	})

	return r
}
