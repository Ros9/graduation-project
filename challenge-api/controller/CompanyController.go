package controller

import (
	"github.com/gin-gonic/gin"
	"graduation-project/challenge-api/service"
)

type CompanyController interface {
	CreateCompany() gin.HandlerFunc
	GetCompany() gin.HandlerFunc
	GetCompanies() gin.HandlerFunc
	UpdateCompany() gin.HandlerFunc
	DeleteCompany() gin.HandlerFunc
}

type companyController struct {
	companyService service.CompanyService
}

func NewCompanyController(companyService service.CompanyService) CompanyController {
	return &companyController{companyService}
}

func (cc *companyController) CreateCompany() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (cc *companyController) GetCompany() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (cc *companyController) GetCompanies() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (cc *companyController) UpdateCompany() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (cc *companyController) DeleteCompany() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
