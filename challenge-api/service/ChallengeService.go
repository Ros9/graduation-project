package service

import (
	"fmt"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/repository"

	"github.com/google/uuid"
)

type ChallengeService interface {
	CreateChallenge(challenge *model.Challenge) (*model.Challenge, error)
	GetChallenge(challengeID string) (*model.Challenge, error)
	GetChallenges() ([]*model.Challenge, error)
	GetChallengesTgResp() ([]*model.ChallengeTelegramResponse, error)
	GetChallengesByUserId(userId string) ([]*model.Challenge, error)
	GetChallengesTgRespByUserId(userId string) ([]*model.ChallengeTelegramResponse, error)
}

type challengeService struct {
	challengeRepository    repository.ChallengeRepository
	attachmentService      AttachmentService
	challengeTagRepository repository.ChallengeTagRepository
	tagRepository          repository.TagRepository
}

func NewChallengeService(challengeRepository repository.ChallengeRepository,
	attachmentService AttachmentService, challengeTagRepository repository.ChallengeTagRepository,
	tagRepository repository.TagRepository) ChallengeService {
	return &challengeService{
		challengeRepository,
		attachmentService,
		challengeTagRepository,
		tagRepository,
	}
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
	for _, tagId := range challenge.TagsIds {
		tag, err := cs.tagRepository.FindTagById(tagId)
		if err != nil {
			fmt.Println("error here =", err.Error())
			continue
		}
		challengeTag := &model.ChallengeTag{
			ChallengeId: challenge.ID,
			TagId:       tag.ID,
		}
		_, err = cs.challengeTagRepository.CreateChallengeTag(challengeTag)
		if err != nil {
			fmt.Println("error =", err.Error())
		}
	}
	return createdChallenge, nil
}

func (cs *challengeService) GetChallenge(challengeID string) (*model.Challenge, error) {
	challenge, err := cs.challengeRepository.FindChallengeById(challengeID)
	if err != nil {
		return nil, err
	}
	challengeExternalId := "challenge_" + challenge.ID
	attachment, err := cs.attachmentService.GetAttachmentByExternalId(challengeExternalId)
	if err != nil {
		fmt.Println("error when get challenge =", err.Error())
	}
	if attachment != nil {
		challenge.ImageUrl = "/assets/image/" + challengeExternalId
	}
	return challenge, nil
}

func (cs *challengeService) GetChallenges() ([]*model.Challenge, error) {
	challenges, err := cs.challengeRepository.FindChallenges()
	if err != nil {
		return nil, err
	}
	for _, challenge := range challenges {
		challengeExternalId := "challenge_" + challenge.ID
		attachment, err := cs.attachmentService.GetAttachmentByExternalId(challengeExternalId)
		if err != nil {
			fmt.Println("error when get challenge =", err.Error())
		}
		if attachment != nil {
			challenge.ImageUrl = "/assets/image/" + challengeExternalId
		}
	}
	return challenges, nil
}

func (cs *challengeService) GetChallengesTgResp() ([]*model.ChallengeTelegramResponse, error) {
	challenges, err := cs.challengeRepository.FindActiveChallenges()
	if err != nil {
		return nil, err
	}
	var challengesTgResponses []*model.ChallengeTelegramResponse
	for _, ch := range challenges {
		challengesResp := model.ChallengeTelegramResponse(*ch)
		challengesTgResponses = append(challengesTgResponses, &challengesResp)
	}
	return challengesTgResponses, nil
}

func (cs *challengeService) GetChallengesByUserId(userId string) ([]*model.Challenge, error) {
	challenges, err := cs.challengeRepository.GetChallengesByUserId(userId)
	if err != nil {
		return nil, err
	}
	for _, challenge := range challenges {
		challengeExternalId := "challenge_" + challenge.ID
		attachment, err := cs.attachmentService.GetAttachmentByExternalId(challengeExternalId)
		if err != nil {
			fmt.Println("error when get challenge =", err.Error())
		}
		if attachment != nil {
			challenge.ImageUrl = "/assets/image/" + challengeExternalId
		}
	}
	return challenges, nil
}

func (cs *challengeService) GetChallengesTgRespByUserId(userId string) ([]*model.ChallengeTelegramResponse, error) {
	challenges, err := cs.challengeRepository.GetChallengesByUserId(userId)
	if err != nil {
		return nil, err
	}
	var challengesTgResponses []*model.ChallengeTelegramResponse
	for _, ch := range challenges {
		challengesResp := model.ChallengeTelegramResponse(*ch)
		challengesTgResponses = append(challengesTgResponses, &challengesResp)
	}
	return challengesTgResponses, nil
}
