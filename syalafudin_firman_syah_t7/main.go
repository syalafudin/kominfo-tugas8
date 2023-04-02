package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id,omitempty"`
	Title  string  `json:"title,omitempty"`
	Author string  `json:"author,omitempty"`
	Price  float32 `json:"price,omitempty"`
}

var books []Book

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(nil)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, book := range books {
		if book.ID == params["id"] {
			var updatedBook Book
			_ = json.NewDecoder(r.Body).Decode(&updatedBook)
			books[index] = updatedBook
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}
	json.NewEncoder(w).Encode(nil)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, book := range books {
		if book.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	router := mux.NewRouter()

	books = append(books, Book{ID: "1", Title: "Clean Code", Author: "Robert C. Martin", Price: 30.0})
	books = append(books, Book{ID: "2", Title: "The Go Programming Language", Author: "Alan A. A. Donovan, Brian W. Kernighan", Price: 25.0})
	books = append(books, Book{ID: "3", Title: "Designing Data-Intensive Applications", Author: "Martin Kleppmann", Price: 35.0})

	router.HandleFunc("/books", getAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBookByID).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
