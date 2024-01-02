package main

import "fmt"

func salaryTypeC(minutes int) float32 {
	return 1000 * (float32(minutes) / 60)
}

func salaryTypeB(minutes int) float32 {
	return (1500 * (float32(minutes) / 60)) * 1.2
}

func salaryTypeA(minutes int) float32 {
	return (3000 * (float32(minutes) / 60)) * 1.56
}

func getSalary(minutes int, salaryType string) float32 {
	var salary float32
	switch salaryType {
	case "A":
		salary = salaryTypeA(minutes)
	case "B":
		salary = salaryTypeB(minutes)
	case "C":
		salary = salaryTypeC(minutes)
	}

	return salary
}

func main() {
	fmt.Println(getSalary(60, "C"))
	fmt.Println(getSalary(60, "B"))
	fmt.Println(getSalary(60, "A"))
}
