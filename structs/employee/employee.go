package main

import (
	"fmt"
	"time"
)

type Person struct {
	id          int
	name        string
	dateOfBirth time.Time
}

type Employee struct {
	id       int
	position string
	person   Person
}

func (e Employee) PrintEmployee() {
	fmt.Printf(
		"Employee %d, %s, %s, %s\n",
		e.id, e.person.name, e.position, e.person.dateOfBirth.Format("02 - 01 - 2006"),
	)
}

func main() {
	person := Person{
		id:          1,
		name:        "churu churu",
		dateOfBirth: time.Now(),
	}
	employee := Employee{
		id:       1,
		position: "food taster",
		person:   person,
	}

	employee.PrintEmployee()
}
