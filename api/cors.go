package api

import (
	"net/http"

	"github.com/go-chi/cors"
)

// CORSMiddleware is a wrapper around the go-chi cors middleware
func CORSMiddleware() func(http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		// AllowOriginFunc: // allow all origins for now,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
}
