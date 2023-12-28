package main

import "fmt"

func main() {
	var age int
	var employeeResponse string
	var yearsOfExperience int
	var salary int

	fmt.Printf("Age: ")
	_, err := fmt.Scan(&age)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	if age < 22 {
		fmt.Print("Not elegible for loan")
		return
	}

	fmt.Print("Are you employeed (y/n): ")
	_, err = fmt.Scan(&employeeResponse)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	if employeeResponse != "y" {
		fmt.Print("Not elegible for loan")
		return
	}

	fmt.Print("Years of experience: ")
	_, err = fmt.Scan(&yearsOfExperience)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	if yearsOfExperience <= 1 {
		fmt.Print("Not elegible for loan")
		return
	}

	fmt.Print("Salary: ")
	_, err = fmt.Scan(&salary)
	if err != nil {
		fmt.Println("Invalid input, exiting")
		return
	}
	if salary <= 100000 {
		fmt.Print("You are elegible for loan free of interest\n")
	} else {
		fmt.Print("You are elegible for a loan with interests\n")
	}
}
