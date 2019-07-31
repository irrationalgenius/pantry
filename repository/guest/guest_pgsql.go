package guestRepository

import (
	"database/sql"
	"log"
	"pantry2/models"
)

//GuestRepository : GuestRepository
type GuestRepository struct{}

//GetGuests : GetGuests
func (g GuestRepository) GetGuests(db *sql.DB, guest models.Guest, guests []models.Guest) ([]models.Guest, error) {
	log.Println("Accessing the Get Guests Repository")

	rows, err := db.Query("select * from pantry.guests")

	if err != nil {
		return []models.Guest{}, err
	}

	for rows.Next() {
		err = rows.Scan(&guest.ID, &guest.Status, &guest.FirstName, &guest.LastName,
			&guest.Gender, &guest.StAddress, &guest.UnitNum, &guest.State, &guest.City, &guest.Zip,
			&guest.TelNum, &guest.Email, &guest.ChildNum, &guest.AdultNum, &guest.IsMember,
			&guest.IsBaptized, &guest.IsEspanol, &guest.IsUnemployed, &guest.IsHomeless,
			&guest.IsFamily, &guest.IsContactOk, &guest.DateEnrolled, &guest.PlaceOfWorship,
			&guest.Notes, &guest.Allergies, &guest.LastDateUpdated)

		guests = append(guests, guest)
	}

	if err != nil {
		return []models.Guest{}, err
	}

	return guests, nil
}
