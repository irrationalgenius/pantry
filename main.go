package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book : Book Struct
type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/book/{id}", getBook).Methods("GET")
	router.HandleFunc("/book", addBook).Methods("POST")
	router.HandleFunc("/book", updateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Geting all books method is invoked")

}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Retrieving a single book method is invoked")

}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Adding a book method is invoked")

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Updating a book method is invoked")

}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Removing a book method is invoked")

}
