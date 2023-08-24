package controllers

import (
	"BankApp/Interfaces"
	"BankApp/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionService interfaces.Icustomer
}

type Date1 struct{
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
}

func InitTransController(transactionService interfaces.Icustomer) TransactionController {
	return TransactionController{transactionService}
}

func (t *TransactionController) CreateCustomer(ctx *gin.Context) {
	var trans *models.Customer
	if err := ctx.ShouldBindJSON(&trans); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newtrans, err := t.TransactionService.CreateCustomer(trans)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newtrans})

}

func (t *TransactionController) GetCustomerById(ctx *gin.Context) {
	id := ctx.Param("id")
	id1, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	val, err := t.TransactionService.GetCustomerById(id1)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})

	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": val})
}

func (t *TransactionController) UpdateCustomerById(ctx *gin.Context) {
	id := ctx.Param("id")
	fv := &models.UpdateModel{}
	if err := ctx.ShouldBindJSON(&fv); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(fv)
	id1, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	res, err := t.TransactionService.UpdateCustomerById(id1, fv)
	if err != nil {
		fmt.Println("error")
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}

func (t *TransactionController) DeleteCustomerById(ctx *gin.Context) {
	id := ctx.Param("id")
	id1, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	res, err := t.TransactionService.DeleteCustomerById(id1)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}

func (t *TransactionController) GetAllCustomerTransaction(ctx *gin.Context) {
	id := ctx.Param("id")
	id1, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	res, err := t.TransactionService.GetAllCustomerTransaction(id1)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": res})
}

func (t *TransactionController) GetAllCustomerTransactionByDate(ctx *gin.Context) {
	customerIDString := ctx.Param("id") // Extract customer ID from URL parameter

	// Parse the customer ID from string to int64
	customerID, err := strconv.ParseInt(customerIDString, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID format"})
		return
	}
    var date *Date1
	if err := ctx.ShouldBindJSON(&date); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// Parse the start date string into a time.Time object
	// startDateString := ctx.Param("date") // Extract start date string from URL parameter
	// startDate, err := time.Parse(time.RFC3339, startDateString)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
	// 	return
	// }

	// Get customer transactions from start date to the present date
	// endDate := time.Now() // Assuming you want transactions up to the current datetime.Time

	// Call the method to fetch transactions based on customer ID and date range
	transactions, err := t.TransactionService.GetAllCustomerTransactionByDate(customerID, date.StartDate, date.EndDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

func (t *TransactionController) CreateTransaction(ctx *gin.Context) {
	var transaction models.CustTransaction

	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	result, err := t.TransactionService.CreateTransaction(&transaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	ctx.JSON(http.StatusOK, result)
}