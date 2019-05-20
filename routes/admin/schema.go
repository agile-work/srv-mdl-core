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
		r.Get("/{schema_id}", controller.GetSchema)
		r.Patch("/{schema_id}", controller.UpdateSchema)
		r.Delete("/{schema_id}", controller.DeleteSchema)
		r.Delete("/{schema_id}/start", controller.CallDeleteSchema)
		r.Get("/{schema_id}/modules", controller.GetAllModulesBySchema)
		r.Post("/{schema_id}/fields", controller.PostField)
		r.Get("/{schema_id}/fields", controller.GetAllFields)
		r.Get("/{schema_id}/fields/{field_id}", controller.GetField)
		r.Patch("/{schema_id}/fields/{field_id}", controller.UpdateField)
		r.Delete("/{schema_id}/fields/{field_id}", controller.DeleteField)
		r.Post("/{schema_id}/fields/{field_id}/validations", controller.PostFieldValidation)
		r.Get("/{schema_id}/fields/{field_id}/validations", controller.GetAllFieldValidations)
		r.Get("/{schema_id}/fields/{field_id}/validations/{field_validation_id}", controller.GetFieldValidation)
		r.Patch("/{schema_id}/fields/{field_id}/validations/{field_validation_id}", controller.UpdateFieldValidation)
		r.Delete("/{schema_id}/fields/{field_id}/validations/{field_validation_id}", controller.DeleteFieldValidation)
		r.Post("/{schema_id}/pages", controller.PostPage)
		r.Get("/{schema_id}/pages", controller.GetAllPages)
		r.Get("/{schema_id}/pages/{page_id}", controller.GetPage)
		r.Patch("/{schema_id}/pages/{page_id}", controller.UpdatePage)
		r.Delete("/{schema_id}/pages/{page_id}", controller.DeletePage)
		r.Post("/{schema_id}/pages/{page_id}/sections", controller.PostSection)
		r.Get("/{schema_id}/pages/{page_id}/sections", controller.GetAllSections)
		r.Get("/{schema_id}/pages/{page_id}/sections/{section_id}", controller.GetSection)
		r.Patch("/{schema_id}/pages/{page_id}/sections/{section_id}", controller.UpdateSection)
		r.Delete("/{schema_id}/pages/{page_id}/sections/{section_id}", controller.DeleteSection)
		r.Post("/{schema_id}/pages/{page_id}/sections/{section_id}/tabs", controller.PostTab)
		r.Get("/{schema_id}/pages/{page_id}/sections/{section_id}/tabs", controller.GetAllTabs)
		r.Get("/{schema_id}/pages/{page_id}/sections/{section_id}/tabs/{tab_id}", controller.GetTab)
		r.Patch("/{schema_id}/pages/{page_id}/sections/{section_id}/tabs/{tab_id}", controller.UpdateTab)
		r.Delete("/{schema_id}/pages/{page_id}/sections/{section_id}/tabs/{tab_id}", controller.DeleteTab)
		r.Post("/{schema_id}/pages/{page_id}/containers/{container_id}/{type}/structures", controller.PostContainerStructure)
		r.Get("/{schema_id}/pages/{page_id}/containers/{container_id}/{type}/structures", controller.GetAllContainerStructures)
		r.Get("/{schema_id}/pages/{page_id}/containers/{container_id}/{type}/structures/{container_structure_id}", controller.GetContainerStructure)
		r.Patch("/{schema_id}/pages/{page_id}/containers/{container_id}/{type}/structures/{container_structure_id}", controller.UpdateContainerStructure)
		r.Delete("/{schema_id}/pages/{page_id}/containers/{container_id}/{type}/structures/{container_structure_id}", controller.DeleteContainerStructure)
		r.Post("/{schema_id}/views", controller.PostView)
		r.Get("/{schema_id}/views", controller.GetAllViews)
		r.Get("/{schema_id}/views/{view_id}", controller.GetView)
		r.Patch("/{schema_id}/views/{view_id}", controller.UpdateView)
		r.Delete("/{schema_id}/views/{view_id}", controller.DeleteView)
	})

	return r
}
