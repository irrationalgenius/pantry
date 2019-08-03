package models

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

//Guest : Captures revelent information from
// people (Guests) which utilize pantry services.
type Guest struct {
	//ID : Unique Identifier
	ID int8 `json:id`

	//DateEnrolled : The date of the first pantry visit
	DateEnrolled time.Time `json:dateenrolled`

	//Status : Current status [A=Active; I=Inactive; W=Waiting; X=Archive]
	Status string `json:status`

	//FirstName : First name
	FirstName string `json:firstname`

	//LastName : Last name
	LastName string `json:lastname`

	//Gender : Gender [M=Male; F=Female]
	Gender string `json:gender`

	//UnitNum : Number of the adress unit: street, apt, etc.
	UnitNum string `json:unitnum`

	//StAddress : Current Street Address
	StAddress string `json:staddress`

	//State : State code
	State string `json:state`

	//City : Current City
	City string `json:city`

	//Zip : 5 digit Zip code
	Zip string `json:zip`

	//Telnum : Current Primary contact number
	TelNum string `json:telnum`

	//Email : Current Primary email address
	Email string `json:email`

	//ChildNum : Total Household child count
	ChildNum int64 `json:chldnum`

	//AdultNum : Total Household adult count
	AdultNum int64 `json:adultnum`

	//PlaceOfWorship : Primary place of worship
	PlaceOfWorship string `json:placeofworship`

	//IsMember : Is a Member of the Laurel Church of Christ
	IsMember string `json:ismember`

	//IsBaptized : Is Baptized into Jesus Christ
	IsBaptized string `json:isbaptized`

	//IsEspanol : Is Spanish speaking, Hispanic
	IsEspanol string `json:isespanol`

	//IsUnemployed : This Guest is looking for employment
	IsUnemployed string `json:isunemployed`

	//IsHomeless : Is the Guest Homeless? Does not have a Home? Stable roof over the head
	IsHomeless string `json:ishomeless`

	//IsFamily : Has at least 2 people that lives with guest
	IsFamily string `json:isfamily`

	//IsContactOk : Is okay for Contact? [0 = No, -1 = Yes]
	IsContactOk string `json:iscontactok`

	//Allergies : Note of known allergies pertaining to this guest
	Allergies string `json:allergies`

	//Notes : Additional remarks
	Notes string `json:notes`

	//LastDateUpdated : Last date of record  update
	LastDateUpdated time.Time `json:lastdateupdated`
}

//GuestRaw :
type GuestRaw struct {
	//ID : Unique Identifier
	ID sql.NullInt64 `json:id`

	//DateEnrolled : The date of the first pantry visit
	DateEnrolled pq.NullTime `json:dateenrolled`

	//Status : Current status [A=Active; I=Inactive; W=Waiting; X=Archive]
	Status sql.NullString `json:status`

	//FirstName : First name
	FirstName sql.NullString `json:firstname`

	//LastName : Last name
	LastName sql.NullString `json:lastname`

	//Gender : Gender [M=Male; F=Female]
	Gender sql.NullString `json:gender`

	//UnitNum : Number of the adress unit: street, apt, etc.
	UnitNum sql.NullString `json:unitnum`

	//StAddress : Current Street Address
	StAddress sql.NullString `json:staddress`

	//State : State code
	State sql.NullString `json:state`

	//City : Current City
	City sql.NullString `json:city`

	//Zip : 5 digit Zip code
	Zip sql.NullString `json:zip`

	//Telnum : Current Primary contact number
	TelNum sql.NullString `json:telnum`

	//Email : Current Primary email address
	Email sql.NullString `json:email`

	//ChildNum : Total Household child count
	ChildNum sql.NullInt64 `json:chldnum`

	//AdultNum : Total Household adult count
	AdultNum sql.NullInt64 `json:adultnum`

	//PlaceOfWorship : Primary place of worship
	PlaceOfWorship sql.NullString `json:placeofworship`

	//IsMember : Is a Member of the Laurel Church of Christ
	IsMember sql.NullString `json:ismember`

	//IsBaptized : Is Baptized into Jesus Christ
	IsBaptized sql.NullString `json:isbaptized`

	//IsEspanol : Is Spanish speaking, Hispanic
	IsEspanol sql.NullString `json:isespanol`

	//IsUnemployed : This Guest is looking for employment
	IsUnemployed sql.NullString `json:isunemployed`

	//IsHomeless : Is the Guest Homeless? Does not have a Home? Stable roof over the head
	IsHomeless sql.NullString `json:ishomeless`

	//IsFamily : Has at least 2 people that lives with guest
	IsFamily sql.NullString `json:isfamily`

	//IsContactOk : Is okay for Contact? [0 = No, -1 = Yes]
	IsContactOk sql.NullString `json:iscontactok`

	//Allergies : Note of known allergies pertaining to this guest
	Allergies sql.NullString `json:allergies`

	//Notes : Additional remarks
	Notes sql.NullString `json:notes`

	//LastDateUpdated : Last date of record  update
	LastDateUpdated pq.NullTime `json:lastdateupdated`
}
