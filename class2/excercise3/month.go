package main

import "fmt"

var numberToMonth = map[int]string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}

func main() {
	var monthNumber int

	_, err := fmt.Scan(&monthNumber)
	if err != nil {
		fmt.Println("Invalid Value")
		return
	}

	month, ok := numberToMonth[monthNumber]
	if !ok {
		fmt.Println("Month number must be from 1 to 12")
		return
	}

	fmt.Println(month)
}
