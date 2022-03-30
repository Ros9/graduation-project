package repository

import (
	"database/sql"
	"github.com/prometheus/common/log"
	"graduation-project/challenge-api/model"
)

type BonusRepository interface {
	CreateBonus(bonus *model.Bonus) (*model.Bonus, error)
	FindBonusById(bonusId string) (*model.Bonus, error)
}

type bonusRepository struct {
	db *sql.DB
}

func NewBonusRepository(db *sql.DB) BonusRepository {
	preQueries := []string{
		`create table bonuses (
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
	return &bonusRepository{db}
}

func (ar *bonusRepository) CreateBonus(bonus *model.Bonus) (*model.Bonus, error) {
	row := ar.db.QueryRow("insert into bonuses (id, name, description) "+
		"values ($1, $2, $3)", &bonus.ID, &bonus.Title, &bonus.Description)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return bonus, nil
}

func (ar *bonusRepository) FindBonusById(bonusId string) (*model.Bonus, error) {
	bonus := &model.Bonus{}
	err := ar.db.QueryRow("select * from bonuses where id = $1", &bonusId).
		Scan(&bonus.ID, &bonus.Title, &bonus.Description)
	if err != nil {
		return nil, err
	}
	return bonus, nil
}
