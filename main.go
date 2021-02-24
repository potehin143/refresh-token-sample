package main

import (
	"./controller"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("It Works")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" //localhost
	}

	fmt.Println(port)

	router := mux.NewRouter()
	router.HandleFunc("/api/user/login",
		controller.Authenticate).Methods("POST")
	router.HandleFunc("/api/user/refresh",
		controller.Refresh).Methods("POST")
	router.HandleFunc("/api/user/{id}",
		controller.GetUser).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
