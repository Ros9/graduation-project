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
	answerRepository         repository.AnswerRepository
	userChallengesRepository repository.UsersChallengesRepository
	challengeRepository      repository.ChallengeRepository
}

func NewAnswerService(answerRepository repository.AnswerRepository,
	userChallengesRepository repository.UsersChallengesRepository, challengeRepository repository.ChallengeRepository) AnswerService {
	return &answerService{answerRepository, userChallengesRepository, challengeRepository}
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
	_, err = cs.userChallengesRepository.CreateUserChallenge(&model.UserChallenge{
		UserId:      answer.UserID,
		ChallengeId: answer.ChallengeID,
	})
	if err != nil {

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
	fmt.Println("\n\n==== ans in service", answer)

	challenge, err := cs.challengeRepository.GetChallengeByAnswer(answer.Answer)
	if err != nil {
		return nil, err
	}

	result := model.Result{Challenge: model.ChallengeTelegramResponse(*challenge)}
	if challenge != (&model.Challenge{}) {
		result.Status = 1
	}

	//todo Здесь еще надо сделать создание записи в таблице answers

	return &result, nil
}
