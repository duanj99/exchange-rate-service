package gateways

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//
//curl -X GET 'https://openexchangerates.org/api/latest.json?app_id=a3d29c1fdc8c4d97bc71264336f8efcd' \
//-H 'Content-Type: application/json'

type OpenExchangeRateResponse struct {
	Timestamp int64              `json:"timestamp"`
	Base      string             `json:"base"`
	Rates     map[string]float64 `json:"rates"`
}

/*
	curl -X GET 'https://openexchangerates.org/api/latest.json?app_id=a3d29c1fdc8c4d97bc71264336f8efcd' \
	    -H 'Content-Type: application/json'
*/
func GetOpenExchangeRate() (response OpenExchangeRateResponse) {
	requestURL := "https://openexchangerates.org/api/latest.json?app_id=a3d29c1fdc8c4d97bc71264336f8efcd"
	resp, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error read: %s\n", err)
	}

	err = json.Unmarshal(body, &response)

	return response
}
