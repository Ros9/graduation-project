package service

import (
	"fmt"
	"github.com/google/uuid"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/repository"
	"io"
	"os"
)

type AttachmentService interface {
	CreateAttachment(attachment *model.Attachment) (*model.Attachment, error)
	GetAttachment(attachmentID string) (*model.Attachment, error)
	GetAttachments() ([]*model.Attachment, error)
}

type attachmentService struct {
	attachmentRepository repository.AttachmentRepository
}

func NewAttachmentService(attachmentRepository repository.AttachmentRepository) AttachmentService {
	return &attachmentService{attachmentRepository}
}

func (as *attachmentService) CreateAttachment(attachment *model.Attachment) (*model.Attachment, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	attachment.ID = id.String()
	createdAttachment, err := as.attachmentRepository.CreateAttachment(attachment)
	if err != nil {
		return nil, err
	}
	out, err := os.Create(attachment.ID)
	if err != nil {
		fmt.Println("error =", err.Error())
		return nil, err
	}
	defer out.Close()
	_, err = io.Copy(out, *attachment.File)
	if err != nil {
		fmt.Println("error =", err.Error())
	}
	createdAttachment.File = attachment.File
	return createdAttachment, nil
}

func (as *attachmentService) GetAttachment(attachmentId string) (*model.Attachment, error) {
	attachment, err := as.attachmentRepository.FindAttachmentById(attachmentId)
	if err != nil {
		return nil, err
	}
	return attachment, nil
}

func (as *attachmentService) GetAttachments() ([]*model.Attachment, error) {
	return []*model.Attachment{}, nil
}
