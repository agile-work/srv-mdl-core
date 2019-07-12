package admin

import (
	controller "github.com/agile-work/srv-mdl-core/controllers/admin/job"
	"github.com/go-chi/chi"
)

// JobRoutes creates the api methods
func JobRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/admin/jobs
	r.Route("/", func(r chi.Router) {
		// Job
		r.Post("/", controller.PostJob)
		r.Get("/", controller.GetAllJobs)
		r.Get("/{job_id}", controller.GetJob)
		r.Patch("/{job_id}", controller.UpdateJob)
		r.Delete("/{job_id}", controller.DeleteJob)

		// Task
		r.Post("/{job_id}/tasks", controller.PostTask)
		r.Get("/{job_id}/tasks", controller.GetAllTasks)
		r.Get("/{job_id}/tasks/{job_task_id}", controller.GetTask)
		r.Patch("/{job_id}/tasks/{job_task_id}", controller.UpdateTask)
		r.Delete("/{job_id}/tasks/{job_task_id}", controller.DeleteTask)

		// r.Post("/{job_id}/instance", controller.PostJobInstance)
		// r.Get("/instances", controller.GetAllJobInstances)
		// r.Get("/{job_instance_id}/instances/tasks/instances", controller.GetAllTaskInstances)
		// r.Get("/followers/available", controller.LoadAllJobFollowersAvailable)
		// r.Post("/{job_id}/followers/{follower_id}/type/{follower_type}", controller.InsertFollowerInJob)
		// r.Get("/{job_id}/followers", controller.LoadAllFollowersByJob)
		// r.Delete("/{job_id}/followers/{follower_id}", controller.RemoveFollowerFromJob)
	})

	return r
}
