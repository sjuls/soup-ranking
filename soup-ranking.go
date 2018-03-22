package main

import (
	"os"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/sjuls/soup-ranking/routes"
	"github.com/sjuls/soup-ranking/dbctx"
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

	log.Fatal(http.ListenAndServe(":" + port, router))
}

