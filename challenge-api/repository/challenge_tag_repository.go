package repository

import (
	"database/sql"
	"fmt"
	"github.com/prometheus/common/log"
	"graduation-project/challenge-api/model"
)

type ChallengeTagRepository interface {
	CreateChallengeTag(challengeTag *model.ChallengeTag) (*model.ChallengeTag, error)
	FindTagsIdsByChallengeId(challengeId string) ([]*model.ChallengeTag, error)
}

type challengeTagRepository struct {
	db *sql.DB
}

func NewChallengeTagRepository(db *sql.DB) ChallengeTagRepository {
	preQueries := []string{
		`create table challenge_tag (
			challenge_id text,
			tag_id text
		)`,
	}
	for _, query := range preQueries {
		row := db.QueryRow(query)
		if row.Err() != nil {
			log.Error(row.Err().Error())
		}
	}
	return &challengeTagRepository{db}
}

func (ct *challengeTagRepository) CreateChallengeTag(challengeTag *model.ChallengeTag) (*model.ChallengeTag, error) {
	row := ct.db.QueryRow("insert into challenge_tag (challenge_id, tag_id) "+
		"values ($1, $2)", &challengeTag.ChallengeId, &challengeTag.TagId)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return challengeTag, nil
}

func (at *challengeTagRepository) FindTagsIdsByChallengeId(challengeId string) ([]*model.ChallengeTag, error) {
	rows, err := at.db.Query("select * from challenge_tag where challenge_id = $1", &challengeId)
	if err != nil {
		return nil, err
	}
	cts := []*model.ChallengeTag{}
	for rows.Next() {
		challengeTag := &model.ChallengeTag{}
		err := rows.Scan(&challengeTag.ChallengeId, &challengeTag.TagId)
		if err != nil {
			fmt.Println("error =", err.Error())
			return nil, err
		}
		cts = append(cts, challengeTag)
	}
	return cts, nil
}
