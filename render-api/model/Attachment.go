package model

import (
	"mime/multipart"
)

type Attachment struct {
	ExternalId string          `json:"external_id"`
	File       *multipart.File `json:"-"`
}
