package main

import "fmt"

func getTax(salary float32) float32 {
	var tax float32
	switch {
	case salary > 150000:
		tax += salary * 0.17
		fallthrough
	case salary > 50000:
		tax += salary * 0.1
	}
	return tax
}

func main() {
	fmt.Println(getTax(10000))
	fmt.Println(getTax(60000))
	fmt.Println(getTax(200000))
}
