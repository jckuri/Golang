package qa

import (
    "net/http"
    httptransport "github.com/go-kit/kit/transport/http"
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
    "fmt"
    "os"
)

func NewServer(svc Service, port string) *http.Server {
    endpoints := Endpoints{
        ReadQuestionEndpoint: MakeReadQuestionEndpoint(svc),
        ReadAllQuestionsEndpoint: MakeReadAllQuestionsEndpoint(svc),
        CreateQuestionEndpoint: MakeCreateQuestionEndpoint(svc),
        UpdateQuestionEndpoint: MakeUpdateQuestionEndpoint(svc),
        DeleteQuestionEndpoint: MakeDeleteQuestionEndpoint(svc),
        DeleteAllQuestionsEndpoint: MakeDeleteAllQuestionsEndpoint(svc),
        ReadQuestionsOfUserEndpoint: MakeReadQuestionsOfUserEndpoint(svc),
        ReadAnswersOfUserEndpoint: MakeReadAnswersOfUserEndpoint(svc),
    }
    r := makeHandlers(endpoints)
    server := &http.Server{
        Addr: port,
        Handler: handlers.CORS(
            handlers.AllowedHeaders([]string{"Content-Type"}),
            handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
        )(r),
    }
    return server
}

func makeHandlers(endpoints Endpoints) *mux.Router {
    r := mux.NewRouter()
    r.Use(commonMiddleware)
    r.Methods("GET").Path("/read_question/{id}").Handler(httptransport.NewServer(endpoints.ReadQuestionEndpoint, DecodeIDRequest, EncodeResponse))
    r.Methods("GET").Path("/read_all_questions").Handler(httptransport.NewServer(endpoints.ReadAllQuestionsEndpoint, DecodeEmptyRequest, EncodeResponse))
    r.Methods("POST").Path("/create_question").Handler(httptransport.NewServer(endpoints.CreateQuestionEndpoint, DecodeQuestionRequest, EncodeResponse))
    r.Methods("PUT").Path("/update_question").Handler(httptransport.NewServer(endpoints.UpdateQuestionEndpoint, DecodeQuestionRequest, EncodeResponse))
    r.Methods("DELETE").Path("/delete_question/{id}").Handler(httptransport.NewServer(endpoints.DeleteQuestionEndpoint, DecodeIDRequest, EncodeResponse))
    r.Methods("DELETE").Path("/delete_all_questions").Handler(httptransport.NewServer(endpoints.DeleteAllQuestionsEndpoint, DecodeEmptyRequest, EncodeResponse))
    r.Methods("GET").Path("/read_questions_of_user/{user}").Handler(httptransport.NewServer(endpoints.ReadQuestionsOfUserEndpoint, DecodeUserRequest, EncodeResponse))
    r.Methods("GET").Path("/read_answers_of_user/{user}").Handler(httptransport.NewServer(endpoints.ReadAnswersOfUserEndpoint, DecodeUserRequest, EncodeResponse))
    return r
}

func commonMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}

func StartServer() {
    svc := NewService()
    if svc == nil {
        fmt.Println("Failed to create service")
        os.Exit(2)
    }
    errc := make(chan error)
    server := NewServer(svc, ":8080")
    go func() {
        fmt.Printf("HTTP service started listening %v\n", server.Addr)
        errc <- server.ListenAndServe()
    }()
    fmt.Printf("exit %v\n", <-errc)
}
