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
	var reviews []models.Review
	paginateData := helpers.Paginate(&param, &reviews)
	ctx.Writer.Header().Set("x-total-count", fmt.Sprintf("%d", paginateData.TotalRecord))

	ctx.JSON(http.StatusOK, reviews)
}

func (ctrl *RAController) GetReview(ctx *gin.Context) {
	id := ctx.Param("id")
	var review models.Review
	database.DB.First(&review, "id=?", id)
	ctx.JSON(http.StatusOK, review)
}
