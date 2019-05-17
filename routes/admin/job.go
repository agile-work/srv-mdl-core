package admin

import (
	controller "github.com/agile-work/srv-mdl-core/controllers/admin"
	"github.com/go-chi/chi"
)

// JobRoutes creates the api methods
func JobRoutes() *chi.Mux {
	r := chi.NewRouter()

	// v1/api/admin/jobs
	r.Route("/", func(r chi.Router) {
		r.Post("/", controller.PostJob)
		r.Get("/", controller.GetAllJobs)
		r.Get("/{job_id}", controller.GetJob)
		r.Patch("/{job_id}", controller.UpdateJob)
		r.Delete("/{job_id}", controller.DeleteJob)
		r.Post("/{job_id}/tasks", controller.PostJobTask)
		r.Get("/{job_id}/tasks", controller.GetAllJobTasks)
		r.Get("/{job_id}/tasks/{job_task_id}", controller.GetJobTask)
		r.Patch("/{job_id}/tasks/{job_task_id}", controller.UpdateJobTask)
		r.Delete("/{job_id}/tasks/{job_task_id}", controller.DeleteJobTask)
		r.Get("/followers/avaible", controller.LoadAllJobFollowersAvaible)
		r.Post("/{job_id}/followers/{follower_id}/type/{follower_type}", controller.InsertFollowerInJob)
		r.Get("/{job_id}/followers", controller.LoadAllFollowersByJob)
		r.Delete("/{job_id}/followers/{follower_id}", controller.RemoveFollowerFromJob)
		r.Post("/{job_id}/instance", controller.PostJobInstance)
	})

	return r
}
