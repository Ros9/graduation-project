package model

type Achivment struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Tags  []Tag  `json:"tags"`
}
