package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/leslie-qiwa/react-admin-demo/infra/database"
	"github.com/leslie-qiwa/react-admin-demo/infra/logger"
	"github.com/leslie-qiwa/react-admin-demo/models"
	"net/http"
	"strconv"
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
	var products []models.Product
	database.DB.Find(&products)
	ctx.Writer.Header().Set("x-total-count", strconv.Itoa(len(products)))

	ctx.JSON(http.StatusOK, products)
}
