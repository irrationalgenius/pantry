package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"pantry2/models"
	"pantry2/repository"
	"pantry2/utils"
	"strconv"

	"github.com/gorilla/mux"
)

// Controller : Controllers
type Controller struct{}

var guests []models.Guest

// GetGuests : GetGuests
func (c Controller) GetGuests(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Retrieving all Guests method was invoked")

		// Initialize a variable with the type of the Error struct
		var error models.Error
		// Initialize an instance of the repository and assign to
		// the guestRepo variable
		guestRepo := repository.GuestRepository{}

		// Invoke the GetGuests method by the guestRepo instance and assign
		// the value to the variables guests and err
		guests, err := guestRepo.GetGuests(db)

		// If an error arises handle it using the SendError function
		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			utils.LogFatal(err)
		}
		//
		// // When successful send the results and status code to the client
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, guests)
	}
}

//GetGuest : GetGuest
func (c Controller) GetGuest(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Retrieving the guest method was invoked")

		// Initialize a variable with the type of the Error struct
		// var error models.Error

		// Retreieve the URL parameters as "r" and insert into a
		// map data type as "params"
		params := mux.Vars(r)

		guestRepo := repository.GuestRepository{}

		// Convert the URL parameter value to an int,
		// rather than a string
		id, _ := strconv.Atoi(params["id"])

		guest, _ := guestRepo.GetGuest(db, id)

		// if err != nil {
		// 	if err == sql.ErrNoRows {
		// 		error.Message = "Record not found"
		// 		utils.SendError(w, http.StatusNotFound, error)
		// 		return
		// 	} else {
		// 		error.Message = "Server error"
		// 		utils.SendError(w, http.StatusInternalServerError, error)
		// 		return
		// 	}
		// }

		// When successful send the results and status code to the client
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, guest)
	}
}

// AddGuest : AddGuest
func (c Controller) AddGuest(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Adding a guest method was invoked")

		// Initialize a variable with the type of the Guest struct
		var guest models.Guest
		// Initialize a variable with the type of int
		// var guestID int
		// Initialize a variable with the type of the Error struct
		var error models.Error

		// Handle the response Body and map values to the hex value
		// of the book var
		json.NewDecoder(r.Body).Decode(&guest)

		// Validate book data, before saving details
		guest, err := repository.GuestValidate(guest)

		if err != nil {
			error.Message = "Validation failed, see error output"
			utils.SendError(w, http.StatusInternalServerError, error)
			utils.LogFatal(err)
		}

		guestRepo := repository.GuestRepository{}
		err = guestRepo.AddGuest(db, guest)

		if err != nil {
			error.Message = "Failed to add Guest, see error output"
			utils.SendError(w, http.StatusInternalServerError, error)
			utils.LogFatal(err)
		}

		// When successful send the results and status code to the client
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, "Success")
	}
}

// UpdateGuest : UpdateGuest
func (c Controller) UpdateGuest(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Updating a guest method was invoked")
		//
		// 		// Initialize a variable with the type of the Book struct
		// 		var book models.Book
		// 		// Initialize a variable with the type of the Error struct
		// 		var error models.Error
		//
		// 		// Retrieves the response body and maps it to the book variable
		// 		json.NewDecoder(r.Body).Decode(&book)
		//
		// 		// Validate book data, before saving details
		// 		if book.ID == 0 || book.Author == "" || book.Title == "" || book.Year == "" {
		// 			error.Message = "Cannot save record with missing data."
		// 			utils.SendError(w, http.StatusBadRequest, error)
		// 			return
		// 		}
		//
		// 		bookRepo := bookRepository.BookRepository{}
		// 		rowsUpdated, err := bookRepo.UpdateBook(db, book)
		//
		// 		if err != nil {
		// 			error.Message = "Server error"
		// 			utils.SendError(w, http.StatusInternalServerError, error)
		// 			return
		// 		}
		//
		// 		// When successful send the results and status code to the client
		// 		w.Header().Set("Content-Type", "text/plain")
		// 		utils.SendSuccess(w, rowsUpdated)
	}
}

// RemoveGuest : RemoveGuest
func (c Controller) RemoveGuest(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Removing a guest method was invoked")
		//
		// 		// Initialize a variable with the type of the Error struct
		// 		var error models.Error
		// 		// Retreieve the URL parameters as "r" and insert into a
		// 		// map data type as "params"
		// 		params := mux.Vars(r)
		//
		// 		bookRepo := bookRepository.BookRepository{}
		//
		// 		id, _ := strconv.Atoi(params["id"])
		//
		// 		rowsDeleted, err := bookRepo.RemoveBook(db, id)
		//
		// 		if err != nil {
		// 			error.Message = "Server error"
		// 			utils.SendError(w, http.StatusInternalServerError, error)
		// 			return
		// 		}
		//
		// 		if rowsDeleted == 0 {
		// 			error.Message = "Record not found"
		// 			utils.SendError(w, http.StatusNotFound, error)
		// 			return
		// 		}
		//
		// 		// When successful send the results and status code to the client
		// 		w.Header().Set("Content-Type", "text/plain")
		// 		utils.SendSuccess(w, rowsDeleted)
	}
}
