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
		// //Pages
		// r.Post("/{schema_code}/pages", controller.PostPage)
		// r.Get("/{schema_code}/pages", controller.GetAllPages)
		// r.Get("/{schema_code}/pages/{page_id}", controller.GetPage)
		// r.Patch("/{schema_code}/pages/{page_id}", controller.UpdatePage)
		// r.Delete("/{schema_code}/pages/{page_id}", controller.DeletePage)
		// r.Post("/{schema_code}/pages/{page_id}/sections", controller.PostSection)
		// r.Get("/{schema_code}/pages/{page_id}/sections", controller.GetAllSections)
		// r.Get("/{schema_code}/pages/{page_id}/sections/{section_id}", controller.GetSection)
		// r.Patch("/{schema_code}/pages/{page_id}/sections/{section_id}", controller.UpdateSection)
		// r.Delete("/{schema_code}/pages/{page_id}/sections/{section_id}", controller.DeleteSection)
		// r.Post("/{schema_code}/pages/{page_id}/sections/{section_id}/tabs", controller.PostTab)
		// r.Get("/{schema_code}/pages/{page_id}/sections/{section_id}/tabs", controller.GetAllTabs)
		// r.Get("/{schema_code}/pages/{page_id}/sections/{section_id}/tabs/{tab_id}", controller.GetTab)
		// r.Patch("/{schema_code}/pages/{page_id}/sections/{section_id}/tabs/{tab_id}", controller.UpdateTab)
		// r.Delete("/{schema_code}/pages/{page_id}/sections/{section_id}/tabs/{tab_id}", controller.DeleteTab)
		// r.Post("/{schema_code}/pages/{page_id}/containers/{container_id}/{type}/structures", controller.PostContainerStructure)
		// r.Get("/{schema_code}/pages/{page_id}/containers/{container_id}/{type}/structures", controller.GetAllContainerStructures)
		// r.Get("/{schema_code}/pages/{page_id}/containers/{container_id}/{type}/structures/{container_structure_id}", controller.GetContainerStructure)
		// r.Patch("/{schema_code}/pages/{page_id}/containers/{container_id}/{type}/structures/{container_structure_id}", controller.UpdateContainerStructure)
		// r.Delete("/{schema_code}/pages/{page_id}/containers/{container_id}/{type}/structures/{container_structure_id}", controller.DeleteContainerStructure)
		// //Views
		// r.Post("/{schema_code}/views", controller.PostView)
		// r.Get("/{schema_code}/views", controller.GetAllViews)
		// r.Get("/{schema_code}/views/{view_id}", controller.GetView)
		// r.Patch("/{schema_code}/views/{view_id}", controller.UpdateView)
		// r.Delete("/{schema_code}/views/{view_id}", controller.DeleteView)
	})

	return r
}
