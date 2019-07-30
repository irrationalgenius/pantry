package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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
		// Initialize a variable with the type of the Error struct
		var error models.Error

		// Retreieve the URL parameters as "r" and insert into a
		// map data type as "params"
		params := mux.Vars(r)

		books = []models.Book{}
		bookRepo := bookRepository.BookRepository{}

		id, _ := strconv.Atoi(params["id"])

		book, err := bookRepo.GetBook(db, book, id)

		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Record not found"
				utils.SendError(w, http.StatusNotFound, error)
				return
			} else {
				error.Message = "Server error"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
		}

		// When successful send the results and status code to the client
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, book)
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
		// Initialize a variable with the type of the Error struct
		var error models.Error

		// Handle the response Body and map values to the hex value
		// of the book var
		json.NewDecoder(r.Body).Decode(&book)

		// Validate book data, before saving details
		if book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "Cannot save record with missing data."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookRepository.BookRepository{}
		bookID, err := bookRepo.AddBook(db, book)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		// When successful send the results and status code to the client
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, bookID)
	}
}

// UpdateBook : UpdateBook
func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Updating a book method is invoked")

		// Initialize a variable with the type of the Book struct
		var book models.Book
		// Initialize a variable with the type of the Error struct
		var error models.Error

		// Retrieves the response body and maps it to the book variable
		json.NewDecoder(r.Body).Decode(&book)

		// Validate book data, before saving details
		if book.ID == 0 || book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "Cannot save record with missing data."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookRepository.BookRepository{}
		rowsUpdated, err := bookRepo.UpdateBook(db, book)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		// When successful send the results and status code to the client
		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsUpdated)
	}
}

// RemoveBook : RemoveBook
func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Removing a book method is invoked")

		// Initialize a variable with the type of the Error struct
		var error models.Error
		// Retreieve the URL parameters as "r" and insert into a
		// map data type as "params"
		params := mux.Vars(r)

		bookRepo := bookRepository.BookRepository{}

		id, _ := strconv.Atoi(params["id"])

		rowsDeleted, err := bookRepo.RemoveBook(db, id)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		if rowsDeleted == 0 {
			error.Message = "Record not found"
			utils.SendError(w, http.StatusNotFound, error)
			return
		}

		// When successful send the results and status code to the client
		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsDeleted)
	}
}
