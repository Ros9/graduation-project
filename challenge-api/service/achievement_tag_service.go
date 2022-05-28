package service

import (
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/repository"
)

type AchievementTagService interface {
	CreateAchievementTag(achievementTag *model.AchievementTag) (*model.AchievementTag, error)
	GetTagsIdsByAchievementId(achievementId string) ([]*model.AchievementTag, error)
}

type achievementTagService struct {
	achievementTagRepository repository.AchievementTagRepository
}

func NewAchievementTagService(achievementTagRepository repository.AchievementTagRepository) AchievementTagService {
	return &achievementTagService{achievementTagRepository}
}

func (cs *achievementTagService) CreateAchievementTag(achievementTag *model.AchievementTag) (*model.AchievementTag, error) {
	createdAchievementTag, err := cs.achievementTagRepository.CreateAchievementTag(achievementTag)
	if err != nil {
		return nil, err
	}
	return createdAchievementTag, nil
}

func (cs *achievementTagService) GetTagsIdsByAchievementId(achievementId string) ([]*model.AchievementTag, error) {
	ats, err := cs.achievementTagRepository.FindTagsIdsByAchievementId(achievementId)
	if err != nil {
		return nil, err
	}
	return ats, nil
}
