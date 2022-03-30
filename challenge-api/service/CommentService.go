package service

import (
	"github.com/google/uuid"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/repository"
)

type CommentService interface {
	CreateComment(comment *model.Comment) (*model.Comment, error)
	GetComment(commentID string) (*model.Comment, error)
	GetComments() ([]*model.Comment, error)
}

type commentService struct {
	commentRepository repository.CommentRepository
}

func NewCommentService(commentRepository repository.CommentRepository) CommentService {
	return &commentService{commentRepository}
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
