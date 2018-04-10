package middleware

import (
	"log"
	"net/http"
	"time"
)

// Log is used to output access of endpoints.
func Log(handler http.Handler) http.Handler {
	logFn := func(w http.ResponseWriter, r *http.Request) {
		s := time.Now()
		handler.ServeHTTP(w, r)
		e := time.Now()
		log.Printf("- %s %s [%d] (%s)", r.Method, r.RequestURI, r.ContentLength, e.Sub(s))
	}

	return http.HandlerFunc(logFn)
}
