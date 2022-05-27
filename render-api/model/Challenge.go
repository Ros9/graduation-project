package model

import "time"

type Challenge struct {
	ID          string     `json:"id"`
	CompanyID   string     `json:"company_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	AnswerCode  string     `json:"answer_code"`
	TagsIds     []string   `json:"tags_ids"`
	ImageUrl    string     `json:"image_url"`
	Tags        []Tag      `json:"tags"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
}

//GetChallengesTgRespByUserId
type ChallengeTelegramResponse struct {
	ID          string     `json:"id"`
	CompanyID   string     `json:"company_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	AnswerCode  string     `json:"-"`
	TagsIds     []string   `json:"-"`
	ImageUrl    string     `json:"-"`
	Tags        []Tag      `json:"-"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
}
