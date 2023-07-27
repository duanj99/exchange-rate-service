package repository_test

import (
	"CurrencyExchangeService/clients"
	"CurrencyExchangeService/logger"
	"CurrencyExchangeService/repository"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestInsertDataMongoDB(t *testing.T) {
	// init repository
	mockLogger := logger.NewLogger()
	mockDB := clients.NewMongoDBClient(mockLogger)
	mockRepository := repository.NewMongoDBRepository(mockLogger, mockDB)

	// mock date
	insertData := repository.ExchangeRate{
		BaseCurrency:    "USD",
		Rates:           map[string]float64{"CAD": 1.3, "EUR": 0.9, "KPP": 100.1},
		InsertTimeStamp: primitive.DateTime(time.Now().Unix()),
	}

	// action
	insertID := mockRepository.AddRates(insertData)
	mockLogger.Info(fmt.Sprintf("inserted ID + %s", insertID))
	_ = mockDB.Disconnect(context.Background())
}

func TestReadMongoDB(t *testing.T) {
	// init the mock repository
	mockLogger := logger.NewLogger()
	mockDB := clients.NewMongoDBClient(mockLogger)
	mockRepository := repository.NewMongoDBRepository(mockLogger, mockDB)

	// action
	result := mockRepository.GetLatestRates()

	// assertion
	//result.BaseCurrency != "USD" {}
	mockLogger.Info(result.BaseCurrency)
	mockLogger.Info(fmt.Sprintf("Rates: %v", result.Rates))
	_ = mockDB.Disconnect(context.Background())
}

func TestReadRangeMongoDB(t *testing.T) {
	// init Test repository
	mockLogger := logger.NewLogger()
	mockDB := clients.NewMongoDBClient(mockLogger)

	mockRepository := repository.NewMongoDBRepository(mockLogger, mockDB)

	// mock data
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -14)

	request := repository.RangeRateRequest{
		StartTime: primitive.DateTime(startDate.Unix()),
		EndTime:   primitive.DateTime(endDate.Unix()),
	}

	// action
	result := mockRepository.GetRangeRates(request)

	// assert
	mockLogger.Info(result[0].BaseCurrency)
	mockLogger.Info(fmt.Sprintf("Rates: %v", result[0].Rates))
	mockLogger.Info(fmt.Sprintf("Number of result: %d", len(result)))

	_ = mockDB.Disconnect(context.Background())
}
