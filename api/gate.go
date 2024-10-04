package api

import (
	"encoding/json"
	"github.com/shopspring/decimal"

	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
)

type GateExchange struct {
	Endpoint string
}

func (e *GateExchange) GetHistoryPrice(coin string, timestamp int64) (historyPrice HistoryPrice, err error) {
	url := fmt.Sprintf("%s?currency_pair=%s&resolution=1&from=%d", GateEndpoint+"/api/v4/spot/candlesticks", coin+"_USDT", timestamp)

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

	var gateRes [][]string
	err = json.Unmarshal(body, &gateRes)
	if err != nil || len(gateRes) == 0 {
		slog.Error("get history price from gate error", "err", err)
		return
	}

	timeSValue, err := decimal.NewFromString(gateRes[0][0])
	if err != nil {
		return
	}

	historyPrice.Timestamp = timeSValue.Mul(decimal.NewFromInt(1000)).String()
	historyPrice.Open = gateRes[0][5]
	historyPrice.High = gateRes[0][3]
	historyPrice.Low = gateRes[0][4]
	historyPrice.Close = gateRes[0][2]

	return
}
