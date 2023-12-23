package repository

import (
	"CurrencyExchangeService/logger"
	"errors"
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

// ExchangeRateCacheRepository - Abstract Interface
type ExchangeRateCacheRepository interface {
	GetLatestRates() (ExchangeRate, error)
	AddRates(ExchangeRate) string
	StopCache()
}

func NewExchangeRateCacheRepository(
	logger *logger.ServiceLogger,
	mongoDBRepo ExchangeRateRepository,
) ExchangeRateCacheRepository {
	newCache := &ExchangeRateCache{
		Rate:        mongoDBRepo.GetLatestRates(),
		mongoDBRepo: mongoDBRepo,
		dbLogger:    logger,
		Stop:        make(chan bool),
	}

	go newCache.updateCachePeriodically(1 * time.Minute)

	return newCache
}

// today, yesterday, the day before yesterday
// database write read lock
// dirty reads/ dirty writes

func (c *ExchangeRateCache) GetLatestRates() (er ExchangeRate, err error) {
	c.dbLogger.Info("Someone reach Cache.GetLatestRates()")
	er = c.Rate
	if c.Rate.BaseCurrency != "USD" {
		return c.Rate, errors.New("cache is Empty")
	}

	return c.Rate, nil
}

func (c *ExchangeRateCache) AddRates(input ExchangeRate) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.dbLogger.Info("Refresh Data in Cache")
	c.Rate = input
	return c.Rate.ToString()
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
