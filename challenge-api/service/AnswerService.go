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
	answerRepository          repository.AnswerRepository
	challengeRepository       repository.ChallengeRepository
	challengeTagRepository    repository.ChallengeTagRepository
	achievementRepository     repository.AchievementRepository
	userTagRepository         repository.UserTagRepository
	achievementTagRepository  repository.AchievementTagRepository
	userAchievementRepository repository.UserAchievementRepository
}

func NewAnswerService(answerRepository repository.AnswerRepository, challengeRepository repository.ChallengeRepository,
	challengeTagRepository repository.ChallengeTagRepository, achievementRepository repository.AchievementRepository,
	userTagRepository repository.UserTagRepository, achievementTagRepository repository.AchievementTagRepository,
	userAchievementRepository repository.UserAchievementRepository) AnswerService {
	return &answerService{
		answerRepository,
		challengeRepository,
		challengeTagRepository,
		achievementRepository,
		userTagRepository,
		achievementTagRepository,
		userAchievementRepository,
	}
}

func (cs *answerService) CreateAnswer(answer *model.Answer) (*model.Answer, error) {
	challenge, err := cs.challengeRepository.FindChallengeById(answer.ChallengeID)
	if err != nil {
		fmt.Println("error when CreateChallenge =", err.Error())
		return nil, err
	}
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	answer.ID = id.String()
	createdAnswer, err := cs.answerRepository.CreateAnswer(answer)
	if err != nil {
		return nil, err
	}
	if answer.Answer == challenge.AnswerCode {
		cts, err := cs.challengeTagRepository.FindTagsIdsByChallengeId(challenge.ID)
		if err != nil {
			fmt.Println("error =", err.Error())
			return nil, err
		}
		for _, challengeTag := range cts {
			userTag := &model.UserTag{
				UserId: answer.UserID,
				TagId:  challengeTag.TagId,
			}
			_, err := cs.userTagRepository.CreateUserTag(userTag)
			if err != nil {
				fmt.Println("error =", err.Error())
				return nil, err
			}
		}
		uts, err := cs.userTagRepository.FindTagsIdsByUserId(answer.UserID)
		if err != nil {
			fmt.Println("error =", err.Error())
			return nil, err
		}
		achievements, err := cs.achievementRepository.FindAchievements()
		if err != nil {
			fmt.Println("error =", err.Error())
			return nil, err
		}
		for _, achievement := range achievements {
			ats, err := cs.achievementTagRepository.FindTagsIdsByAchievementId(achievement.ID)
			if err != nil {
				fmt.Println("error =", err.Error())
				return nil, err
			}
			cnt := 0
			for _, achievementTag := range ats {
				for _, userTag := range uts {
					if achievementTag.TagId == userTag.TagId {
						cnt++
					}
				}
			}
			if cnt == len(ats) {
				userAchievement := &model.UserAchievement{
					UserId:        answer.UserID,
					AchievementId: achievement.ID,
				}
				_, err := cs.userAchievementRepository.CreateUserAchievement(userAchievement)
				if err != nil {
					fmt.Println("error =", err.Error())
					return nil, err
				}
			}
		}
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
