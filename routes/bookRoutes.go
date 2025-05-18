package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-book-api/models"
	"net/http"
)

var books []models.Book

func InitBooks() {
	books = append(books, models.Book{ID: "1", Title: "1984", Author: "George Orwell"})
	books = append(books, models.Book{ID: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee"})
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range books {
		if item.ID == params["id"] {
			var updatedBook models.Book
			_ = json.NewDecoder(r.Body).Decode(&updatedBook)
			updatedBook.ID = item.ID
			books[i] = updatedBook
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}
	http.NotFound(w, r)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range books {
		if item.ID == params["id"] {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}