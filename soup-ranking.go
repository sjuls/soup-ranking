package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sjuls/soup-ranking/dbctx"
	"github.com/sjuls/soup-ranking/middleware"
	"github.com/sjuls/soup-ranking/score"
	"github.com/sjuls/soup-ranking/slack"
	"github.com/sjuls/soup-ranking/status"
)

var (
	middlewares = [](func(handler http.Handler) http.Handler){
		middleware.Log,
		middleware.CORS,
	}
)

func main() {
	port := os.Getenv("PORT")
	database := os.Getenv("DATABASE_URL")
	slackToken := os.Getenv("SLACK_TOKEN")

	if err := dbctx.Init(&database); err != nil {
		panic(err)
	}

	router := mux.NewRouter().StrictSlash(true)
	routes := []func(router *mux.Router){
		status.AddRoute,
		score.AddRoute,
		slack.AddRoute(slackToken),
	}

	registerRoutes(router, routes)

	log.Fatal(http.ListenAndServe(":"+port, applyMiddleware(router)))
}

func registerRoutes(router *mux.Router, routes []func(router *mux.Router)) {
	for _, route := range routes {
		route(router)
	}
}

func applyMiddleware(handler http.Handler) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
