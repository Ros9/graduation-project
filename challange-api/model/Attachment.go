package model

import "mime/multipart"

type Attachment struct {
	ID   string `json:"id"`
	File multipart.File
}
