package model

type Answer struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	ChallengeID string `json:"challenge_id"`
	Answer      string `json:"answer"`
	Status      string `json:"status"`
}

type Result struct {
	Status    int                       `json:"status"`
	Challenge ChallengeTelegramResponse `json:"challenge"`
}
