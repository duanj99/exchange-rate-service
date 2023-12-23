package scheduler_test

import (
	"CurrencyExchangeService/logger"
	mock "CurrencyExchangeService/mock"
	"CurrencyExchangeService/scheduler"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGateway(t *testing.T) {

	// mock
	sysLogger := logger.NewLogger()
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()
	database := mock.NewMockExchangeRateRepository(mockCtrl)

	// run
	// expect database.AddRates to be called 2 times
	// all database.AddRates call returns string "abcd"

	// database repository addRates function expected to be call twice
	// if the addRates get called either less or greater than twice, it throws error

	// scheduler(database Repo, svcLogger)
	// instead of passing in real DB Repo, in test, we passed in the Mock Repo.

	// Whenever mock repo.addRates function is reached, it mock the function is executed and
	// return 'abcd'
	database.EXPECT().AddRates(gomock.Any()).Return("abcd").Times(2)
	scheduler.NewScheduler(database, sysLogger, true)

}

// go get github.com/go-co-op/gocron
// go install github.com/golang/mock/mockgen@v1.5.0

// mockgen -source=repository/mongo_repository.go -destination=mock/repository/mongo_repository.go -package=mock
//mockgen -source=gateways/exchange_rate_gateway.go -destination=mock/gateways/exchange_rate_gateway.go -package=mock
//
//mockgen -source=repository/cache_repository.go -destination=mock/repository/cache_repository.go -package=mock

//mockgen -source=controller/exchange_rate_controller.go -destination=mock/controller/exchange_rate_controller.go -package=mock
//exchange_rate_controller.go
