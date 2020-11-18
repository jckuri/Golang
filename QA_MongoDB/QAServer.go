package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
    "context"
    "errors"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type QA struct {
    Id string `bson:"id"`
    Question string `bson:"question"`
    QUser string `bson:"quser"`
    Answer string `bson:"answer"`
    AUser string `bson:"auser"`
}

var questionId int = 0

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    collection = client.Database("qa").Collection("qas")
}

func MongoExperiment() {
    qa1 := QA {"1", "Where's Waldo?", "jckuri", "", ""}
    createQA(&qa1)
    qas, err := getAll()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("All questions:")
    for _, qa := range qas {
        fmt.Printf("%+v\n", qa)
    }  
    qa1 = QA {"1", "Where's Waldo?", "jckuri", "Here.", "ccedano"}
    updateQA(&qa1)    
    qas, err = getAll()
    if err != nil {
        fmt.Println(err)
    }
    for _, qa := range qas {
        fmt.Printf("%+v\n", qa)
        err = deleteQA(qa.Id)
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Printf("Question %v was deleted.\n", qa.Id)
        }
        
    }    
}

func getAll() ([]*QA, error) {
    filter := bson.D{{}}
    return filterQAs(filter)
}

func getQA(id string) (*QA, error) {
    filter := bson.D{primitive.E{Key: "id", Value: id}}
    qas, err := filterQAs(filter)
    if err != nil {
        return nil, err
    } else {
        if len(qas) == 0 {
            return nil, fmt.Errorf("Question id %v not found.", id)
        } else {
            return qas[0], nil
        }
    }
}

func getQuestionsOfUser(quser string) ([]*QA, error) {
    filter := bson.D{primitive.E{Key: "quser", Value: quser}}
    return filterQAs(filter)
}

func getAnswersOfUser(auser string) ([]*QA, error) {
    filter := bson.D{primitive.E{Key: "auser", Value: auser}}
    return filterQAs(filter)
}


func filterQAs(filter interface{}) ([]*QA, error) {
    var qas []*QA
    cur, err := collection.Find(ctx, filter)
    if err != nil {
        return qas, err
    }
    for cur.Next(ctx) {
        var qa QA
        err := cur.Decode(&qa)
        if err != nil {
            return qas, err
        }
        qas = append(qas, &qa)
    }
    if err := cur.Err(); err != nil {
        return qas, err
    }
    cur.Close(ctx)
    if len(qas) == 0 {
        return qas, mongo.ErrNoDocuments
    }
    return qas, nil
}

func createQA(qa *QA) error {
    _, err := collection.InsertOne(ctx, qa)
    return err
}

func updateQA(qa *QA) error {
    filter := bson.D{primitive.E{Key: "id", Value: qa.Id}}
    update := bson.D{primitive.E{Key: "$set", Value: bson.D{
        primitive.E{Key: "question", Value: qa.Question},
        primitive.E{Key: "quser", Value: qa.QUser},
        primitive.E{Key: "answer", Value: qa.Answer},
        primitive.E{Key: "auser", Value: qa.AUser},
    }}}
    return collection.FindOneAndUpdate(ctx, filter, update).Decode(qa)
}

func deleteQA(Id string) error {
    filter := bson.D{primitive.E{Key: "id", Value: Id}}
    res, err := collection.DeleteOne(ctx, filter)
    if err != nil {
        return err
    }
    if res.DeletedCount == 0 {
        return errors.New("No questions were deleted")
    }
    return nil
}



func QAsToString(qas []*QA) string {
    s := ""
    for _, qa := range qas {
        s += fmt.Sprintf("%+v ", qa)
    }
    return s
}

func ReadQuestion(w http.ResponseWriter, r *http.Request) {
    questionId := mux.Vars(r)["id"]
    qa, err := getQA(questionId)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "ERROR: Question ID does not exist.")
    } else {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "%+v", qa)
    }
}

func ReadAllQuestions(w http.ResponseWriter, r *http.Request) {
    qas, err := getAll()
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "ERROR: Questions could not be read.")
    } else {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "%+v", QAsToString(qas))
    }
}

func ReadQuestionsOfUser(w http.ResponseWriter, r *http.Request) {
    quser := mux.Vars(r)["user"]
    qas, err := getQuestionsOfUser(quser)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "ERROR: No questions of user found.")
    } else {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "%+v", QAsToString(qas))
    }
}

func ReadAnswersOfUser(w http.ResponseWriter, r *http.Request) {
    auser := mux.Vars(r)["user"]
    qas, err := getQuestionsOfUser(auser)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "ERROR: No answers of user found.")
    } else {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "%+v", QAsToString(qas))
    }
}

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
    reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "ERROR: Body of request must contain a question description.")
        return
    }
    var question QA
    json.Unmarshal(reqBody, &question)
    questionId ++
    qa := QA {Id: strconv.Itoa(questionId), Question: question.Question, QUser: question.QUser, Answer: "", AUser: ""}
    err = createQA(&qa)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "ERROR: Question could not be created.")
    } else {
        w.WriteHeader(http.StatusCreated)
        fmt.Fprintf(w, "%+v", qa)    
    }
}

func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
    reqBody, err := ioutil.ReadAll(r.Body)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "ERROR: Body of request must contain a question description.")
        return
    }
    var question QA
    json.Unmarshal(reqBody, &question)
    questionId := mux.Vars(r)["id"]
    _, err = getQA(questionId)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "ERROR: Question ID does not exist.")
    } else {
        qa := QA {Id: questionId, Question: question.Question, QUser: question.QUser, Answer: question.Answer, AUser: question.AUser}
        err = updateQA(&qa)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            fmt.Fprintf(w, "ERROR: Question could not be updated.")
        } else {
            w.WriteHeader(http.StatusOK)
            fmt.Fprintf(w, "%+v", qa)    
        }
    }
}

func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
    questionId := mux.Vars(r)["id"]
    qa, err := getQA(questionId)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "ERROR: Question ID does not exist.")
    } else {
        err = deleteQA(questionId)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            fmt.Fprintf(w, "ERROR: Question could not be deleted.")
        } else {
            w.WriteHeader(http.StatusOK)
            fmt.Fprintf(w, "%+v", qa)
        }
    }
}

func main() {
    //MongoExperiment()
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
