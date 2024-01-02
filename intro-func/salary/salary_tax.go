package salary

import "fmt"

func GetTax(salary float64) float64 {
	var tax float64
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
	fmt.Println(GetTax(10000))
	fmt.Println(GetTax(60000))
	fmt.Println(GetTax(200000))
}
