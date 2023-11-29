package controller

import (
	"CurrencyExchangeService/logger"
	"CurrencyExchangeService/repository"
)

type ExchangeRateController struct {
	logger   *logger.ServiceLogger
	cache    *repository.ExchangeRateCache
	database repository.ExchangeRateRepository
}

func NewController(
	dbLogger *logger.ServiceLogger,
	rateCache *repository.ExchangeRateCache,
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
	return c.cache.GetLatestRates()
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
