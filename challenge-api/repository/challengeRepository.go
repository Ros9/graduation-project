package repository

import (
	"database/sql"
	"github.com/prometheus/common/log"
	"graduation-project/challenge-api/model"
)

type ChallengeRepository interface {
	CreateChallenge(challenge *model.Challenge) (*model.Challenge, error)
	FindChallengeById(challengeId string) (*model.Challenge, error)
}

type challengeRepository struct {
	db *sql.DB
}

func NewChallengeRepository(db *sql.DB) ChallengeRepository {
	preQueries := []string{
		`create table challenges (
			id text,
			company_id text,
			title text,
			description text,
			answer_code text,
			start_time timestamp,
			end_time timestamp
		)`,
	}
	for _, query := range preQueries {
		row := db.QueryRow(query)
		if row.Err() != nil {
			log.Error(row.Err().Error())
		}
	}
	return &challengeRepository{db}
}

func (cr *challengeRepository) CreateChallenge(challenge *model.Challenge) (*model.Challenge, error) {
	row := cr.db.QueryRow("insert into challenges (id, company_id, title, description, answer_code, start_time, end_time) "+
		"values ($1, $2, $3, $4, $5, $6, $7)",
		&challenge.ID, &challenge.CompanyID, &challenge.Title, &challenge.Description, &challenge.AnswerCode, &challenge.StartDate, &challenge.EndDate)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return challenge, nil
}

func (cr *challengeRepository) FindChallengeById(challengeId string) (*model.Challenge, error) {
	challenge := &model.Challenge{}
	err := cr.db.QueryRow("select * from challenges where id = $1", &challengeId).
		Scan(&challenge.ID, &challenge.CompanyID, &challenge.Title, &challenge.Description, &challenge.AnswerCode, &challenge.StartDate, &challenge.EndDate)
	if err != nil {
		return nil, err
	}
	return challenge, nil
}
