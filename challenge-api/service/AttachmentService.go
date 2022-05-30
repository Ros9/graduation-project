package service

import (
	"fmt"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/repository"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type AttachmentService interface {
	CreateAttachment(attachment *model.Attachment) (*model.Attachment, error)
	GetAttachment(attachmentID string) (*model.Attachment, error)
	GetAttachments() ([]*model.Attachment, error)
	GetAttachmentByExternalId(challengeId string) (*model.Attachment, error)
	CreateAttachmentFromTelegram(attachmentLinkReq *model.AttachmentLinkReq) (*model.Attachment, error)
}

type attachmentService struct {
	attachmentRepository repository.AttachmentRepository
	httpClt              http.Client
}

func NewAttachmentService(attachmentRepository repository.AttachmentRepository, httpClt http.Client) AttachmentService {
	return &attachmentService{attachmentRepository, httpClt}
}

func (as *attachmentService) CreateAttachment(attachment *model.Attachment) (*model.Attachment, error) {
	createdAttachment, err := as.attachmentRepository.CreateAttachment(attachment)
	if err != nil {
		return nil, err
	}
	fileName := "../render-api/assets/image/" + attachment.ExternalId
	out, err := os.Create(fileName)
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

func (as *attachmentService) CreateAttachmentFromTelegram(attachmentLinkReq *model.AttachmentLinkReq) (*model.Attachment, error) {
	attachment := &model.Attachment{
		ExternalId: attachmentLinkReq.ExternalId,
	}
	createdAttachment, err := as.attachmentRepository.CreateAttachment(attachment)
	if err != nil {
		return nil, err
	}
	resp, err := as.httpClt.Get(attachmentLinkReq.Link)
	if err != nil {
		fmt.Println("error =", err.Error())
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error =", err.Error())
		return nil, err
	}
	fileName := "../render-api/assets/image/" + attachment.ExternalId
	err = ioutil.WriteFile(fileName, data, 0666)
	if err != nil {
		fmt.Println("error =", err.Error())
		return nil, err
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

func (as *attachmentService) GetAttachmentByExternalId(challengeId string) (*model.Attachment, error) {
	return as.attachmentRepository.FindAttachmentByExternalId(challengeId)
}
