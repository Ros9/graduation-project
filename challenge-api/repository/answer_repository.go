package repository

import (
	"database/sql"
	"fmt"
	"graduation-project/challenge-api/model"

	"github.com/prometheus/common/log"
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
			status integer,
			created_at timestamp
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
	fmt.Println("\n\nans rep", answer)
	row := ar.db.QueryRow("insert into answers (id, user_id, challenge_id, answer, status, created_at) "+
		"values ($1, $2, $3, $4, $5, now())", &answer.ID, &answer.UserID, &answer.ChallengeID, &answer.Answer, &answer.Status)
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
