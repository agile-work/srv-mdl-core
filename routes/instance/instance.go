package instance

import (
	controller "github.com/agile-work/srv-mdl-core/controllers/instance"
	"github.com/go-chi/chi"
)

// Routes creates the api methods
func Routes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/core/instance/
	r.Route("/", func(r chi.Router) {
		r.Post("/schemas/{schema_code}", controller.PostSchemaInstance)
		r.Get("/schemas/{schema_code}", controller.GetAllSchemaInstances)
		r.Get("/schemas/{schema_code}/{instance_id}", controller.GetSchemaInstance)
		r.Patch("/schemas/{schema_code}/{instance_id}", controller.UpdateSchemaInstance)
		r.Delete("/schemas/{schema_code}/{instance_id}", controller.DeleteSchemaInstance)
		// r.Post("/schemas/{schema_code}/{instance_id}/permissions", controller.PostPermission)
		// r.Get("/schemas/{schema_code}/{instance_id}/permissions", controller.GetAllPermissions)
		// r.Patch("/schemas/{schema_code}/{instance_id}/permissions/{permission_id}", controller.UpdatePermission)
		// r.Delete("/schemas/{schema_code}/{instance_id}/permissions/{permission_id}", controller.DeletePermission)
		// resources
		r.Get("/resources", controller.GetAllResources)
		r.Get("/resources/{resource_id}", controller.GetResource)
		r.Patch("/resources/{resource_id}", controller.UpdateResource)
		// lookups
		r.Get("/lookups/{lookup_code}", controller.GetLookupInstance)
	})

	return r
}
