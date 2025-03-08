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
	history := t.history.GetHistory(t.Symbol)
	return t.history.transformData(history)
}

func (t *Ticker) OptionChain() OptionData {
	optionChain := t.option.GetOptionChain(t.Symbol)
	return t.option.transformData(optionChain)
}

func (t *Ticker) OptionChainByExpiration(expiration string) OptionData {
	optionChain := t.option.GetOptionChainByExpiration(t.Symbol, expiration)
	return t.option.transformData(optionChain)
}

func (t *Ticker) ExpirationDates() []string {
	expirationDates := t.option.GetExpirationDates(t.Symbol)
	return expirationDates
}
