package models

import "time"

//Guest : Captures revelent information from
// people (Guests) which utilize pantry services.
type Guest struct {
	//ID : Unique Identifier
	ID int8 `json:id`

	//Status : Current status [A=Active; I=Inactive; W=Waiting; X=Archive]
	Status string `json:status`

	//FirstName : First name
	FirstName string `json:firstname`

	//LastName : Last name
	LastName string `json:lastname`

	//Gender : Gender [M=Male; F=Female]
	Gender string `json:gender`

	//StAddress : Current Street Address
	StAddress string `json:staddress`

	//UnitNum : Number of the adress unit: street, apt, etc.
	UnitNum int8 `json:unitnum`

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
	ChildNum int16 `json:chldnum`

	//AdultNum : Total Household adult count
	AdultNum int16 `json:adultnum`

	//IsMember : Is a Member of the Laurel Church of Christ
	IsMember int8 `json:ismember`

	//IsBaptized : Is Baptized into Jesus Christ
	IsBaptized int8 `json:isbaptized`

	//IsEspanol : Is Spanish speaking, Hispanic
	IsEspanol int8 `json:isespanol`

	//IsUnemployed : This Guest is looking for employment
	IsUnemployed int8 `json:isunemployed`

	//IsHomeless : Is the Guest Homeless? Does not have a Home? Stable roof over the head
	IsHomeless int8 `json:ishomeless`

	//IsFamily : Has at least 2 people that lives with guest
	IsFamily int8 `json:isfamily`

	//IsContactOk : Is okay for Contact? [0 = No, -1 = Yes]
	IsContactOk int8 `json:iscontactok`

	//DateEnrolled : The date of the first pantry visit
	DateEnrolled time.Time `json:dateenrolled`

	//PlaceOfWorship : Primary place of worship
	PlaceOfWorship string `json:placeofworship`

	//Notes : Additional remarks
	Notes string `json:notes`

	//Allergies : Note of known allergies pertaining to this guest
	Allergies string `json:allergies`

	//LastDateUpdated : Last date of record  update
	LastDateUpdated time.Time `json:lastdateupdated`
}
