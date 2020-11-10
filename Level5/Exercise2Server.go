/*
Statement
Create a RESTful API implementing all the CRUD Actions, and integrate using gorilla/mux

Create a Product struct with the following attributes: [ID string, Code string, Name string, Price float64] 

Create a ProductInventory struct with the following attributes: [Product Product, Quantity int] 

Then to simulate a table in memory, create a var inventory as a []ProductInventory

Now, Create the functions [Add, Update, Delete, Get] that will be executed against inventory. For example: Add will add a new ProductInventory into the inventory ([]ProductInventory) 

Topics to Practice: 
net/http, RESTful, []byte, log pkg, gorilla/mux
*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

type Product struct {
    ID string
    Code string
    Name string
    Price float64
}

type ProductInventory struct {
    Product Product
    Quantity int
}

var inventory []ProductInventory

func NewProductInventory(ID string, Code string, Name string, Price float64, Quantity int) ProductInventory {
    return ProductInventory {Product {ID, Code, Name, Price}, Quantity}
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%+v", inventory)
}

func FindProduct(id string) (ProductInventory, int) {
    for id1, pi := range inventory {
        if id == pi.Product.ID {
            return inventory[id1], id1
        }
    }     
    return ProductInventory {}, -1
}

func GetOneProduct(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    pi, index := FindProduct(id)
    if index != -1 {
        fmt.Fprintf(w, "%+v", pi)
    } else {
        fmt.Fprintf(w, "ERROR: Product ID does not exist.")
    }
}

type RawProduct struct {
    ID, Code, Name, Price, Quantity string
}

func RawProductToProductInventory(rawProduct RawProduct) ProductInventory {
    var price float64
    var quantity int64
    var convErr error
    price, convErr = strconv.ParseFloat(rawProduct.Price, 64)
    quantity, convErr = strconv.ParseInt(rawProduct.Quantity, 10, 64)
    if convErr != nil {
        panic(convErr)
    }
    return NewProductInventory(rawProduct.ID, rawProduct.Code, rawProduct.Name, price, int(quantity)) 
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
    fmt.Println("CreateProduct")
    reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Fprintf(w, "ERROR: Body of request must contain a product description.")
    }
    var rawProduct RawProduct
    json.Unmarshal(reqBody, &rawProduct)
    pi := RawProductToProductInventory(rawProduct)
    _, index := FindProduct(pi.Product.ID)
    if index == -1 {
        inventory = append(inventory, pi)
        fmt.Fprintf(w, "Product created: %+v", pi)
    } else {
        fmt.Fprintf(w, "ERROR: Product ID already exists.")
    }
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
    fmt.Println("UpdateProduct")
    reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Fprintf(w, "ERROR: Body of request must contain a product description.")
    }
    var rawProduct RawProduct
    json.Unmarshal(reqBody, &rawProduct)
    pi := RawProductToProductInventory(rawProduct)
    _, index := FindProduct(pi.Product.ID)
    if index != -1 {
        inventory[index] = pi
        fmt.Fprintf(w, "Product updated: %+v", pi)
    } else {
        fmt.Fprintf(w, "ERROR: Product ID does not exist.")
    }
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
    fmt.Println("DeleteProduct")
    productId := mux.Vars(r)["id"]
    pi2, index := FindProduct(productId)
    if index == -1 {
        fmt.Fprintf(w, "ERROR: Product ID does not exist.")
    } else {
        inventory = append(inventory[:index], inventory[index + 1:]...)
        fmt.Fprintf(w, "Product deleted: %+v", pi2)
    }
}

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/Inventory", GetAllProducts).Methods("GET")
    router.HandleFunc("/Inventory/{id}", GetOneProduct).Methods("GET")
    router.HandleFunc("/CreateProduct", CreateProduct).Methods("POST")
    router.HandleFunc("/UpdateProduct", UpdateProduct).Methods("PUT")
    router.HandleFunc("/DeleteProduct/{id}", DeleteProduct).Methods("DELETE")
    fmt.Println("Listening requests...")
    log.Fatal(http.ListenAndServe(":8080", router))
}

/*
OUTPUT: 

$ go run Exercise2Server.go 
Listening requests...
CreateProduct
CreateProduct
CreateProduct
UpdateProduct
DeleteProduct
CreateProduct
CreateProduct
CreateProduct
UpdateProduct
DeleteProduct

*/
