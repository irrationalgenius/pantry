package models

//ConfigParam : ConfigParam
type ConfigParam struct {
	//ID : Unique Identifier
	ID int8 `json:id`

	//Name : Property Name
	Name string `json:name`

	//Category : Property Category
	Category string `json:category`

	//Value : Value of the Property
	Value string `json:value`
}
