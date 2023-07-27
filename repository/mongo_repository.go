package repository

import (
	"CurrencyExchangeService/logger"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ExchangeRate - Table Schma
type ExchangeRate struct {
	BaseCurrency    string             `bson:"bc,omitempty"`
	Rates           map[string]float64 `bson:"rates,omitempty"`
	InsertTimeStamp primitive.DateTime `bson:"ts,omitempty"`
}

func (e *ExchangeRate) ToString() string {
	return fmt.Sprintf("%+v", e)
}

// RangeRateRequest -
type RangeRateRequest struct {
	StartTime primitive.DateTime
	EndTime   primitive.DateTime
}

// ExchangeRateRepository - Abstract Interface
type ExchangeRateRepository interface {
	GetLatestRates() ExchangeRate
	GetRangeRates(RangeRateRequest) []ExchangeRate
	AddRates(ExchangeRate) string
}

// MongoDB -
type MongoDB struct {
	mongoClient *mongo.Client
	dbLogger    *logger.ServiceLogger
}

func NewMongoDBRepository(
	logger *logger.ServiceLogger, dbClient *mongo.Client,
) ExchangeRateRepository {
	//Google: best practice to implementa an interface in Go
	var mongoImpl ExchangeRateRepository

	mongoImpl = &MongoDB{
		mongoClient: dbClient,
		dbLogger:    logger,
	}
	
	// coupled
	return mongoImpl
}

func (db *MongoDB) getCollection() *mongo.Collection {
	collection := db.mongoClient.
		Database("exchange_rate_service").
		Collection("exchange_rates")
	return collection
}

func (db *MongoDB) GetLatestRates() ExchangeRate {
	filter := bson.M{}
	result := ExchangeRate{}
	// find options, timestamp descending
	findOptions := &options.FindOneOptions{
		Sort: bson.D{{"ts", -1}},
	}

	coll := db.getCollection()
	err := coll.FindOne(context.Background(), filter, findOptions).Decode(&result)
	if err != nil {
		db.dbLogger.Fatal(fmt.Sprintf("error GetLatestRates: %s", err.Error()))
		return ExchangeRate{}
	}
	return result
}

func (db *MongoDB) GetRangeRates(
	request RangeRateRequest,
) []ExchangeRate {
	filter := bson.D{
		{Key: "ts", Value: bson.D{{Key: "$gte", Value: request.StartTime}}},
		{Key: "ts", Value: bson.D{{Key: "$lte", Value: request.EndTime}}},
	}
	result := []ExchangeRate{}

	cur, err := db.getCollection().Find(context.Background(), filter)
	if err != nil {
		db.dbLogger.Fatal(fmt.Sprintf("error GetRangeRates: %s", err.Error()))
		return result
	}

	if err := cur.All(context.Background(), &result); err != nil {
		db.dbLogger.Fatal(fmt.Sprintf("error GetRangeRates.DeconstructData: %s", err.Error()))
	}
	return result

}

func (db *MongoDB) AddRates(rate ExchangeRate) string {
	coll := db.getCollection()
	result, err := coll.InsertOne(context.Background(), rate)
	if err != nil {
		db.dbLogger.Fatal(fmt.Sprintf("error insert data: %s", err.Error()))
	}
	return string(fmt.Sprintf("%v", result.InsertedID))
}
