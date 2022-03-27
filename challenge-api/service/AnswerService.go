package service

import (
	"github.com/google/uuid"
	"graduation-project/challenge-api/model"
)

type AnswerService interface {
	CreateAnswer(answer *model.Answer) (*model.Answer, error)
	GetAnswer(answerID string) (*model.Answer, error)
	GetAnswers() ([]*model.Answer, error)
}

type answerService struct {
}

func NewAnswerService() AnswerService {
	return &answerService{}
}

func (cs *answerService) CreateAnswer(answer *model.Answer) (*model.Answer, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	answer.ID = id.String()
	return answer, nil
}

func (cs *answerService) GetAnswer(answerId string) (*model.Answer, error) {
	return &model.Answer{}, nil
}

func (cs *answerService) GetAnswers() ([]*model.Answer, error) {
	return []*model.Answer{}, nil
}
