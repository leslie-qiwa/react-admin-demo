package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/leslie-qiwa/react-admin-demo/helpers"
	"github.com/leslie-qiwa/react-admin-demo/infra/database"
	"github.com/leslie-qiwa/react-admin-demo/infra/logger"
	"github.com/leslie-qiwa/react-admin-demo/models"
	"net/http"
)

func (ctrl *RAController) CreateCustomer(ctx *gin.Context) {
	cus := new(models.Customer)

	err := ctx.ShouldBindJSON(&cus)
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = database.DB.Create(&cus).Error
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &cus)
}

func (ctrl *RAController) GetCustomers(ctx *gin.Context) {
	query, err := parseQueryPagination(ctx)
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	param := helpers.Param{
		DB:     database.DB,
		Offset: query.offset,
		Limit:  query.limit,
	}
	if query.order != "" {
		param.OrderBy = "id " + query.order
	}
	var cus []models.Customer
	paginateData := helpers.Paginate(&param, &cus)
	ctx.Writer.Header().Set("x-total-count", fmt.Sprintf("%d", paginateData.TotalRecord))

	ctx.JSON(http.StatusOK, cus)
}

func (ctrl *RAController) GetCustomer(ctx *gin.Context) {
	id := ctx.Param("id")
	var cu models.Customer
	database.DB.First(&cu, "id=?", id)
	ctx.JSON(http.StatusOK, cu)
}
