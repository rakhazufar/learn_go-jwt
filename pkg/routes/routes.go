package routes

import (
	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-jwt/pkg/controllers/authcontrollers"
	"github.com/rakhazufar/go-jwt/pkg/controllers/productcontroller"
	"github.com/rakhazufar/go-jwt/pkg/middlewares"
)

func Authentication(router *mux.Router) {
	router.HandleFunc("/login/", authcontrollers.Login).Methods("POST")
	router.HandleFunc("/register/", authcontrollers.Register).Methods("POST")
	router.HandleFunc("/logout/", authcontrollers.Logout).Methods("GET")
}


func Products (router *mux.Router) {
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products/", productcontroller.GetProduct).Methods("GET")
	api.Use(middlewares.JWTMiddleware)
} 