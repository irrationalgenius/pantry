package repository

import (
	"database/sql"
	"errors"
	"pantry/models"
	"time"
)

//GuestRepository : GuestRepository
type GuestRepository struct{}

//GetGuest : GetGuest
func (g GuestRepository) GetGuest(db *sql.DB, id int) (models.Guest, error) {

	// Initialize an instance of the GuestRaw struct
	var guestRaw models.GuestRaw

	sqlGuestRawGet := `SELECT id, date_enrolled, status, first_name, last_name, gender,
			unit_num, st_address, state, city, zip, tel_num, email,
			count_children, count_adults, worship_place, is_member, is_baptized,
			is_espanol, is_unemployed, is_homeless, is_family,
			is_contact_ok, allergies, notes, last_date_updated
		FROM pantry.guests
		WHERE id = $1`

	rows, err := db.Query(sqlGuestRawGet, id)

	if err != nil {
		errorMsg := `[ERROR] Issue occured while retrieving data from the database.`
		return models.Guest{}, errors.New(errorMsg)
	}

	for rows.Next() {
		err = rows.Scan(&guestRaw.ID, &guestRaw.DateEnrolled, &guestRaw.Status, &guestRaw.FirstName, &guestRaw.LastName, &guestRaw.Gender,
			&guestRaw.UnitNum, &guestRaw.StAddress, &guestRaw.State, &guestRaw.City, &guestRaw.Zip, &guestRaw.TelNum, &guestRaw.Email,
			&guestRaw.ChildNum, &guestRaw.AdultNum, &guestRaw.PlaceOfWorship, &guestRaw.IsMember, &guestRaw.IsBaptized,
			&guestRaw.IsEspanol, &guestRaw.IsUnemployed, &guestRaw.IsHomeless, &guestRaw.IsFamily,
			&guestRaw.IsContactOk, &guestRaw.Allergies, &guestRaw.Notes, &guestRaw.LastDateUpdated)
	}

	if err != nil {
		errorMsg := `[ERROR] Issue occured while assigning data from the database.`
		return models.Guest{}, errors.New(errorMsg)
	}

	// After the data is retrieved from the database, it must be
	// cleaned, meaning NULL values set to Golang defaults
	// (Go has no idea what a NULL is, and neither do I :|)
	guest := guestClean(guestRaw)

	return guest, nil
}

//GetGuests : GetGuests
func (g GuestRepository) GetGuests(db *sql.DB) ([]models.Guest, int, error) {

	// Initialize an instance of the GuestRaw struct
	var guestRaw models.GuestRaw
	// Initialize an slice instance of the GuestRaw struct
	var guestsRaw []models.GuestRaw

	sqlGuestsRawGet := `SELECT id, date_enrolled, status, first_name, last_name, gender,
			unit_num, st_address, state, city, zip, tel_num, email,
			count_children, count_adults, worship_place, is_member, is_baptized,
			is_espanol, is_unemployed, is_homeless, is_family,
			is_contact_ok, allergies, notes, last_date_updated
		FROM pantry.guests`

	rows, err := db.Query(sqlGuestsRawGet)

	if err != nil {
		errorMsg := `[ERROR] Issue occured while retrieving data from the database.`
		return []models.Guest{}, 0, errors.New(errorMsg)
	}

	for rows.Next() {
		err = rows.Scan(&guestRaw.ID, &guestRaw.DateEnrolled, &guestRaw.Status, &guestRaw.FirstName, &guestRaw.LastName, &guestRaw.Gender,
			&guestRaw.UnitNum, &guestRaw.StAddress, &guestRaw.State, &guestRaw.City, &guestRaw.Zip, &guestRaw.TelNum, &guestRaw.Email,
			&guestRaw.ChildNum, &guestRaw.AdultNum, &guestRaw.PlaceOfWorship, &guestRaw.IsMember, &guestRaw.IsBaptized,
			&guestRaw.IsEspanol, &guestRaw.IsUnemployed, &guestRaw.IsHomeless, &guestRaw.IsFamily,
			&guestRaw.IsContactOk, &guestRaw.Allergies, &guestRaw.Notes, &guestRaw.LastDateUpdated)

		guestsRaw = append(guestsRaw, guestRaw)
	}

	if err != nil {
		errorMsg := `[ERROR] Issue occured while assigning data from the database.`
		return []models.Guest{}, 0, errors.New(errorMsg)
	}

	// After the data is retrieved from the database, it must be
	// cleaned, meaning NULL values set to Golang defaults
	// (Go has no idea what a NULL is, and neither do I :|)
	guests := guestsClean(guestsRaw)

	// Let's count the items in the group. This will be used for logging
	// purposes.
	guestsSize := len(guests)

	return guests, guestsSize, nil
}

//AddGuest : AddGuest
func (g GuestRepository) AddGuest(db *sql.DB, guest models.Guest) (int8, error) {

	// First thing first: validating the data, if the data does not meet
	// requirements, an error is issued back to the client
	guest, err := guestValidate(guest)

	if err != nil {
		return 0, err
	}

	sqlGuestAdd := `INSERT INTO pantry.guests(
			status, first_name, last_name, gender,
			unit_num, st_address, state, city, zip, tel_num, email,
			count_children, count_adults, worship_place, is_member, is_baptized,
			is_espanol, is_unemployed, is_homeless, is_family,
			is_contact_ok, allergies, notes)
		VALUES($1, $2, $3, $4,
		  $5, $6, $7, $8, $9, $10, $11,
		  $12, $13, $14, $15, $16,
		  $17, $18, $19, $20,
		  $21, $22, $23)
		RETURNING id`

	result := db.QueryRow(sqlGuestAdd,
		guest.Status, guest.FirstName, guest.LastName, guest.Gender,
		guest.UnitNum, guest.StAddress, guest.State, guest.City, guest.Zip, guest.TelNum, guest.Email,
		guest.ChildNum, guest.AdultNum, guest.PlaceOfWorship, guest.IsMember, guest.IsBaptized,
		guest.IsEspanol, guest.IsUnemployed, guest.IsHomeless, guest.IsFamily,
		guest.IsContactOk, guest.Allergies, guest.Notes)

	err = result.Scan(&guest.ID)

	if err != nil {
		errorMsg := `[ERROR] Issue occured while setting Guests variables in db.`
		return 0, errors.New(errorMsg)
	}

	return guest.ID, nil
}

//UpdateGuest : UpdateGuest
func (g GuestRepository) UpdateGuest(db *sql.DB, guest models.Guest) error {

	// First thing first: validating the data, if the data does not meet
	// requirements, an error is issued back to the client
	guest, err := guestValidate(guest)

	if err != nil {
		return err
	}

	// After validation, we need to set the current date in which this
	// record was updated. It is good record-keeping habits to keep the
	// most current details of every Guest.
	guest.LastDateUpdated = time.Now()

	sqlGuestUpdate := `UPDATE pantry.guests SET
			status = $1, first_name = $2, last_name = $3, gender = $4,
			unit_num = $5, st_address = $6, state = $7, city = $8,
			zip = $9, tel_num = $10, email = $11, count_children = $12,
			count_adults = $13, worship_place = $14, is_member = $15,
			is_baptized = $16, is_espanol = $17, is_unemployed = $18,
			is_homeless = $19, is_family = $20, is_contact_ok = $21,
			allergies = $22, notes = $23, last_date_updated = $24
		WHERE id = $25`

	_, err = db.Exec(sqlGuestUpdate,
		guest.Status, guest.FirstName, guest.LastName, guest.Gender,
		guest.UnitNum, guest.StAddress, guest.State, guest.City, guest.Zip, guest.TelNum, guest.Email,
		guest.ChildNum, guest.AdultNum, guest.PlaceOfWorship, guest.IsMember, guest.IsBaptized,
		guest.IsEspanol, guest.IsUnemployed, guest.IsHomeless, guest.IsFamily,
		guest.IsContactOk, guest.Allergies, guest.Notes, guest.LastDateUpdated, guest.ID)

	if err != nil {
		return err
	}

	return nil
}

//RemoveGuest : RemoveGuest
func (g GuestRepository) RemoveGuest(db *sql.DB) error {

	return nil
}

// ********************* Helper Functions ********************* //

func guestClean(guestRaw models.GuestRaw) models.Guest {
	var guest models.Guest

	// The following elements must have values from
	// the database, otherwise a error will result
	if guestRaw.ID.Valid == true {
		guestID := int8(guestRaw.ID.Int64)
		guest.ID = guestID
	}
	if guestRaw.DateEnrolled.Valid == true {
		guest.DateEnrolled = guestRaw.DateEnrolled.Time
	} else {
		guest.DateEnrolled = time.Now()
	}
	if guestRaw.Status.Valid == true {
		guest.Status = guestRaw.Status.String
	}
	if guestRaw.FirstName.Valid == true {
		guest.FirstName = guestRaw.FirstName.String
	}
	if guestRaw.LastName.Valid == true {
		guest.LastName = guestRaw.LastName.String
	}
	if guestRaw.LastDateUpdated.Valid == true {
		guest.LastDateUpdated = guestRaw.LastDateUpdated.Time
	} else {
		guest.LastDateUpdated = time.Now()
	}

	// Elements can have values or be empty
	if guestRaw.Gender.Valid == true {
		guest.Gender = guestRaw.Gender.String
	} else {
		guest.Gender = ""
	}
	if guestRaw.UnitNum.Valid == true {
		guest.UnitNum = guestRaw.UnitNum.String
	} else {
		guest.UnitNum = ""
	}
	if guestRaw.StAddress.Valid == true {
		guest.StAddress = guestRaw.StAddress.String
	} else {
		guest.StAddress = ""
	}
	if guestRaw.State.Valid == true {
		guest.State = guestRaw.State.String
	} else {
		guest.State = ""
	}
	if guestRaw.City.Valid == true {
		guest.City = guestRaw.City.String
	} else {
		guest.City = ""
	}
	if guestRaw.Zip.Valid == true {
		guest.Zip = guestRaw.Zip.String
	} else {
		guest.Zip = ""
	}
	if guestRaw.TelNum.Valid == true {
		guest.TelNum = guestRaw.TelNum.String
	} else {
		guest.TelNum = ""
	}
	if guestRaw.Email.Valid == true {
		guest.Email = guestRaw.Email.String
	} else {
		guest.Email = ""
	}
	if guestRaw.ChildNum.Valid == true {
		// childNum := int8(guestRaw.ChildNum.Int64)
		guest.ChildNum = guestRaw.ChildNum.Int64
	} else {
		guest.ChildNum = 0
	}
	if guestRaw.AdultNum.Valid == true {
		// adultNum := int8(guestRaw.AdultNum.Int64)
		guest.AdultNum = guestRaw.AdultNum.Int64
	} else {
		guest.AdultNum = 0
	}
	if guestRaw.PlaceOfWorship.Valid == true {
		guest.PlaceOfWorship = guestRaw.PlaceOfWorship.String
	} else {
		guest.PlaceOfWorship = ""
	}
	if guestRaw.IsMember.Valid == true {
		guest.IsMember = guestRaw.IsMember.String
	} else {
		guest.IsMember = ""
	}
	if guestRaw.IsEspanol.Valid == true {
		guest.IsEspanol = guestRaw.IsEspanol.String
	} else {
		guest.IsEspanol = ""
	}
	if guestRaw.IsBaptized.Valid == true {
		guest.IsBaptized = guestRaw.IsBaptized.String
	} else {
		guest.IsBaptized = ""
	}
	if guestRaw.IsUnemployed.Valid == true {
		guest.IsUnemployed = guestRaw.IsUnemployed.String
	} else {
		guest.IsUnemployed = ""
	}
	if guestRaw.IsHomeless.Valid == true {
		guest.IsHomeless = guestRaw.IsHomeless.String
	} else {
		guest.IsHomeless = ""
	}
	if guestRaw.IsFamily.Valid == true {
		guest.IsFamily = guestRaw.IsFamily.String
	} else {
		guest.IsFamily = ""
	}
	if guestRaw.IsContactOk.Valid == true {
		guest.IsContactOk = guestRaw.IsContactOk.String
	} else {
		guest.IsContactOk = ""
	}
	if guestRaw.Allergies.Valid == true {
		guest.Allergies = guestRaw.Allergies.String
	} else {
		guest.Allergies = ""
	}
	if guestRaw.Notes.Valid == true {
		guest.Notes = guestRaw.Notes.String
	} else {
		guest.Notes = ""
	}

	return guest
}

func guestsClean(guestsRaw []models.GuestRaw) []models.Guest {

	var guests []models.Guest

	for _, guestRaw := range guestsRaw {

		guest := guestClean(guestRaw)

		guests = append(guests, guest)
	}

	return guests
}

func guestValidate(guest models.Guest) (models.Guest, error) {

	if guest.FirstName == "" {
		validationMsg := "[ERROR] First name cannot be empty."
		return models.Guest{}, errors.New(validationMsg)
	}
	if guest.LastName == "" {
		validationMsg := "[ERROR] Last name cannot be empty."
		return models.Guest{}, errors.New(validationMsg)
	}
	if guest.StAddress == "" && guest.IsHomeless == "N" {
		validationMsg := "[ERROR] Street Address cannot be empty, unless guest is homeless."
		return models.Guest{}, errors.New(validationMsg)
	}

	return guest, nil
}
