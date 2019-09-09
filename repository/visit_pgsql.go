package repository

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"pantry-api/models"
	"strconv"
	"time"
)

// VisitRepository : VisitRepository
type VisitRepository struct{}

//GetGuestVisit : GetGuestVisit
func (v VisitRepository) GetGuestVisit(db *sql.DB, id int, vid int) (models.Visit, error) {

	err := visitCheckID(db, id)

	if err != nil {
		return models.Visit{}, err
	}

	// Initialize an instance of the VisitRaw struct
	var visitRaw models.VisitRaw

	sqlVisitRawGet := `SELECT id, guest_id, date_visit, date_visit_next,
			notes, last_date_updated
		FROM visits
		WHERE guest_id = $1
			AND id = $2`

	row := db.QueryRow(sqlVisitRawGet, id, vid)

	err = row.Scan(&visitRaw.ID, &visitRaw.GuestID, &visitRaw.DateofVisitLast, &visitRaw.DateofVisitNext,
		&visitRaw.Notes, &visitRaw.LastDateUpdated)

	if err != nil {
		return models.Visit{}, err
	}

	// After the data is retrieved from the database, it must be
	// cleaned, meaning NULL values set to Golang defaults
	// (Go has no idea what a NULL is, and neither do I :|)
	visit := visitClean(visitRaw)

	return visit, nil
}

//GetGuestVisits : GetGuestVisits
func (v VisitRepository) GetGuestVisits(db *sql.DB, id int) ([]models.Visit, int, error) {

	// Initialize an instance of the VisitRaw struct
	var visitRaw models.VisitRaw
	// Initialize an slice instance of the VisitRaw struct
	var visitsRaw []models.VisitRaw

	sqlVisitRawGet := `SELECT id, guest_id, date_visit, date_visit_next,
			notes, last_date_updated
		FROM visits
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

//GetGuestVisitsArchive : GetGuestVisitsArchive
func (v VisitRepository) GetGuestVisitsArchive(db *sql.DB, id int) ([]models.Visit, int, error) {

	// Initialize an instance of the VisitRaw struct
	var visitRaw models.VisitRaw
	// Initialize an slice instance of the VisitRaw struct
	var visitsRaw []models.VisitRaw

	sqlVisitRawGet := `SELECT id, guest_id, date_visit, date_visit_next,
			notes, last_date_updated
		FROM visits_archive
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
func (v VisitRepository) AddGuestVisit(db *sql.DB, guest models.Guest, visit models.Visit) (time.Time, error) {

	// First thing first: validating the data, if the data does not meet
	// requirements, an error is issued back to the client
	if guest.ID == 0 {
		errorMsg := `[ERROR] Guest ID value is not set, adding visit failed.`
		return time.Time{}, errors.New(errorMsg)
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
		return time.Time{}, err
	}

	visit, err = visitCalcVisistNext(visit)

	if err != nil {
		return time.Time{}, err
	}

	sqlVisitAdd := `INSERT INTO visits(
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
		return time.Time{}, err
	}

	return visit.DateofVisitNext, nil
}

//UpdateGuestVisit : UpdateGuestVisit
func (v VisitRepository) UpdateGuestVisit(db *sql.DB, guest models.Guest, visit models.Visit) error {

	visitID := int(visit.ID)

	err := visitCheckID(db, visitID)

	if err != nil {
		return err
	}

	visit.LastDateUpdated = time.Now()

	sqlGuestVisitUpd := `UPDATE visits SET
			date_visit = $1, date_visit_next = $2,
			notes = $3, last_date_updated = $4
		WHERE id = $5
			AND guest_id = $6`

	_, err = db.Exec(sqlGuestVisitUpd,
		visit.DateofVisitLast, visit.DateofVisitNext, visit.Notes, visit.LastDateUpdated,
		visit.ID, visit.GuestID)

	if err != nil {
		return err
	}

	return nil
}

//ArchiveGuestVisit : ArchiveGuestVisit
func (v VisitRepository) ArchiveGuestVisit(db *sql.DB, visit models.Visit) error {

	var visitRaw models.VisitRaw
	var archiveDateLast time.Time = time.Now()
	// Set archiveMethod to "D" denoting this record was moved using
	// the application.
	var archiveMethod = "D"

	// We are going to create a transaction because we have a few statements
	// to execute. To ensure every step completes successfully without
	// anyone doing halfway work, and not telling us
	tx, _ := db.Begin()

	sqlVisitRawGet := `SELECT id, guest_id, date_visit, date_visit_next,
			notes, last_date_updated
		FROM visits
		WHERE id = $1
		  AND guest_id = $2`

	row := tx.QueryRow(sqlVisitRawGet, visit.ID, visit.GuestID)

	err := row.Scan(&visitRaw.ID, &visitRaw.GuestID, &visitRaw.DateofVisitLast, &visitRaw.DateofVisitNext,
		&visitRaw.Notes, &visitRaw.LastDateUpdated)

	if err != nil {
		_ = tx.Rollback()
		return err
	}

	visit = visitClean(visitRaw)

	sqlVisitAdd := `INSERT INTO visits_archive(
	  	guest_id,
			date_visit,
			date_visit_next,
	  	notes,
			last_date_updated,
			archive_last_date_updated,
			archive_method)
		VALUES($1, $2, $3, $4, $5, $6, $7)`

	_, err = tx.Exec(sqlVisitAdd,
		visit.GuestID, visit.DateofVisitLast, visit.DateofVisitNext,
		visit.Notes, visit.LastDateUpdated, archiveDateLast, archiveMethod)

	if err != nil {
		_ = tx.Rollback()
		return err
	}

	sqlVisitRemove := `DELETE FROM visits WHERE id = $1 AND guest_id = $2`

	_, err = tx.Exec(sqlVisitRemove, visit.ID, visit.GuestID)

	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_ = tx.Commit()

	return nil
}

//UnarchiveGuestVisit : UnarchiveGuestVisit
func (v VisitRepository) UnarchiveGuestVisit(db *sql.DB, visit models.Visit) error {

	var visitRaw models.VisitRaw

	tx, _ := db.Begin()

	sqlVisitRawGet := `SELECT id, guest_id, date_visit, date_visit_next,
			notes, last_date_updated
		FROM visits_archive
		WHERE id = $1
		  AND guest_id = $2`

	row := tx.QueryRow(sqlVisitRawGet, visit.ID, visit.GuestID)

	err := row.Scan(&visitRaw.ID, &visitRaw.GuestID, &visitRaw.DateofVisitLast, &visitRaw.DateofVisitNext,
		&visitRaw.Notes, &visitRaw.LastDateUpdated)

	if err != nil {
		_ = tx.Rollback()
		// return errors.New("[DEBUG] Error #1: Getting from Archive")
		return err
	}

	visit = visitClean(visitRaw)

	visit.LastDateUpdated = time.Now()

	sqlArchVisitAdd := `INSERT INTO visits(
			id,
	  	guest_id,
			date_visit,
			date_visit_next,
	  	notes,
			last_date_updated)
		VALUES($1, $2, $3, $4, $5, $6)`

	_, err = tx.Exec(sqlArchVisitAdd,
		visit.ID, visit.GuestID, visit.DateofVisitLast, visit.DateofVisitNext,
		visit.Notes, visit.LastDateUpdated)

	if err != nil {
		_ = tx.Rollback()
		// return errors.New("[DEBUG] Error #2: Inserting into Visits table")
		return err
	}

	sqlArchVisitRemove := `DELETE FROM visits_archive WHERE id = $1 AND guest_id = $2`

	_, err = tx.Exec(sqlArchVisitRemove, visit.ID, visit.GuestID)

	if err != nil {
		_ = tx.Rollback()
		// return errors.New("[DEBUG] Error #3: Removing from Archive")
		return err
	}

	_ = tx.Commit()

	return nil
}

// ********************* Helper Functions ********************* //

func visitValidate(db *sql.DB, guest models.Guest, visit models.Visit) (bool, error) {

	// Declare a time variable to store the guests latest next visit to the pantry
	var nextVisitCheck time.Time

	// Retrieves the latest next visit date set from the previous visit, if this is
	// a new Guest, the default value of the visit will be 0001-01-01.
	sqlVisitValidate := `select coalesce(max(date_visit_next), '0001-01-01')::date date_visit_next
			from visits where guest_id = $1`

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

func visitCheckID(db *sql.DB, id int) error {

	var visitID int

	sqlVisitGetID := `SELECT id FROM visits WHERE id = $1`

	row := db.QueryRow(sqlVisitGetID, id)

	err := row.Scan(&visitID)

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("[ERROR] Guest visit does not exist")
		}
		return err
	}

	return nil
}

func visitCalcVisistNext(visit models.Visit) (models.Visit, error) {

	// Set the visit interval by the variable value from the environment
	currentVisitInterval, err := strconv.Atoi(os.Getenv("app_visit_interval"))

	if err != nil {
		return models.Visit{}, err
	}

	// Add the newly set currentVisitInterval to the Guest's visit record for saving.
	visit.DateofVisitNext = visit.DateofVisitLast.AddDate(0, 0, currentVisitInterval)

	return visit, nil
}
