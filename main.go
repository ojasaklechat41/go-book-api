package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"go-book-api/routes"
)

func main() {
	routes.InitBooks()

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the book club"))
	}).Methods("GET")

	r.HandleFunc("/books", routes.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", routes.GetBook).Methods("GET")
	r.HandleFunc("/books", routes.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", routes.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", routes.DeleteBook).Methods("DELETE")

	log.Println("Server Started at :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

