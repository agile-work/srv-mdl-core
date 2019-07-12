package admin

import (
	controller "github.com/agile-work/srv-mdl-core/controllers/admin/dataset"
	"github.com/go-chi/chi"
)

// DatasetRoutes creates the api methods
func DatasetRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/core/admin/datasets
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostDataset)
		r.Get("/", controller.GetAllDatasets)
		r.Get("/{dataset_code}", controller.GetDataset)
		r.Patch("/{dataset_code}", controller.UpdateDataset)
		r.Delete("/{dataset_code}", controller.DeleteDataset)
		r.Post("/{dataset_code}/options", controller.AddDatasetOption)
		r.Patch("/{dataset_code}/options/{option_code}", controller.UpdateDatasetOption)
		r.Delete("/{dataset_code}/options/{option_code}", controller.DeleteDatasetOption)
		r.Post("/{dataset_code}/order", controller.UpdateDatasetOrder)
		r.Patch("/{dataset_code}/query", controller.UpdateDatasetQuery)
		r.Patch("/{dataset_code}/{param_type}/{param_code}", controller.UpdateDatasetDynamicParam)
	})

	return r
}
