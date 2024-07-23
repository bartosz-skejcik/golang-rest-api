package main

import (
	"log"
	"net/http"
	"restapi/internal/handler"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Printf("%s: %s", r.Method, r.RequestURI)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.HelloHandler)
	r.HandleFunc("/users", handler.UsersHandler).Methods("GET", "POST")
	r.HandleFunc("/users/{id}", handler.UsersByIdHandler).Methods("GET")
	r.Use(loggingMiddleware)
	http.Handle("/", r)
	log.Print("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
