package repository

import (
	"database/sql"
	"fmt"
	"github.com/prometheus/common/log"
	"graduation-project/challenge-api/model"
)

type AchievementRepository interface {
	CreateAchievement(achievement *model.Achievement) (*model.Achievement, error)
	FindAchievementById(achievementId string) (*model.Achievement, error)
	FindAchievements() ([]*model.Achievement, error)
}

type achievementRepository struct {
	db *sql.DB
}

func NewAchievementRepository(db *sql.DB) AchievementRepository {
	preQueries := []string{
		`create table achievements (
			id text,
			title text,
			description text
		)`,
	}
	for _, query := range preQueries {
		row := db.QueryRow(query)
		if row.Err() != nil {
			log.Error(row.Err().Error())
		}
	}
	return &achievementRepository{db}
}

func (ar *achievementRepository) CreateAchievement(achievement *model.Achievement) (*model.Achievement, error) {
	row := ar.db.QueryRow("insert into achievements (id, name, description) "+
		"values ($1, $2, $3)", &achievement.ID, &achievement.Title, &achievement.Description)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return achievement, nil
}

func (ar *achievementRepository) FindAchievementById(achievementId string) (*model.Achievement, error) {
	achievement := &model.Achievement{}
	err := ar.db.QueryRow("select * from achievements where id = $1", &achievementId).
		Scan(&achievement.ID, &achievement.Title, &achievement.Description)
	if err != nil {
		return nil, err
	}
	return achievement, nil
}

func (at *achievementRepository) FindAchievements() ([]*model.Achievement, error) {
	rows, err := at.db.Query("select * from achievements")
	if err != nil {
		return nil, err
	}
	achievements := []*model.Achievement{}
	for rows.Next() {
		achievement := &model.Achievement{}
		err := rows.Scan(&achievement.ID, &achievement.Title, &achievement.Description)
		if err != nil {
			fmt.Println("error =", err.Error())
			return nil, err
		}
		achievements = append(achievements, achievement)
	}
	return achievements, nil
}
