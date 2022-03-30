package repository

import (
	"database/sql"
	"github.com/prometheus/common/log"
	"graduation-project/challenge-api/model"
)

type TagRepository interface {
	CreateTag(tag *model.Tag) (*model.Tag, error)
	FindTagById(tagId string) (*model.Tag, error)
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
	row := tr.db.QueryRow("insert into tags (id, name, description) "+
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
