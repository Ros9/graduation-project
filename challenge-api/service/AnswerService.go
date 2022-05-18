package service

import (
	"fmt"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/repository"

	"github.com/google/uuid"
)

type AnswerService interface {
	CreateAnswer(answer *model.Answer) (*model.Answer, error)
	GetAnswer(answerID string) (*model.Answer, error)
	GetAnswers() ([]*model.Answer, error)
	PostAnswerFromTelegram(answer *model.Answer) (*model.Result, error)
}

type answerService struct {
	answerRepository    repository.AnswerRepository
	challengeRepository repository.ChallengeRepository
}

func NewAnswerService(answerRepository repository.AnswerRepository, challengeRepository repository.ChallengeRepository) AnswerService {
	return &answerService{answerRepository, challengeRepository}
}

func (cs *answerService) CreateAnswer(answer *model.Answer) (*model.Answer, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	answer.ID = id.String()

	fmt.Println("\n\nans service", answer)
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

func (cs *answerService) PostAnswerFromTelegram(answer *model.Answer) (*model.Result, error) {
	challenge, err := cs.challengeRepository.GetActiveChallengeByAnswer(answer.Answer)
	if err != nil {
		answer.Status = 0
		cs.CreateAnswer(answer)
		return nil, err
	}

	result := model.Result{Challenge: model.ChallengeTelegramResponse(*challenge)}
	if challenge != (&model.Challenge{}) {
		result.Status = 1
	}

	answer.ChallengeID = challenge.ID
	answer.Status = result.Status

	cs.CreateAnswer(answer)
	//todo Здесь еще надо сделать создание записи в таблице answers

	return &result, nil
}
