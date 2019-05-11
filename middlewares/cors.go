package middlewares

import (
	"net/http"

	"github.com/go-chi/cors"
)

// Cors defines the security parameters to access to api
func Cors() *cors.Cors {
	cors := cors.New(cors.Options{
		AllowOriginFunc:  allowOriginFunc,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Language", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	return cors
}

// allowOriginFunc defines a list of hosts authorized to access the api
func allowOriginFunc(r *http.Request, origin string) bool {
	if origin == "http://localhost:8081" {
		return true
	}

	return false
}
