package service

import (
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/repository"

	"github.com/google/uuid"
)

type ChallengeService interface {
	CreateChallenge(challenge *model.Challenge) (*model.Challenge, error)
	GetChallenge(challengeID string) (*model.Challenge, error)
	GetChallenges() ([]*model.Challenge, error)
	GetChallengesByUserId(userId string) ([]*model.Challenge, error)
}

type challengeService struct {
	challengeRepository repository.ChallengeRepository
}

func NewChallengeService(challengeRepository repository.ChallengeRepository) ChallengeService {
	return &challengeService{challengeRepository}
}

func (cs *challengeService) CreateChallenge(challenge *model.Challenge) (*model.Challenge, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	challenge.ID = id.String()
	createdChallenge, err := cs.challengeRepository.CreateChallenge(challenge)
	if err != nil {
		return nil, err
	}
	return createdChallenge, nil
}

func (cs *challengeService) GetChallenge(challengeID string) (*model.Challenge, error) {
	challenge, err := cs.challengeRepository.FindChallengeById(challengeID)
	if err != nil {
		return nil, err
	}
	return challenge, nil
}

func (cs *challengeService) GetChallenges() ([]*model.Challenge, error) {
	return cs.challengeRepository.FindChallenges()
}

func (cs *challengeService) GetChallengesByUserId(userId string) ([]*model.Challenge, error) {
	return cs.challengeRepository.GetChallengesByUserId(userId)
}
