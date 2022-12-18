package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/leslie-qiwa/react-admin-demo/infra/database"
	"github.com/leslie-qiwa/react-admin-demo/models"
	"net/http"
)

// GetNormalData get normal data if added pagination see example_controller
func (ctrl *RAController) GetNormalData(ctx *gin.Context) {
	var categories []models.Category
	db := database.GetDB()
	db.Find(&categories)
	ctx.JSON(http.StatusOK, gin.H{"data": categories})

}

// GetForeignRelationData Get Foreign Data with Preload
func (ctrl *RAController) GetForeignRelationData(ctx *gin.Context) {
	var articles []models.Article
	db := database.GetDB()
	db.Preload("Category").Find(&articles)
	ctx.JSON(http.StatusOK, &articles)

}
