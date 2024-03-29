package service

import (
	interfaces "BankApp/Interfaces"
	"BankApp/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Bank1 struct {
	ctx  context.Context
	coll *mongo.Collection
}

func InitBank(collection *mongo.Collection, ctx context.Context) interfaces.IBank {
	return &Bank1{ctx, collection}
}

func (b *Bank1) CreateBank(bank *models.Bank) (*mongo.InsertOneResult, error) {
	indexModel := []mongo.IndexModel{
		{
			Keys:    bson.M{"bank_id": 1}, // 1 for ascending, -1 for descending
			Options: options.Index().SetUnique(true),
		},
	}
	_, err := b.coll.Indexes().CreateMany(b.ctx, indexModel)
	if err != nil {
		return nil, err
	}
	result, err := b.coll.InsertOne(b.ctx, bank)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *Bank1) GetBankById(id int64) (*models.Bank, error) {
	filter := bson.M{"bank_id": id}
	var bank models.Bank
	err := b.coll.FindOne(b.ctx, filter).Decode(&bank)
	if err != nil {
		return nil, err
	}
	return &bank, nil
}

func (b *Bank1) UpdateBankById(id int64, banks *models.Bank) (*mongo.UpdateResult, error) {
	iv := bson.M{"bank_id": id}
	fv := bson.M{"$set": &banks}
	result, err := b.coll.UpdateOne(b.ctx, iv, fv)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *Bank1) DeleteBankById(id int64) (*mongo.DeleteResult, error) {
	filter := bson.M{"bank_id": id}
	result, err := b.coll.DeleteOne(b.ctx, filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *Bank1) GetAllCustomerBank() ([]primitive.M, error) {
	pipeline := []bson.M{
		{
			"$lookup": bson.M{
				"from":         "customer",
				"localField":   "bank_id",
				"foreignField": "bank_id",
				"as":           "customers",
			},
		},
	}
	res, err := b.coll.Aggregate(b.ctx, pipeline)
	if err != nil {
		return nil, err
	}
	var results []bson.M
	if err := res.All(b.ctx, &results); err != nil {
		return nil, err
	}
	fmt.Println(results)
	return results, nil
}

func (b *Bank1) GetAllBankTransaction(id int64) ([]interface{}, error) {
	pipeline := []bson.M{
		{
			"$match": bson.M{"bank_id": id},
		},
		{"$lookup": bson.M{
			"from":         "customer",
			"localField":   "bank_id",
			"foreignField": "bank_id",
			"as":           "transactionsBank",
		},
		},
	}
	var bank []bson.M
	res, err := b.coll.Aggregate(b.ctx, pipeline)
	if err != nil {
		return nil, err
	}
	if err := res.All(b.ctx, &bank); err != nil {
		return nil, err
	}
	var p []interface{}
	for _, v := range bank {
		for _, v1 := range v["transactionsBank"].(primitive.A) {
			p = append(p, v1.(primitive.M)["transaction"])
		}
	}
	return p, nil
}

func (b *Bank1) GetAllBankTransDate(id int64, date1 string, date2 string) ([]interface{}, error) {
	layout := "2006-01-02 15:04:05.999999 -0700 MST"

	start,_ := time.Parse(layout, date1)
	end,_ := time.Parse(layout, date2)

	pipeline := []bson.M{
		{
			"$match": bson.M{"bank_id": id},
		},
		{"$lookup": bson.M{
			"from":         "customer",
			"localField":   "bank_id",
			"foreignField": "bank_id",
			"as":           "transactionsBank",
		},
		},
	}
	var bank []bson.M
	res, err := b.coll.Aggregate(b.ctx, pipeline)
	if err != nil {
		return nil, err
	}
	if err := res.All(b.ctx, &bank); err != nil {
		return nil, err
	}
	var x []interface{}
	var p []interface{}
	for _, v := range bank {
		for _, v1 := range v["transactionsBank"].(primitive.A) {
			p = append(p, v1.(primitive.M)["transaction"])
		}
	}
	for _, t := range p {
			t1 := t.(primitive.A)
			for i := 0; i < len(t1); i++ {
				date := t1[i].(primitive.M)["date"].(primitive.DateTime).Time()
				if date.After(start) && date.Before(end){
					
					x = append(x, t1[i])
				}
			}
		}
	return x, nil
}

