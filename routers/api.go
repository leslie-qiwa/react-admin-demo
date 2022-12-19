package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/leslie-qiwa/react-admin-demo/controllers"
)

func ExamplesRoutes(route *gin.Engine) {
	var ctrl controllers.RAController

	ra := route.Group("/v1/")
	ra.GET("categories", ctrl.GetCategories)
	ra.POST("category", ctrl.CreateCategory)
	ra.GET("products", ctrl.GetProducts)
	ra.POST("product", ctrl.CreateProduct)
	ra.POST("customer", ctrl.CreateCustomer)
	ra.GET("customers", ctrl.GetCustomers)
	ra.GET("customers/:id", ctrl.GetCustomer)
	ra.POST("command", ctrl.CreateCommand)
	ra.GET("commands", ctrl.GetCommands)
	ra.GET("commands/:id", ctrl.GetCommand)
	ra.POST("review", ctrl.CreateReview)
	ra.GET("reviews", ctrl.GetReviews)
	ra.GET("reviews/:id", ctrl.GetReview)
}
