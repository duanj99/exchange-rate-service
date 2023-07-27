package scheduler

import (
	"CurrencyExchangeService/gateways"
	"CurrencyExchangeService/logger"
	"CurrencyExchangeService/repository"
	"fmt"
	"github.com/go-co-op/gocron"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func NewScheduler(
	database repository.ExchangeRateRepository,
	svcLogger *logger.ServiceLogger,
	isTest bool,
) {
	// -> scheduler -> call API -> read data -> data transformation -> update Repo
	s := gocron.NewScheduler(time.Local)
	job, err := s.Every(5).Seconds().Tag("UpdateRate").Do(UpdateDatabase, database, svcLogger)

	if err != nil {
		svcLogger.Fatal(fmt.Sprintf("error running script: %+v", err))
	}

	svcLogger.Info(fmt.Sprintf("job scheduled: %+v", job.Tags()))

	if isTest {
		svcLogger.Info("Scheduler blocking and run ")
		s.StartAsync()

		time.Sleep(8 * time.Second)
		svcLogger.Info("Shut down the Scheduler ")
		s.Stop()
	} else {
		s.StartAsync()
	}
}

func UpdateDatabase(
	database repository.ExchangeRateRepository, svcLogger *logger.ServiceLogger,
) {
	newRates := gateways.GetOpenExchangeRate()

	svcLogger.Info(fmt.Sprintf("get latest rate from OpenExchange %+v", newRates))

	exchangeRate := repository.ExchangeRate{
		BaseCurrency:    newRates.Base,
		Rates:           newRates.Rates,
		InsertTimeStamp: primitive.DateTime(newRates.Timestamp),
	}

	result := database.AddRates(exchangeRate)

	svcLogger.Info(fmt.Sprintf("Scheduler Update Database +%s", result))

}
