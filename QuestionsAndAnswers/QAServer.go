/*
Golang Development Program - Level 6 Final Project

Statement
You are to design the backend side of a system for the following business idea.
We want to build a site called QuestionsAndAnswers.com that will compete with Quora/Stackoverflow and others with 1 major difference. We will only allow 1 answer per question. If someone thinks they have a better answer, they will have to update the existing answer for that question instead of adding another answer. In essence, each question can only have 0 or 1 answer.
The backend should support the following operations:
Get one question by its ID
Get a list of all questions
Get all the questions created by a given user
Create a new question
Update an existing question (the statement and/or the answer)
Delete an existing question
No user tracking or security needed for this version. 
Database design is up to you.
We would like to receive code that runs, so remember to focus on the MVP functionality. You can document what’s missing that you wish you had more time for? Please think about the different problems you might encounter if the business idea is successful. This would include considerations such as increased load, increased data, and an upvoting feature.
*/

package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
)

type QA struct {
    Id, Question, QUser, Answer, AUser string
}

var questionId int = 0
var qas []QA

func FindQuestion(id string) int {
    for i, question := range qas {
        if id == question.Id {
            return i
        }
    }    
    return -1
}

func ReadQuestion(w http.ResponseWriter, r *http.Request) {
    questionId := mux.Vars(r)["id"]
    questionIndex := FindQuestion(questionId)
    if questionIndex == -1 {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "ERROR: Question ID does not exist.")
    } else {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "%+v", qas[questionIndex])
    }  
}

func ReadAllQuestions(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "%+v", qas)
}

func ReadQuestionsOfUser(w http.ResponseWriter, r *http.Request) {
    user := mux.Vars(r)["user"]
    w.WriteHeader(http.StatusOK)
    questions := make([]QA, 0)
    for _, qa := range qas {
        if user == qa.QUser {
            questions = append(questions, qa)
        }
    }
    fmt.Fprintf(w, "%+v", questions)
}

func ReadAnswersOfUser(w http.ResponseWriter, r *http.Request) {
    user := mux.Vars(r)["user"]
    w.WriteHeader(http.StatusOK)
    answers := make([]QA, 0)
    for _, qa := range qas {
        if user == qa.AUser {
            answers = append(answers, qa)
        }
    }
    fmt.Fprintf(w, "%+v", answers)
}

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
    reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Fprintf(w, "ERROR: Body of request must contain a question description.")
    }
    var question QA
    json.Unmarshal(reqBody, &question)
    questionId ++
    qa := QA {Id: strconv.Itoa(questionId), Question: question.Question, QUser: question.QUser, Answer: "", AUser: ""}
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "%+v", qa)    
    qas = append(qas, qa)
}

func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
    reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Fprintf(w, "ERROR: Body of request must contain a question description.")
    }
    var question QA
    json.Unmarshal(reqBody, &question)
    questionId := mux.Vars(r)["id"]
    questionIndex := FindQuestion(questionId)
    if questionIndex == -1 {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "ERROR: Question ID does not exist.")
    } else {
        qa := QA {Id: questionId, Question: question.Question, QUser: question.QUser, Answer: question.Answer, AUser: question.AUser}
        qas[questionIndex] = qa
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "%+v", qa)    
    }
}

func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
    questionId := mux.Vars(r)["id"]
    questionIndex := FindQuestion(questionId)
    if questionIndex == -1 {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "ERROR: Question ID does not exist.")
    } else {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "%+v", qas[questionIndex])
        qas = append(qas[:questionIndex], qas[questionIndex+1:]...)
    }
}

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/ReadQuestion/{id}", ReadQuestion).Methods("GET")
    router.HandleFunc("/ReadAllQuestions", ReadAllQuestions).Methods("GET")
    router.HandleFunc("/ReadQuestionsOfUser/{user}", ReadQuestionsOfUser).Methods("GET")
    router.HandleFunc("/ReadAnswersOfUser/{user}", ReadAnswersOfUser).Methods("GET")
    router.HandleFunc("/CreateQuestion", CreateQuestion).Methods("POST")
    router.HandleFunc("/UpdateQuestion/{id}", UpdateQuestion).Methods("PUT")
    router.HandleFunc("/DeleteQuestion/{id}", DeleteQuestion).Methods("DELETE")
    fmt.Println("Listening requests...")
    log.Fatal(http.ListenAndServe(":8080", router))
}
