package model

type Answer struct {
	ID          string `json:"id"`
	ChallengeID string `json:"challenge_id"`
	UserID      string `json:"user_id"`
	Answer      string `json:"answer"`
	Status      string `json:"status"`
}
