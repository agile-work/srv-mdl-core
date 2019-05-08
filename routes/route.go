package routes

import (
	"net/http"

	module "github.com/agile-work/srv-mdl-shared"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

//Routes module endpoints
func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.HandleFunc("/core/*", func(rw http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusOK)
		response := module.Response{}
		response.Code = http.StatusOK
		response.Data = "Core Service - Response"
		render.JSON(rw, r, response)
	})
	return r
}
