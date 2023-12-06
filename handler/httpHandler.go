package handler

import (
	"CurrencyExchangeService/config"
	"CurrencyExchangeService/controller"
	"CurrencyExchangeService/logger"
	"CurrencyExchangeService/repository"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
	"time"
)

// Application Define an application struct to hold the dependencies for our HTTP handlers, helpers,
// and middleware. At the moment this only contains a copy of the config struct and a
// logger, but it will grow to include a lot more as our build progresses.
type Application struct {
	Config     config.Config
	Logger     *logger.ServiceLogger
	Controller *controller.ExchangeRateController
}

func (app *Application) Routes() *httprouter.Router {
	// Initialize a new httprouter router instance.
	router := httprouter.New()

	// Register the relevant methods, URL patterns and handler functions for our
	// endpoints using the HandlerFunc() method. Note that http.MethodGet and
	// http.MethodPost are constants which equate to the strings "GET" and "POST"
	// respectively.
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/latestRate", app.getExchangeRateHandler)
	router.HandlerFunc(http.MethodGet, "/v1/rangeRate", app.getRangeRateHandler)

	// Return the httprouter instance.
	return router
}

// healthcheckHandler :Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *Application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	app.Logger.Info("Someone is calling the service")
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.Config.Env)
	fmt.Fprintf(w, "version: %s\n", app.Config.AppVersion)
}

// getExchangeRateHandler :Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *Application) getExchangeRateHandler(w http.ResponseWriter, r *http.Request) {
	app.Logger.Info("Someone is calling the server.getExchangeRateHandler")
	//var response = "{\"base_currency\": \"USD\", \"rates\": {\"CAD\": 1.3}}"

	respStruct := app.Controller.GetLatestRate()
	response := respStruct.ToString()
	w.Header().Set("Content-Type", "application/json")

	//fmt.Fprintln(w, response)

	b, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	// Write the JSON as the HTTP response body.
	fmt.Fprintln(w, string(b))
}

type httpRangeInput struct {
	StarDate string `json:"startDate"`
	EndDate  string `json:"endDate"`
}

// curl -d '{"startDate":"2023-06-01","endDate":"2023-07-22"}' -X GET http://localhost:8080/v1/rangeRate -H "Content-Type: application/json"
func (app *Application) getRangeRateHandler(w http.ResponseWriter, r *http.Request) {

	var inputVar httpRangeInput
	app.Logger.Info(fmt.Sprintf(" HTTP Raw Request %+v", r))
	b, err := io.ReadAll(r.Body)
	if err != nil {
		app.Logger.Info("readAll" + err.Error())
	}

	err = json.Unmarshal(b, &inputVar)
	app.Logger.Info(fmt.Sprintf(" HTTP Formatted Input %+v", inputVar))
	if err != nil {
		app.Logger.Info("Unmarshal" + err.Error())
	}

	app.Logger.Info(fmt.Sprintf("%+v", inputVar))

	startTime, _ := time.Parse("2006-01-02", inputVar.StarDate)
	endTime, _ := time.Parse("2006-01-02", inputVar.EndDate)
	input := repository.RangeRateRequest{
		StartTime: primitive.DateTime(startTime.UnixMilli()),
		EndTime:   primitive.DateTime(endTime.UnixMilli()),
	}

	app.Logger.Info(fmt.Sprintf("%+v", input))
	respStruct := app.Controller.GetRangeRates(input)

	//var jsonResult []httpRangeResponse
	//for _, resp := range respStruct {
	//	app.Logger.Info(resp.InsertTimeStamp.Time().String())
	//	jr := httpRangeResponse{
	//		BaseCurrency:    resp.BaseCurrency,
	//		Rates:           resp.Rates,
	//		InsertTimeStamp: time.Unix(int64(resp.InsertTimeStamp), 0).Format("2006-01-02"),
	//	}
	//
	//	jsonResult = append(jsonResult, jr)
	//}

	b, err = json.Marshal(respStruct)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	// Write the JSON as the HTTP response body.
	fmt.Fprintln(w, string(b))
}
