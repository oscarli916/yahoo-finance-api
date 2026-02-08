package yahoofinanceapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/url"
	"strconv"
	"strings"
)

// YahooSearchResponse represents the raw Yahoo Finance search API response
type YahooSearchResponse struct {
	Quotes []struct {
		Symbol    string `json:"symbol"`
		ShortName string `json:"shortname"`
		LongName  string `json:"longname"`
		QuoteType string `json:"quoteType"`
		Exchange  string `json:"exchange"`
		ExchDisp  string `json:"exchDisp"`
		TypeDisp  string `json:"typeDisp"`
	} `json:"quotes"`
	News     []any `json:"news"`
	Lists    []any `json:"lists"`
	Research struct {
		Reports []any `json:"reports"`
	} `json:"reports"`
}

// SearchData represents the transformed search results
type SearchData struct {
	Results []SearchResult
}

// SearchResult represents a symbol search result from Yahoo Finance
type SearchResult struct {
	Symbol    string `json:"symbol"`
	Name      string `json:"name"`
	ShortName string `json:"shortname"`
	LongName  string `json:"longname"`
	Type      string `json:"type"`
	Exchange  string `json:"exchange"`
	ExchDisp  string `json:"exchDisp"`
}

// SearchParams holds configurable parameters for Yahoo Finance search API
type SearchParams struct {
	// Core search parameters
	Query       string // Search query (symbol, company name, etc.)
	QuotesCount int    // Number of quote results to return (default: 10, max: 20)
	NewsCount   int    // Number of news results (default: 0, not needed for symbol search)
	ListsCount  int    // Number of list results (default: 0, not needed)

	// Query optimization parameters
	EnableFuzzyQuery           bool // Enable fuzzy matching (default: false for exact matching)
	EnableEnhancedTrivialQuery bool // Enable better trivial query handling (default: true)
	EnableCccBoost             bool // Enable cryptocurrency boost (default: true)
	EnablePrivateCompany       bool // Include private companies (default: true)

	// Feature flags
	EnableResearchReports bool // Enable research reports data (default: true)
	EnableCulturalAssets  bool // Enable cultural assets data (default: true)
	EnableLogoUrl         bool // Enable logo URLs (default: true)
	EnableNavLinks        bool // Enable navigation links (default: true)
	EnableCb              bool // Enable clickback (default: false)
	EnableLists           bool // Enable lists in results (default: false)

	// Query IDs (Yahoo Finance internal query routing)
	QuotesQueryId     string // Default: "tss_match_phrase_query"
	MultiQuoteQueryId string // Default: "multi_quote_single_token_query"
	NewsQueryId       string // Default: "news_cie_vespa" (not used when newsCount=0)

	// Recommendation
	RecommendCount int // Number of recommendations (default: 5)

	// Language
	Lang string // Language code (default: "en-US")
}

// DefaultSearchParams returns the default search parameters optimized for symbol lookup
func DefaultSearchParams(query string, quotesCount int) SearchParams {
	return SearchParams{
		Query:                      query,
		QuotesCount:                quotesCount,
		NewsCount:                  0,
		ListsCount:                 0,
		EnableFuzzyQuery:           false,
		EnableEnhancedTrivialQuery: true,
		EnableCccBoost:             true,
		EnablePrivateCompany:       true,
		EnableResearchReports:      true,
		EnableCulturalAssets:       true,
		EnableLogoUrl:              true,
		EnableNavLinks:             true,
		EnableCb:                   false,
		EnableLists:                false,
		QuotesQueryId:              "tss_match_phrase_query",
		MultiQuoteQueryId:          "multi_quote_single_token_query",
		NewsQueryId:                "news_cie_vespa",
		RecommendCount:             5,
		Lang:                       "en-US",
	}
}

// buildSearchURL constructs the Yahoo Finance search URL with parameters
func buildSearchURL(params SearchParams) string {
	baseURL := fmt.Sprintf("%s/v1/finance/search", BASE_URL)

	// Build query parameters
	values := url.Values{}
	values.Set("q", params.Query)
	values.Set("lang", params.Lang)
	values.Set("quotesCount", strconv.Itoa(params.QuotesCount))
	values.Set("newsCount", strconv.Itoa(params.NewsCount))
	values.Set("listsCount", strconv.Itoa(params.ListsCount))
	values.Set("enableFuzzyQuery", strconv.FormatBool(params.EnableFuzzyQuery))
	values.Set("quotesQueryId", params.QuotesQueryId)
	values.Set("multiQuoteQueryId", params.MultiQuoteQueryId)
	values.Set("enableCb", strconv.FormatBool(params.EnableCb))
	values.Set("enableNavLinks", strconv.FormatBool(params.EnableNavLinks))
	values.Set("enableEnhancedTrivialQuery", strconv.FormatBool(params.EnableEnhancedTrivialQuery))
	values.Set("enableResearchReports", strconv.FormatBool(params.EnableResearchReports))
	values.Set("enableCulturalAssets", strconv.FormatBool(params.EnableCulturalAssets))
	values.Set("enableLogoUrl", strconv.FormatBool(params.EnableLogoUrl))
	values.Set("enableLists", strconv.FormatBool(params.EnableLists))
	values.Set("recommendCount", strconv.Itoa(params.RecommendCount))
	values.Set("enableCccBoost", strconv.FormatBool(params.EnableCccBoost))
	values.Set("enablePrivateCompany", strconv.FormatBool(params.EnablePrivateCompany))
	if params.NewsQueryId != "" {
		values.Set("newsQueryId", params.NewsQueryId)
	}

	return fmt.Sprintf("%s?%s", baseURL, values.Encode())
}

// Search holds the HTTP client for symbol searching
type Search struct {
	client *Client
}

// newSearch initializes the Search struct with an HTTP client
func newSearch() *Search {
	return &Search{client: getClient()}
}

// GetSearchResults searches for investment symbols by query using Yahoo Finance's public search API
// API endpoint: https://query2.finance.yahoo.com/v1/finance/search
//
// This is a convenience function that uses default search parameters.
// For custom parameters, use GetSearchResultsWithOptions.
//
// Parameters:
//   - query: Search query (symbol, company name, etc.)
//   - limit: Maximum number of results to return (max 20)
//
// Returns:
//   - YahooSearchResponse: Raw search response from Yahoo Finance
//   - error: Error if request fails or query is invalid
func (s *Search) GetSearchResults(query string, limit int) (YahooSearchResponse, error) {
	params := DefaultSearchParams(query, limit)
	return s.GetSearchResultsWithOptions(params)
}

// GetSearchResultsWithOptions searches for investment symbols using custom parameters
//
// Parameters:
//   - params: SearchParams with custom configuration
//
// Returns:
//   - YahooSearchResponse: Raw search response from Yahoo Finance
//   - error: Error if request fails or query is invalid
func (s *Search) GetSearchResultsWithOptions(params SearchParams) (YahooSearchResponse, error) {
	// Validate query
	params.Query = strings.TrimSpace(params.Query)
	if params.Query == "" {
		return YahooSearchResponse{}, fmt.Errorf("query cannot be empty")
	}

	// Validate and cap quotes count
	if params.QuotesCount <= 0 {
		params.QuotesCount = 10
	}
	if params.QuotesCount > 20 {
		params.QuotesCount = 20
	}

	// Build request URL
	searchURL := buildSearchURL(params)

	// Make the HTTP GET request using the client
	resp, err := s.client.Get(searchURL, url.Values{})
	if err != nil {
		slog.Error("Failed to search symbols", "err", err)
		return YahooSearchResponse{}, fmt.Errorf("failed to search symbols: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return YahooSearchResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return YahooSearchResponse{}, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse JSON response
	var searchResponse YahooSearchResponse
	if err := json.Unmarshal(body, &searchResponse); err != nil {
		return YahooSearchResponse{}, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return searchResponse, nil
}

// transformData converts YahooSearchResponse to SearchData
func (s *Search) transformData(data YahooSearchResponse) SearchData {
	results := make([]SearchResult, 0, len(data.Quotes))
	for _, quote := range data.Quotes {
		// Get name from shortname or longname
		name := quote.ShortName
		if name == "" {
			name = quote.LongName
		}

		results = append(results, SearchResult{
			Symbol:    quote.Symbol,
			Name:      name,
			ShortName: quote.ShortName,
			LongName:  quote.LongName,
			Type:      quote.QuoteType,
			Exchange:  quote.Exchange,
			ExchDisp:  quote.ExchDisp,
		})
	}
	return SearchData{Results: results}
}
