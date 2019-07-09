package admin

import (
	controller "github.com/agile-work/srv-mdl-core/controllers/admin/bpm"
	"github.com/go-chi/chi"
)

// ContentRoutes creates the api methods
func WorkflowRoutes() *chi.Mux {
	r := chi.NewRouter()

	// api/v1/core/admin/workflow
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostWorkflow)
		r.Get("/", controller.GetAllWorkflowss)
		r.Get("/{workflow_code}", controller.GetWorkflow)
		r.Patch("/{workflow_code}", controller.UpdateWorkflow)
		r.Delete("/{workflow_code}", controller.DeleteWorkflow)
		// Definition Steps
		r.Post("/", controller.PostStep)
		r.Get("/", controller.GetAllSteps)
		r.Get("/{step_code}", controller.GetStep)
		r.Patch("/{step_code}", controller.UpdateStep)
		r.Delete("/{step_code}", controller.DeleteStep)
		// Instance Workflow
		r.Post("/{workflow_code}/instances", controller.PostWorkflowInstance)
		r.Get("/{workflow_code}/instances", controller.GetAllWorkflowInstances)
		r.Get("/{workflow_code}/instances/{instance_id}", controller.GetWorkflowInstance)
		r.Patch("/{workflow_code}/instances/{instance_id}", controller.UpdateWorkflowInstance)
		r.Delete("/{workflow_code}/instances/{instance_id}", controller.DeleteWorkflowInstance)
		// Instance Steps
		r.Post("/{workflow_code}/instances/{instance_id}/steps", controller.PostStepInstance)
		r.Get("/{workflow_code}/instances/{instance_id}/steps", controller.GetAllStepInstances)
		r.Get("/{workflow_code}/instances/{instance_id}/steps/{step_id}", controller.GetStepInstance)
		r.Patch("/{workflow_code}/instances/{instance_id}/steps/{step_id}", controller.UpdateStepInstance)
		r.Delete("/{workflow_code}/instances/{instance_id}/steps/{step_id}", controller.DeleteStepInstance)
	})

	return r