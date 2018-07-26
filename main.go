package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/handlers"
	"rest-api-go/store"
)

func main() {
	router := store.NewRouter()
	allowedOrigins := handlers.AllowedOrigins([]string{"*"}) 
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(allowedOrigins, allowedMethods)(router))) 
}