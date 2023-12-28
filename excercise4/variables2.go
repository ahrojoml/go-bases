package main

import "fmt"

func main() {
	var lastName string = "Smith"
	var age int = 35               // variable cannot be type string
	boolean := false               // Should define a better variable name, booleans are true or false not string
	var salary string = "45857.90" // should be a string
	var firstName string = "Mary"

	fmt.Println(lastName, age, boolean, salary, firstName)
}
