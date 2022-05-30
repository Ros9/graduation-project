package models

import (
	"mime/multipart"
)

type Achievement struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Tags  []Tag  `json:"tags"`
}

type Answer struct {
	ID          string `json:"id"`
	ChallengeID string `json:"challenge_id"`
	UserID      string `json:"user_id"`
	Answer      string `json:"answer"`
	Status      int    `json:"status"`
}

type Attachment struct {
	ID   string `json:"id"`
	File multipart.File
}

type Bonus struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Challenge struct {
	ID          string   `json:"id"`
	CompanyID   string   `json:"company_id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	AnswerCode  string   `json:"answer_code"`
	TagsIds     []string `json:"tags_ids"`
	ImageUrl    string   `json:"image_url"`
	Tags        []Tag    `json:"tags"`
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
}

type Comment struct {
	ID              string       `json:"id"`
	ChallengeID     string       `json:"challenge_id"`
	UserId          string       `json:"user_id"`
	ParentCommentID string       `json:"parent_comment_id"`
	Text            string       `json:"text"`
	Attachments     []Attachment `json:"attachments"`
}

type Company struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Email       string `json:"email"`
	ImageUrl    string `json:"image_url"`
}

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
	Telegram string `json:"telegram"`
	IsAdmin  int    `json:"is_admin"`
}

type Result struct {
	Status    int       `json:"status"`
	Challenge Challenge `json:"challenge"`
}

type AttachmentLinkReq struct {
	ExternalId string `json:"external_id"`
	Link       string `json:"link"`
}
