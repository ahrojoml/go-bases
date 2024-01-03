/*
Algunas tiendas ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.

La empresa tiene 3 tipos de productos: Small, Medium y Large. (Se espera que sean muchos más)


Y los costos adicionales son:

Small: solo tiene el costo del producto
Medium: el precio del producto + un 3% de mantenerlo en la tienda
Large: el precio del producto + un 6% de mantenerlo en la tienda, y adicional a eso $2500 de costo de envío.

El porcentaje de mantenerlo en la tienda es en base al precio del producto.

El costo de mantener el producto en stock en la tienda es un porcentaje del precio del producto.


Se requiere una función factory que reciba el tipo de producto y el precio y retorne una interfaz Product que tenga el método Price.


Se debe poder ejecutar el método Price y que el método me devuelva el precio total en base al costo del producto y los adicionales en caso que los tenga
*/

package main

import (
	"errors"
	"fmt"
)

type Product interface {
	Price() float64
}

type Small struct {
	productPrice float64
}

func (s Small) Price() float64 {
	return s.productPrice
}

type Medium struct {
	productPrice float64
}

func (m Medium) Price() float64 {
	return m.productPrice * 1.03
}

type Large struct {
	productPrice float64
}

func (l Large) Price() float64 {
	return l.productPrice*1.03 + 2500
}

const (
	small = iota
	medium
	large
)

func GetProduct(productType int, price float64) (Product, error) {
	switch productType {
	case small:
		return Small{productPrice: price}, nil
	case medium:
		return Medium{productPrice: price}, nil
	case large:
		return Large{productPrice: price}, nil
	default:
		return Small{}, errors.New("Product type not found")
	}
}

func main() {
	var product Product
	product, _ = GetProduct(small, 4000)
	fmt.Println(product.Price())

	product, _ = GetProduct(medium, 4000)
	fmt.Println(product.Price())

	product, _ = GetProduct(large, 4000)
	fmt.Println(product.Price())

}
