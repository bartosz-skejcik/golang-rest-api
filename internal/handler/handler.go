package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"restapi/internal/functions"
	"restapi/internal/models"
	"strconv"

	"github.com/gorilla/mux"
)

type UsersResponse struct {
	models.Response
	Users []models.User `json:"users"`
}

type SingleUserResponse struct {
	models.Response
	User models.User `json:"user"`
}

// ErrorHandler sends an error response with a given status code
func ErrorHandler(w http.ResponseWriter, message string, code int) {
	response := models.Response{Message: message}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(response)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := models.HelloResponse{Message: "Hello Jello!"}
	json.NewEncoder(w).Encode(response)
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		var users models.Users
		err := functions.GetAllUsers(&users)
		if err != nil {
			ErrorHandler(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := UsersResponse{Response: models.Response{Message: "ok"}, Users: users}
		json.NewEncoder(w).Encode(response)

	case http.MethodPost:
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			ErrorHandler(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		err := functions.CreateUser(&user)
		if err != nil {
			ErrorHandler(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := SingleUserResponse{Response: models.Response{Message: "User created successfully"}, User: user}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)

	default:
		ErrorHandler(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func UsersByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, conversionErr := strconv.Atoi(vars["id"])
	if conversionErr != nil {
		log.Printf("Failed to convert string (%s) to number: %s", vars["id"], conversionErr.Error())
		ErrorHandler(w, "Wrong type of id", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := functions.GetUserById(&user, userId)
	if err != nil {
		log.Printf(err.Error())
		ErrorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := SingleUserResponse{Response: models.Response{Message: "ok"}, User: user}

	json.NewEncoder(w).Encode(response)
}
