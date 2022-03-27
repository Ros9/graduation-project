package controller

import (
	"github.com/gin-gonic/gin"
	"graduation-project/challange-api/service"
)

type CompanyController interface {
	CreateCompany() gin.HandlerFunc
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
