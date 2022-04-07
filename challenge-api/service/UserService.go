package service

import (
	"fmt"
	"github.com/google/uuid"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/repository"
)

type UserService interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUser(userID string) (*model.User, error)
	GetUsers() ([]*model.User, error)
}

type userService struct {
	userRepository            repository.UserRepository
	usersChallengesRepository repository.UsersChallengesRepository
	challengeRepository       repository.ChallengeRepository
}

func NewUserService(userRepository repository.UserRepository,
	usersChallengesRepository repository.UsersChallengesRepository,
	challengeRepository repository.ChallengeRepository) UserService {
	return &userService{
		userRepository,
		usersChallengesRepository,
		challengeRepository,
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

func (cs *userService) GetUser(userID string) (*model.User, error) {
	user, err := cs.userRepository.FindUserById(userID)
	if err != nil {
		return nil, err
	}
	ucs, err := cs.usersChallengesRepository.FindChallengesByUserId(user.ID)
	fmt.Println()
	for _, value := range ucs {
		fmt.Println(value)
	}
	fmt.Println()
	for _, uc := range ucs {
		userChallenge, err := cs.challengeRepository.FindChallengeById(uc.ChallengeId)
		if err != nil {
			fmt.Println("err =", err.Error())
		}
		user.Challenges = append(user.Challenges, userChallenge)
	}
	return user, nil
}

func (cs *userService) GetUsers() ([]*model.User, error) {
	return []*model.User{}, nil
}
