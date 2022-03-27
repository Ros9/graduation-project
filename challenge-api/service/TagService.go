package service

import (
	"github.com/google/uuid"
	"graduation-project/challenge-api/model"
)

type TagService interface {
	CreateTag(tag *model.Tag) (*model.Tag, error)
	GetTag(tagID string) (*model.Tag, error)
	GetTags() ([]*model.Tag, error)
}

type tagService struct {
}

func NewTagService() TagService {
	return &tagService{}
}

func (cs *tagService) CreateTag(tag *model.Tag) (*model.Tag, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	tag.ID = id.String()
	return tag, nil
}

func (cs *tagService) GetTag(tagID string) (*model.Tag, error) {
	return &model.Tag{}, nil
}

func (cs *tagService) GetTags() ([]*model.Tag, error) {
	return []*model.Tag{}, nil
}
