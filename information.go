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

// YahooTickerInfo --> Struct to hold key metadata about the ticker
type YahooTickerInfo struct {
	Symbol             string `json:"symbol"`
	ShortName          string `json:"shortName"`
	LongName           string `json:"longName"`
	Currency           string `json:"currency"`
	ExchangeName       string `json:"exchangeName"`
	MarketState        string `json:"marketState"`
	RegularMarketPrice struct {
		Raw float64 `json:"raw"`
		Fmt string  `json:"fmt"`
	} `json:"regularMarketPrice"`
}

// Information holds the HTTP client
type Information struct {
	client *Client
}

// newInformation initializes the Information struct with an HTTP client
func newInformation() *Information {
	return &Information{client: getClient()}
}

// GetTickerInfo fetches metadata information for a given ticker
func (i *Information) GetTickerInfo(symbol string) (YahooTickerInfo, error) {
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
