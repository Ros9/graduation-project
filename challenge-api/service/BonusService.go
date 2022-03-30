package service

import (
	"github.com/google/uuid"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/repository"
)

type BonusService interface {
	CreateBonus(bonus *model.Bonus) (*model.Bonus, error)
	GetBonus(bonusID string) (*model.Bonus, error)
	GetBonuses() ([]*model.Bonus, error)
}

type bonusService struct {
	bonusRepository repository.BonusRepository
}

func NewBonusService(bonusRepository repository.BonusRepository) BonusService {
	return &bonusService{bonusRepository}
}

func (cs *bonusService) CreateBonus(bonus *model.Bonus) (*model.Bonus, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	bonus.ID = id.String()
	createdBonus, err := cs.bonusRepository.CreateBonus(bonus)
	if err != nil {
		return nil, err
	}
	return createdBonus, nil
}

func (cs *bonusService) GetBonus(bonusID string) (*model.Bonus, error) {
	bonus, err := cs.bonusRepository.FindBonusById(bonusID)
	if err != nil {
		return nil, err
	}
	return bonus, nil
}

func (cs *bonusService) GetBonuses() ([]*model.Bonus, error) {
	return []*model.Bonus{}, nil
}
