package controller

import "github.com/gin-gonic/gin"

type CompanyController interface {
	CreateCompany() gin.HandlerFunc
}

type companyController struct {
}

func NewCompanyController() CompanyController {
	return &companyController{}
}

func (cc *companyController) CreateCompany() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
