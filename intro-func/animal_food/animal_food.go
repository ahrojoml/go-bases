package main

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

var food = map[string]float32{
	dog:       10,
	cat:       5,
	hamster:   0.250,
	tarantula: 0.150,
}

func animal(animalType string) (func(float32) float32, error) {
	switch animalType {
	case dog:
		return dogFood, nil
	case cat:
		return catFood, nil
	case hamster:
		return hamsterFood, nil
	case tarantula:
		return tarantulaFood, nil
	}
	return nil, errors.New(fmt.Sprintf("Animal %s was not found", animalType))
}

func dogFood(required float32) float32 {
	return required - food[dog]
}

func catFood(required float32) float32 {
	return required - food[cat]
}

func hamsterFood(required float32) float32 {
	return required - food[hamster]
}

func tarantulaFood(required float32) float32 {
	return required - food[tarantula]
}

func main() {
	var amount float32
	animalDog, err := animal(dog)
	if err != nil {
		fmt.Println(err.Error())
	}

	animalCat, err := animal(cat)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = animal("rat")
	if err != nil {
		fmt.Println(err.Error())
	}

	amount += animalDog(20)
	amount += animalCat(10)

	fmt.Println(amount)
}
