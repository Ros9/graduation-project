package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"graduation-project/challenge-api/model"
	"graduation-project/challenge-api/service"
	"io/ioutil"
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
		jsonData, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.JSON(500, err.Error())
		}
		company := &model.Company{}
		err = json.Unmarshal(jsonData, company)
		if err != nil {
			context.JSON(404, err.Error())
		}
		createdCompany, err := cc.companyService.CreateCompany(company)
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, createdCompany)
	}
}

func (cc *companyController) GetCompany() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		company, err := cc.companyService.GetCompany(id)
		if err != nil {
			context.JSON(404, err.Error())
		}
		context.JSON(200, company)
	}
}

func (cc *companyController) GetCompanies() gin.HandlerFunc {
	return func(context *gin.Context) {
		companies, err := cc.companyService.GetCompanies()
		if err != nil {
			context.JSON(500, err.Error())
		}
		context.JSON(200, companies)
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
