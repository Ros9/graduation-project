package repository

import (
	"database/sql"
	"fmt"
	"github.com/prometheus/common/log"
	"graduation-project/challenge-api/model"
)

type UserAchievementRepository interface {
	CreateUserAchievement(userAchievement *model.UserAchievement) (*model.UserAchievement, error)
	FindAchievementIdsByUserId(userId string) ([]*model.UserAchievement, error)
}

type userAchievementRepository struct {
	db *sql.DB
}

func NewUserAchievementRepository(db *sql.DB) UserAchievementRepository {
	preQueries := []string{
		`create table user_achievement (
			user_id text,
			achievement_id text
		)`,
	}
	for _, query := range preQueries {
		row := db.QueryRow(query)
		if row.Err() != nil {
			log.Error(row.Err().Error())
		}
	}
	return &userAchievementRepository{db}
}

func (ut *userAchievementRepository) CreateUserAchievement(userAchievement *model.UserAchievement) (*model.UserAchievement, error) {
	row := ut.db.QueryRow("insert into user_achievement (user_id, achievement_id) "+
		"values ($1, $2)", &userAchievement.UserId, &userAchievement.AchievementId)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return userAchievement, nil
}

func (ut *userAchievementRepository) FindAchievementIdsByUserId(userId string) ([]*model.UserAchievement, error) {
	rows, err := ut.db.Query("select * from user_achievement where user_id = $1", &userId)
	if err != nil {
		return nil, err
	}
	uas := []*model.UserAchievement{}
	for rows.Next() {
		userAchievement := &model.UserAchievement{}
		err := rows.Scan(&userAchievement.UserId, &userAchievement.AchievementId)
		if err != nil {
			fmt.Println("error =", err.Error())
			return nil, err
		}
		uas = append(uas, userAchievement)
	}
	return uas, nil
}
