package model

type Company struct {
	ID          string `json:"id"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Email       string `json:"email"`
}
