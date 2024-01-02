package animalfood

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

var food = map[string]float64{
	dog:       10,
	cat:       5,
	hamster:   0.250,
	tarantula: 0.150,
}

func Animal(animalType string) (func(float64) float64, error) {
	switch animalType {
	case dog:
		return DogFood, nil
	case cat:
		return CatFood, nil
	case hamster:
		return HamsterFood, nil
	case tarantula:
		return TarantulaFood, nil
	}
	return nil, errors.New(fmt.Sprintf("Animal %s was not found", animalType))
}

func DogFood(required float64) float64 {
	amount := required - food[dog]
	food[dog] = max(food[dog], 0)
	return amount
}

func CatFood(required float64) float64 {
	amount := required - food[cat]
	food[cat] = max(food[cat], 0)
	return amount
}

func HamsterFood(required float64) float64 {
	amount := required - food[hamster]
	food[hamster] = max(food[hamster], 0)
	return amount
}

func TarantulaFood(required float64) float64 {
	amount := required - food[tarantula]
	food[tarantula] = max(food[tarantula], 0)
	return amount
}

func main() {
	var amount float64
	animalDog, err := Animal(dog)
	if err != nil {
		fmt.Println(err.Error())
	}

	animalCat, err := Animal(cat)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = Animal("rat")
	if err != nil {
		fmt.Println(err.Error())
	}

	amount += animalDog(20)
	amount += animalCat(10)

	fmt.Println(amount)
}
