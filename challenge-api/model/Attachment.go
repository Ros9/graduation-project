package model

import "mime/multipart"

type Attachment struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	File  multipart.File
}
