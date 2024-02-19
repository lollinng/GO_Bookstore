package main

import (
	"go-bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisteredBookStoreRoutes(r) // Accessing the routes

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
