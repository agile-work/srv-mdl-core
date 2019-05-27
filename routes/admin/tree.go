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
		r.Get("/{tree_code}", controller.GetTree)
		r.Patch("/{tree_code}", controller.UpdateTree)
		r.Delete("/{tree_code}", controller.DeleteTree)
		// Units
		r.Post("/{tree_code}/units", controller.PostTreeUnit)
		r.Get("/{tree_code}/units", controller.GetAllTreeUnits)
		r.Get("/{tree_code}/units/{tree_unit_code}", controller.GetTreeUnit)
		r.Patch("/{tree_code}/units/{tree_unit_code}", controller.UpdateTreeUnit)
		r.Delete("/{tree_code}/units/{tree_unit_code}", controller.DeleteTreeUnit)
		// Permissions
		r.Post("/{tree_code}/units/{tree_unit_code}/permissions", controller.PostTreeUnitPermission)
		r.Delete("/{tree_code}/units/{tree_unit_code}/permissions/{permission_id}", controller.DeleteTreeUnitPermission)
		// Levels
		r.Post("/{tree_code}/levels", controller.PostTreeLevel)
		r.Get("/{tree_code}/levels", controller.GetAllTreeLevels)
		r.Get("/{tree_code}/levels/{tree_level_code}", controller.GetTreeLevel)
		r.Patch("/{tree_code}/levels/{tree_level_code}", controller.UpdateTreeLevel)
		r.Delete("/{tree_code}/levels/{tree_level_code}", controller.DeleteTreeLevel)
	})

	return r
}
