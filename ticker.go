package yahoofinanceapi

import (
	"fmt"
	"io"
	"log/slog"
	"net/url"
)

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

func (t *Ticker) OptionChain() {
	t.option.GetOptionChain(t.Symbol)
}

type Option struct {
	client *Client
}

func NewOption() *Option {
	return &Option{client: GetClient()}
}

func (o *Option) GetOptionChain(symbol string) {
	endpoint := fmt.Sprintf("%s/v7/finance/options/%s", BASE_URL, symbol)
	resp, err := o.client.Get(endpoint, url.Values{})
	if err != nil {
		slog.Error("Failed to get option chain", "err", err)
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Failed to read response body", "err", err)
		return
	}

	fmt.Println(string(data))
}
