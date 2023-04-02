package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Book adalah model data untuk item buku
type Book struct {
	gorm.Model
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float32 `json:"price"`
}

// buat variabel untuk menyimpan koneksi database
var db *gorm.DB

// inisialisasi koneksi database
func initDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=bookstore port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// auto migrate model Book
	db.AutoMigrate(&Book{})
}

// Get all books
func getAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []Book
	db.Find(&books)
	json.NewEncoder(w).Encode(books)
}

// Get book by ID
func getBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book Book
	db.First(&book, params["id"])
	json.NewEncoder(w).Encode(book)
}

// Create a book
func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	db.Create(&book)
	json.NewEncoder(w).Encode(book)
}

// Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book Book
	db.First(&book, params["id"])
	_ = json.NewDecoder(r.Body).Decode(&book)
	db.Save(&book)
	json.NewEncoder(w).Encode(book)
}

// Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book Book
	db.Delete(&book, params["id"])
	json.NewEncoder(w).Encode(book)
}

func main() {
	// inisialisasi koneksi database
	initDB()

	router := mux.NewRouter()

	// set up routes
	router.HandleFunc("/books", getAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBookByID).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	// start the server
	log.Fatal(http.ListenAndServe(":8000", router))
}
