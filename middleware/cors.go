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
	allowedHeaders = []string{
		"X-Requested-With",
		"Content-Type",
	}
	allowedMethods = []string{
		"GET",
		"HEAD",
		"POST",
		"PUT",
		"OPTIONS",
	}
)

// CORS wraps handlers used to add CORS headers to responses.
func CORS(handler http.Handler) http.Handler {
	corsOrigin := handlers.AllowedOrigins(allowedOrigins)
	corsHeaders := handlers.AllowedHeaders(allowedHeaders)
	corsMethods := handlers.AllowedMethods(allowedMethods)

	return handlers.CORS(corsHeaders, corsOrigin, corsMethods)(handler)
}
