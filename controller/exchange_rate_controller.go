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
