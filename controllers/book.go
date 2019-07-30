package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"udemy-bookstore/models"
	"udemy-bookstore/repository/book"
	"udemy-bookstore/utils"

	"github.com/gorilla/mux"
)

// Controller : Controllers
type Controller struct{}

var books []models.Book

// GetBooks : GetBooks
func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Retrieving all books method is invoked")

		// Initialize a variable with the type of the Book struct
		var book models.Book
		// Initialize a variable with the type of the Error struct
		var error models.Error
		// Assign an empty slice to the global []books data type
		books = []models.Book{}
		// Initialize an instance of the BookRepository and assign to
		// the bookRepo variable
		bookRepo := bookRepository.BookRepository{}

		// Invoke the GetBooks method by the bookRepo instance and assign
		// the value to the variables books and err
		books, err := bookRepo.GetBooks(db, book, books)

		// If an error arises handle it using the SendError function
		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		// When successful send the results and status code to the client
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, books)
	}
}

// GetBook : GetBook
func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Retrieving the book method is invoked")

		// Initialize a variable with the type of the Book struct
		var book models.Book

		// Retreieve the URL parameters as "r" and insert into a
		// map data type as "params"
		params := mux.Vars(r)

		// Search the database for this parameter value
		row := db.QueryRow("select * from books where id=$1", params["id"])

		// Insert the values from the database into the hex value
		// of the book vars
		err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)

		// See logFatal() function
		utils.LogFatal(err)

		// Convert the book data type to a json object and send the response
		// to the client as "w"
		json.NewEncoder(w).Encode(book)
	}
}

// AddBook : AddBook
func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Adding a book method is invoked")

		// Initialize a variable with the type of the Book struct
		var book models.Book

		// Initialize a variable with the type of int
		var bookID int

		// Handle the response Body and map values to the hex value
		// of the book var
		json.NewDecoder(r.Body).Decode(&book)

		// Insert values received from the client into the database
		// Return the id of the insert into the hex value location for bookID
		err := db.QueryRow("insert into books(title, author, year) values($1, $2, $3) RETURNING id",
			book.Title, book.Author, book.Year).Scan(&bookID)

		// See logFatal() function
		utils.LogFatal(err)

		// Convert the bookID var to a json object and send the response
		// to the client as "w"
		json.NewEncoder(w).Encode(bookID)
	}
}

// UpdateBook : UpdateBook
func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Updating a book method is invoked")

		// Initialize a variable with the type of the Book struct
		var book models.Book

		// Retrieves the response body and maps it to the book variable
		json.NewDecoder(r.Body).Decode(&book)

		// Update values received from the client to the record in the database
		// Return the id of the update into the hex value location for bookID
		result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id",
			&book.Title, &book.Author, &book.Year, &book.ID)

		// See logFatal() function
		utils.LogFatal(err)

		// Get the number of rows affected for the update clause
		rowUpdated, err := result.RowsAffected()

		// See logFatal() function
		utils.LogFatal(err)

		// Convert the rowUpdated var to a json object and send the response
		// to the client as "w"
		json.NewEncoder(w).Encode(rowUpdated)
	}
}

// RemoveBook : RemoveBook
func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Removing a book method is invoked")

		// Retreieve the URL parameters as "r" and insert into a
		// map data type as "params"
		params := mux.Vars(r)

		// Remove from the database a record matching this parameter value
		row, err := db.Exec("delete from books where id = $1", params["id"])

		// See logFatal() function
		utils.LogFatal(err)

		// Get the number of rows affected for the delete clause
		rowDeleted, err := row.RowsAffected()

		// See logFatal() function
		utils.LogFatal(err)

		// Convert the rowDeleted var to a json object and send the response
		// to the client as "w"
		json.NewEncoder(w).Encode(rowDeleted)
	}
}
