package model

type UserAnswer struct {
	AnswerID    string `json:"answer_id"`
	ChallengeID string `json:"challenge_id"`
	UserID      string `json:"user_id"`
	Answer      string `json:"answer"`
	Status      string `json:"status"`
}
