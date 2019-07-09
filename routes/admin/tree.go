package admin

import (
	controller "github.com/agile-work/srv-mdl-core/controllers/admin/tree"
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
		r.Post("/{tree_code}/units", controller.PostUnit)
		r.Get("/{tree_code}/units", controller.GetAllUnits)
		r.Get("/{tree_code}/units/{unit_code}", controller.GetUnit)
		r.Patch("/{tree_code}/units/{unit_code}", controller.UpdateUnit)
		r.Delete("/{tree_code}/units/{unit_code}", controller.DeleteUnit)
		// Levels
		r.Post("/{tree_code}/levels", controller.PostLevel)
		r.Get("/{tree_code}/levels", controller.GetAllLevels)
		r.Get("/{tree_code}/levels/{level_code}", controller.GetLevel)
		r.Patch("/{tree_code}/levels/{level_code}", controller.UpdateLevel)
		r.Delete("/{tree_code}/levels/{level_code}", controller.DeleteLevel)
	})

	return r
}
