package service

import (
	"errors"
	"fmt"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/repository"
	"graduation-project/challenge-api/utils"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUser(userID string) (*model.User, error)
	GetUsers() ([]*model.User, error)
	GetTokenForUser(login, password string) (string, error)
	GetUserByTelegram(userTelegram string) (*model.UserTelegram, error)
}

type userService struct {
	userRepository            repository.UserRepository
	challengeRepository       repository.ChallengeRepository
	attachmentService         AttachmentService
	userAchievementRepository repository.UserAchievementRepository
	achievementRepository     repository.AchievementRepository
}

func NewUserService(userRepository repository.UserRepository, challengeRepository repository.ChallengeRepository,
	attachmentService AttachmentService, userAchievementRepository repository.UserAchievementRepository,
	achievementRepository repository.AchievementRepository) UserService {
	return &userService{
		userRepository,
		challengeRepository,
		attachmentService,
		userAchievementRepository,
		achievementRepository,
	}
}

func (cs *userService) CreateUser(user *model.User) (*model.User, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	user.ID = id.String()
	createdUser, err := cs.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

//TODO
func (cs *userService) GetUser(userID string) (*model.User, error) {
	user, err := cs.userRepository.FindUserById(userID)
	if err != nil {
		return nil, err
	}
	ucs, err := cs.challengeRepository.GetChallengesByUserId(user.ID)
	for _, challenge := range ucs {
		challengeExternalId := "challenge" + challenge.ID
		attachment, err := cs.attachmentService.GetAttachmentByExternalId(challengeExternalId)
		if err != nil {
			fmt.Println("error when get challenge =", err.Error())
		}
		if attachment != nil {
			challenge.ImageUrl = "/assets/image/" + challengeExternalId
		}
	}
	user.Challenges = append(user.Challenges, ucs...)
	uas, err := cs.userAchievementRepository.FindAchievementIdsByUserId(user.ID)
	if err != nil {
		fmt.Println("error =", err.Error())
		return nil, err
	}
	for _, userAchievement := range uas {
		achievement, err := cs.achievementRepository.FindAchievementById(userAchievement.AchievementId)
		if err != nil {
			fmt.Println("error =", err.Error())
			continue
		}
		achievementExternalId := "achievement_" + achievement.ID
		attachment, err := cs.attachmentService.GetAttachmentByExternalId(achievementExternalId)
		if err != nil {
			fmt.Println(err.Error())
		}
		if attachment != nil {
			achievement.ImageUrl = "/assets/image/" + achievementExternalId
		}
		user.Achievements = append(user.Achievements, achievement)
	}
	return user, nil
}

func (cs *userService) GetUserByTelegram(userTelegram string) (*model.UserTelegram, error) {
	fmt.Println("\n\n==== service 1", userTelegram)
	user, err := cs.userRepository.FindUserByTelegram(userTelegram)
	if err != nil {
		return nil, err
	}
	fmt.Println("\n\n==== service 2", user)

	userChallenges, err := cs.challengeRepository.GetChallengesByUserId(user.ID)
	userResponse := model.UserTelegram{
		ID:       user.ID,
		Username: user.Username,
		Password: "",
		Email:    user.Email,
		Telegram: user.Telegram,
		IsAdmin:  user.IsAdmin,
	}
	for _, challenge := range userChallenges {
		ch := model.ChallengeTelegramResponse(*challenge)
		userResponse.Challenges = append(userResponse.Challenges, &ch)
	}
	fmt.Println("\n\n==== service 3", userResponse)

	return &userResponse, err
}

func (cs *userService) GetUsers() ([]*model.User, error) {
	return []*model.User{}, nil
}

func (cs *userService) GetTokenForUser(login, password string) (string, error) {
	user, err := cs.userRepository.FindUserByUsername(login)
	if err != nil {
		return "", err
	}
	if user.Password != password {
		fmt.Println("user.Password =", user.Password)
		fmt.Println("password =", password)
		return "", errors.New("invalid user or password")
	}
	return utils.GetToken(user.ID)
}
