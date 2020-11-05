/*
Simple Inventory
Statement
Create a Product struct with an attribute ID and Name, and an Inventory struct that will contain a map[string]Products. Create a method with a pointer of Inventory as a receiver and It will expect a product to be added into the map of products. 
Required: Add validations into the Add method, to check that the ID can’t be empty and also not duplicated ID

Topics to Practice:
struct, composition, methods, pointer, error handling
*/

package main

import (
	"fmt"
)

type Product struct {
	ID, Name string
}

type Inventory struct {
	products map[string]Product
}

func (inventory *Inventory) AddProduct(product Product) string {
	if product.ID == "" {return "ERROR: Empty ID."}
	_, ok := inventory.products[product.ID]
	if ok == false {
		inventory.products[product.ID] = product
		return "OK: Product added."
	} else {
		return "ERROR: Duplicated ID."
	}
}

func main() {
	products := []Product {
		Product {"123", "Soap"},
		Product {"124", "Perfume"},
		Product {"125", "Shampoo"},
		Product {"126", "Deodorant"},
		Product {"123", "Duplicated 1"},
		Product {"124", "Duplicated 2"},
		Product {"", "Empty 1"},
		Product {"", "Empty 2"},
		Product {"127", "Spray"},
	}
	inventory := &Inventory{make(map[string]Product)}
	for _, product := range products {
		result := inventory.AddProduct(product)
		fmt.Println(result)
	}
	fmt.Println("\nPRODUCTS:")
	for id, product := range inventory.products {
		fmt.Printf("ID: %v, Product: %v\n", id, product.Name)
	}
}

/*
OUTPUT:

OK: Product added.
OK: Product added.
OK: Product added.
OK: Product added.
ERROR: Duplicated ID.
ERROR: Duplicated ID.
ERROR: Empty ID.
ERROR: Empty ID.
OK: Product added.

PRODUCTS:
ID: 124, Product: Perfume
ID: 125, Product: Shampoo
ID: 126, Product: Deodorant
ID: 127, Product: Spray
ID: 123, Product: Soap
*/
