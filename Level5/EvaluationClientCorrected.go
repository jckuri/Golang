package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "net/http"
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
    fmt.Printf("URL: %v\n", url)
    fmt.Printf("Request: %v\n", req)
    req.Header.Set("Content-Type", "application/json; charset=utf-8")
    client := &http.Client{}
    resp, err := client.Do(req)
    fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
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
    fmt.Printf("URL: %v\n", url)
    fmt.Printf("Request: %v\n", req)
    req.Header.Set("Content-Type", "application/json; charset=utf-8")
    client := &http.Client{}
    resp, err := client.Do(req)
    fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    return string(body)
}

type RawTask struct {
    ID, Description string
}

var tasks []RawTask = []RawTask {RawTask{"", "Program Exercises"}, RawTask{"", "Take exam"}, RawTask{"", "Brush teeth"}}

func ReadAllTasks() {
     response := Request("/ReadAllTasks", http.MethodGet, "")
     fmt.Printf("RESPONSE: %v\n\n", response)
}

func ReadTask(id string) {
     path := fmt.Sprintf("/ReadTask/%v", id)
     response := Request(path, http.MethodGet, "")
     fmt.Printf("RESPONSE: %v\n\n", response)
}

func CreateTask(id string, description string) {
     values := make(map[string]string)
     values["Description"] = description
     response := JsonRequest("/CreateTask", http.MethodPost, values)
     fmt.Printf("RESPONSE: %v\n\n", response)
}

func UpdateTask(id string, description string) {
     values := make(map[string]string)
     values["Description"] = description
     path := fmt.Sprintf("/UpdateTask/%v", id)
     response := JsonRequest(path, http.MethodPut, values)
     fmt.Printf("RESPONSE: %v\n\n", response)
}

func DeleteTask(id string) {
     path := fmt.Sprintf("/DeleteTask/%v", id)
     response := Request(path, http.MethodDelete, "")
     fmt.Printf("RESPONSE: %v\n\n", response)
}

func main() {
    ReadAllTasks()
    for _, task := range tasks {
        CreateTask(task.ID, task.Description)
    }
    UpdateTask("3", "Brush MY TEETH")
    UpdateTask("3000000", "Undefined Task")
    ReadAllTasks()
    DeleteTask("3")
    DeleteTask("3000000")
    ReadAllTasks()
    ReadTask("2")
    ReadTask("3000000")
}

/*
OUTPUT:

$ go run EvaluationClientCorrected.go 
URL: http://127.0.0.1:8080/ReadAllTasks
Request: &{GET http://127.0.0.1:8080/ReadAllTasks HTTP/1.1 1 1 map[] {""} 0x64cda0 2 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 200 OK
RESPONSE: {Tasks:[]}

URL: http://127.0.0.1:8080/CreateTask
Request: &{POST http://127.0.0.1:8080/CreateTask HTTP/1.1 1 1 map[] {{"Description":"Program Exercises"}} 0x64cda0 35 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 201 Created
RESPONSE: {ID:1 Description:Program Exercises}

URL: http://127.0.0.1:8080/CreateTask
Request: &{POST http://127.0.0.1:8080/CreateTask HTTP/1.1 1 1 map[] {{"Description":"Take exam"}} 0x64cda0 27 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 201 Created
RESPONSE: {ID:2 Description:Take exam}

URL: http://127.0.0.1:8080/CreateTask
Request: &{POST http://127.0.0.1:8080/CreateTask HTTP/1.1 1 1 map[] {{"Description":"Brush teeth"}} 0x64cda0 29 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 201 Created
RESPONSE: {ID:3 Description:Brush teeth}

URL: http://127.0.0.1:8080/UpdateTask/3
Request: &{PUT http://127.0.0.1:8080/UpdateTask/3 HTTP/1.1 1 1 map[] {{"Description":"Brush MY TEETH"}} 0x64cda0 32 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 200 OK
RESPONSE: {ID:3 Description:Brush MY TEETH}

URL: http://127.0.0.1:8080/UpdateTask/3000000
Request: &{PUT http://127.0.0.1:8080/UpdateTask/3000000 HTTP/1.1 1 1 map[] {{"Description":"Undefined Task"}} 0x64cda0 32 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 400 Bad Request
RESPONSE: ERROR: Task ID does not exist.

URL: http://127.0.0.1:8080/ReadAllTasks
Request: &{GET http://127.0.0.1:8080/ReadAllTasks HTTP/1.1 1 1 map[] {""} 0x64cda0 2 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 200 OK
RESPONSE: {Tasks:[{ID:1 Description:Program Exercises} {ID:2 Description:Take exam} {ID:3 Description:Brush MY TEETH}]}

URL: http://127.0.0.1:8080/DeleteTask/3
Request: &{DELETE http://127.0.0.1:8080/DeleteTask/3 HTTP/1.1 1 1 map[] {""} 0x64cda0 2 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 200 OK
RESPONSE: {ID:3 Description:Brush MY TEETH}

URL: http://127.0.0.1:8080/DeleteTask/3000000
Request: &{DELETE http://127.0.0.1:8080/DeleteTask/3000000 HTTP/1.1 1 1 map[] {""} 0x64cda0 2 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 400 Bad Request
RESPONSE: ERROR: Task ID does not exist.

URL: http://127.0.0.1:8080/ReadAllTasks
Request: &{GET http://127.0.0.1:8080/ReadAllTasks HTTP/1.1 1 1 map[] {""} 0x64cda0 2 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 200 OK
RESPONSE: {Tasks:[{ID:1 Description:Program Exercises} {ID:2 Description:Take exam}]}

URL: http://127.0.0.1:8080/ReadTask/2
Request: &{GET http://127.0.0.1:8080/ReadTask/2 HTTP/1.1 1 1 map[] {""} 0x64cda0 2 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 200 OK
RESPONSE: {ID:2 Description:Take exam}

URL: http://127.0.0.1:8080/ReadTask/3000000
Request: &{GET http://127.0.0.1:8080/ReadTask/3000000 HTTP/1.1 1 1 map[] {""} 0x64cda0 2 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 400 Bad Request
RESPONSE: ERROR: Task ID does not exist.



$ go run EvaluationClientCorrected.go 
URL: http://127.0.0.1:8080/ReadAllTasks
Request: &{GET http://127.0.0.1:8080/ReadAllTasks HTTP/1.1 1 1 map[] {""} 0x64cda0 2 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 200 OK
RESPONSE: {Tasks:[{ID:1 Description:Program Exercises} {ID:2 Description:Take exam}]}

URL: http://127.0.0.1:8080/CreateTask
Request: &{POST http://127.0.0.1:8080/CreateTask HTTP/1.1 1 1 map[] {{"Description":"Program Exercises"}} 0x64cda0 35 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 201 Created
RESPONSE: {ID:4 Description:Program Exercises}

URL: http://127.0.0.1:8080/CreateTask
Request: &{POST http://127.0.0.1:8080/CreateTask HTTP/1.1 1 1 map[] {{"Description":"Take exam"}} 0x64cda0 27 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 201 Created
RESPONSE: {ID:5 Description:Take exam}

URL: http://127.0.0.1:8080/CreateTask
Request: &{POST http://127.0.0.1:8080/CreateTask HTTP/1.1 1 1 map[] {{"Description":"Brush teeth"}} 0x64cda0 29 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 201 Created
RESPONSE: {ID:6 Description:Brush teeth}

URL: http://127.0.0.1:8080/UpdateTask/3
Request: &{PUT http://127.0.0.1:8080/UpdateTask/3 HTTP/1.1 1 1 map[] {{"Description":"Brush MY TEETH"}} 0x64cda0 32 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 400 Bad Request
RESPONSE: ERROR: Task ID does not exist.

URL: http://127.0.0.1:8080/UpdateTask/3000000
Request: &{PUT http://127.0.0.1:8080/UpdateTask/3000000 HTTP/1.1 1 1 map[] {{"Description":"Undefined Task"}} 0x64cda0 32 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 400 Bad Request
RESPONSE: ERROR: Task ID does not exist.

URL: http://127.0.0.1:8080/ReadAllTasks
Request: &{GET http://127.0.0.1:8080/ReadAllTasks HTTP/1.1 1 1 map[] {""} 0x64cda0 2 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 200 OK
RESPONSE: {Tasks:[{ID:1 Description:Program Exercises} {ID:2 Description:Take exam} {ID:4 Description:Program Exercises} {ID:5 Description:Take exam} {ID:6 Description:Brush teeth}]}

URL: http://127.0.0.1:8080/DeleteTask/3
Request: &{DELETE http://127.0.0.1:8080/DeleteTask/3 HTTP/1.1 1 1 map[] {""} 0x64cda0 2 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 400 Bad Request
RESPONSE: ERROR: Task ID does not exist.

URL: http://127.0.0.1:8080/DeleteTask/3000000
Request: &{DELETE http://127.0.0.1:8080/DeleteTask/3000000 HTTP/1.1 1 1 map[] {""} 0x64cda0 2 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 400 Bad Request
RESPONSE: ERROR: Task ID does not exist.

URL: http://127.0.0.1:8080/ReadAllTasks
Request: &{GET http://127.0.0.1:8080/ReadAllTasks HTTP/1.1 1 1 map[] {""} 0x64cda0 2 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 200 OK
RESPONSE: {Tasks:[{ID:1 Description:Program Exercises} {ID:2 Description:Take exam} {ID:4 Description:Program Exercises} {ID:5 Description:Take exam} {ID:6 Description:Brush teeth}]}

URL: http://127.0.0.1:8080/ReadTask/2
Request: &{GET http://127.0.0.1:8080/ReadTask/2 HTTP/1.1 1 1 map[] {""} 0x64cda0 2 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 200 OK
RESPONSE: {ID:2 Description:Take exam}

URL: http://127.0.0.1:8080/ReadTask/3000000
Request: &{GET http://127.0.0.1:8080/ReadTask/3000000 HTTP/1.1 1 1 map[] {""} 0x64cda0 2 [] false 127.0.0.1:8080 map[] map[] <nil> map[]   <nil> <nil> <nil> 0xc0000160f0}
HTTP Response Status: 400 Bad Request
RESPONSE: ERROR: Task ID does not exist.

*/
