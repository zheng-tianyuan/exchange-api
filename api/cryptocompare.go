package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"strconv"
)

type CompareExchange struct {
	Endpoint string
}

type CompareHistoryRes struct {
	Response   string `json:"Response"`
	Message    string `json:"Message"`
	HasWarning bool   `json:"HasWarning"`
	Type       int    `json:"Type"`
	RateLimit  struct {
	} `json:"RateLimit"`
	Data struct {
		Aggregated bool `json:"Aggregated"`
		TimeFrom   int  `json:"TimeFrom"`
		TimeTo     int  `json:"TimeTo"`
		Data       []struct {
			Time             int     `json:"time"`
			High             float64 `json:"high"`
			Low              float64 `json:"low"`
			Open             float64 `json:"open"`
			Volumefrom       float64 `json:"volumefrom"`
			Volumeto         float64 `json:"volumeto"`
			Close            float64 `json:"close"`
			ConversionType   string  `json:"conversionType"`
			ConversionSymbol string  `json:"conversionSymbol"`
		} `json:"Data"`
	} `json:"Data"`
}

func (e *CompareExchange) GetHistoryPrice(coin string, timestamp int64) (historyPrice HistoryPrice, err error) {
	url := fmt.Sprintf("%s&fsym=%s&tsym=USD&limit=1&toTs=%d", e.Endpoint+"/data/v2/histoday", coin, timestamp)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var compareRes CompareHistoryRes
	err = json.Unmarshal(body, &compareRes)
	if err != nil || compareRes.Response != "Success" {
		slog.Error("get history price from bybit error", "err", err)
		return
	}

	historyPrice.Timestamp = strconv.Itoa(compareRes.Data.Data[0].Time)
	historyPrice.Open = fmt.Sprintf("%.8f", compareRes.Data.Data[0].Open)
	historyPrice.Close = fmt.Sprintf("%.8f", compareRes.Data.Data[0].Close)
	historyPrice.Low = fmt.Sprintf("%.8f", compareRes.Data.Data[0].Low)
	historyPrice.High = fmt.Sprintf("%.8f", compareRes.Data.Data[0].High)

	return
}
