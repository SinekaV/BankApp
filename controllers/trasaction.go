package controllers

import (
	interfaces "BankApp/Interfaces"
	"BankApp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransferController struct{
	TransactionService interfaces.TransactionInterface
}

func InitTransferController(transactionService interfaces.TransactionInterface) TransferController{
	return TransferController{transactionService}
}

func (t *TransferController) TransferMoney(ctx *gin.Context){
	var trans *models.Transaction
	if err := ctx.ShouldBindJSON(&trans); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status":"fail", "message":err.Error()})
	}
	res, err := t.TransactionService.TransferMoney(trans.From, trans.To, trans.Amount)
	if err!=nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"status":"fail", "message":err.Error()})
	}
	ctx.JSON(http.StatusAccepted, gin.H{"status":"success", "message":res})
}