package utils

import "log"

//Tester : Tester
type Tester struct{}

//TestFunc : TestFunc
func (l Tester) TestFunc() {
	log.Println("This package works")
}
