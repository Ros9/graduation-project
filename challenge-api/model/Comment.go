package model

type Comment struct {
	ID          string `json:"id"`
	ChallengeID string `json:"challenge_id"`
	UserId      string `json:"user_id"`
	Description string `json:"description"`
}
