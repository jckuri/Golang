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

type Task struct {
    ID int `json:"id,omitempty"`
    Description string `json:"description,omitempty"`
}

type RawTask struct {
    ID, Description string
}

type Todo struct {
    Tasks []Task
}

var tasks Todo
var taskID int = 0

func FindTask(ID int) int {
    for id, task := range tasks.Tasks {
        if ID == task.ID {
            return id
        }
    }
    return -1
}

func ReadAllTasks(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "%+v", tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
    reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Fprintf(w, "ERROR: Body of request must contain a task description.")
    }
    var rawTask RawTask
    json.Unmarshal(reqBody, &rawTask)
    taskID ++
    task := Task {ID: taskID, Description: rawTask.Description}
    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "%+v", task)    
    tasks.Tasks = append(tasks.Tasks, task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
    reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Fprintf(w, "ERROR: Body of request must contain a task description.")
    }
    var rawTask RawTask
    json.Unmarshal(reqBody, &rawTask)
    taskID := mux.Vars(r)["id"]
    intTaskID, err := strconv.Atoi(taskID)
    if err != nil {
        panic(err)
    }
    taskIDFound := FindTask(intTaskID)
    if taskIDFound == -1 {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "ERROR: Task ID does not exist.")
    } else {
        task := Task {ID: intTaskID, Description: rawTask.Description}
        tasks.Tasks[taskIDFound] = task
        w.WriteHeader(204)
        fmt.Fprintf(w, "%+v", task)    
    }
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
    taskID := mux.Vars(r)["id"]
    intTaskID, err := strconv.Atoi(taskID)
    if err != nil {
        panic(err)
    }
    taskIDFound := FindTask(intTaskID)
    if taskIDFound == -1 {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "ERROR: Task ID does not exist.")
    } else {
        w.WriteHeader(202)
        fmt.Fprintf(w, "%+v", tasks.Tasks[taskIDFound])
        tasks.Tasks = append(tasks.Tasks[:taskIDFound], tasks.Tasks[taskIDFound+1:]...)
    }
}

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/ReadAllTasks", ReadAllTasks).Methods("GET")
    router.HandleFunc("/CreateTask", CreateTask).Methods("POST")
    router.HandleFunc("/UpdateTask/{id}", UpdateTask).Methods("PUT")
    router.HandleFunc("/DeleteTask/{id}", DeleteTask).Methods("DELETE")
    fmt.Println("Listening requests...")
    log.Fatal(http.ListenAndServe(":8080", router))
}
