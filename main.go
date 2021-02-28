package main

import (
	"./app"
	"./controller"
	"./helper"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("It Works")

	port := app.GetParameters().ServerPort()

	router := mux.NewRouter()
	router.Use(helper.JwtAuthentication)
	router.HandleFunc("/api/user/register",
		controller.Register).Methods("POST")
	router.HandleFunc("/api/user/login",
		controller.Login).Methods("POST")
	router.HandleFunc("/api/user/refresh",
		controller.Refresh).Methods("POST")
	router.HandleFunc("/api/user/logout",
		controller.Logout).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
