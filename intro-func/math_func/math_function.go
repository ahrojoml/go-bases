package main

import "fmt"

const (
	Addition       = "+"
	Subtraction    = "-"
	Multiplication = "*"
	Division       = "/"
)

func calculate(val1, val2 float64, operation string) float64 {
	switch operation {
	case Addition:
		return val1 + val2
	case Subtraction:
		return val1 - val2
	case Multiplication:
		return val1 * val2
	case Division:
		if val2 != 0 {
			return val1 / val2
		}
	}
	return 0
}

func main() {
	fmt.Println(calculate(6, 2, Addition))
	fmt.Println(calculate(6, 2, Subtraction))
	fmt.Println(calculate(6, 2, Multiplication))
	fmt.Println(calculate(6, 2, Division))
}
