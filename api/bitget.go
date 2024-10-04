package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
)

type BitGetExchange struct {
	Endpoint string
}

type BitGetRes struct {
	Code        string     `json:"code"`
	Msg         string     `json:"msg"`
	RequestTime int64      `json:"requestTime"`
	Data        [][]string `json:"data"`
}

func (e *BitGetExchange) GetHistoryPrice(coin string, timestamp int64) (historyPrice HistoryPrice, err error) {
	timestamp = timestamp * 1000
	endTime := timestamp + 60000
	url := fmt.Sprintf("%s?symbol=%s&granularity=1min&endTime=%d&limit=1", BitgetEndpoint+"/api/v2/spot/market/history-candles", coin+"USDT", endTime)

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

	var bitGetRes BitGetRes
	err = json.Unmarshal(body, &bitGetRes)
	if err != nil || bitGetRes.Msg != "success" || len(bitGetRes.Data) == 0 {
		slog.Error("get history price from bitget error", "err", err, "msg", bitGetRes.Msg)
		return
	}

	historyPrice.Timestamp = bitGetRes.Data[0][0]
	historyPrice.Open = bitGetRes.Data[0][1]
	historyPrice.High = bitGetRes.Data[0][2]
	historyPrice.Low = bitGetRes.Data[0][3]
	historyPrice.Close = bitGetRes.Data[0][4]

	return
}
