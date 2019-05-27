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
		r.Post("/{schema_code}", controller.PostInstance)
		r.Get("/{schema_code}", controller.GetAllInstances)
		r.Get("/{schema_code}/{instance_id}", controller.GetInstance)
		r.Patch("/{schema_code}/{instance_id}", controller.UpdateInstance)
		r.Delete("/{schema_code}/{instance_id}", controller.DeleteInstance)
		r.Post("/{schema_code}/{instance_id}/permissions", controller.PostPermission)
		r.Get("/{schema_code}/{instance_id}/permissions", controller.GetAllPermissions)
		r.Patch("/{schema_code}/{instance_id}/permissions/{permission_id}", controller.UpdatePermission)
		r.Delete("/{schema_code}/{instance_id}/permissions/{permission_id}", controller.DeletePermission)
		// resources
		r.Get("/resources", controller.GetAllResources)
		r.Get("/resources/{resource_id}", controller.GetResource)
		r.Patch("/resources/{resource_id}", controller.UpdateResource)
	})

	return r
}
