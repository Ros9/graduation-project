package service

import (
	"github.com/google/uuid"
	"graduation-project/challenge-api/model"
)

type CommentService interface {
	CreateComment(comment *model.Comment) (*model.Comment, error)
	GetComment(commentID string) (*model.Comment, error)
	GetComments() ([]*model.Comment, error)
}

type commentService struct {
}

func NewCommentService() CommentService {
	return &commentService{}
}

func (cs *commentService) CreateComment(comment *model.Comment) (*model.Comment, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	comment.ID = id.String()
	return comment, nil
}

func (cs *commentService) GetComment(commentID string) (*model.Comment, error) {
	return &model.Comment{}, nil
}

func (cs *commentService) GetComments() ([]*model.Comment, error) {
	return []*model.Comment{}, nil
}
