package controller_test

import (
	"CurrencyExchangeService/controller"
	"CurrencyExchangeService/logger"
	mock "CurrencyExchangeService/mock"
	"CurrencyExchangeService/repository"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExchangeRateController_GetLatestRate_CacheError(t *testing.T) {
	var svcLogger = logger.NewLogger()
	//var mongoClient = clients.NewMongoDBClient(svcLogger)
	mockCtrl := gomock.NewController(t)
	var database = mock.NewMockExchangeRateRepository(mockCtrl)

	//var cache = repository.NewExchangeRateCacheRepository(
	//	svcLogger,
	//	database,
	//)
	var cache = mock.NewMockExchangeRateCacheRepository(mockCtrl)

	var svcController = controller.NewController(
		svcLogger,
		cache,
		database,
	)

	// Controller - getLatestRate()
	// Step 1: Call Repo Mongo database:
	//			mock repo.getLastest() -> "ab"
	// Step 2: Call Cache Repository
	//			mock repo.cacheGetLaster() -> "cd"
	//			mock repo.cacheGetLaster() -> throw error
	// Step 3: if 1, 2 failed, threw error
	// 			if 1, 2, succeed, return "abcd"

	mockCacheReturn := repository.ExchangeRate{
		BaseCurrency: "EUR",
	}

	mockDBReturn := repository.ExchangeRate{
		BaseCurrency: "USD",
		Rates:        map[string]float64{"EUR": 0.9},
	}

	cache.EXPECT().GetLatestRates().Return(mockCacheReturn, errors.New("cache is Empty")).Times(1)
	database.EXPECT().GetLatestRates().Return(mockDBReturn).Times(1)
	// Unit Test: Arrange

	// Take Action
	latestRate := svcController.GetLatestRate()
	svcLogger.Info("read latest rate: " + latestRate.ToString())
	//time.Sleep(10 * time.Second)

	// latestRate.BaseCurrency : Actual Value
	// "USD": Expected Value
	assert.Equal(t, "USD", latestRate.BaseCurrency)
}

func TestExchangeRateController_GetLatestRate_CacheNoError(t *testing.T) {
	var svcLogger = logger.NewLogger()
	//var mongoClient = clients.NewMongoDBClient(svcLogger)
	mockCtrl := gomock.NewController(t)
	var database = mock.NewMockExchangeRateRepository(mockCtrl)

	//var cache = repository.NewExchangeRateCacheRepository(
	//	svcLogger,
	//	database,
	//)
	var cache = mock.NewMockExchangeRateCacheRepository(mockCtrl)

	var svcController = controller.NewController(
		svcLogger,
		cache,
		database,
	)

	mockCacheReturn := repository.ExchangeRate{
		BaseCurrency: "EUR",
	}

	mockDBReturn := repository.ExchangeRate{
		BaseCurrency: "USD",
		Rates:        map[string]float64{"EUR": 0.9},
	}

	// Unit Test: Arrange
	cache.EXPECT().GetLatestRates().Return(mockCacheReturn, nil).Times(1)
	database.EXPECT().GetLatestRates().Return(mockDBReturn).Times(0)

	// Take Action
	latestRate := svcController.GetLatestRate()
	svcLogger.Info("read latest rate: " + latestRate.ToString())
	//time.Sleep(10 * time.Second)

	// latestRate.BaseCurrency : Actual Value
	// "USD": Expected Value
	assert.Equal(t, "EUR", latestRate.BaseCurrency)
}
