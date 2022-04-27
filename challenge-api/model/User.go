package model

type User struct {
	ID         string       `json:"id"`
	Login      string       `json:"login"`
	Email      string       `json:"email"`
	Name       string       `json:"name"`
	Surname    string       `json:"surname"`
	Password   string       `json:"password"`
	Telegram   string       `json:"telegram"`
	Challenges []*Challenge `json:"challenges"`
}

type UserTelegram struct {
	ID         string                       `json:"id"`
	Login      string                       `json:"login"`
	Email      string                       `json:"email"`
	Name       string                       `json:"name"`
	Surname    string                       `json:"surname"`
	Password   string                       `json:"-"`
	Telegram   string                       `json:"telegram"`
	Challenges []*ChallengeTelegramResponse `json:"challenges"`
}
