package model

type User struct {
	ID         string       `json:"id"`
	Username   string       `json:"username"`
	Password   string       `json:"password"`
	Email      string       `json:"email"`
	Telegram   string       `json:"telegram"`
	Challenges []*Challenge `json:"challenges"`
}

type UserTelegram struct {
	ID         string                       `json:"id"`
	Username   string                       `json:"username"`
	Password   string                       `json:"-"`
	Email      string                       `json:"email"`
	Telegram   string                       `json:"telegram"`
	Challenges []*ChallengeTelegramResponse `json:"challenges"`
}
