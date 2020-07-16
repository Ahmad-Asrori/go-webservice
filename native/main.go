package main

import (
	"carakan-apps/Middleware"
	"carakan-apps/configuration"
	"carakan-apps/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.Use(Middleware.LoggerMiddleware)
	router.Use(Middleware.JWTMiddleware)

	router.HandleFunc("/auth", controller.Authentication).Methods("GET")
	router.HandleFunc("/account/all", controller.GetAllAccount).Methods("GET")
	router.HandleFunc("/account", controller.GetAccountById).Methods("GET")
	router.HandleFunc("/account/insert", controller.InsertAccount).Methods("POST")
	router.HandleFunc("/account/update", controller.UpdateAccount).Methods("PUT")
	router.HandleFunc("/account/delete", controller.DeleteAccount).Methods("DELETE")

	http.ListenAndServe(configuration.SERVER_ADDR, router)
}