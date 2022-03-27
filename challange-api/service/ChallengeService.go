package service

import (
	"github.com/google/uuid"
	"graduation-project/challange-api/model"
)

type ChallengeService interface {
	CreateChallenge(challenge *model.Challenge) (*model.Challenge, error)
	GetChallenge(challengeID string) (*model.Challenge, error)
	GetChallenges() ([]*model.Challenge, error)
}

type challengeService struct {
}

func NewChallengeService() ChallengeService {
	return &challengeService{}
}

func (cs *challengeService) CreateChallenge(challenge *model.Challenge) (*model.Challenge, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	challenge.ID = id.String()
	return challenge, nil
}

func (cs *challengeService) GetChallenge(challengeID string) (*model.Challenge, error) {
	return &model.Challenge{}, nil
}

func (cs *challengeService) GetChallenges() ([]*model.Challenge, error) {
	return []*model.Challenge{}, nil
}
