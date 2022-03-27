package service

import (
	"github.com/google/uuid"
	"graduation-project/challenge-api/model"
)

type CompanyService interface {
	CreateCompany(answer *model.Company) (*model.Company, error)
	GetCompany(answerID string) (*model.Company, error)
	GetCompanies() ([]*model.Company, error)
}

type companyService struct {
}

func NewCompanyService() CompanyService {
	return &companyService{}
}

func (cs *companyService) CreateCompany(userAnswer *model.Company) (*model.Company, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	userAnswer.ID = id.String()
	return userAnswer, nil
}

func (cs *companyService) GetCompany(companyID string) (*model.Company, error) {
	return &model.Company{}, nil
}

func (cs *companyService) GetCompanies() ([]*model.Company, error) {
	return []*model.Company{}, nil
}
