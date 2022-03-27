package model

type Achievement struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Tags  []Tag  `json:"tags"`
}
