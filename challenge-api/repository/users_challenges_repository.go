package repository

import (
	"database/sql"
	"fmt"
	"graduation-project/challenge-api/model"

	"github.com/prometheus/common/log"
)

type UsersChallengesRepository interface {
	CreateUserChallenge(userChallenge *model.UserChallenge) (*model.UserChallenge, error)
	FindChallengesByUserId(userId string) ([]*model.UserChallenge, error)
}

type usersChallengesRepository struct {
	db *sql.DB
}

func NewUsersChallengesRepository(db *sql.DB) UsersChallengesRepository {
	preQueries := []string{
		`create table users_challenges (
			user_id text,
			challenge_id text,
			users_answer text,
			status int,
			created_at timestamp
		)`,
	}
	for _, query := range preQueries {
		row := db.QueryRow(query)
		if row.Err() != nil {
			log.Error(row.Err().Error())
		}
	}
	return &usersChallengesRepository{db}
}

func (ur *usersChallengesRepository) CreateUserChallenge(userChallenge *model.UserChallenge) (*model.UserChallenge, error) {
	row := ur.db.QueryRow("insert into users_challenges (user_id, challenge_id) "+
		"values ($1, $2)", &userChallenge.UserId, &userChallenge.ChallengeId)
	if row.Err() != nil {
		return nil, row.Err()
	}
	fmt.Println()
	fmt.Println("userChallenge =", *userChallenge)
	fmt.Println()
	return userChallenge, nil
}

func (ur *usersChallengesRepository) FindChallengesByUserId(userId string) ([]*model.UserChallenge, error) {
	rows, err := ur.db.Query("select * from users_challenges where user_id = $1", userId)
	if err != nil {

	}
	usersChallenges := []*model.UserChallenge{}
	for rows.Next() {
		if rows.Err() != nil {
			fmt.Println("error =", err.Error())
		}
		userChallenge := &model.UserChallenge{}
		err = rows.Scan(&userChallenge.UserId, &userChallenge.ChallengeId)
		if err != nil {
			fmt.Println("error =", err.Error())
		}
		fmt.Println(*userChallenge)
		usersChallenges = append(usersChallenges, userChallenge)
	}
	return usersChallenges, nil
}
