package service

import (
	"github.com/google/uuid"
	"graduation-project/challenge-api/model"
)

type BonusService interface {
	CreateBonus(bonus *model.Bonus) (*model.Bonus, error)
	GetBonus(bonusID string) (*model.Bonus, error)
	GetBonuses() ([]*model.Bonus, error)
}

type bonusService struct {
}

func NewBonusService() BonusService {
	return &bonusService{}
}

func (cs *bonusService) CreateBonus(userAnswer *model.Bonus) (*model.Bonus, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	userAnswer.ID = id.String()
	return userAnswer, nil
}

func (cs *bonusService) GetBonus(bonusID string) (*model.Bonus, error) {
	return &model.Bonus{}, nil
}

func (cs *bonusService) GetBonuses() ([]*model.Bonus, error) {
	return []*model.Bonus{}, nil
}
