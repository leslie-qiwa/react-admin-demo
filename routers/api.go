package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/leslie-qiwa/react-admin-demo/controllers"
)

func ExamplesRoutes(route *gin.Engine) {
	var ctrl controllers.RAController
	v1 := route.Group("/v1/examples")
	v1.GET("test/", ctrl.GetExampleData)
	v1.POST("test/", ctrl.CreateExample)
	v1.GET("test/paginated", ctrl.GetExamplePaginated)
	v1.GET("test/relational", ctrl.GetHasManyRelationUserData)
	v1.GET("test/card", ctrl.GetHasManyRelationCreditCardData)
	v1.GET("test/user", ctrl.GetUserDetails)

	ra := route.Group("/v1/")
	ra.GET("category/", ctrl.GetCategories)
	ra.POST("category/", ctrl.CreateCategory)
	ra.GET("product/", ctrl.GetProducts)
	ra.POST("product/", ctrl.CreateProduct)
	ra.GET("customer/", ctrl.GetCustomers)
	ra.POST("customer/", ctrl.CreateCustomer)
}
