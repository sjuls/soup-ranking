package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/sjuls/soup-ranking/dbctx"
	"github.com/sjuls/soup-ranking/routes"
	"log"
	"net/http"
	"os"
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

	corsHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})

	corsOrigin:=handlers.AllowedOrigins([]string{
		"https://junesoup.surge.sh",
		"http://junesoup.surge.sh",
		"https://soup-ranking.herokuapp.com",
		"http://soup-ranking.herokuapp.com",
	})
	corsMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(corsHeaders, corsOrigin, corsMethods)(router)))
}
