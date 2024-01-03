package main

import (
	"errors"
	"fmt"
)

type Product struct {
	id          int
	name        string
	price       int
	description string
	category    string
}

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	for _, product := range Products {
		fmt.Println(product)
	}
}

var Products = []Product{
	{
		id:          1,
		name:        "cookie",
		price:       200,
		description: "delicious",
		category:    "food",
	},
	{
		id:          2,
		name:        "ruler",
		price:       300,
		description: "to make straight lines",
		category:    "tools",
	},
}

func getById(id int) (Product, error) {
	for _, product := range Products {
		if product.id == id {
			return product, nil
		}
	}
	return Product{}, errors.New(fmt.Sprintf("product %d was not found", id))
}

func main() {
	var product Product = Product{
		id:          3,
		name:        "lemon",
		price:       400,
		description: "acid",
	}

	product.GetAll()
	_, err := getById(3)
	if err != nil {
		fmt.Println(err.Error())
	}
	product.Save()

	product.GetAll()
	prod2, err := getById(3)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(prod2)
}
