package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func getMin(grades ...int) (int, error) {
	if len(grades) == 0 {
		return 0, errors.New("No grades given")
	}
	lowestGrade := math.MaxInt
	for _, grade := range grades {
		lowestGrade = min(lowestGrade, grade)
	}
	return lowestGrade, nil
}

func getMax(grades ...int) (int, error) {
	if len(grades) == 0 {
		return 0, errors.New("No grades given")
	}
	maxGrade := 0
	for _, grade := range grades {
		maxGrade = max(maxGrade, grade)
	}
	return maxGrade, nil
}

func getMean(grades ...int) (int, error) {
	if len(grades) == 0 {
		return 0, errors.New("No grades given")
	}
	var mean int
	for _, grade := range grades {
		mean += grade
	}
	return mean / len(grades), nil
}

func getOperation(operation string) (func(...int) (int, error), error) {
	switch operation {
	case minimum:
		return getMin, nil
	case maximum:
		return getMax, nil
	case average:
		return getMean, nil
	}
	return nil, errors.New("No valid operation given")
}

func main() {
	op, err := getOperation(minimum)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(op(2, 3, 3, 4, 10, 2, 4, 5))

	op, err = getOperation(maximum)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(op(2, 3, 3, 4, 10, 2, 4, 5))

	op, err = getOperation(average)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(op(2, 3, 3, 4, 10, 2, 4, 5))
}
