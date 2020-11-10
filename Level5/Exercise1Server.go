/*
Statement
Create a simple endpoint to just return a message with the method of the Request. The function should handle at least GET and POST

Topics to Practice: 
net/http, RESTful, []byte, log pkg
*/

package main

import (
    "log"
    "net/http"
    "fmt"
    "encoding/json"
)

func DecodeBody(r *http.Request) string {
    dec := json.NewDecoder(r.Body)
    dec.DisallowUnknownFields()
    var s string
    err := dec.Decode(&s)
    if err != nil {
        return fmt.Sprintf("ERROR: %v", err)
    }
    return s
}

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    switch r.Method {
    case "GET":
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"message": "GET called"}`))
        fmt.Printf("GET called. Body: %v\n", DecodeBody(r))
    case "POST":
        w.WriteHeader(http.StatusCreated)
        w.Write([]byte(`{"message": "POST called"}`))
        fmt.Printf("POST called. Body: %v\n", DecodeBody(r))
    case "PUT":
        w.WriteHeader(http.StatusAccepted)
        w.Write([]byte(`{"message": "PUT called"}`))
        fmt.Printf("PUT called. Body: %v\n", DecodeBody(r))
    case "DELETE":
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"message": "DELETE called"}`))
        fmt.Printf("DELETE called. Body: %v\n", DecodeBody(r))
    default:
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "Not found"}`))
        fmt.Printf("UNKNOWN called. Body: %v\n", DecodeBody(r))
    }
}

func main() {
    http.HandleFunc("/Exercise1", Handler)
    fmt.Println("Listening requests...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
$ go run Exercise1Server.go 
Listening requests...
GET called. Body: Hello GET.
POST called. Body: Hello POST.
PUT called. Body: Hello PUT.
DELETE called. Body: Hello DELETE.
UNKNOWN called. Body: Hello PATCH.

*/
