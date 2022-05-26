package model

type Comment struct {
	ID              string       `json:"id"`
	ChallengeID     string       `json:"challenge_id"`
	UserId          string       `json:"user_id"`
	Description     string       `json:"text"`
	ParentCommentID string       `json:"parent_comment_id"`
	Attachments     []Attachment `json:"attachments"`
}
