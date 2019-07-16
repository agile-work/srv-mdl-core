package admin

import (
	controller "github.com/agile-work/srv-mdl-core/controllers/admin/module"
	"github.com/go-chi/chi"
)

// ModuleRoutes creates the api methods
func ModuleRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/core/admin/modules
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.RegisterModule)
		r.Get("/", controller.GetAllModules)
		r.Get("/{module_code}", controller.GetModule)
		r.Patch("/{module_code}", controller.UpdateModule)
		r.Delete("/{module_code}", controller.DeleteModule)

		r.Post("/{module_code}/instances", controller.AddModuleInstance)
		r.Delete("/{module_code}/instances", controller.DeleteModuleInstance)
	})

	return r
}
