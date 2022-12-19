package database

import (
	"github.com/leslie-qiwa/react-admin-demo/models"
)

//Add list of model add for migrations
var migrationModels = []interface{}{&models.Category{}, &models.Product{}, &models.Customer{}, &models.Command{},
	&models.Basket{}, &models.Review{}}
