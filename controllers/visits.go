package controllers

import (
	"database/sql"
	"log"
	"pantry/models"
	"pantry/repository"
	"pantry/utils"
)

// VisitController : VisitController
type VisitController struct{}

var visits []models.Visit

//GetAllVisits : GetAllVisits
func (v VisitController) GetAllVisits(db *sql.DB, guest models.Guest) ([]models.Visit, error) {
	log.Println("Invoking the Get Guest Visits method")

	// db.Query(`select `)

	return []models.Visit{}, nil
}

//AddVisit : AddVisit
func (v VisitController) AddVisit(db *sql.DB, guest models.Guest, visit models.Visit) error {

	visitRepo := repository.VisitRepository{}
	err := visitRepo.AddVisit(db, guest, visit)

	if err != nil {
		// utils.SendError(w, http.StatusInternalServerError, error)
		utils.LogFatal(err)
	}

	return nil
}
