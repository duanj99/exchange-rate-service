package repository

import (
	"CurrencyExchangeService/logger"
	"sync"
	"time"
)

type ExchangeRateCache struct {
	Rate        ExchangeRate
	mongoDBRepo ExchangeRateRepository
	dbLogger    *logger.ServiceLogger

	Stop chan bool
	mu   sync.RWMutex //A RWMutex is a reader/writer mutual exclusion lock.
}

func NewExchangeRateCacheRepository(
	logger *logger.ServiceLogger,
	mongoDBRepo ExchangeRateRepository,
) *ExchangeRateCache {
	newCache := &ExchangeRateCache{
		Rate:        mongoDBRepo.GetLatestRates(),
		mongoDBRepo: mongoDBRepo,
		dbLogger:    logger,
		Stop:        make(chan bool),
	}

	go newCache.updateCachePeriodically(2 * time.Minute)

	return newCache
}

// today, yesterday, the day before yesterday
// database write read lock
// dirty reads/ dirty writes

func (c *ExchangeRateCache) GetLatestRates() ExchangeRate {
	return c.Rate
}

func (c *ExchangeRateCache) AddRates(input ExchangeRate) interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.dbLogger.Info("Refresh Data in Cache")
	c.Rate = input
	return c.Rate
}

func (c *ExchangeRateCache) updateCachePeriodically(interval time.Duration) {
	for {
		select {
		case <-c.Stop:
			c.dbLogger.Info("received stop signal from Chain, Stop updateCachePeriodically")
			return
		default:
			latestRate := c.mongoDBRepo.GetLatestRates()
			c.AddRates(latestRate)
		}
		time.Sleep(interval)
	}
}

func (c *ExchangeRateCache) StopCache() {
	c.dbLogger.Info("received stop signal")
	c.Stop <- true
	close(c.Stop)
}
