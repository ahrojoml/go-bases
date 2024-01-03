package main

import "fmt"

func SalaryTypeC(minutes int) float64 {
	return 1000 * (float64(minutes) / 60)
}

func SalaryTypeB(minutes int) float64 {
	return (1500 * (float64(minutes) / 60)) * 1.2
}

func SalaryTypeA(minutes int) float64 {
	return (3000 * (float64(minutes) / 60)) * 1.50
}

func GetSalary(minutes int, salaryType string) float64 {
	var salary float64
	switch salaryType {
	case "A":
		salary = SalaryTypeA(minutes)
	case "B":
		salary = SalaryTypeB(minutes)
	case "C":
		salary = SalaryTypeC(minutes)
	}

	return salary
}

func main() {
	fmt.Println(GetSalary(60, "C"))
	fmt.Println(GetSalary(60, "B"))
	fmt.Println(GetSalary(60, "A"))
}
