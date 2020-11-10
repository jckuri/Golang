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
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "bytes"
)

var baseUrl = "http://127.0.0.1:8080"

func Request(path string, method string, data string) string {
    url := fmt.Sprintf("%v%v", baseUrl, path)
    json, err := json.Marshal(data)
    if err != nil {
        panic(err)
    }
    req, err := http.NewRequest(method, url, bytes.NewBuffer(json))
    if err != nil {
        panic(err)
    }
    req.Header.Set("Content-Type", "application/json; charset=utf-8")
    client := &http.Client{}
    resp, err := client.Do(req)
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    return string(body)
}

func JsonRequest(path string, method string, values map[string]string) string {
    jsonValue, err0 := json.Marshal(values)
    if err0 != nil {
        panic(err0)
    }
    url := fmt.Sprintf("%v%v", baseUrl, path)
    req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonValue))
    if err != nil {
        panic(err)
    }
    req.Header.Set("Content-Type", "application/json; charset=utf-8")
    client := &http.Client{}
    resp, err := client.Do(req)
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    return string(body)
}

func ConsultInventory() {
    body := Request("/Inventory", http.MethodGet, "")
    fmt.Println("Inventory:", body)
}

func ConsultOneProduct(id string) {
    body := Request(fmt.Sprintf("/Inventory/%v", id), http.MethodGet, "")
    fmt.Println("Lookup:", body)
}

type ProductInventory map[string]string

func NewProductInventory(ID string, Code string, Name string, Price string, Quantity string) ProductInventory {
    return map[string]string {"ID": ID, "Code": Code, "Name": Name, "Price": Price, "Quantity": Quantity}
}

func CreateProduct(product map[string]string) {
    body := JsonRequest("/CreateProduct", http.MethodPost, product)
    fmt.Println("CreateProduct Result:", body)
}

func UpdateProduct(product map[string]string) {
    body := JsonRequest("/UpdateProduct", http.MethodPut, product)
    fmt.Println("UpdateProduct Result:", body)
}

func DeleteProduct(id string) {
    body := Request(fmt.Sprintf("/DeleteProduct/%v", id), http.MethodDelete, "")
    fmt.Println("DeleteProduct Result:", body)
}

var inventory []ProductInventory = []ProductInventory {
    NewProductInventory("17418", "QRUEIWQ", "T-shirt", "10.5", "2"), 
    NewProductInventory("34781", "QERPUOR", "Shoes", "90.5", "4"),
    NewProductInventory("37282", "WUDJSJS", "Pants", "30.5", "5"),
}

func main() {
    ConsultInventory()
    ConsultOneProduct("34781")
    for _, pi := range inventory {
        CreateProduct(pi)
    }
    pi := inventory[2]
    pi["Name"] = "Pair of Pants"
    UpdateProduct(pi)
    ConsultInventory()
    DeleteProduct("17418")
    ConsultInventory()
}

/*
OUTPUT:

$ go run Exercise2Client.go 
Inventory: []
Lookup: ERROR: Product ID does not exist.
CreateProduct Result: Product created: {Product:{ID:17418 Code:QRUEIWQ Name:T-shirt Price:10.5} Quantity:2}
CreateProduct Result: Product created: {Product:{ID:34781 Code:QERPUOR Name:Shoes Price:90.5} Quantity:4}
CreateProduct Result: Product created: {Product:{ID:37282 Code:WUDJSJS Name:Pants Price:30.5} Quantity:5}
UpdateProduct Result: Product updated: {Product:{ID:37282 Code:WUDJSJS Name:Pair of Pants Price:30.5} Quantity:5}
Inventory: [{Product:{ID:17418 Code:QRUEIWQ Name:T-shirt Price:10.5} Quantity:2} {Product:{ID:34781 Code:QERPUOR Name:Shoes Price:90.5} Quantity:4} {Product:{ID:37282 Code:WUDJSJS Name:Pair of Pants Price:30.5} Quantity:5}]
DeleteProduct Result: Product deleted: {Product:{ID:17418 Code:QRUEIWQ Name:T-shirt Price:10.5} Quantity:2}
Inventory: [{Product:{ID:34781 Code:QERPUOR Name:Shoes Price:90.5} Quantity:4} {Product:{ID:37282 Code:WUDJSJS Name:Pair of Pants Price:30.5} Quantity:5}]

$ go run Exercise2Client.go 
Inventory: [{Product:{ID:34781 Code:QERPUOR Name:Shoes Price:90.5} Quantity:4} {Product:{ID:37282 Code:WUDJSJS Name:Pair of Pants Price:30.5} Quantity:5}]
Lookup: {Product:{ID:34781 Code:QERPUOR Name:Shoes Price:90.5} Quantity:4}
CreateProduct Result: Product created: {Product:{ID:17418 Code:QRUEIWQ Name:T-shirt Price:10.5} Quantity:2}
CreateProduct Result: ERROR: Product ID already exists.
CreateProduct Result: ERROR: Product ID already exists.
UpdateProduct Result: Product updated: {Product:{ID:37282 Code:WUDJSJS Name:Pair of Pants Price:30.5} Quantity:5}
Inventory: [{Product:{ID:34781 Code:QERPUOR Name:Shoes Price:90.5} Quantity:4} {Product:{ID:37282 Code:WUDJSJS Name:Pair of Pants Price:30.5} Quantity:5} {Product:{ID:17418 Code:QRUEIWQ Name:T-shirt Price:10.5} Quantity:2}]
DeleteProduct Result: Product deleted: {Product:{ID:17418 Code:QRUEIWQ Name:T-shirt Price:10.5} Quantity:2}
Inventory: [{Product:{ID:34781 Code:QERPUOR Name:Shoes Price:90.5} Quantity:4} {Product:{ID:37282 Code:WUDJSJS Name:Pair of Pants Price:30.5} Quantity:5}]

*/
