package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	user_handler "api/handlers/user.handler"
)

func main() {
	fmt.Println("Start")

	r := mux.NewRouter()

	r.HandleFunc("/add", user_handler.Add).Methods("POST")
	r.HandleFunc("/get", user_handler.Get).Methods("GET")
	r.HandleFunc("/get/{id}", user_handler.GetById).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))
}
