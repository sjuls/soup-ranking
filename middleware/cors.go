package middleware

import (
	"net/http"

	"github.com/gorilla/handlers"
)

var (
	allowedOrigins = []string{
		"https://junesoup.surge.sh",
		"http://junesoup.surge.sh",
		"https://soup-ranking.herokuapp.com",
		"http://soup-ranking.herokuapp.com",
	}
)

// CORS wraps handlers used to add CORS headers to responses.
func CORS(handler http.Handler) http.Handler {
	corsHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	corsOrigin := handlers.AllowedOrigins(allowedOrigins)
	corsMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	return handlers.CORS(corsHeaders, corsOrigin, corsMethods)(handler)
}
