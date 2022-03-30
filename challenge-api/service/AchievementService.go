package service

import (
	"github.com/google/uuid"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/repository"
)

type AchievementService interface {
	CreateAchievement(userAnswer *model.Achievement) (*model.Achievement, error)
	GetAchievement(achievementID string) (*model.Achievement, error)
	GetAchievements() ([]*model.Achievement, error)
}

type achievementService struct {
	achievementsRepository repository.AchievementRepository
}

func NewAchievementService(achievementsRepository repository.AchievementRepository) AchievementService {
	return &achievementService{achievementsRepository}
}

func (cs *achievementService) CreateAchievement(achievement *model.Achievement) (*model.Achievement, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	achievement.ID = id.String()
	createdAchievement, err := cs.achievementsRepository.CreateAchievement(achievement)
	if err != nil {
		return nil, err
	}
	return createdAchievement, nil
}

func (cs *achievementService) GetAchievement(achievementId string) (*model.Achievement, error) {
	achievement, err := cs.achievementsRepository.FindAchievementById(achievementId)
	if err != nil {
		return nil, err
	}
	return achievement, nil
}

func (cs *achievementService) GetAchievements() ([]*model.Achievement, error) {
	return []*model.Achievement{}, nil
}
