package service

import (
	"github.com/google/uuid"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/repository"
)

type AnswerService interface {
	CreateAnswer(answer *model.Answer) (*model.Answer, error)
	GetAnswer(answerID string) (*model.Answer, error)
	GetAnswers() ([]*model.Answer, error)
}

type answerService struct {
	answerRepository repository.AnswerRepository
}

func NewAnswerService(answerRepository repository.AnswerRepository) AnswerService {
	return &answerService{answerRepository}
}

func (cs *answerService) CreateAnswer(answer *model.Answer) (*model.Answer, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	answer.ID = id.String()
	createdAnswer, err := cs.answerRepository.CreateAnswer(answer)
	if err != nil {
		return nil, err
	}
	return createdAnswer, nil
}

func (cs *answerService) GetAnswer(answerId string) (*model.Answer, error) {
	answer, err := cs.answerRepository.FindAnswerById(answerId)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

func (cs *answerService) GetAnswers() ([]*model.Answer, error) {
	return []*model.Answer{}, nil
}
