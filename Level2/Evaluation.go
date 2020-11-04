/*
A store sells two types of products: books and games. Each product has the fields name (string) and price (float). Define the following functionalities:
Product has the following methods:
A method to print the information of each product (type, name, and price)
A method to apply a discount ratio to the price
The store should be able to apply custom discounts based on the type of product: 10% discount for books and 20% discount for games.
*/

package main

import (
	"fmt"
)

type Book struct {
	name string
	price float64
}

type Game struct {
	name string
	price float64
}

type Product interface {
	PrintInformation()
	ApplyDiscountRatio() 
}

func (book *Book) PrintInformation() {
	fmt.Printf("(Type: %T, Name: %v, Price: %v)\n", book, book.name, book.price)
}

func (book *Book) ApplyDiscountRatio() {
	book.price *= (1. - 0.10)
}

func (game *Game) PrintInformation() {
	fmt.Printf("(Type: %T, Name: %v, Price: %v)\n", game, game.name, game.price)
}

func (game *Game) ApplyDiscountRatio() {
	game.price *= (1. - 0.20)
}

func main() {
	products := make([] Product, 4)
	products[0] = &Book{"Harry Potter", 20.}
	products[1] = &Book{"Cortex and Mind", 50.}
	products[2] = &Game{"God of War", 30.}
	products[3] = &Game{"Street Fighter 4", 7.}
	for i := range products {
		products[i].ApplyDiscountRatio()
		products[i].PrintInformation()
	}
	fmt.Println("MORE")
	products[3] = nil
	for i := range products {
		switch products[i].(type) {
			case *Book:
				products[i].ApplyDiscountRatio()
			case *Game:
				products[i].ApplyDiscountRatio()
			default:
				fmt.Println("Unknown type")
		}
		if products[i] != nil {products[i].PrintInformation()}
	}
}
