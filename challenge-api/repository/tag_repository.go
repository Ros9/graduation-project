package repository

import (
	"database/sql"
	"fmt"
	"graduation-project/challenge-api/model"

	"github.com/prometheus/common/log"
)

type TagRepository interface {
	CreateTag(tag *model.Tag) (*model.Tag, error)
	FindTagById(tagId string) (*model.Tag, error)
	FindTags() ([]*model.Tag, error)
}

type tagRepository struct {
	db *sql.DB
}

func NewTagRepository(db *sql.DB) TagRepository {
	preQueries := []string{
		`create table tags (
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
	return &tagRepository{db}
}

func (tr *tagRepository) CreateTag(tag *model.Tag) (*model.Tag, error) {
	row := tr.db.QueryRow("insert into tags (id, title, description) "+
		"values ($1, $2, $3)", &tag.ID, &tag.Title, &tag.Description)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return tag, nil
}

func (tr *tagRepository) FindTagById(tagId string) (*model.Tag, error) {
	tag := &model.Tag{}
	err := tr.db.QueryRow("select * from tags where id = $1", &tagId).
		Scan(&tag.ID, &tag.Title, &tag.Description)
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func (tr *tagRepository) FindTags() ([]*model.Tag, error) {
	rows, err := tr.db.Query("select * from tags order by id")
	if err != nil {
		return nil, err
	}
	tags := []*model.Tag{}
	for rows.Next() {
		tag := &model.Tag{}
		err := rows.Scan(&tag.ID, &tag.Title, &tag.Description)
		if err != nil {
			fmt.Println("error =", err.Error())
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}
