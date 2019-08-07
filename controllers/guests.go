package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"pantry/models"
	"pantry/repository"
	"pantry/utils"
	"strconv"

	"github.com/gorilla/mux"
)

// GuestController : GuestController
type GuestController struct{}

var guests []models.Guest

// GetGuests : GetGuests
func (g GuestController) GetGuests(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Invoking the Get all Guests Controller")

		// Initialize an instance of the repository and assign to
		// the guestRepo variable
		guestRepo := repository.GuestRepository{}

		// Invoke the GetGuests method by the guestRepo instance and assign
		// the value to the variables guests and err
		guests, err := guestRepo.GetGuests(db)

		// If an error arises handle it using the SendError function
		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, err)
			utils.LogFatal(err)
		}
		//
		// // When successful send the results and status code to the client
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, guests)
	}
}

//GetGuest : GetGuest
func (g GuestController) GetGuest(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Invoking the Get a speific Guest Controller")

		// Initialize a variable with the type of the Error struct
		var error models.Error

		// Retreieve the URL parameters as "r" and insert into a
		// map data type as "params"
		params := mux.Vars(r)

		// Create an instance of the Guest Repo to manage incoming data
		guestRepo := repository.GuestRepository{}

		// Convert the URL parameter value to an int,
		// rather than a string
		id, _ := strconv.Atoi(params["id"])

		guest, err := guestRepo.GetGuest(db, id)

		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Error: The Record was not found"
				utils.SendError(w, http.StatusNotFound, error)
				return
			} else {
				utils.SendError(w, http.StatusInternalServerError, err)
				return
			}
		}

		// When successful send the results and status code to the client
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, guest)
	}
}

// AddGuest : AddGuest
func (g GuestController) AddGuest(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Invoking the Guest Adding Controller")

		// Initialize a variable with the type of the Guest struct
		var guest models.Guest
		// Initialize a variable with the type of the Visit struct
		var visit models.Visit

		// Handle the response Body and map values to the hex value
		// of the book var
		json.NewDecoder(r.Body).Decode(&guest)

		guestRepo := repository.GuestRepository{}
		guestID, err := guestRepo.AddGuest(db, guest)

		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, err)
			utils.LogFatal(err)
		}

		// Assigns the newly returned ID from the database to the Guest ID
		// in the current object
		guest.ID = guestID

		// Dummy note used to save with the record to ensure correct saving
		// of database vars, used for debugging
		// visit.Notes = "This is a dummy note to save with the visit record."

		visitRepo := repository.VisitRepository{}
		err = visitRepo.AddVisit(db, guest, visit)

		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, err)
			utils.LogFatal(err)
		}

		// When successful send the results and status code to the client
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, "Success")
	}
}

// UpdateGuest : UpdateGuest
func (g GuestController) UpdateGuest(db *sql.DB) http.HandlerFunc {
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
func (g GuestController) RemoveGuest(db *sql.DB) http.HandlerFunc {
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
