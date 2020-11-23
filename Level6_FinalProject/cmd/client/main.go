package main

import (
    "fmt"
    "net/http"
    "qa/pkg"
)

var baseUrl = "http://127.0.0.1:8080"

func GetUrl(path string) string {
    return fmt.Sprintf("%v%v", baseUrl, path)
}

func ReadQuestion(id string) {
    path := fmt.Sprintf("/read_question/%v", id)
    response := qa.Request(GetUrl(path), http.MethodGet, "")
    fmt.Printf("RESPONSE: %v\n\n", response)
}

func ReadAllQuestions() {
    response := qa.Request(GetUrl("/read_all_questions"), http.MethodGet, "")
    fmt.Printf("RESPONSE: %v\n\n", response)
}

func CreateQuestion(qa1 qa.QA) {
    values := make(map[string]interface{})
    values["qa"] = qa1
    response := qa.JsonRequest(GetUrl("/create_question"), http.MethodPost, values)
    fmt.Printf("RESPONSE: %v\n\n", response)
}

func UpdateQuestion(qa1 qa.QA) {
    values := make(map[string]interface{})
    values["qa"] = qa1
    response := qa.JsonRequest(GetUrl("/update_question"), http.MethodPut, values)
    fmt.Printf("RESPONSE: %v\n\n", response)
}

func DeleteQuestion(id string) {
    path := fmt.Sprintf("/delete_question/%v", id)
    response := qa.Request(GetUrl(path), http.MethodDelete, "")
    fmt.Printf("RESPONSE: %v\n\n", response)
}

func DeleteAllQuestions() {
    response := qa.Request(GetUrl("/delete_all_questions"), http.MethodDelete, "")
    fmt.Printf("RESPONSE: %v\n\n", response)
}

func ReadQuestionsOfUser(user string) {
    path := fmt.Sprintf("/read_questions_of_user/%v", user)
    response := qa.Request(GetUrl(path), http.MethodGet, "")
    fmt.Printf("RESPONSE: %v\n\n", response)
}

func ReadAnswersOfUser(user string) {
    path := fmt.Sprintf("/read_answers_of_user/%v", user)
    response := qa.Request(GetUrl(path), http.MethodGet, "")
    fmt.Printf("RESPONSE: %v\n\n", response)
}

func main() {
    DeleteAllQuestions()
    ReadAllQuestions()
    CreateQuestion(qa.QA {"", "Where are we?", "jckuri", "", ""})
    CreateQuestion(qa.QA {"", "What are we doing?", "ccedano", "", ""})
    ReadAllQuestions()
    UpdateQuestion(qa.QA {"1", "Where are we?", "jckuri", "We are in Latin America.", "ccedano"})
    UpdateQuestion(qa.QA {"2", "What are we doing?", "ccedano", "We are programming a project.", "jckuri"})
    ReadAllQuestions()
    ReadQuestion("2")
    CreateQuestion(qa.QA {"", "Where's Waldo?", "tpeycere", "", ""})
    UpdateQuestion(qa.QA {"3", "Where's Waldo?", "tpeycere", "Here.", "jckuri"})
    ReadQuestionsOfUser("ccedano")
    ReadAnswersOfUser("jckuri")
    CreateQuestion(qa.QA {"", "Who are we?", "tpeycere", "", ""})
    DeleteQuestion("4")
    ReadAllQuestions()
}
