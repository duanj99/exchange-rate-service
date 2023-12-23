package controller

import (
	"CurrencyExchangeService/logger"
	"CurrencyExchangeService/repository"
)

type ExchangeRateController struct {
	logger   *logger.ServiceLogger
	cache    repository.ExchangeRateCacheRepository
	database repository.ExchangeRateRepository
}

func NewController(
	dbLogger *logger.ServiceLogger,
	rateCache repository.ExchangeRateCacheRepository,
	database repository.ExchangeRateRepository,
) *ExchangeRateController {
	dbLogger.Info("Initialize Controller")
	return &ExchangeRateController{
		logger:   dbLogger,
		cache:    rateCache,
		database: database,
	}
}

func (c *ExchangeRateController) GetLatestRate() repository.ExchangeRate {
	c.logger.Info("read data from GetLatestRate Controller")

	result, err := c.cache.GetLatestRates()

	// Return a default Value if error
	if err != nil {
		result = c.database.GetLatestRates()
	}

	return result
}

func (c *ExchangeRateController) GetRangeRates(request repository.RangeRateRequest) []repository.ExchangeRate {
	c.logger.Info("read data from GetRangeRates Controller")
	return c.database.GetRangeRates(request)
}

// root: main.go
//	-> app.Route()
//		-> app.getRangeRateHandler() decode request from HTTP, call controller, and encode response
//			-> controller(): control business logic
//				-> repository(): control data access.
