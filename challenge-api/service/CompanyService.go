package service

import (
	"errors"
	"github.com/google/uuid"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/repository"
	"graduation-project/challenge-api/utils"
)

type CompanyService interface {
	CreateCompany(answer *model.Company) (*model.Company, error)
	GetCompany(companyId string) (*model.Company, error)
	GetCompanies() ([]*model.Company, error)
	GetTokenForCompany(login, password string) (string, error)
}

type companyService struct {
	companyRepository repository.CompanyRepository
}

func NewCompanyService(companyRepository repository.CompanyRepository) CompanyService {
	return &companyService{companyRepository}
}

func (cs *companyService) CreateCompany(company *model.Company) (*model.Company, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	company.ID = id.String()
	createdCompany, err := cs.companyRepository.CreateCompany(company)
	if err != nil {
		return nil, err
	}
	return createdCompany, nil
}

func (cs *companyService) GetCompany(companyID string) (*model.Company, error) {
	company, err := cs.companyRepository.FindCompanyById(companyID)
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (cs *companyService) GetCompanies() ([]*model.Company, error) {
	return []*model.Company{}, nil
}

func (cs *companyService) GetTokenForCompany(login, password string) (string, error) {
	company, err := cs.companyRepository.FindCompanyByLogin(login)
	if err != nil {
		return "", err
	}
	if company.Password != password {
		return "", errors.New("invalid company or password")
	}
	return utils.GetToken(company.ID)
}
