package qa

import (
    "context"
    "github.com/go-kit/kit/endpoint"
    "qa/pkg"
)

type Endpoints struct {
    ReadQuestionEndpoint endpoint.Endpoint
    ReadAllQuestionsEndpoint endpoint.Endpoint
    CreateQuestionEndpoint endpoint.Endpoint
    UpdateQuestionEndpoint endpoint.Endpoint
    DeleteQuestionEndpoint endpoint.Endpoint
    DeleteAllQuestionsEndpoint endpoint.Endpoint
    ReadQuestionsOfUserEndpoint endpoint.Endpoint
    ReadAnswersOfUserEndpoint endpoint.Endpoint
}

func MakeReadQuestionEndpoint(srv Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(IDRequest)
        qa1, err := srv.ReadQuestion(ctx, req.Id)
        if err != nil {
            return ReadQuestionResponse{qa.QA {}, err.Error()}, nil
        }
        return ReadQuestionResponse{qa1, ""}, nil
    }
}

func MakeReadAllQuestionsEndpoint(srv Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        _ = request.(EmptyRequest)
        qas, err := srv.ReadAllQuestions(ctx)
        if err != nil {
            return ReadQuestionsResponse{nil, err.Error()}, nil
        }
        return ReadQuestionsResponse{qas, ""}, nil
    }
}

func MakeCreateQuestionEndpoint(srv Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(QuestionRequest)
        qa1, err := srv.CreateQuestion(ctx, req.QA)
        if err != nil {
            return CreateQuestionResponse{qa.QA {}, err.Error()}, nil
        }
        return CreateQuestionResponse{qa1, ""}, nil
    }
}

func MakeUpdateQuestionEndpoint(srv Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(QuestionRequest)
        err := srv.UpdateQuestion(ctx, req.QA)
        if err != nil {
            return EmptyResponse {err.Error()}, nil
        } else {
            return EmptyResponse {}, err
        }
    }
}

func MakeDeleteQuestionEndpoint(srv Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(IDRequest)
        err := srv.DeleteQuestion(ctx, req.Id)
        if err != nil {
            return EmptyResponse {err.Error()}, nil
        } else {
            return EmptyResponse {}, err
        }
    }
}

func MakeDeleteAllQuestionsEndpoint(srv Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        _ = request.(EmptyRequest)
        err := srv.DeleteAllQuestions(ctx)
        if err != nil {
            return EmptyResponse {err.Error()}, nil
        } else {
            return EmptyResponse {}, err
        }
    }
}

func MakeReadQuestionsOfUserEndpoint(srv Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(UserRequest)
        qas, err := srv.ReadQuestionsOfUser(ctx, req.User)
        if err != nil {
            return ReadQuestionsResponse{nil, err.Error()}, nil
        }
        return ReadQuestionsResponse{qas, ""}, nil
    }
}

func MakeReadAnswersOfUserEndpoint(srv Service) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(UserRequest)
        qas, err := srv.ReadAnswersOfUser(ctx, req.User)
        if err != nil {
            return ReadQuestionsResponse{nil, err.Error()}, nil
        }
        return ReadQuestionsResponse{qas, ""}, nil
    }
}
