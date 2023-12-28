package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	employees["Federico"] = 25
	delete(employees, "Pedro")

	// I assume we want the value of employees older than 21 after we make changes in db
	var olderEmployees int = 0
	for _, age := range employees {
		if age > 21 {
			olderEmployees += 1
		}
	}

	fmt.Printf("%d employees are older than 21\n", olderEmployees)
}
