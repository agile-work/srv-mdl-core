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

		r.Patch("/{group_code}/permissions/widgets", controller.UpdatePermissionWidgets)
		r.Patch("/{group_code}/permissions/widgets/{widget_code}", controller.UpdatePermissionWidget)
		r.Delete("/{group_code}/permissions/widgets/{widget_code}", controller.DeletePermissionWidget)
		r.Patch("/{group_code}/permissions/processes", controller.UpdatePermissionProcesses)
		r.Patch("/{group_code}/permissions/processes/{process_code}", controller.UpdatePermissionProcess)
		r.Delete("/{group_code}/permissions/processes/{process_code}", controller.DeletePermissionProcess)
		r.Patch("/{group_code}/permissions/schemas/{schema_code}", controller.UpdatePermissionSchema)
		r.Delete("/{group_code}/permissions/schemas/{schema_code}", controller.DeletePermissionSchema)
		r.Patch("/{group_code}/permissions/schemas/{schema_code}/instances", controller.UpdatePermissionSchemaInstance)
		r.Patch("/{group_code}/permissions/schemas/{schema_code}/fields", controller.UpdatePermissionSchemaField)
		r.Patch("/{group_code}/permissions/schemas/{schema_code}/modules/{module_code}", controller.UpdatePermissionModule)
		r.Delete("/{group_code}/permissions/schemas/{schema_code}/modules/{module_code}", controller.DeletePermissionModule)
		r.Patch("/{group_code}/permissions/schemas/{schema_code}/modules/{module_code}/instances", controller.UpdatePermissionModuleInstance)
		r.Patch("/{group_code}/permissions/schemas/{schema_code}/modules/{module_code}/fields", controller.UpdatePermissionModuleField)
		r.Patch("/{group_code}/permissions/schemas/{schema_code}/modules/{module_code}/features/{feature_code}", controller.UpdatePermissionModuleFeature)
	})

	return r
}
