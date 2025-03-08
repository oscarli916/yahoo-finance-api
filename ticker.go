package yahoofinanceapi

type Ticker struct {
	Symbol  string
	history *History
	option  *Option
}

func NewTicker(symbol string) *Ticker {
	h := NewHistory()
	o := NewOption()
	return &Ticker{Symbol: symbol, history: h, option: o}
}

func (t *Ticker) History(query HistoryQuery) map[string]PriceData {
	t.history.SetQuery(query)
	data := t.history.GetHistory(t.Symbol)
	return data
}

func (t *Ticker) OptionChain() OptionData {
	data := t.option.GetOptionChain(t.Symbol)
	return data
}
