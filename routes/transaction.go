package routes

import (
	"BankApp/controllers"

	"github.com/gin-gonic/gin"
)

func TransRoute(r *gin.Engine, controller controllers.TransactionController){
	r.POST("/transfer", controllers.TransferMoney)
}