package repository

import (
	"database/sql"
	"fmt"
	"graduation-project/challenge-api/model"

	"github.com/prometheus/common/log"
)

type ChallengeRepository interface {
	CreateChallenge(challenge *model.Challenge) (*model.Challenge, error)
	FindChallengeById(challengeId string) (*model.Challenge, error)
	FindChallenges() ([]*model.Challenge, error)
	GetChallengesByUserId(userId string) ([]*model.Challenge, error)
	GetChallengeByAnswer(answer string) (*model.Challenge, error)
	GetActiveChallengeByAnswer(answer string) (*model.Challenge, error)
	FindActiveChallenges() ([]*model.Challenge, error)
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

func (cr *challengeRepository) FindChallenges() ([]*model.Challenge, error) {
	q := "select * from challenges"
	rows, err := cr.db.Query(q)
	if err != nil {
		fmt.Println("error =", err.Error())
	}
	challenges := []*model.Challenge{}
	for rows.Next() {
		challenge := &model.Challenge{}
		err := rows.Scan(&challenge.ID, &challenge.CompanyID, &challenge.Title, &challenge.Description,
			&challenge.AnswerCode, &challenge.StartDate, &challenge.EndDate)
		if err != nil {
			fmt.Println("error =", err.Error())
			return nil, err
		}
		challenges = append(challenges, challenge)
	}
	return challenges, nil
}

func (cr *challengeRepository) FindActiveChallenges() ([]*model.Challenge, error) {
	q := "select * from challenges where end_time > now()"
	rows, err := cr.db.Query(q)
	if err != nil {
		fmt.Println("error =", err.Error())
	}
	challenges := []*model.Challenge{}
	for rows.Next() {
		challenge := &model.Challenge{}
		err := rows.Scan(&challenge.ID, &challenge.CompanyID, &challenge.Title, &challenge.Description,
			&challenge.AnswerCode, &challenge.StartDate, &challenge.EndDate)
		if err != nil {
			fmt.Println("error =", err.Error())
			return nil, err
		}
		challenges = append(challenges, challenge)
	}
	return challenges, nil
}

//TODO
func (cr *challengeRepository) GetChallengesByUserId(userId string) ([]*model.Challenge, error) {
	q := fmt.Sprintf("select c.* from answers a join challenges c on c.id = a.challenge_id and a.status = 1 where a.user_id = '%s'", userId)
	rows, err := cr.db.Query(q)
	if err != nil {
		fmt.Println("error =", err.Error())
	}
	challenges := []*model.Challenge{}
	for rows.Next() {
		challenge := &model.Challenge{}
		err := rows.Scan(&challenge.ID, &challenge.CompanyID, &challenge.Title, &challenge.Description,
			&challenge.AnswerCode, &challenge.StartDate, &challenge.EndDate)
		if err != nil {
			fmt.Println("error =", err.Error())
			return nil, err
		}
		challenges = append(challenges, challenge)
	}
	fmt.Println("challenges of userId = ", userId, " ===> ")
	for _, challenge := range challenges {
		fmt.Println(challenge)
	}
	return challenges, nil
}

func (cr *challengeRepository) GetChallengeByAnswer(answer string) (*model.Challenge, error) {
	challenge := &model.Challenge{}
	err := cr.db.QueryRow("select * from challenges where answer_code = $1", &answer).
		Scan(&challenge.ID, &challenge.CompanyID, &challenge.Title, &challenge.Description, &challenge.AnswerCode, &challenge.StartDate, &challenge.EndDate)
	if err != nil {
		return nil, err
	}
	return challenge, nil
}

func (cr *challengeRepository) GetActiveChallengeByAnswer(answer string) (*model.Challenge, error) {
	challenge := &model.Challenge{}
	err := cr.db.QueryRow("select * from challenges where answer_code = $1 and end_time > now()", &answer).
		Scan(&challenge.ID, &challenge.CompanyID, &challenge.Title, &challenge.Description, &challenge.AnswerCode, &challenge.StartDate, &challenge.EndDate)
	if err != nil {
		return nil, err
	}
	return challenge, nil
}
