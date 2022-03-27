package service

import (
	"github.com/google/uuid"
	"graduation-project/challenge-api/model"
)

type AchievementService interface {
	CreateAchievement(userAnswer *model.Achievement) (*model.Achievement, error)
	GetAchievement(achievementID string) (*model.Achievement, error)
	GetAchievements() ([]*model.Achievement, error)
}

type achievementService struct {
}

func NewAchievementService() AchievementService {
	return &achievementService{}
}

func (cs *achievementService) CreateAchievement(achievement *model.Achievement) (*model.Achievement, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	achievement.ID = id.String()
	return achievement, nil
}

func (cs *achievementService) GetAchievement(achievementID string) (*model.Achievement, error) {
	return &model.Achievement{}, nil
}

func (cs *achievementService) GetAchievements() ([]*model.Achievement, error) {
	return []*model.Achievement{}, nil
}
