package repository_test

import (
	"CurrencyExchangeService/clients"
	"CurrencyExchangeService/logger"
	"CurrencyExchangeService/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestCacheRunning(t *testing.T) {
	// init repository
	mockLogger := logger.NewLogger()
	mockDB := clients.NewMongoDBClient(mockLogger)
	mockRepository := repository.NewMongoDBRepository(mockLogger, mockDB)
	testCache := repository.NewExchangeRateCacheRepository(
		mockLogger,
		mockRepository,
	)

	// action
	rate := testCache.GetLatestRates()
	mockLogger.Info(rate.ToString())

	mockRepository.AddRates(
		repository.ExchangeRate{
			BaseCurrency:    "USD",
			Rates:           map[string]float64{"CAD": 1.3, "EUR": 0.9, "KPP": 121.1},
			InsertTimeStamp: primitive.DateTime(time.Now().Unix()),
		},
	)

	time.Sleep(10 * time.Second)
	testCache.StopCache()
	newRate := testCache.GetLatestRates()
	mockLogger.Info(newRate.ToString())
	time.Sleep(5 * time.Second)
}
