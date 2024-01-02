package main

import "fmt"

func getMean(grades ...int) (mean int) {
	for _, grade := range grades {
		mean += grade
	}
	return mean / len(grades)
}

func main() {
	fmt.Println(getMean(4, 6, 2, 7))
	fmt.Println(getMean(2, 7, 3))
}
