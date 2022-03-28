package service

import (
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
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
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
	return &model.User{}, nil
}

func (cs *userService) GetUsers() ([]*model.User, error) {
	return []*model.User{}, nil
}
