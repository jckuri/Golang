package qa

import (
    "context"
    "qa/pkg"
)

type Service interface {
    ReadQuestion(ctx context.Context, id string) (qa.QA, error)
    ReadAllQuestions(ctx context.Context) ([]qa.QA, error)
    CreateQuestion(ctx context.Context, qa1 qa.QA) (qa.QA, error)
    UpdateQuestion(ctx context.Context, qa1 qa.QA) error
    DeleteQuestion(ctx context.Context, id string) error
    DeleteAllQuestions(ctx context.Context) error
    ReadQuestionsOfUser(ctx context.Context, quser string) ([]qa.QA, error)
    ReadAnswersOfUser(ctx context.Context, auser string) ([]qa.QA, error)
}

type QAService struct {}

func NewService() Service {
    return QAService {}
}

func (QAService) ReadQuestion(ctx context.Context, id string) (qa.QA, error) {
    qa1, err := qa.GetQA(id)
    if err != nil {
        return qa.QA {}, err
    } else {
        return qa1, nil
    }
}

func (QAService) ReadAllQuestions(ctx context.Context) ([]qa.QA, error) {
    qas, err := qa.GetAllQA()
    if err != nil {
        return nil, err
    } else {
        return qas, nil
    }
}

func (QAService) CreateQuestion(ctx context.Context, qa1 qa.QA) (qa.QA, error) {
    qa2, err := qa.CreateQA(qa1)
    if err != nil {
        return qa2, err
    } else {
        return qa2, nil
    }
}

func (QAService) UpdateQuestion(ctx context.Context, qa1 qa.QA) error {
    return qa.UpdateQA(qa1)
}

func (QAService) DeleteQuestion(ctx context.Context, id string) error {
    return qa.DeleteQA(id)
}

func (QAService) DeleteAllQuestions(ctx context.Context) error {
    return qa.DeleteAllQAs()
}

func (QAService) ReadQuestionsOfUser(ctx context.Context, quser string) ([]qa.QA, error) {
    qas, err := qa.GetQuestionsOfUser(quser)
    if err != nil {
        return nil, err
    } else {
        return qas, nil
    }
}

func (QAService) ReadAnswersOfUser(ctx context.Context, quser string) ([]qa.QA, error) {
    qas, err := qa.GetAnswersOfUser(quser)
    if err != nil {
        return nil, err
    } else {
        return qas, nil
    }
}
