package models

import "time"

//Visit : Visit
type Visit struct {
	//ID : Unique Identifier
	ID int8 `json:id`

	//GuestID : Pointer to a Guest "ID"
	GuestID int8 `json:guestid`

	//DateofVisitLast : Official date of pantry visit
	DateofVisitLast time.Time `json:dateofvisitlast`

	//DateofVisitNext : Official next date of pantry visit
	DateofVisitNext time.Time `json:dateofvisitnext`

	//Notes : Additional remarks
	Notes string `json:notes`

	//LastDateUpdated : Last date of record  update
	LastDateUpdated time.Time `json:lastdateupdated`
}
