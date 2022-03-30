package repository

import (
	"database/sql"
	"github.com/prometheus/common/log"
	"graduation-project/challenge-api/model"
)

type AnswerRepository interface {
	CreateAnswer(answer *model.Answer) (*model.Answer, error)
	FindAnswerById(answerId string) (*model.Answer, error)
}

type answerRepository struct {
	db *sql.DB
}

func NewAnswerRepository(db *sql.DB) AnswerRepository {
	preQueries := []string{
		`create table answers (
			id text,
			user_id text,
			challenge_id text,
			answer text,
			status text
		)`,
	}
	for _, query := range preQueries {
		row := db.QueryRow(query)
		if row.Err() != nil {
			log.Error(row.Err().Error())
		}
	}
	return &answerRepository{db}
}

func (ar *answerRepository) CreateAnswer(answer *model.Answer) (*model.Answer, error) {
	row := ar.db.QueryRow("insert into companies (id, user_id, challenge_id, answer, status) "+
		"values ($1, $2, $3, $4, $5)", &answer.ID, &answer.UserID, &answer.ChallengeID, &answer.Answer, &answer.Status)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return answer, nil
}

func (ar *answerRepository) FindAnswerById(answerId string) (*model.Answer, error) {
	answer := &model.Answer{}
	err := ar.db.QueryRow("select * from answers where id = $1", &answerId).
		Scan(&answer.ID, &answer.UserID, &answer.ChallengeID, &answer.Answer, &answer.Status)
	if err != nil {
		return nil, err
	}
	return answer, nil
}
