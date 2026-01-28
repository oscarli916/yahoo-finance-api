package yahoofinanceapi

import (
	"strings"
	"testing"
)

func TestGetSearchResults_BasicSearch(t *testing.T) {
	// Skip in unit tests (requires network)
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	s := newSearch()
	results, err := s.GetSearchResults("AAPL", 10)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(results.Quotes) == 0 {
		t.Error("expected at least one result")
	}

	// Check first result contains AAPL
	found := false
	for _, r := range results.Quotes {
		if r.Symbol == "AAPL" {
			found = true
			if r.ShortName == "" && r.LongName == "" {
				t.Error("expected name to be set")
			}
		}
	}
	if !found {
		t.Error("expected to find AAPL in results")
	}
}

func TestGetSearchResults_EmptyQuery(t *testing.T) {
	s := newSearch()
	_, err := s.GetSearchResults("", 10)

	if err == nil {
		t.Error("expected error for empty query")
	}

	if !strings.Contains(err.Error(), "query cannot be empty") {
		t.Errorf("expected 'query cannot be empty' error, got: %v", err)
	}
}

func TestGetSearchResults_VietnameseStock(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	s := newSearch()
	results, err := s.GetSearchResults("VCB", 10)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(results.Quotes) == 0 {
		t.Error("expected at least one result for VCB")
	}
}

func TestGetSearchResults_Cryptocurrency(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	s := newSearch()
	results, err := s.GetSearchResults("Bitcoin", 10)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(results.Quotes) == 0 {
		t.Error("expected at least one result for Bitcoin")
	}

	// Check for BTC-USD symbol
	found := false
	for _, r := range results.Quotes {
		if r.Symbol == "BTC-USD" {
			found = true
			if r.QuoteType != "CRYPTOCURRENCY" {
				t.Errorf("expected type CRYPTOCURRENCY, got: %s", r.QuoteType)
			}
		}
	}
	if !found {
		t.Log("BTC-USD not found in results (this is acceptable)")
	}
}

func TestGetSearchResults_Limit(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	s := newSearch()
	limit := 5
	results, err := s.GetSearchResults("AAPL", limit)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(results.Quotes) > limit {
		t.Errorf("expected max %d results, got: %d", limit, len(results.Quotes))
	}
}

func TestGetSearchResults_MaxLimitEnforcement(t *testing.T) {
	s := newSearch()
	params := DefaultSearchParams("AAPL", 50) // Try to set limit above 20

	results, err := s.GetSearchResultsWithOptions(params)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	// Should cap at 20
	if len(results.Quotes) > 20 {
		t.Errorf("expected max 20 results, got: %d", len(results.Quotes))
	}
}

func TestGetSearchResults_DefaultLimit(t *testing.T) {
	s := newSearch()
	params := DefaultSearchParams("AAPL", 0) // Set invalid limit

	results, err := s.GetSearchResultsWithOptions(params)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	// Should default to 10
	if len(results.Quotes) > 10 {
		t.Errorf("expected max 10 results (default), got: %d", len(results.Quotes))
	}
}

func TestGetSearchResults_CompanyName(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	s := newSearch()
	results, err := s.GetSearchResults("Apple", 10)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(results.Quotes) == 0 {
		t.Error("expected at least one result for 'Apple'")
	}

	// Check for AAPL
	found := false
	for _, r := range results.Quotes {
		if r.Symbol == "AAPL" {
			found = true
			// Check name contains Apple
			name := r.ShortName
			if name == "" {
				name = r.LongName
			}
			if !strings.Contains(strings.ToLower(name), "apple") {
				t.Logf("Note: AAPL name doesn't contain 'apple': %s", name)
			}
		}
	}
	if !found {
		t.Error("expected to find AAPL in results for 'Apple' search")
	}
}

func TestDefaultSearchParams(t *testing.T) {
	params := DefaultSearchParams("test", 10)

	if params.Query != "test" {
		t.Errorf("expected query 'test', got: %s", params.Query)
	}

	if params.QuotesCount != 10 {
		t.Errorf("expected QuotesCount 10, got: %d", params.QuotesCount)
	}

	if params.NewsCount != 0 {
		t.Errorf("expected NewsCount 0, got: %d", params.NewsCount)
	}

	if params.EnableFuzzyQuery {
		t.Error("expected EnableFuzzyQuery to be false by default")
	}

	if !params.EnableEnhancedTrivialQuery {
		t.Error("expected EnableEnhancedTrivialQuery to be true by default")
	}
}

func TestBuildSearchURL(t *testing.T) {
	params := SearchParams{
		Query:        "AAPL",
		QuotesCount:  5,
		NewsCount:    0,
		ListsCount:   0,
		EnableFuzzyQuery: false,
		Lang:         "en-US",
	}

	url := buildSearchURL(params)

	if !strings.Contains(url, "q=AAPL") {
		t.Error("expected URL to contain query parameter")
	}

	if !strings.Contains(url, "quotesCount=5") {
		t.Error("expected URL to contain quotesCount parameter")
	}

	if !strings.Contains(url, "enableFuzzyQuery=false") {
		t.Error("expected URL to contain enableFuzzyQuery parameter")
	}

	if !strings.Contains(url, "lang=en-US") {
		t.Error("expected URL to contain lang parameter")
	}
}

func TestTransformSearchData(t *testing.T) {
	s := newSearch()

	yahooResp := YahooSearchResponse{
		Quotes: []struct {
			Symbol    string `json:"symbol"`
			ShortName string `json:"shortname"`
			LongName  string `json:"longname"`
			QuoteType string `json:"quoteType"`
			Exchange  string `json:"exchange"`
			ExchDisp  string `json:"exchDisp"`
			TypeDisp  string `json:"typeDisp"`
		}{
			{
				Symbol:    "AAPL",
				ShortName: "Apple Inc.",
				QuoteType: "EQUITY",
				Exchange:  "NAS",
				ExchDisp:  "NASDAQ",
			},
		},
	}

	data := s.transformData(yahooResp)

	if len(data.Results) != 1 {
		t.Errorf("expected 1 result, got: %d", len(data.Results))
	}

	if data.Results[0].Symbol != "AAPL" {
		t.Errorf("expected symbol AAPL, got: %s", data.Results[0].Symbol)
	}

	if data.Results[0].Name != "Apple Inc." {
		t.Errorf("expected name 'Apple Inc.', got: %s", data.Results[0].Name)
	}

	if data.Results[0].Type != "EQUITY" {
		t.Errorf("expected type EQUITY, got: %s", data.Results[0].Type)
	}
}

func TestGetSearchResults_ETF(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	s := newSearch()
	results, err := s.GetSearchResults("SPY", 10)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(results.Quotes) == 0 {
		t.Error("expected at least one result for SPY")
	}

	// Check for SPY ETF
	found := false
	for _, r := range results.Quotes {
		if r.Symbol == "SPY" {
			found = true
			if r.QuoteType != "ETF" {
				t.Logf("Note: SPY type is '%s' (expected ETF)", r.QuoteType)
			}
		}
	}
	if !found {
		t.Error("expected to find SPY in results")
	}
}

func TestGetSearchResults_CustomParams(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	s := newSearch()
	params := DefaultSearchParams("AAPL", 5)
	params.EnableFuzzyQuery = true
	params.Lang = "vi-VN"

	results, err := s.GetSearchResultsWithOptions(params)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(results.Quotes) == 0 {
		t.Error("expected at least one result with custom params")
	}
}

func TestGetSearchResults_TrailingWhitespaceQuery(t *testing.T) {
	s := newSearch()
	params := DefaultSearchParams("  AAPL  ", 5)

	results, err := s.GetSearchResultsWithOptions(params)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	// Query should be trimmed
	if len(results.Quotes) == 0 {
		t.Error("expected at least one result with trimmed query")
	}
}

func TestGetSearchResults_InvalidSearch(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	s := newSearch()
	// Use a very unlikely search term
	results, err := s.GetSearchResults("xyz123abc456def789", 10)

	if err != nil {
		t.Fatalf("expected no error even for invalid search, got: %v", err)
	}

	// May return 0 results, but should not error
	if len(results.Quotes) == 0 {
		t.Log("No results found for invalid search (expected behavior)")
	}
}

func TestGetSearchResults_MultipleExchanges(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	s := newSearch()
	results, err := s.GetSearchResults("Toyota", 10)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(results.Quotes) == 0 {
		t.Error("expected at least one result for Toyota")
	}

	// Check for Toyota symbols from different exchanges
	exchanges := make(map[string]bool)
	for _, r := range results.Quotes {
		if r.Exchange != "" {
			exchanges[r.Exchange] = true
		}
	}

	if len(exchanges) > 1 {
		t.Logf("Found symbols from multiple exchanges: %v", exchanges)
	}
}

func TestTicker_Search(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	ticker := NewTicker("AAPL")
	results, err := ticker.Search("Bitcoin", 10)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(results) == 0 {
		t.Error("expected at least one result")
	}
}

func TestTicker_SearchWithOptions(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	ticker := NewTicker("AAPL")
	params := DefaultSearchParams("VCB", 5)
	params.EnableFuzzyQuery = true

	results, err := ticker.SearchWithOptions(params)

	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(results) == 0 {
		t.Error("expected at least one result for VCB")
	}
}
