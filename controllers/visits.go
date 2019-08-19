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

// VisitController : VisitController
type VisitController struct{}

var visits []models.Visit

//GetGuestVisit : GetGuestVisit
func (v VisitController) GetGuestVisit(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Invoking the Get Guest Visits method")

		// db.Query(`select `)
	}
}

//GetGuestVisits : GetGuestVisits
func (v VisitController) GetGuestVisits(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Retreieve the URL parameters as "r" and insert into a
		// map data type as "params"
		params := mux.Vars(r)

		visitRepo := repository.VisitRepository{}

		// Convert the URL parameter value to an int,
		// rather than a string
		id, _ := strconv.Atoi(params["id"])

		visits, visitsSize, err := visitRepo.GetGuestVisits(db, id)

		if err != nil {
			log.Println(err.Error())
			utils.SendError(w, http.StatusInternalServerError, err.Error())
		} else {
			//Int8 is converted to int64 then to a string to output Guest ID
			visitsSize := strconv.FormatInt(int64(visitsSize), 10)

			getSuccessMsg := `[INFO] %s Visit(s) successfully retrieved.`
			getSuccessMsg = fmt.Sprintf(getSuccessMsg, visitsSize)

			log.Println(getSuccessMsg)

			// When successful send the results and status code to the client
			w.Header().Set("Content-Type", "application/json")
			utils.SendSuccess(w, visits)
			utils.SendSuccess(w, getSuccessMsg)
		}
	}
}

//AddGuestVisit : AddGuestVisit
func (v VisitController) AddGuestVisit(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Initialize a variable with the type of the Visit struct
		var visit models.Visit
		// Initialize a variable with the type of the Guest struct
		var guest models.Guest

		// Handle the response Body and map values to the hex value
		// of the guest var
		response := json.NewDecoder(r.Body)

		response.Decode(&guest)
		response.Decode(&visit)

		visitRepo := repository.VisitRepository{}
		err := visitRepo.AddGuestVisit(db, guest, visit)

		if err != nil {
			log.Println(err.Error())
			utils.SendError(w, http.StatusInternalServerError, err.Error())
		} else {
			addSuccessMsg := `[INFO] %s %s's visit is successfully saved. Next visit on %s`
			addSuccessMsg = fmt.Sprintf(addSuccessMsg, guest.FirstName, guest.LastName, visit.DateofVisitNext)

			log.Println(addSuccessMsg)

			w.Header().Set("Content-Type", "text/plain")
			utils.SendSuccess(w, addSuccessMsg)
		}
	}
}

//UpdateGuestVisit : UpdateGuestVisit
func (v VisitController) UpdateGuestVisit(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Invoking the Get Guest Visits method")

		// db.Query(`select `)
	}
}

//UpdateGuestVisits : UpdateGuestVisits
func (v VisitController) UpdateGuestVisits(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Invoking the Get Guest Visits method")

		// db.Query(`select `)
	}
}

//ArchiveGuestVisit : ArchiveGuestVisit
func (v VisitController) ArchiveGuestVisit(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Invoking the Get Guest Visits method")

		// db.Query(`select `)
	}
}
