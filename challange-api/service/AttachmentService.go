package service

import (
	"github.com/google/uuid"
	"graduation-project/challange-api/model"
)

type AttachmentService interface {
	CreateAttachment(attachment *model.Attachment) (*model.Attachment, error)
	GetAttachment(attachmentID string) (*model.Attachment, error)
	GetAttachments() ([]*model.Attachment, error)
}

type attachmentService struct {
}

func NewAttachmentService() AttachmentService {
	return &attachmentService{}
}

func (cs *attachmentService) CreateAttachment(attachment *model.Attachment) (*model.Attachment, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	attachment.ID = id.String()
	return attachment, nil
}

func (cs *attachmentService) GetAttachment(attachmentID string) (*model.Attachment, error) {
	return &model.Attachment{}, nil
}

func (cs *attachmentService) GetAttachments() ([]*model.Attachment, error) {
	return []*model.Attachment{}, nil
}
