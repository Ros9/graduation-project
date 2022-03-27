package service

import (
	"github.com/google/uuid"
	"graduation-project/challange-api/model"
)

type UserService interface {
	CreateUser(answer *model.User) (*model.User, error)
	GetUser(answerID string) (*model.User, error)
	GetUsers() ([]*model.User, error)
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (cs *userService) CreateUser(userAnswer *model.User) (*model.User, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	userAnswer.ID = id.String()
	return userAnswer, nil
}

func (cs *userService) GetUser(userID string) (*model.User, error) {
	return &model.User{}, nil
}

func (cs *userService) GetUsers() ([]*model.User, error) {
	return []*model.User{}, nil
}
