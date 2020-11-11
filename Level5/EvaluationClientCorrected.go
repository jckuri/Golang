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

var tasks []RawTask = []RawTask {RawTask{"", "Program Exercises"}, RawTask{"", "Take exam"}, RawTask{"", "Brush teeth"}, RawTask{"", "Brush teeth at afternoon"}, RawTask{"", "Brush teeth at night"}}

func ReadAllTasks() {
     response := Request("/ReadAllTasks", http.MethodGet, "")
     fmt.Printf("RESPONSE: %v\n\n", response)
}

func ReadTask(id string) {
     path := fmt.Sprintf("/ReadTask/%v", id)
     response := Request(path, http.MethodGet, "")
     fmt.Printf("RESPONSE: %v\n\n", response)
}

func SearchTasks(substring string) {
     path := fmt.Sprintf("/SearchTasks/%v", substring)
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
    SearchTasks("Brush")
    DeleteTask("3")
    DeleteTask("3000000")
    ReadAllTasks()
    ReadTask("2")
    ReadTask("3000000")
    SearchTasks("Brush")
}
