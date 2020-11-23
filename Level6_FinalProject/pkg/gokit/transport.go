package qa

import (
	"context"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"qa/pkg"
)

// Encode Response:

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
    return json.NewEncoder(w).Encode(response)
}

// Empty Request:

type EmptyRequest struct {
}

func DecodeEmptyRequest(ctx context.Context, r *http.Request) (interface{}, error) {
    var req EmptyRequest
    return req, nil
}

type EmptyResponse struct {
    Err string `json:"err,omitempty"`
}

// ID Request:

type IDRequest struct {
    Id string `json:"id"`
}

func DecodeIDRequest(ctx context.Context, r *http.Request) (interface{}, error) {
    var req IDRequest
    req.Id = mux.Vars(r)["id"]
    return req, nil
}

// User Request:

type UserRequest struct {
    User string `json:"id"`
}

func DecodeUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
    var req UserRequest
    req.User = mux.Vars(r)["user"]
    return req, nil
}

// QuestionRequest:

type QuestionRequest struct {
    QA qa.QA `json:"qa"`
    Err  string `json:"err,omitempty"`
}

func DecodeQuestionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
    var req QuestionRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        return nil, err
    }
    return req, nil
}

// GetQA:

type ReadQuestionResponse struct {
    QA qa.QA `json:"qa"`
    Err  string `json:"err,omitempty"`
}

// GetAll:

type ReadQuestionsResponse struct {
    QAs []qa.QA `json:"qas"`
    Err  string `json:"err,omitempty"`
}

// CreateQuestion:

type CreateQuestionResponse struct {
    QA qa.QA `json:"qa"`
    Err  string `json:"err,omitempty"`
}
