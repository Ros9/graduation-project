package service

import (
	"github.com/google/uuid"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/repository"
)

type TagService interface {
	CreateTag(tag *model.Tag) (*model.Tag, error)
	GetTag(tagID string) (*model.Tag, error)
	GetTags() ([]*model.Tag, error)
}

type tagService struct {
	tagRepository repository.TagRepository
}

func NewTagService(tagRepository repository.TagRepository) TagService {
	return &tagService{tagRepository}
}

func (ts *tagService) CreateTag(tag *model.Tag) (*model.Tag, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	tag.ID = id.String()
	createdTag, err := ts.tagRepository.CreateTag(tag)
	if err != nil {
		return nil, err
	}
	return createdTag, nil
}

func (ts *tagService) GetTag(tagID string) (*model.Tag, error) {
	tag, err := ts.tagRepository.FindTagById(tagID)
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func (ts *tagService) GetTags() ([]*model.Tag, error) {
	tags, err := ts.tagRepository.FindTags()
	if err != nil {
		return nil, err
	}
	return tags, nil
}
