package qa

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "bytes"
)

var baseUrl = "http://127.0.0.1:8080"

func Request(url string, method string, data string) string {
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

func JsonRequest(url string, method string, values map[string]interface{}) string {
    jsonValue, err0 := json.Marshal(values)
    if err0 != nil {
        panic(err0)
    }
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
