package admin

import (
	controller "github.com/agile-work/srv-mdl-core/controllers/admin"
	"github.com/go-chi/chi"
)

// SchemaRoutes creates the api methods
func SchemaRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/admin/schemas
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostSchema)
		r.Get("/", controller.GetAllSchemas)
		r.Get("/{schema_code}", controller.GetSchema)
		r.Patch("/{schema_code}", controller.UpdateSchema)
		r.Delete("/{schema_code}", controller.DeleteSchema)
		r.Delete("/{schema_code}/start", controller.CallDeleteSchema)
		r.Get("/{schema_code}/modules", controller.GetAllModulesBySchema)
		//fields
		r.Post("/{schema_code}/fields", controller.PostField)
		r.Get("/{schema_code}/fields", controller.GetAllFields)
		r.Get("/{schema_code}/fields/{field_code}", controller.GetField)
		r.Patch("/{schema_code}/fields/{field_code}", controller.UpdateField)
		r.Delete("/{schema_code}/fields/{field_code}", controller.DeleteField)
	})

	return r
}
