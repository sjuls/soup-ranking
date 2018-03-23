package main

import (
	"github.com/gorilla/mux"
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

	log.Fatal(http.ListenAndServe(":"+port, router))
}
