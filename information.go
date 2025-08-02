package yahoofinanceapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/url"
)

// YahooInfoResponse --> Struct to hold the result from the Yahoo Finance quoteSummary endpoint
type YahooInfoResponse struct {
	QuoteSummary struct {
		Result []struct {
			Price YahooTickerInfo `json:"price"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"quoteSummary"`
}

// PriceValue represents a price value with raw and formatted representations
type PriceValue struct {
	Raw     float64 `json:"raw"`
	Fmt     string  `json:"fmt"`
	LongFmt string  `json:"longFmt,omitempty"`
}

// YahooTickerInfo --> Struct to hold key metadata about the ticker
type YahooTickerInfo struct {
	MaxAge                     int         `json:"maxAge"`
	PreMarketChange            *PriceValue `json:"preMarketChange"`
	PreMarketPrice             *PriceValue `json:"preMarketPrice"`
	PreMarketSource            string      `json:"preMarketSource"`
	PostMarketChangePercent    *PriceValue `json:"postMarketChangePercent"`
	PostMarketChange           *PriceValue `json:"postMarketChange"`
	PostMarketTime             int64       `json:"postMarketTime"`
	PostMarketPrice            *PriceValue `json:"postMarketPrice"`
	PostMarketSource           string      `json:"postMarketSource"`
	RegularMarketChangePercent *PriceValue `json:"regularMarketChangePercent"`
	RegularMarketChange        *PriceValue `json:"regularMarketChange"`
	RegularMarketTime          int64       `json:"regularMarketTime"`
	PriceHint                  *PriceValue `json:"priceHint"`
	RegularMarketPrice         *PriceValue `json:"regularMarketPrice"`
	RegularMarketDayHigh       *PriceValue `json:"regularMarketDayHigh"`
	RegularMarketDayLow        *PriceValue `json:"regularMarketDayLow"`
	RegularMarketVolume        *PriceValue `json:"regularMarketVolume"`
	AverageDailyVolume10Day    *PriceValue `json:"averageDailyVolume10Day"`
	AverageDailyVolume3Month   *PriceValue `json:"averageDailyVolume3Month"`
	RegularMarketPreviousClose *PriceValue `json:"regularMarketPreviousClose"`
	RegularMarketSource        string      `json:"regularMarketSource"`
	RegularMarketOpen          *PriceValue `json:"regularMarketOpen"`
	StrikePrice                *PriceValue `json:"strikePrice"`
	OpenInterest               *PriceValue `json:"openInterest"`
	Exchange                   string      `json:"exchange"`
	ExchangeName               string      `json:"exchangeName"`
	ExchangeDataDelayedBy      int         `json:"exchangeDataDelayedBy"`
	MarketState                string      `json:"marketState"`
	QuoteType                  string      `json:"quoteType"`
	Symbol                     string      `json:"symbol"`
	UnderlyingSymbol           *string     `json:"underlyingSymbol"`
	ShortName                  string      `json:"shortName"`
	LongName                   string      `json:"longName"`
	Currency                   string      `json:"currency"`
	QuoteSourceName            string      `json:"quoteSourceName"`
	CurrencySymbol             string      `json:"currencySymbol"`
	FromCurrency               *string     `json:"fromCurrency"`
	ToCurrency                 *string     `json:"toCurrency"`
	LastMarket                 *string     `json:"lastMarket"`
	Volume24Hr                 *PriceValue `json:"volume24Hr"`
	VolumeAllCurrencies        *PriceValue `json:"volumeAllCurrencies"`
	CirculatingSupply          *PriceValue `json:"circulatingSupply"`
	MarketCap                  *PriceValue `json:"marketCap"`
}

// Information holds the HTTP client
type Information struct {
	client *Client
}

// newInformation initializes the Information struct with an HTTP client
func newInformation() *Information {
	return &Information{client: getClient()}
}

// GetInfo fetches metadata information for a given ticker
func (i *Information) GetInfo(symbol string) (YahooTickerInfo, error) {
	// Prepare URL parameters to request the "price" module
	params := url.Values{}
	params.Add("modules", "price")

	// Build the endpoint URL for the Yahoo Finance quoteSummary API
	endpoint := fmt.Sprintf("%s/v10/finance/quoteSummary/%s", BASE_URL, symbol)

	// Make the HTTP GET request using the client
	resp, err := i.client.Get(endpoint, params)
	if err != nil {
		slog.Error("Failed to get ticker info", "err", err)
		return YahooTickerInfo{}, err
	}

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return YahooTickerInfo{}, fmt.Errorf("failed to read response body: %w", err)
	}

	// Unmarshal the JSON response into the YahooInfoResponse struct
	var infoResponse YahooInfoResponse
	if err := json.Unmarshal(bodyBytes, &infoResponse); err != nil {
		return YahooTickerInfo{}, fmt.Errorf("failed to decode info JSON: %w", err)
	}

	// Check if the result array is empty
	if len(infoResponse.QuoteSummary.Result) == 0 {
		return YahooTickerInfo{}, fmt.Errorf("no info found for symbol: %s", symbol)
	}

	// Return the ticker price information
	return infoResponse.QuoteSummary.Result[0].Price, nil
}
