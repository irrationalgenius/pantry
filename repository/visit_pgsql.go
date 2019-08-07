package repository

import (
	"database/sql"
	"errors"
	"log"
	"pantry/models"
	"strconv"
	"time"
)

// VisitRepository : VisitRepository
type VisitRepository struct{}

//AddVisit : AddVisit
func (v VisitRepository) AddVisit(db *sql.DB, guest models.Guest, visit models.Visit) error {
	log.Println("Accessing the Add Visits Function in the Visits Repository")

	if guest.ID == 0 {
		return errors.New("Guest ID value is not set, adding visit failed")
	}

	// Set the current visit to the current day
	visit.DateofVisitLast = time.Now()
	// Set the Last date updated variable
	visit.LastDateUpdated = time.Now()
	// Calculate the next visit, adding 42 days to the current visit
	// and store in the *Next variable.
	visit.DateofVisitNext = visit.DateofVisitLast.AddDate(0, 0, 42)

	sqlVisitAdd := `INSERT INTO pantry.visits(
	  	guest_id,
			date_visit,
			date_visit_next,
	  	notes,
			last_date_updated)
		VALUES($1, $2, $3, $4, $5)`

	result, err := db.Exec(sqlVisitAdd,
		guest.ID, visit.DateofVisitLast, visit.DateofVisitNext,
		visit.Notes, visit.LastDateUpdated)

	if err != nil {
		return err
	}

	affectedInt, _ := result.RowsAffected()

	affected := strconv.FormatInt(affectedInt, 10)

	visitSuccessMsg := "% visit(s) has been successfully saved for %s %s"

	log.Printf(visitSuccessMsg, affected, guest.FirstName, guest.LastName)

	return nil
}

func visitValidate(visit models.Visit) error {

	// If the last visit of the guest is within the application parameter
	// for adding a visit, then the application should error.

	return nil
}
