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

func ReadAllQuestions() {
     response := Request("/ReadAllQuestions", http.MethodGet, "")
     fmt.Printf("RESPONSE: %v\n\n", response)
}

func ReadQuestion(id string) {
     path := fmt.Sprintf("/ReadQuestion/%v", id)
     response := Request(path, http.MethodGet, "")
     fmt.Printf("RESPONSE: %v\n\n", response)
}

func ReadQuestionsOfUser(user string) {
     path := fmt.Sprintf("/ReadQuestionsOfUser/%v", user)
     response := Request(path, http.MethodGet, "")
     fmt.Printf("RESPONSE: %v\n\n", response)
}

func ReadAnswersOfUser(user string) {
     path := fmt.Sprintf("/ReadAnswersOfUser/%v", user)
     response := Request(path, http.MethodGet, "")
     fmt.Printf("RESPONSE: %v\n\n", response)
}

func CreateQuestion(question string, quser string) {
     values := make(map[string]string)
     values["Question"] = question
     values["QUser"] = quser
     response := JsonRequest("/CreateQuestion", http.MethodPost, values)
     fmt.Printf("RESPONSE: %v\n\n", response)
}

func UpdateQuestion(id, question, quser, answer, auser string) {
     values := make(map[string]string)
     values["Id"] = id
     values["Question"] = question
     values["QUser"] = quser
     values["Answer"] = answer
     values["AUser"] = auser
     path := fmt.Sprintf("/UpdateQuestion/%v", id)
     response := JsonRequest(path, http.MethodPut, values)
     fmt.Printf("RESPONSE: %v\n\n", response)
}

func DeleteQuestion(id string) {
     path := fmt.Sprintf("/DeleteQuestion/%v", id)
     response := Request(path, http.MethodDelete, "")
     fmt.Printf("RESPONSE: %v\n\n", response)
}

func main() {
    ReadAllQuestions()
    CreateQuestion("Where are we?", "jckuri")
    CreateQuestion("What are we doing?", "ccedano")
    ReadAllQuestions()
    UpdateQuestion("1", "Where are we?", "jckuri", "We are in Latin America.", "ccedano")
    UpdateQuestion("2", "What are we doing?", "ccedano", "We are programming a project.", "jckuri")
    ReadAllQuestions()
    ReadQuestion("2")
    CreateQuestion("Where's Waldo?", "tpeycere")
    UpdateQuestion("3", "Where's Waldo?", "tpeycere", "Here.", "jckuri")
    ReadQuestionsOfUser("ccedano")
    ReadAnswersOfUser("jckuri")
    CreateQuestion("Who are we?", "tpeycere")
    DeleteQuestion("4")
    ReadAllQuestions()
}
