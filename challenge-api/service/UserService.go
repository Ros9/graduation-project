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
	userRepository      repository.UserRepository
	challengeRepository repository.ChallengeRepository
	attachmentService   AttachmentService
}

func NewUserService(userRepository repository.UserRepository,
	challengeRepository repository.ChallengeRepository,
	attachmentService AttachmentService) UserService {
	return &userService{
		userRepository,
		challengeRepository,
		attachmentService,
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
		fmt.Println("attachment = ", *attachment)
	}
	user.Challenges = append(user.Challenges, ucs...)
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
