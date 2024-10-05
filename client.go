package exchange_api

import (
	"exchange_api/api"
)

type ExchangeType int

const (
	Binance ExchangeType = iota
	Okex
	Gate
	Bitget
	Bybit
	Compare
)

const (
	BinanceEndpoint string = "https://api.binance.com"
	OkexEndpoint    string = "https://www.okx.com"
	GateEndpoint    string = "https://api.gateio.ws"
	BitgetEndpoint  string = "https://api.bitget.com"
	BybitEndpoint   string = "https://api.bybit.com"
	CompareEndpoint string = "https://min-api.cryptocompare.com"
)

type Exchange interface {
	GetHistoryPrice(coin string, timestamp int64) (api.HistoryPrice, error)
}

func NewExchange(exchangeType ExchangeType) Exchange {
	switch exchangeType {
	case Binance:
		return &api.BinanceExchange{
			Endpoint: BinanceEndpoint,
		}
	case Okex:
		return &api.OkexExchange{
			Endpoint: OkexEndpoint,
		}
	case Bitget:
		return &api.BitGetExchange{
			Endpoint: BitgetEndpoint,
		}
	case Bybit:
		return &api.BybitExchange{
			Endpoint: BybitEndpoint,
		}
	case Gate:
		return &api.GateExchange{
			Endpoint: GateEndpoint,
		}
	case Compare:
		return &api.CompareExchange{
			Endpoint: CompareEndpoint,
		}

	default:
		return nil
	}
}
