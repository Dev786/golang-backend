package main

import (
	"log"
	"net/http"

	"../../handlers/registration"

	"github.com/gorilla/mux"
)

func main() {
	app := mux.NewRouter()
	app.HandleFunc("/register", registerAndLogin.Register).Methods("POST")
	app.HandleFunc("/login", registerAndLogin.Login).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", app))
}
