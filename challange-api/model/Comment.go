package model

type Comment struct {
	ID              string       `json:"id"`
	ChallengeID     string       `json:"challenge_id"`
	UserId          string       `json:"user_id"`
	ParentCommentID string       `json:"parent_comment_id"`
	Text            string       `json:"text"`
	Attachments     []Attachment `json:"attachments"`
}
