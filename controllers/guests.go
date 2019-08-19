package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

//GetGuest : GetGuest
func (g GuestController) GetGuest(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Retreieve the URL parameters as "r" and insert into a
		// map data type as "params"
		params := mux.Vars(r)

		// Create an instance of the Guest Repo to manage incoming data
		guestRepo := repository.GuestRepository{}

		// Convert the URL parameter value to an int,
		// rather than a string
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, err)
			return
		}

		guest, err := guestRepo.GetGuest(db, id)

		if err != nil {
			log.Println(err.Error())
			utils.SendError(w, http.StatusInternalServerError, err.Error())
		} else {

			getSuccessMsg := `[INFO] %s %s's record is successfully retrieved.`
			getSuccessMsg = fmt.Sprintf(getSuccessMsg, guest.FirstName, guest.LastName)

			log.Println(getSuccessMsg)

			// When successful send the results and status code to the client
			w.Header().Set("Content-Type", "application/json")
			utils.SendSuccess(w, guest)
			utils.SendSuccess(w, getSuccessMsg)
		}
	}
}

// GetGuests : GetGuests
func (g GuestController) GetGuests(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Initialize an instance of the Guest struct
		var guests []models.Guest

		// Initialize an instance of the repository and assign to
		// the guestRepo variable
		guestRepo := repository.GuestRepository{}

		// Invoke the GetGuests method by the guestRepo instance and assign
		// the value to the variables guests and err
		guests, guestsSize, err := guestRepo.GetGuests(db)

		// If an error arises handle it using the SendError function
		if err != nil {
			log.Println(err.Error())
			utils.SendError(w, http.StatusInternalServerError, err.Error())
		} else {
			//Int8 is converted to int64 then to a string to output Guest ID
			guestsSize := strconv.FormatInt(int64(guestsSize), 10)

			getSuccessMsg := `[INFO] %s Guests successfully retrieved.`
			getSuccessMsg = fmt.Sprintf(getSuccessMsg, guestsSize)

			log.Println(getSuccessMsg)

			// When successful send the results and status code to the client
			w.Header().Set("Content-Type", "application/json")
			utils.SendSuccess(w, guests)
			utils.SendSuccess(w, getSuccessMsg)
		}
	}
}

// AddGuest : AddGuest
func (g GuestController) AddGuest(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Initialize a variable with the type of the Guest struct
		var guest models.Guest
		// Initialize a variable with the type of the Visit struct
		// var visit models.Visit

		// Handle the response Body and map values to the hex value
		// of the guest var
		json.NewDecoder(r.Body).Decode(&guest)

		//
		guestRepo := repository.GuestRepository{}
		guestID, err := guestRepo.AddGuest(db, guest)

		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, err)
			utils.LogFatal(err)
		}

		// Assigns the newly returned ID from the database to the Guest ID
		// in the current object
		guest.ID = guestID

		// visitRepo := repository.VisitRepository{}
		// err = visitRepo.AddGuestVisit(db, guest, visit)

		if err != nil {
			log.Println(err.Error())
			utils.SendError(w, http.StatusInternalServerError, err.Error())
		} else {
			//Int8 is converted to int64 then to a string to output Guest ID
			guestIDStr := strconv.FormatInt(int64(guest.ID), 10)

			addSuccessMsg := `[INFO] %s %s is successfully saved with the Guest ID of %s`
			addSuccessMsg = fmt.Sprintf(addSuccessMsg, guest.FirstName, guest.LastName, guestIDStr)

			log.Println(addSuccessMsg)

			w.Header().Set("Content-Type", "text/plain")
			utils.SendSuccess(w, addSuccessMsg)
		}
	}
}

// UpdateGuest : UpdateGuest
func (g GuestController) UpdateGuest(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Initialize a variable with the type of the Guest struct
		var guest models.Guest

		// Retrieves the response body and maps it to the guest struct
		json.NewDecoder(r.Body).Decode(&guest)

		//
		guestRepo := repository.GuestRepository{}
		err := guestRepo.UpdateGuest(db, guest)

		// If any errors write to the env log and return message to client,
		// otherwise send a successful operation
		if err != nil {
			log.Println(err.Error())
			utils.SendError(w, http.StatusInternalServerError, err.Error())
		} else {
			updateSuccessMsg := "[INFO] %s %s's information is successfully updated."
			updateSuccessMsg = fmt.Sprintf(updateSuccessMsg, guest.FirstName, guest.LastName)

			log.Println(updateSuccessMsg)

			w.Header().Set("Content-Type", "text/plain")
			utils.SendSuccess(w, updateSuccessMsg)
		}
	}
}

// UpdateGuests : UpdateGuests
func (g GuestController) UpdateGuests(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Initialize a variable with the type of the Guest struct
		// var guest models.Guest
		//
		// // Retrieves the response body and maps it to the guest struct
		// json.NewDecoder(r.Body).Decode(&guest)
		//
		// //
		// guestRepo := repository.GuestRepository{}
		// err := guestRepo.UpdateGuest(db, guest)
		//
		// // If any errors write to the env log and return message to client,
		// // otherwise send a successful operation
		// if err != nil {
		// 	log.Println(err.Error())
		// 	utils.SendError(w, http.StatusInternalServerError, err.Error())
		// } else {
		// 	updateSuccessMsg := "[INFO] %s %s's information is successfully updated."
		// 	updateSuccessMsg = fmt.Sprintf(updateSuccessMsg, guest.FirstName, guest.LastName)
		//
		// 	log.Println(updateSuccessMsg)
		//
		// 	w.Header().Set("Content-Type", "text/plain")
		// 	utils.SendSuccess(w, updateSuccessMsg)
		// }
	}
}

// ArchiveGuest : ArchiveGuest
func (g GuestController) ArchiveGuest(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Retreieve the URL parameters as "r" and insert into a
		// map data type as "params"
		params := mux.Vars(r)

		id, _ := strconv.Atoi(params["id"])

		guestRepo := repository.GuestRepository{}

		err := guestRepo.ArchiveGuest(db, id)

		if err != nil {
			log.Println(err.Error())
			utils.SendError(w, http.StatusInternalServerError, err.Error())
		} else {
			// When successful send the results and status code to the client
			w.Header().Set("Content-Type", "text/plain")
			utils.SendSuccess(w, "Guest was successfully archived.")
		}
	}
}
