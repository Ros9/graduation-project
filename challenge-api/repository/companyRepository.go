package repository

import (
	"database/sql"
	"github.com/prometheus/common/log"
	"graduation-project/challenge-api/model"
)

type CompanyRepository interface {
	CreateCompany(company *model.Company) (*model.Company, error)
	FindCompanyById(companyId string) (*model.Company, error)
	FindCompanyByLogin(login string) (*model.Company, error)
}

type companyRepository struct {
	db *sql.DB
}

func NewCompanyRepository(db *sql.DB) CompanyRepository {
	preQueries := []string{
		`create table companies (
			id text,
			login text,
			password text,
			name text,
			description text,
			email text
		)`,
	}
	for _, query := range preQueries {
		row := db.QueryRow(query)
		if row.Err() != nil {
			log.Error(row.Err().Error())
		}
	}
	return &companyRepository{db}
}

func (cr *companyRepository) CreateCompany(company *model.Company) (*model.Company, error) {
	row := cr.db.QueryRow("insert into companies (id, logn, password, name, description, email) "+
		"values ($1, $2, $3, $4, $5, $6)", &company.ID, &company.Login, &company.Password, &company.Name, &company.Description, &company.Email)
	if row.Err() != nil {
		return nil, row.Err()
	}
	return company, nil
}

func (cr *companyRepository) FindCompanyById(companyId string) (*model.Company, error) {
	company := &model.Company{}
	err := cr.db.QueryRow("select * from companies where id = $1", &companyId).
		Scan(&company.ID, &company.Login, &company.Password, &company.Name, &company.Description, &company.Email)
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (cr *companyRepository) FindCompanyByLogin(login string) (*model.Company, error) {
	company := &model.Company{}
	err := cr.db.QueryRow("select * from companies where login = $1", &login).
		Scan(&company.ID, &company.Login, &company.Password, &company.Name, &company.Description, &company.Email)
	if err != nil {
		return nil, err
	}
	return company, nil
}
