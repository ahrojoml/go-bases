package getgrades

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

func GetMin(grades ...int) (int, error) {
	if len(grades) == 0 {
		return 0, errors.New("No grades given")
	}
	lowestGrade := math.MaxInt
	for _, grade := range grades {
		lowestGrade = min(lowestGrade, grade)
	}
	return lowestGrade, nil
}

func GetMax(grades ...int) (int, error) {
	if len(grades) == 0 {
		return 0, errors.New("No grades given")
	}
	maxGrade := 0
	for _, grade := range grades {
		maxGrade = max(maxGrade, grade)
	}
	return maxGrade, nil
}

func GetMean(grades ...int) (int, error) {
	if len(grades) == 0 {
		return 0, errors.New("No grades given")
	}
	var mean int
	for _, grade := range grades {
		mean += grade
	}
	return mean / len(grades), nil
}

func GetOperation(operation string) (func(...int) (int, error), error) {
	switch operation {
	case minimum:
		return GetMin, nil
	case maximum:
		return GetMax, nil
	case average:
		return GetMean, nil
	}
	return nil, errors.New("No valid operation given")
}

func main() {
	op, err := GetOperation(minimum)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(op(2, 3, 3, 4, 10, 2, 4, 5))

	op, err = GetOperation(maximum)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(op(2, 3, 3, 4, 10, 2, 4, 5))

	op, err = GetOperation(average)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(op(2, 3, 3, 4, 10, 2, 4, 5))
}
