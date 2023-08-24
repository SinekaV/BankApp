package interfaces

import (
	"BankApp/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Icustomer interface {
	CreateCustomer(*models.Customer) (*mongo.InsertOneResult, error)
	GetCustomerById(int64) (*models.Customer, error)
	UpdateCustomerById(int64, *models.UpdateModel) (*mongo.UpdateResult, error)
	DeleteCustomerById(int64) (*mongo.DeleteResult, error)
	GetAllCustomerTransaction(int64) (*[]models.CustTransaction, error)
	GetAllCustomerTransactionByDate(int64, time.Time, time.Time) (*[]models.CustTransaction, error)
	CreateTransaction(*models.CustTransaction) (*mongo.InsertOneResult, error)
}