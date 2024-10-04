package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
)

type BybitExchange struct {
	Endpoint string
}

type BybitHistoryRes struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Symbol   string     `json:"symbol"`
		Category string     `json:"category"`
		List     [][]string `json:"list"`
	} `json:"result"`
	RetExtInfo struct {
	} `json:"retExtInfo"`
	Time int64 `json:"time"`
}

func (e *BybitExchange) GetHistoryPrice(coin string, timestamp int64) (historyPrice HistoryPrice, err error) {
	url := fmt.Sprintf("%s&symbol=%s&interval=1&start=%d&limit=1", BybitEndpoint+"/v5/market/kline?category=inverse", coin+"USD", timestamp*1000)

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

	var byteRes BybitHistoryRes
	err = json.Unmarshal(body, &byteRes)
	if err != nil || byteRes.RetCode != 0 || len(byteRes.Result.List) == 0 || byteRes.RetMsg != "OK" {
		slog.Error("get history price from bybit error", "err", err, "msg", byteRes.RetMsg)
		return
	}

	historyPrice.Timestamp = byteRes.Result.List[0][0]
	historyPrice.Open = byteRes.Result.List[0][1]
	historyPrice.High = byteRes.Result.List[0][2]
	historyPrice.Low = byteRes.Result.List[0][3]
	historyPrice.Close = byteRes.Result.List[0][4]

	return
}
