package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
)

type OkexExchange struct {
	Endpoint string
}

type OkexHistoryPriceRes struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data [][]string `json:"data"`
}

func (o *OkexExchange) GetHistoryPrice(coin string, timestamp int64) (historyPrice HistoryPrice, err error) {
	url := fmt.Sprintf("%s?instId=%s&limit=1&after=%d", OkexEndpoint+"/api/v5/market/history-index-candles", coin+"-USDT", timestamp*1000)
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

	var okexRes OkexHistoryPriceRes
	err = json.Unmarshal(body, &okexRes)
	if err != nil || len(okexRes.Data) == 0 {
		slog.Error("get history price from okex error", "err", err, "msg", okexRes.Msg)
		return
	}

	historyPrice.Timestamp = okexRes.Data[0][0]
	historyPrice.Open = okexRes.Data[0][1]
	historyPrice.High = okexRes.Data[0][2]
	historyPrice.Low = okexRes.Data[0][3]
	historyPrice.Close = okexRes.Data[0][4]

	return
}
