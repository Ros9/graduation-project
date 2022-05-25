package repository

import (
	"database/sql"
	"github.com/prometheus/common/log"
	"graduation-project/challenge-api/model"
)

type AttachmentRepository interface {
	CreateAttachment(attachment *model.Attachment) (*model.Attachment, error)
	FindAttachmentById(attachmentId string) (*model.Attachment, error)
	FindAttachmentByExternalId(title string) (*model.Attachment, error)
}

type attachmentRepository struct {
	db *sql.DB
}

func NewAttachmentRepository(db *sql.DB) AttachmentRepository {
	preQueries := []string{
		`create table attachments (
			external_id text
		)`,
	}
	for _, query := range preQueries {
		row := db.QueryRow(query)
		if row.Err() != nil {
			log.Error(row.Err().Error())
		}
	}
	return &attachmentRepository{db}
}

func (ar *attachmentRepository) CreateAttachment(attachment *model.Attachment) (*model.Attachment, error) {
	//challenge_(challenge_id)
	//company_(company_id)
	//achievement_(achievement_id)
	row := ar.db.QueryRow("insert into attachments (external_id) "+
		"values ($1)", &attachment.ExternalId)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return attachment, nil
}

func (ar *attachmentRepository) FindAttachmentById(attachmentId string) (*model.Attachment, error) {
	attachment := &model.Attachment{}
	err := ar.db.QueryRow("select * from attachments where id = $1", &attachmentId).
		Scan(&attachment.ExternalId)
	if err != nil {
		return nil, err
	}
	return attachment, nil
}

func (ar *attachmentRepository) FindAttachmentByExternalId(externalId string) (*model.Attachment, error) {
	attachment := &model.Attachment{}
	err := ar.db.QueryRow("select * from attachments where external_id = $1", &externalId).
		Scan(&attachment.ExternalId)
	if err != nil {
		return nil, err
	}
	return attachment, nil
}
