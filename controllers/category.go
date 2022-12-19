package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/leslie-qiwa/react-admin-demo/infra/database"
	"github.com/leslie-qiwa/react-admin-demo/infra/logger"
	"github.com/leslie-qiwa/react-admin-demo/models"
	"net/http"
	"strconv"
)

func (ctrl *RAController) CreateCategory(ctx *gin.Context) {
	cat := new(models.Category)

	err := ctx.ShouldBindJSON(&cat)
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = database.DB.Create(&cat).Error
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &cat)
}

func (ctrl *RAController) GetCategories(ctx *gin.Context) {
	var cats []models.Category
	database.DB.Find(&cats)
	ctx.Writer.Header().Set("x-total-count", strconv.Itoa(len(cats)))
	ctx.JSON(http.StatusOK, cats)
}
