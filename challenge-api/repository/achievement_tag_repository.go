package repository

import (
	"database/sql"
	"fmt"
	"github.com/prometheus/common/log"
	"graduation-project/challenge-api/model"
)

type AchievementTagRepository interface {
	CreateAchievementTag(achievementTag *model.AchievementTag) (*model.AchievementTag, error)
	FindTagsIdsByAchievementId(achievementId string) ([]*model.AchievementTag, error)
}

type achievementTagRepository struct {
	db *sql.DB
}

func NewAchievementTagRepository(db *sql.DB) AchievementTagRepository {
	preQueries := []string{
		`create table achievement_tag (
			achievement_id text,
			tag_id text
		)`,
	}
	for _, query := range preQueries {
		row := db.QueryRow(query)
		if row.Err() != nil {
			log.Error(row.Err().Error())
		}
	}
	return &achievementTagRepository{db}
}

func (at *achievementTagRepository) CreateAchievementTag(achievementTag *model.AchievementTag) (*model.AchievementTag, error) {
	row := at.db.QueryRow("insert into achievement_tag (achievement_id, tag_id) "+
		"values ($1, $2)", &achievementTag.AchievementId, &achievementTag.TagId)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return achievementTag, nil
}

func (at *achievementTagRepository) FindTagsIdsByAchievementId(achievementId string) ([]*model.AchievementTag, error) {
	rows, err := at.db.Query("select * from achievement_tag where achievement_id = $1", &achievementId)
	if err != nil {
		return nil, err
	}
	ats := []*model.AchievementTag{}
	for rows.Next() {
		achievementTag := &model.AchievementTag{}
		err := rows.Scan(&achievementTag.AchievementId, &achievementTag.TagId)
		if err != nil {
			fmt.Println("error =", err.Error())
			return nil, err
		}
		ats = append(ats, achievementTag)
	}
	return ats, nil
}
