package main

import (
	"errors"
	"fmt"
)

type MinimumNotReachedError struct{}

func (e MinimumNotReachedError) Error() string {
	return "Error: the salary entered does not reach the taxable minimum"
}

type MinimumNotReachedError2 struct{}

func (e MinimumNotReachedError2) Error() string {
	return "Error: salary is less than 10000"
}

func (e MinimumNotReachedError2) New(msg string) error {
	return MinimumNotReachedError2{}
}

func GetTax(salary float64) (float64, error) {
	if salary < 10000 {
		return 0, fmt.Errorf("the minimum taxable amount is 150,000 and the salary entered is: %.2f", salary) //MinimumNotReachedError2{}
	}
	var tax float64
	switch {
	case salary > 150000:
		tax += salary * 0.17
		fallthrough
	case salary > 50000:
		tax += salary * 0.1
	}
	return tax, nil
}

func MonthlySalary(workedHours float64, hourlyPay float64) (float64, error) {
	if workedHours < 80 {
		return 0, errors.New("Error: the worker cannot have worked less than 80 hours per month")
	}

	var salary float64 = workedHours * hourlyPay

	if salary >= 150000 {
		salary = salary * 0.9
	}

	return salary, nil
}

func main() {
	var salary, tax, workedHours float64

	salary = 9999
	workedHours = 3

	_, err := MonthlySalary(salary, workedHours)
	if err != nil {
		fmt.Println(err.Error())
		// return
	}

	salary = 60000
	workedHours = -2
	_, err = MonthlySalary(salary, workedHours)
	if err != nil {
		fmt.Println(err.Error())
	}

	salary = 200000
	workedHours = 90
	tax, err = MonthlySalary(salary, workedHours)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(tax)
}
