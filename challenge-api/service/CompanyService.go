package service

import (
	"fmt"
	"github.com/google/uuid"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/repository"
)

type CompanyService interface {
	CreateCompany(answer *model.Company) (*model.Company, error)
	GetCompany(companyId string) (*model.Company, error)
	GetCompanies() ([]*model.Company, error)
}

type companyService struct {
	companyRepository repository.CompanyRepository
	attachmentService AttachmentService
}

func NewCompanyService(companyRepository repository.CompanyRepository, attachmentService AttachmentService) CompanyService {
	return &companyService{companyRepository, attachmentService}
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
	challengeExternalId := "company_" + company.ID
	attachment, err := cs.attachmentService.GetAttachmentByExternalId(challengeExternalId)
	if err != nil {
		fmt.Println("error when get challenge =", err.Error())
	}
	if attachment != nil {
		company.ImageUrl = "/assets/image/" + challengeExternalId
	}
	return company, nil
}

func (cs *companyService) GetCompanies() ([]*model.Company, error) {
	return []*model.Company{}, nil
}
