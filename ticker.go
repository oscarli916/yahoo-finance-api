package yahoofinanceapi

import "sort"

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

// get the latest Quote from the History, fetches new history in the process
func (t *Ticker) Quote() (PriceData, error) {
	history, err := t.history.GetHistory(t.Symbol)
	if err != nil {
		return PriceData{}, err
	}
	transformedData := t.history.transformData(history)

	dates := make([]string, 0, len(transformedData))
	for date := range transformedData {
		dates = append(dates, date)
	}
	sort.Strings(dates)

	latestDate := dates[len(dates)-1]
	latestPriceData := transformedData[latestDate]
	return latestPriceData, nil
}

func (t *Ticker) History(query HistoryQuery) (map[string]PriceData, error) {
	t.history.SetQuery(query)
	history, err := t.history.GetHistory(t.Symbol)
	if err != nil {
		return nil, err
	}
	return t.history.transformData(history), nil
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
