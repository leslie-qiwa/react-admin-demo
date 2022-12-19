package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/leslie-qiwa/react-admin-demo/infra/database"
	"github.com/leslie-qiwa/react-admin-demo/infra/logger"
	"github.com/leslie-qiwa/react-admin-demo/models"
	"net/http"
	"strconv"
)

func (ctrl *RAController) CreateReview(ctx *gin.Context) {
	review := new(models.Review)

	err := ctx.ShouldBindJSON(&review)
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = database.DB.Create(&review).Error
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &review)
}

func (ctrl *RAController) GetReviews(ctx *gin.Context) {
	var reviews []models.Review
	database.DB.Find(&reviews)
	ctx.Writer.Header().Set("x-total-count", strconv.Itoa(len(reviews)))
	ctx.JSON(http.StatusOK, reviews)
}

func (ctrl *RAController) GetReview(ctx *gin.Context) {
	id := ctx.Param("id")
	var review models.Review
	database.DB.First(&review, "id=?", id)
	ctx.JSON(http.StatusOK, review)
}
