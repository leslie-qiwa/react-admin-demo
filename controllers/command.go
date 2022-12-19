package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/leslie-qiwa/react-admin-demo/infra/database"
	"github.com/leslie-qiwa/react-admin-demo/infra/logger"
	"github.com/leslie-qiwa/react-admin-demo/models"
	"net/http"
	"strconv"
)

func (ctrl *RAController) CreateCommand(ctx *gin.Context) {
	cmd := new(models.Command)

	err := ctx.ShouldBindJSON(&cmd)
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = database.DB.Create(&cmd).Error
	if err != nil {
		logger.Errorf("error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &cmd)
}

func (ctrl *RAController) GetCommands(ctx *gin.Context) {
	var cmds []models.Command
	database.DB.Find(&cmds)
	ctx.Writer.Header().Set("x-total-count", strconv.Itoa(len(cmds)))
	ctx.JSON(http.StatusOK, cmds)
}

func (ctrl *RAController) GetCommand(ctx *gin.Context) {
	id := ctx.Param("id")
	var cmd models.Command
	database.DB.Preload("Baskets").First(&cmd, "id=?", id)
	ctx.JSON(http.StatusOK, cmd)
}
