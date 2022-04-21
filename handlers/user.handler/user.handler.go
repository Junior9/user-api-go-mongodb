package userhandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	m "api/models"

	user_service "api/services/user.service"

	"github.com/gorilla/mux"
)

func Add(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Add user")

	var user m.User
	_ = json.NewDecoder(req.Body).Decode(&user)

	_, err := user_service.Add(user)

	if err != nil {
		json.NewEncoder(w).Encode("Error")
	}

	json.NewEncoder(w).Encode("ok")
}

func Get(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Get users")

	users, _ := user_service.Get()

	json.NewEncoder(w).Encode(users)
}

func GetById(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Get user")
	params := mux.Vars(req)
	id := params["id"]

	users, _ := user_service.GetById(id)

	json.NewEncoder(w).Encode(users)
}
