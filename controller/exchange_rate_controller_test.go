package controller_test

import (
	"CurrencyExchangeService/clients"
	"CurrencyExchangeService/controller"
	"CurrencyExchangeService/logger"
	"CurrencyExchangeService/repository"
	"testing"
	"time"
)

func TestExchangeRateController_GetLatestRate(t *testing.T) {
	var svcLogger = logger.NewLogger()
	var mongoClient = clients.NewMongoDBClient(svcLogger)
	var database = repository.NewMongoDBRepository(
		svcLogger,
		mongoClient,
	)

	var cache = repository.NewExchangeRateCacheRepository(
		svcLogger,
		database,
	)

	var svcController = controller.NewController(
		svcLogger,
		cache,
		database,
	)

	latestRate := svcController.GetLatestRate()
	svcLogger.Info("read latest rate: " + latestRate.ToString())
	time.Sleep(10 * time.Second)
}
