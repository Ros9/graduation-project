package repository

import (
	"database/sql"
	"github.com/prometheus/common/log"
	"graduation-project/challenge-api/model"
)

type AttachmentRepository interface {
	CreateAttachment(attachment *model.Attachment) (*model.Attachment, error)
	FindAttachmentById(attachmentId string) (*model.Attachment, error)
	FindAttachmentByChallengeId(title string) (*model.Attachment, error)
}

type attachmentRepository struct {
	db *sql.DB
}

func NewAttachmentRepository(db *sql.DB) AttachmentRepository {
	preQueries := []string{
		`create table attachments (
			id text,
			title text
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
	row := ar.db.QueryRow("insert into attachments (id, title) "+
		"values ($1, $2)", &attachment.ID, &attachment.Title)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return attachment, nil
}

func (ar *attachmentRepository) FindAttachmentById(attachmentId string) (*model.Attachment, error) {
	attachment := &model.Attachment{}
	err := ar.db.QueryRow("select * from attachments where id = $1", &attachmentId).
		Scan(&attachment.ID, &attachment.Title)
	if err != nil {
		return nil, err
	}
	return attachment, nil
}

func (ar *attachmentRepository) FindAttachmentByChallengeId(title string) (*model.Attachment, error) {
	attachment := &model.Attachment{}
	err := ar.db.QueryRow("select * from attachments where title = $1", &title).
		Scan(&attachment.ID, &attachment.Title)
	if err != nil {
		return nil, err
	}
	return attachment, nil
}
