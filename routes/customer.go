package routes

import (
	"BankApp/controllers"

	"github.com/gin-gonic/gin"
)

func CustRoute(router *gin.Engine, controller controllers.TransactionController) {
	router.POST("/customer", controller.CreateCustomer)
	router.GET("/customer/:id", controller.GetCustomerById)
	router.PUT("/customer/:id", controller.UpdateCustomerById)
	router.DELETE("/customer/:id", controller.DeleteCustomerById)
	router.GET("/customertrans/:id", controller.GetAllCustomerTransaction)
	router.POST("/customertrans/:id", controller.GetAllCustomerTransactionByDate)
	router.POST("/customertrans", controller.CreateTransaction)
	router.GET("/transsum/:id", controller.GetAllTransactionSum)

}