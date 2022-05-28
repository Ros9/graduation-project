package repository

import (
	"database/sql"
	"fmt"
	"github.com/prometheus/common/log"
	"graduation-project/challenge-api/model"
)

type UserTagRepository interface {
	CreateUserTag(userTag *model.UserTag) (*model.UserTag, error)
	FindTagsIdsByUserId(userId string) ([]*model.UserTag, error)
}

type userTagRepository struct {
	db *sql.DB
}

func NewUserTagRepository(db *sql.DB) UserTagRepository {
	preQueries := []string{
		`create table user_tag (
			user_id text,
			tag_id text
		)`,
	}
	for _, query := range preQueries {
		row := db.QueryRow(query)
		if row.Err() != nil {
			log.Error(row.Err().Error())
		}
	}
	return &userTagRepository{db}
}

func (ut *userTagRepository) CreateUserTag(userTag *model.UserTag) (*model.UserTag, error) {
	row := ut.db.QueryRow("insert into user_tag (user_id, tag_id) "+
		"values ($1, $2)", &userTag.UserId, &userTag.TagId)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return userTag, nil
}

func (ut *userTagRepository) FindTagsIdsByUserId(userId string) ([]*model.UserTag, error) {
	rows, err := ut.db.Query("select * from user_tag where user_id = $1", &userId)
	if err != nil {
		return nil, err
	}
	uts := []*model.UserTag{}
	for rows.Next() {
		userTag := &model.UserTag{}
		err := rows.Scan(&userTag.UserId, &userTag.TagId)
		if err != nil {
			fmt.Println("error =", err.Error())
			return nil, err
		}
		uts = append(uts, userTag)
	}
	return uts, nil
}
