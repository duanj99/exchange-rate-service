package main

import (
	"CurrencyExchangeService/clients"
	"CurrencyExchangeService/config"
	"CurrencyExchangeService/controller"
	"CurrencyExchangeService/handler"
	"CurrencyExchangeService/logger"
	"CurrencyExchangeService/repository"
	"flag"
	"fmt"
	"net/http"

	"time"
)

func main() {
	// Declare an instance of the config struct.
	var cfg config.Config

	// Read the value of the port and env command-line flags into the config struct. We
	// default to using the port number 4000 and the environment "development" if no
	// corresponding flags are provided.
	flag.IntVar(&cfg.Port, "port", 8080, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	cfg.AppVersion = "1.00"

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

	//go scheduler.NewScheduler(database, svcLogger, false)
	// Start the gRPC Server
	svcLogger.Info(fmt.Sprintf("starting gRPC %s server on %s", cfg.Env, "8081"))
	go handler.InitGRPCHandler(svcLogger, svcController)

	// Declare an instance of the application struct, containing the config struct and
	// the logger.
	app := &handler.Application{
		Config:     cfg,
		Logger:     svcLogger,
		Controller: svcController,
	}

	// Declare an HTTP server with some sensible timeout settings, which listens on the
	// port provided in the config struct and uses the servemux we created above as the
	// handler.
	svcLogger.Info(fmt.Sprintf("localhost:%d", cfg.Port))
	srv := &http.Server{
		Addr:         fmt.Sprintf("localhost:%d", cfg.Port),
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start the HTTP server.a
	svcLogger.Info(fmt.Sprintf("starting %s server on %s", cfg.Env, srv.Addr))
	err := srv.ListenAndServe()
	svcLogger.Fatal(err.Error())
}
