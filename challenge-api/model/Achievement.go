package model

type Achievement struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Tags        []Tag  `json:"tags"`
}
