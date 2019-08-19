package repository

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"pantry/models"
	"strconv"
	"time"
)

// VisitRepository : VisitRepository
type VisitRepository struct{}

//GetGuestVisits : GetGuestVisits
func (v VisitRepository) GetGuestVisits(db *sql.DB, id int) ([]models.Visit, int, error) {

	// Initialize an instance of the VisitRaw struct
	var visitRaw models.VisitRaw
	// Initialize an slice instance of the VisitRaw struct
	var visitsRaw []models.VisitRaw

	sqlVisitRawGet := `SELECT id, guest_id, date_visit, date_visit_next,
			notes, last_date_updated
		FROM pantry.visits
		WHERE guest_id = $1`

	rows, err := db.Query(sqlVisitRawGet, id)

	if err != nil {
		errorMsg := `[ERROR] Issue occured while retrieving data from the database.`
		return []models.Visit{}, 0, errors.New(errorMsg)
	}

	for rows.Next() {
		err = rows.Scan(&visitRaw.ID, &visitRaw.GuestID, &visitRaw.DateofVisitLast, &visitRaw.DateofVisitNext,
			&visitRaw.Notes, &visitRaw.LastDateUpdated)

		visitsRaw = append(visitsRaw, visitRaw)
	}

	if err != nil {
		errorMsg := `[ERROR] Issue occured while assigning data from the database.`
		return []models.Visit{}, 0, errors.New(errorMsg)
	}

	// After the data is retrieved from the database, it must be
	// cleaned, meaning NULL values set to Golang defaults
	// (Go has no idea what a NULL is, and neither do I :|)
	visits := visitsClean(visitsRaw)

	// Let's count the items in the group. This will be used for logging
	// purposes.
	visitsSize := len(visits)

	return visits, visitsSize, nil
}

//AddGuestVisit : AddGuestVisit
func (v VisitRepository) AddGuestVisit(db *sql.DB, guest models.Guest, visit models.Visit) error {

	// First thing first: validating the data, if the data does not meet
	// requirements, an error is issued back to the client
	if guest.ID == 0 {
		errorMsg := `[ERROR] Guest ID value is not set, adding visit failed.`
		return errors.New(errorMsg)
	}

	// Set the current visit to the current day
	if visit.DateofVisitLast.IsZero() {
		visit.DateofVisitLast = time.Now()
	}

	// Set the Last date updated variable
	visit.LastDateUpdated = time.Now()

	// Before adding this visit, ensure no previous visits are within the
	// pantry timing constrictions.
	visitCheck, err := visitValidate(db, guest, visit)

	if err != nil || visitCheck == false {
		return err
	}

	// Set the visit interval by the variable value from the environment
	currentVisitInterval, err := strconv.Atoi(os.Getenv("APP_VISIT_INTERVAL"))

	if err != nil {
		return err
	}

	// Add the newly set currentVisitInterval to the Guest's visit record for saving.
	visit.DateofVisitNext = visit.DateofVisitLast.AddDate(0, 0, currentVisitInterval)

	sqlVisitAdd := `INSERT INTO pantry.visits(
	  	guest_id,
			date_visit,
			date_visit_next,
	  	notes,
			last_date_updated)
		VALUES($1, $2, $3, $4, $5)`

	_, err = db.Exec(sqlVisitAdd,
		guest.ID, visit.DateofVisitLast, visit.DateofVisitNext,
		visit.Notes, visit.LastDateUpdated)

	if err != nil {
		return err
	}

	return nil
}

// ********************* Helper Functions ********************* //

func visitValidate(db *sql.DB, guest models.Guest, visit models.Visit) (bool, error) {

	// Declare a time variable to store the guests latest next visit to the pantry
	var nextVisitCheck time.Time

	// Retrieves the latest next visit date set from the previous visit, if this is
	// a new Guest, the default value of the visit will be 0001-01-01.
	sqlVisitValidate := `select coalesce(max(date_visit_next), '0001-01-01')::date date_visit_next
			from pantry.visits where guest_id = $1`

	result := db.QueryRow(sqlVisitValidate, guest.ID)

	err := result.Scan(&nextVisitCheck)

	// If a record is not returned for this guest or if this is a new guest then the
	// system returns to generate a new visit.
	if err != nil {
		if err == sql.ErrNoRows || nextVisitCheck.Before(visit.DateofVisitLast) {
			logMsg := `[INFO] No visit record was found for %s %s. Welcome to the Pantry. :)`
			log.Printf(logMsg, guest.FirstName, guest.LastName)

			return true, nil
		}
		// Otherwise if a more serious error, this will fail and return the error.
		return false, err
	}

	// Validates if the latest next visit comes after today, if it does then
	// the system returns the message below, otherwise this check passes, this
	// function returns true and a new visit is generated.
	if nextVisitCheck.After(visit.DateofVisitLast) {
		errorMsg := `[INFO] The system has on record a recent visit. Visit was not added`
		return false, errors.New(errorMsg)
	}

	logMsg := `[INFO] No visit record was found for %s %s. Welcome to the Pantry. :)`
	log.Printf(logMsg, guest.FirstName, guest.LastName)

	return true, nil
}

func visitClean(visitRaw models.VisitRaw) models.Visit {
	var visit models.Visit

	// The following elements must have values from
	// the database, otherwise a error will result
	if visitRaw.ID.Valid == true {
		visitID := int8(visitRaw.ID.Int64)
		visit.ID = visitID
	}
	if visitRaw.GuestID.Valid == true {
		guestID := int8(visitRaw.GuestID.Int64)
		visit.GuestID = guestID
	}
	if visitRaw.DateofVisitLast.Valid == true {
		visit.DateofVisitLast = visitRaw.DateofVisitLast.Time
	} else {
		visit.DateofVisitLast = time.Time{}
	}
	if visitRaw.DateofVisitNext.Valid == true {
		visit.DateofVisitNext = visitRaw.DateofVisitNext.Time
	} else {
		visit.DateofVisitNext = time.Time{}
	}
	if visitRaw.Notes.Valid == true {
		visit.Notes = visitRaw.Notes.String
	} else {
		visit.Notes = ""
	}
	if visitRaw.LastDateUpdated.Valid == true {
		visit.LastDateUpdated = visitRaw.LastDateUpdated.Time
	} else {
		visit.LastDateUpdated = time.Time{}
	}

	return visit
}

func visitsClean(visitsRaw []models.VisitRaw) []models.Visit {

	var visits []models.Visit

	for _, visitRaw := range visitsRaw {

		visit := visitClean(visitRaw)

		visits = append(visits, visit)
	}

	return visits
}
