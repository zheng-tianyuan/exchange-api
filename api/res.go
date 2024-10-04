package api

type Exchange interface {
	GetHistoryPrice(coin string, timestamp int64) (HistoryPrice, error)
}

type HistoryPrice struct {
	Timestamp string `json:"timestamp"`
	Open      string `json:"open"`
	Close     string `json:"close"`
	High      string `json:"high"`
	Low       string `json:"low"`
}
