package mock

//
// The Purpose of this main file is to demonstrate the ability
// to use basic web functionailty with go using mock data.
//
// This was the first model before adding the database in the
// main file.
//
// To reuse this file set package mock to package main and move
// into project root
//

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
	router.HandleFunc("/book", addBook).Methods("POST")
	router.HandleFunc("/book", updateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Geting all books method is invoked")

	json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Retrieving a single book method is invoked")

	// Retrieves the parameters received from the client request
	params := mux.Vars(r)

	// Prints the retrieved parameter to the console
	log.Println(params)

	// Converts the string to an integer by the params map id value and
	// places the value in a new variable of "i", discards the err var
	i, _ := strconv.Atoi(params["id"])

	// Golang doc for the Range Type
	// URL: https://tour.golang.org/moretypes/16
	//   When ranging over a slice, two values are returned for each iteration.
	//   The first is the index, and the second is a copy
	//   of the element at that index.
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

	// Create a new variable book as of type Book
	var book Book

	// Handle the Body of the response, and decode into a pointer of the
	// book variable just created.
	json.NewDecoder(r.Body).Decode(&book)

	// Take the books array and add the newly retrieved book to the existing
	// books collection
	books = append(books, book)

	// Take all the data and encode into the json format, and sen as a
	// server request to the client
	json.NewEncoder(w).Encode(books)

	// Print to the terminal
	log.Println(books)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Updating a book method is invoked")

	var book Book

	json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}
	}

	json.NewEncoder(w).Encode(books)

}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Removing a book method is invoked")

	// Retrieves the parameters received from the client request
	params := mux.Vars(r)

	// Converts the string to an integer by the params map id value and
	// places the value in a new variable of "id", discards the err var
	id, _ := strconv.Atoi(params["id"])

	// Golang doc for the Range Type
	// URL: https://tour.golang.org/moretypes/16
	//   When ranging over a slice, two values are returned for each iteration.
	//   The first is the index, and the second is a copy
	//   of the element at that index.
	// Loops through all the books, until the value of "i" is matched
	for i, item := range books {
		if item.ID == id {

			// Create a new slice without the matching id
			books = append(books[:i], books[i+1:]...)
		}
	}

	// Send the response as a Json encoded object by the book variable
	json.NewEncoder(w).Encode(books)

}
