// main.go

package main

import (
    "log"
    "net/http"
    "user-auth-app/database"
    "user-auth-app/handlers"
    "github.com/gorilla/mux"
)

func main() {
    // Connect to the database
    database.ConnectDatabase()

    // Create a new router
    r := mux.NewRouter()

    // Define the routes
    r.HandleFunc("/register", handlers.Register).Methods("GET", "POST")
    r.HandleFunc("/login", handlers.Login).Methods("GET", "POST")
    r.HandleFunc("/success", handlers.Success).Methods("GET")

    // Start the server
    log.Println("Server started at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
