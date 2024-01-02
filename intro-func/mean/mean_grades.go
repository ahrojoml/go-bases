package mean

import "fmt"

func GetMean(grades ...int) (mean int) {
	for _, grade := range grades {
		mean += grade
	}
	return mean / len(grades)
}

func main() {
	fmt.Println(GetMean(4, 6, 2, 7))
	fmt.Println(GetMean(2, 7, 3))
}
