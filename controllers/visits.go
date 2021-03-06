package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pantry-api/models"
	"pantry-api/repository"
	"pantry-api/utils"
	"strconv"

	"github.com/gorilla/mux"
)

// VisitController : VisitController
type VisitController struct{}

var visits []models.Visit

//GetGuestVisit : GetGuestVisit
func (v VisitController) GetGuestVisit(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Retreieve the URL parameters as "r" and insert into a
		// map data type as "params"
		params := mux.Vars(r)

		// Return the Guest ID and store in id
		id, _ := strconv.Atoi(params["id"])
		// Return the Visit ID and store in vid
		vid, _ := strconv.Atoi(params["vid"])

		visitRepo := repository.VisitRepository{}

		visit, err := visitRepo.GetGuestVisit(db, id, vid)

		if err != nil {
			log.Println(err.Error())
			utils.SendError(w, http.StatusInternalServerError, err.Error())
		} else {

			getSuccessMsg := `[INFO] Visit successfully retrieved.`
			log.Println(getSuccessMsg)

			// When successful send the results and status code to the client
			w.Header().Set("Content-Type", "application/json")
			utils.SendSuccess(w, visit)
			utils.SendSuccess(w, getSuccessMsg)
		}
	}
}

//GetGuestVisits : GetGuestVisits
func (v VisitController) GetGuestVisits(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Retreieve the URL parameters as "r" and insert into a
		// map data type as "params"
		params := mux.Vars(r)

		// Convert the URL parameter value to an int,
		// rather than a string
		id, _ := strconv.Atoi(params["id"])

		visitRepo := repository.VisitRepository{}

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
		dateofVisitNext, err := visitRepo.AddGuestVisit(db, guest, visit)

		if err != nil {
			log.Println(err.Error())
			utils.SendError(w, http.StatusInternalServerError, err.Error())
		} else {
			addSuccessMsg := `[INFO] %s %s's visit is successfully saved. Next visit on %s`
			addSuccessMsg = fmt.Sprintf(addSuccessMsg, guest.FirstName, guest.LastName, dateofVisitNext)

			log.Println(addSuccessMsg)

			w.Header().Set("Content-Type", "text/plain")
			utils.SendSuccess(w, addSuccessMsg)
		}
	}
}

//UpdateGuestVisit : UpdateGuestVisit
func (v VisitController) UpdateGuestVisit(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var guest models.Guest
		var visit models.Visit

		data := json.NewDecoder(r.Body)
		data.Decode(&guest)
		data.Decode(&visit)

		visitRepo := repository.VisitRepository{}
		err := visitRepo.UpdateGuestVisit(db, guest, visit)

		// If any errors write to the env log and return message to client,
		// otherwise send a successful operation
		if err != nil {
			log.Println(err.Error())
			utils.SendError(w, http.StatusInternalServerError, err.Error())
		} else {
			updateSuccessMsg := "[INFO] %s %s's visit information is successfully updated."
			updateSuccessMsg = fmt.Sprintf(updateSuccessMsg, guest.FirstName, guest.LastName)

			log.Println(updateSuccessMsg)

			w.Header().Set("Content-Type", "text/plain")
			utils.SendSuccess(w, updateSuccessMsg)
		}
	}
}

//ArchiveGuestVisit : ArchiveGuestVisit
func (v VisitController) ArchiveGuestVisit(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Retreieve the URL parameters as "r" and insert into a
		// map data type as "params"
		params := mux.Vars(r)

		// Initialize a variable with the type of the Guest struct
		var guest models.Guest
		// Initialize a variable with the type of the Visit struct
		var visit models.Visit

		data := json.NewDecoder(r.Body)
		data.Decode(&guest)
		data.Decode(&visit)

		// Get value of input params and store in local vars
		// id, _ := strconv.Atoi(params["id"])
		// vid, _ := strconv.Atoi(params["vid"])
		do := params["do"]

		visitRepo := repository.VisitRepository{}

		if do == "A" {
			err := visitRepo.ArchiveGuestVisit(db, visit)

			if err != nil {
				log.Println(err.Error())
				utils.SendError(w, http.StatusInternalServerError, err.Error())
			} else {
				// When successful send the results and status code to the client
				w.Header().Set("Content-Type", "text/plain")
				utils.SendSuccess(w, "Visit is successfully Archived.")
			}
		}

		if do == "U" {
			err := visitRepo.UnarchiveGuestVisit(db, visit)

			if err != nil {
				log.Println(err.Error())
				utils.SendError(w, http.StatusInternalServerError, err.Error())
			} else {
				// When successful send the results and status code to the client
				w.Header().Set("Content-Type", "text/plain")
				utils.SendSuccess(w, "Visit is successfully Unarchived.")
			}
		}
	}
}
