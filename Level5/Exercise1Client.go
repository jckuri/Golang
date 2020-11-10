/*
Statement
Create a simple endpoint to just return a message with the method of the Request. The function should handle at least GET and POST

Topics to Practice: 
net/http, RESTful, []byte, log pkg
*/

package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "bytes"
)

var url = "http://127.0.0.1:8080/Exercise1"

func Request(method string, data string) {
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
    fmt.Println(string(body))
}

func main() {
    Request(http.MethodGet, "Hello GET.")
    Request(http.MethodPost, "Hello POST.")
    Request(http.MethodPut, "Hello PUT.")
    Request(http.MethodDelete, "Hello DELETE.")
    Request(http.MethodPatch, "Hello PATCH.")
}

/*
OUTPUT:

go run Exercise1Client.go 
{"message": "GET called"}
{"message": "POST called"}
{"message": "PUT called"}
{"message": "DELETE called"}
{"message": "Not found"}


*/
