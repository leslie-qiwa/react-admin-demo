package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/leslie-qiwa/react-admin-demo/infra/database"
	"github.com/leslie-qiwa/react-admin-demo/models"
	"net/http"
)

// SelectedFiledFetch fields fetch from defining new struct
type SelectedFiledFetch struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

func (ctrl *RAController) GetSelectedFieldData(ctx *gin.Context) {
	var selectData []SelectedFiledFetch
	database.DB.Model(&models.Article{}).Find(&selectData)
	ctx.JSON(http.StatusOK, selectData)

}
