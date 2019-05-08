package auth

import (
	"net/http"

	"github.com/agile-work/srv-mdl-core/services"
	"github.com/go-chi/render"
)

// Login endpoint to get user credentials and return token
func Login(w http.ResponseWriter, r *http.Request) {
	response := services.Login(r)

	render.Status(r, response.Code)
	render.JSON(w, r, response)
}
