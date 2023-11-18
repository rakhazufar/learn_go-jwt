package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-jwt/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.Authentication(r)
	routes.Products(r)

	log.Fatal(http.ListenAndServe(":9010", r))
}