package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"
	"go-book-api/models"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", GetBook).Methods("GET")
	router.HandleFunc("/books", CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")
	return router
}

func TestGetBooks(t *testing.T) {
	InitBooks()
	req, _ := http.NewRequest("GET", "/books", nil)
	rr := httptest.NewRecorder()

	SetupRouter().ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", status)
	}

	var resp []models.Book
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatal("could not decode response")
	}

	if len(resp) != 2 {
		t.Errorf("expected 2 books, got %d", len(resp))
	}
}

func TestGetBook(t *testing.T) {
	InitBooks()
	req, _ := http.NewRequest("GET", "/books/1", nil)
	rr := httptest.NewRecorder()

	SetupRouter().ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", status)
	}

	var resp models.Book
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatal("could not decode response")
	}

	if resp.ID != "1" {
		t.Errorf("expected book ID 1, got %s", resp.ID)
	}
}

func TestCreateBook(t *testing.T) {
	books = []models.Book{}
	newBook := models.Book{ID: "3", Title: "Book 3", Author: "Author"}
	body, _ := json.Marshal(newBook)

	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	SetupRouter().ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code 201, got %v", status)
	}

	var resp models.Book
	json.Unmarshal(rr.Body.Bytes(), &resp)

	if resp.Title != "Book 3" {
		t.Errorf("expected book title 'Book 3', got %s", resp.Title)
	}
}

func TestUpdateBook(t *testing.T) {
    books = []models.Book{{ID: "1", Title: "Old Title", Author: "Author"}}
    updated := models.Book{Title: "New Title", Author: "Author"}
    body, _ := json.Marshal(updated)

    req, _ := http.NewRequest("PUT", "/books/1", bytes.NewBuffer(body))
    rr := httptest.NewRecorder()

    SetupRouter().ServeHTTP(rr, req)

    if rr.Code != http.StatusOK {
        t.Fatalf("expected status 200, got %v", rr.Code)
    }

    var result models.Book
    json.Unmarshal(rr.Body.Bytes(), &result)

    if result.Title != "New Title" {
        t.Errorf("expected New Title, got %s", result.Title)
    }
}

func TestDeleteBook(t *testing.T) {
    books = []models.Book{{ID: "1", Title: "Book to Delete", Author: "X"}}
    req, _ := http.NewRequest("DELETE", "/books/1", nil)
    rr := httptest.NewRecorder()

    SetupRouter().ServeHTTP(rr, req)

    if rr.Code != http.StatusNoContent {
        t.Errorf("expected status 204, got %v", rr.Code)
    }

    if len(books) != 0 {
        t.Errorf("expected 0 books, got %d", len(books))
    }
}
