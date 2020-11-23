// https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/
// https://www.digitalocean.com/community/tutorials/how-to-use-go-with-mongodb-using-the-mongodb-go-driver

package qa

import (
    "fmt"
    "log"
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

var QuestionId int = 0

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
    fmt.Println("MongoDB initialized.")
    UpdateNewQuestionId()
    //ShowAllQuestions()
}

func UpdateNewQuestionId() {
    qas, err := GetAllQA()
    if err == nil {
        QuestionId = len(qas)
        fmt.Printf("QuestionId=%v\n", QuestionId)
    } else {
        fmt.Println(err)
    }
}

func ShowAllQuestions() {
    qas, err := GetAllQA()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("All questions:")
    for _, qa := range qas {
        fmt.Printf("%+v\n", qa)
    }
    qa, _ := GetQA("1")
    fmt.Printf("GetQA(1): %+v\n", qa)  
}

func GetAllQA() ([]QA, error) {
    filter := bson.D{{}}
    return FilterQAs(filter)
}

func GetQA(id string) (QA, error) {
    filter := bson.D{primitive.E{Key: "id", Value: id}}
    qas, err := FilterQAs(filter)
    if err != nil {
        return QA {}, err
    } else {
        if len(qas) == 0 {
            return QA {}, fmt.Errorf("Question id %v not found.", id)
        } else {
            return qas[0], nil
        }
    }
}

func GetQuestionsOfUser(quser string) ([]QA, error) {
    filter := bson.D{primitive.E{Key: "quser", Value: quser}}
    return FilterQAs(filter)
}

func GetAnswersOfUser(auser string) ([]QA, error) {
    filter := bson.D{primitive.E{Key: "auser", Value: auser}}
    return FilterQAs(filter)
}


func FilterQAs(filter interface{}) ([]QA, error) {
    var qas []QA
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
        qas = append(qas, qa)
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

func CreateQA(qa QA) (QA, error) {
    QuestionId += 1
    qa.Id = strconv.Itoa(QuestionId)
    _, err := collection.InsertOne(ctx, qa)
    return qa, err
}

func UpdateQA(qa QA) error {
    filter := bson.D{primitive.E{Key: "id", Value: qa.Id}}
    update := bson.D{primitive.E{Key: "$set", Value: bson.D{
        primitive.E{Key: "question", Value: qa.Question},
        primitive.E{Key: "quser", Value: qa.QUser},
        primitive.E{Key: "answer", Value: qa.Answer},
        primitive.E{Key: "auser", Value: qa.AUser},
    }}}
    return collection.FindOneAndUpdate(ctx, filter, update).Decode(qa)
}

func DeleteQA(Id string) error {
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

func DeleteAllQAs() error {
    filter := bson.D{{}}
    res, err := collection.DeleteMany(ctx, filter)
    if err != nil {
        return err
    }
    QuestionId = 0
    if res.DeletedCount == 0 {
        return errors.New("No questions were deleted")
    }
    return nil
}
