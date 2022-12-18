package database

import (
	"github.com/leslie-qiwa/react-admin-demo/models"
)

//Add list of model add for migrations
//var migrationModels = []interface{}{&ex_models.Example{}, &model.Example{}, &model.Address{})}
var migrationModels = []interface{}{&models.Category{}, &models.Product{}}
