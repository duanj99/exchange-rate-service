package scheduler_test

import (
	"CurrencyExchangeService/logger"
	mocks "CurrencyExchangeService/mocks/repository"
	"CurrencyExchangeService/scheduler"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGateway(t *testing.T) {

	// mock
	sysLogger := logger.NewLogger()
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()
	database := mocks.NewMockExchangeRateRepository(mockCtrl)

	// run
	// expect database.AddRates to be called 2 times
	// all database.AddRates call returns string "abcd"
	database.EXPECT().AddRates(gomock.Any()).Return("abcd").Times(2)
	scheduler.NewScheduler(database, sysLogger, true)

}

// go get github.com/go-co-op/gocron
// go install github.com/golang/mock/mockgen@v1.5.0

// mockgen -source=repository/mongo_repository.go -destination=mocks/repository/mongo_repository.go -package=mocks
//mockgen -source=gateways/exchange_rate_gateway.go -destination=mocks/gateways/exchange_rate_gateway.go -package=mocks
//
//mockgen -source=repository/cache_repository.go -destination=mocks/repository/cache_repository.go -package=mocks

//mockgen -source=controller/exchange_rate_controller.go -destination=mocks/controller/exchange_rate_controller.go -package=mocks
//exchange_rate_controller.go
