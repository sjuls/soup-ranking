package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sjuls/soup-ranking/dbctx"
	"github.com/sjuls/soup-ranking/routes"
)

var (
	allowedOrigins = []string{
		"https://junesoup.surge.sh",
		"http://junesoup.surge.sh",
		"https://soup-ranking.herokuapp.com",
		"http://soup-ranking.herokuapp.com",
	}
)

func main() {
	port := os.Getenv("PORT")
	database := os.Getenv("DATABASE_URL")

	if err := dbctx.Init(&database); err != nil {
		panic(err)
	}

	router := mux.NewRouter().StrictSlash(true)
	routes.AddStatus(router)
	routes.AddScore(router)

	log.Fatal(http.ListenAndServe(":"+port, wrapCORS(router)))
}

func wrapCORS(router *mux.Router) http.Handler {
	corsHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	corsOrigin := handlers.AllowedOrigins(allowedOrigins)
	corsMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	return handlers.CORS(corsHeaders, corsOrigin, corsMethods)(router)
}
