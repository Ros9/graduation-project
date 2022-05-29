package model

type Comment struct {
	ID          string `json:"id"`
	ChallengeID string `json:"challenge_id"`
	UserId      string `json:"user_id"`
	Username    string `json:"username"`
	Description string `json:"description"`
}
