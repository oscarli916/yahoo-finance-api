package yahoofinanceapi

type Ticker struct {
	Symbol  string
	history History
}

func NewTicker(symbol string) *Ticker {
	return &Ticker{Symbol: symbol, history: History{}}
}

func (t *Ticker) History(query HistoryQuery) map[string]PriceData {
	t.history.SetQuery(query)
	data := t.history.GetHistory(t.Symbol)
	return data
}
