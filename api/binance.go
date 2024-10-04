package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
)

type BinanceExchange struct {
	Endpoint string
}

func (e *BinanceExchange) GetHistoryPrice(coin string, timestamp int64) (historyPrice HistoryPrice, err error) {

	params := fmt.Sprintf("?symbol=%s&interval=%s&limit=%d&startTime=%d", coin+"USDT", "1m", 1, timestamp*1000)
	resp, err := http.Get(e.Endpoint + "/api/v3/klines" + params)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Error reading response body", "err", err)
		return
	}

	var k [][]interface{}
	err = json.Unmarshal(body, &k)
	if len(k) == 0 || err != nil {
		return
	}
	klineData := k[0]
	openTime := klineData[0].(float64)
	historyPrice.Timestamp = strconv.FormatInt(int64(openTime), 10)
	historyPrice.Open = klineData[1].(string)
	historyPrice.High = klineData[2].(string)
	historyPrice.Low = klineData[3].(string)
	historyPrice.Close = klineData[4].(string)

	return
}
