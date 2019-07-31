package repository

import (
	"database/sql"
	"fmt"
	"log"
	"pantry2/models"
	"time"
)

// VisitRepository : VisitRepository
type VisitRepository struct{}

func calcVisitNext(day time.Time) time.Time {

	return day
}

//AddVisit : AddVisit
func (v VisitRepository) AddVisit(db *sql.DB, guest models.Guest, visit models.Visit) error {
	log.Println("Accessing the Add Visits Repository")

	result, err := db.Exec(`INSERT INTO pantry.visits(
                          guest_id, date_visit, date_visit_next,
                          notes, last_date_updated)
                        VALUES($1, $2, $3, $4, $5)`)

	if err != nil {
		return err
	}

	result.RowsAffected()

	fmt.Println("Visit is Successfully Saved.")

	return nil
}
