package service

import (
	"fmt"
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
	achievementsRepository   repository.AchievementRepository
	achievementTagRepository repository.AchievementTagRepository
	tagRepository            repository.TagRepository
	attachmentService        AttachmentService
}

func NewAchievementService(achievementsRepository repository.AchievementRepository,
	achievementTagRepository repository.AchievementTagRepository, tagRepository repository.TagRepository, attachmentService AttachmentService) AchievementService {
	return &achievementService{achievementsRepository, achievementTagRepository,
		tagRepository, attachmentService}
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
	for _, tagId := range achievement.TagsIds {
		tag, err := cs.tagRepository.FindTagById(tagId)
		if err != nil {
			fmt.Println("error =", err.Error())
			return nil, err
		}
		achievementTag := &model.AchievementTag{
			AchievementId: achievement.ID,
			TagId:         tag.ID,
		}
		_, err = cs.achievementTagRepository.CreateAchievementTag(achievementTag)
		if err != nil {
			fmt.Println("error =", err.Error())
		}
	}
	return createdAchievement, nil
}

func (cs *achievementService) GetAchievement(achievementId string) (*model.Achievement, error) {
	achievement, err := cs.achievementsRepository.FindAchievementById(achievementId)
	if err != nil {
		return nil, err
	}
	achievementExternalId := "achievement_" + achievement.ID
	attachment, err := cs.attachmentService.GetAttachmentByExternalId(achievementExternalId)
	if err != nil {
		fmt.Println("error =", err.Error())
		return nil, err
	}
	if attachment != nil {
		achievement.ImageUrl = "/assets/image/" + achievementExternalId
	}
	ats, err := cs.achievementTagRepository.FindTagsIdsByAchievementId(achievementId)
	if err != nil {
		return nil, err
	}
	for _, achievementTag := range ats {
		tag, err := cs.tagRepository.FindTagById(achievementTag.TagId)
		if err != nil {
			fmt.Println("error =", err.Error())
		}
		achievement.Tags = append(achievement.Tags, *tag)
	}
	return achievement, nil
}

func (cs *achievementService) GetAchievements() ([]*model.Achievement, error) {
	achievements, err := cs.achievementsRepository.FindAchievements()
	if err != nil {
		fmt.Println("error =", err.Error())
		return nil, err
	}
	for _, achievement := range achievements {
		achievementExternalId := "achievement_" + achievement.ID
		attachment, err := cs.attachmentService.GetAttachmentByExternalId(achievementExternalId)
		if err != nil {
			fmt.Println("error =", err.Error())
			return nil, err
		}
		if attachment != nil {
			achievement.ImageUrl = "/assets/image/" + achievementExternalId
		}
	}
	return achievements, nil
}
