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
	answerRepository         repository.AnswerRepository
	userChallengesRepository repository.UsersChallengesRepository
}

func NewAnswerService(answerRepository repository.AnswerRepository,
	userChallengesRepository repository.UsersChallengesRepository) AnswerService {
	return &answerService{answerRepository, userChallengesRepository}
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
