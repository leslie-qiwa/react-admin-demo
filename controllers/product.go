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

func (ctrl *RAController) CreateProduct(ctx *gin.Context) {
	product := new(models.Product)

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = database.DB.Create(&product).Error
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &product)
}

func (ctrl *RAController) GetProducts(ctx *gin.Context) {
	query, err := parseQueryPagination(ctx)
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var products []models.Product
	param := mkPaginateParam(query)
	paginateData := helpers.Paginate(param, &products)

	ctx.Writer.Header().Set("x-total-count", fmt.Sprintf("%d", paginateData.TotalRecord))
	ctx.JSON(http.StatusOK, products)
}
