package exchange_api

import (
	"testing"
)

func TestGetHistoryPrice(t *testing.T) {
	exchange := NewExchange(Binance)

	historyPrice, err := exchange.GetHistoryPrice("BTC", 1638048000)
	if err != nil {
		t.Fatalf("failed to get history price: %v", err)
	}

	// 使用 t.Logf 打印历史价格
	t.Logf("历史价格: %+v\n", historyPrice)
}
