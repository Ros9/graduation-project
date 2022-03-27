package model

type Challenge struct {
	ID          string       `json:"id"`
	SponsorId   string       `json:"sponsor_id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	AnswerCode  string       `json:"answer_code"`
	Tags        []Tag        `json:"tags"`
	Attachments []Attachment `json:"attachments"`
	StartDate   string       `json:"start_date"`
	EndDate     string       `json:"end_date"`
}
