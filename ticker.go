package yahoofinanceapi

import (
	"fmt"
	"sort"
)

type Ticker struct {
	Symbol      string
	history     *History
	option      *Option
	information *Information
	search      *Search
}

// NewTicker creates a new Ticker instance for the given symbol.
// It initializes the history, option, and information components needed to fetch
// historical price data, options data, and ticker information.
func NewTicker(symbol string) *Ticker {
	h := newHistory()
	o := newOption()
	i := newInformation()
	s := newSearch()
	return &Ticker{Symbol: symbol, history: h, option: o, information: i, search: s}
}

// Quote returns the latest PriceData for the Ticker's symbol.
// This is a convenience wrapper around the History function. It fetches the historical
// price data for the symbol, sorts the dates, and returns the most recent entry.
// If you need more control or access to the full historical data, use the History method directly.
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

// Info retrieves the ticker information for the Ticker's symbol.
// It returns a YahooTickerInfo struct containing metadata such as the symbol, name, currency, and market state.
// If no information is found, it returns an error.
func (t *Ticker) Info() (YahooTickerInfo, error) {
	info, err := t.information.GetInfo(t.Symbol)
	if err != nil {
		return YahooTickerInfo{}, err
	}

	if info.Symbol != t.Symbol {
		return YahooTickerInfo{}, fmt.Errorf("symbol mismatch: expected %s, got %s", t.Symbol, info.Symbol)
	}
	return info, nil
}

// History retrieves the historical price data for the Ticker's symbol based on the provided query.
// It returns a map of date strings to PriceData structs.
// The query can specify the range, interval, and other parameters for the historical data.
func (t *Ticker) History(query HistoryQuery) (map[string]PriceData, error) {
	t.history.SetQuery(query)
	history, err := t.history.GetHistory(t.Symbol)
	if err != nil {
		return nil, err
	}
	return t.history.transformData(history), nil
}

// OptionChain retrieves the option chain for the Ticker's symbol.
// It returns an OptionData struct containing the options available for the ticker.
// If no options are found, it returns an empty OptionData struct.
func (t *Ticker) OptionChain() OptionData {
	optionChain := t.option.GetOptionChain(t.Symbol)
	return t.option.transformData(optionChain)
}

// OptionChainByExpiration retrieves the option chain for the Ticker's symbol filtered by a specific expiration date.
// It returns an OptionData struct containing the options available for the ticker on that expiration date.
// If no options are found for the specified expiration, it returns an empty OptionData struct.
func (t *Ticker) OptionChainByExpiration(expiration string) OptionData {
	optionChain := t.option.GetOptionChainByExpiration(t.Symbol, expiration)
	return t.option.transformData(optionChain)
}

// ExpirationDates retrieves a list of available expiration dates for options on the Ticker's symbol.
// It returns a slice of strings representing the expiration dates.
func (t *Ticker) ExpirationDates() []string {
	expirationDates := t.option.GetExpirationDates(t.Symbol)
	return expirationDates
}

// Search searches for investment symbols by query using Yahoo Finance's public search API.
// This is a convenience function that uses default search parameters.
//
// Parameters:
//   - query: Search query (symbol, company name, etc.)
//   - limit: Maximum number of results to return (max 20, default 10)
//
// Returns:
//   - []SearchResult: List of search results
//   - error: Error if request fails or query is invalid
func (t *Ticker) Search(query string, limit int) ([]SearchResult, error) {
	searchResponse, err := t.search.GetSearchResults(query, limit)
	if err != nil {
		return nil, err
	}
	data := t.search.transformData(searchResponse)
	return data.Results, nil
}

// SearchWithOptions searches for investment symbols using custom parameters.
//
// Parameters:
//   - params: SearchParams with custom configuration
//
// Returns:
//   - []SearchResult: List of search results
//   - error: Error if request fails or query is invalid
func (t *Ticker) SearchWithOptions(params SearchParams) ([]SearchResult, error) {
	searchResponse, err := t.search.GetSearchResultsWithOptions(params)
	if err != nil {
		return nil, err
	}
	data := t.search.transformData(searchResponse)
	return data.Results, nil
}
