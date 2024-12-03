package main

import (
  "fmt"
  "log"
  "net/http"
  "simple-api/handlers"
  "simple-api/db"
  "github.com/gorilla/mux"
)


func main() {
	db.ConnectDB()
	db.InitDB()

	router := mux.NewRouter()
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
