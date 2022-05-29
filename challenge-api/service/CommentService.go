package service

import (
	"fmt"
	"github.com/google/uuid"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/repository"
)

type CommentService interface {
	CreateComment(comment *model.Comment) (*model.Comment, error)
	GetComment(commentID string) (*model.Comment, error)
	GetComments() ([]*model.Comment, error)
	GetCommentsByChallengeId(challengeId string) ([]*model.Comment, error)
}

type commentService struct {
	commentRepository repository.CommentRepository
	userRepository    repository.UserRepository
}

func NewCommentService(commentRepository repository.CommentRepository, userRepository repository.UserRepository) CommentService {
	return &commentService{commentRepository, userRepository}
}

func (cs *commentService) CreateComment(comment *model.Comment) (*model.Comment, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	comment.ID = id.String()
	createdComment, err := cs.commentRepository.CreateComment(comment)
	if err != nil {
		return nil, err
	}
	return createdComment, nil
}

func (cs *commentService) GetComment(commentID string) (*model.Comment, error) {
	comment, err := cs.commentRepository.FindCommentById(commentID)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (cs *commentService) GetComments() ([]*model.Comment, error) {
	return []*model.Comment{}, nil
}

func (cs *commentService) GetCommentsByChallengeId(challengeId string) ([]*model.Comment, error) {
	comments, err := cs.commentRepository.FindCommentsByChallengeId(challengeId)
	if err != nil {
		fmt.Println("error =", err.Error())
		return nil, err
	}
	for _, comment := range comments {
		user, err := cs.userRepository.FindUserById(comment.UserId)
		if err != nil {
			fmt.Println("error =", err.Error())
			continue
		}
		comment.Username = user.Username
	}
	return comments, nil
}
