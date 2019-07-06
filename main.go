package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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

	books = append(books,
		Book{ID: 1, Title: "Golang Pointers", Author: "Mr. Pointers", Year: "2010"},
		Book{ID: 2, Title: "Goroutines", Author: "Mr. Goroutines", Year: "2011"},
		Book{ID: 3, Title: "Golang Routers", Author: "Mr. Router", Year: "2012"},
		Book{ID: 4, Title: "Golang Concurrency", Author: "Mr. Concurrency", Year: "2013"},
		Book{ID: 5, Title: "Golang Best Practices", Author: "Mr. Golang", Year: "2014"},
	)

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/book/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	// log.Println("Geting all books method is invoked")

	json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request) {
	// log.Println("Geting a single book method is invoked")

	// Retrieves the parameters received from the request
	params := mux.Vars(r)

	// Prints the retrieved parameter to the console
	log.Println(params)

	// Converts the string to an integer by the params map id value and
	// places the value in a new variable of "i"
	i, _ := strconv.Atoi(params["id"])

	// Loops through all the books, until the value of "i" is matched
	for _, book := range books {
		if book.ID == i {
			// Send the response as a Json encoded object by the Book pointer
			json.NewEncoder(w).Encode(book)
		}
	}

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
