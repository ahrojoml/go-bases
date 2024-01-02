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
	amount := required - food[dog]
	food[dog] = max(food[dog], 0)
	return amount
}

func catFood(required float32) float32 {
	amount := required - food[cat]
	food[cat] = max(food[cat], 0)
	return amount
}

func hamsterFood(required float32) float32 {
	amount := required - food[hamster]
	food[hamster] = max(food[hamster], 0)
	return amount
}

func tarantulaFood(required float32) float32 {
	amount := required - food[tarantula]
	food[tarantula] = max(food[tarantula], 0)
	return amount
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
