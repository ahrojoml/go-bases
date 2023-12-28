package main

import "fmt"

func main() {
	var firstName string // can't start with numbers
	var lastName string
	var age int              // variable type should be after variable name
	lastName = "6"           // variable is already instantiated no need to use the implicit statement
	var driverLicense = true // use camelCase
	var personHeight int     // cant use spaces
	childsNumber := 2

	fmt.Println(firstName, lastName, age, lastName, driverLicense, personHeight, childsNumber)
}
