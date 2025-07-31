package yahoofinanceapi

import (
	"testing"
)

func TestNewTicker(t *testing.T) {
	ticker := NewTicker("AAPL")
	if ticker == nil {
		t.Fatal("NewTicker returned nil")
	}
	if ticker.Symbol != "AAPL" {
		t.Errorf("Expected symbol 'AAPL', got '%s'", ticker.Symbol)
	}
}

func TestQuoteValidSymbol(t *testing.T) {
	ticker := NewTicker("AAPL")
	quote, err := ticker.Quote()
	if err != nil {
		t.Fatalf("Quote returned error: %v", err)
	}
	if quote.Close == 0 {
		t.Error("Quote returned PriceData with zero Close")
	}
}

func TestQuoteInvalidSymbol(t *testing.T) {
	ticker := NewTicker("INVALID_SYMBOL_123")
	_, err := ticker.Quote()
	if err == nil {
		t.Error("Expected error for invalid symbol, got nil")
	}
}

func TestGetInfoValidSymbol(t *testing.T) {
	ticker := NewTicker("AAPL")
	info, err := ticker.GetInfo()
	if err != nil {
		t.Fatalf("GetInfo returned error: %v", err)
	}
	if info.Symbol != "AAPL" {
		t.Errorf("Expected symbol 'AAPL', got '%s'", info.Symbol)
	}
	if info.ShortName == "" {
		t.Error("Expected non-empty ShortName")
	}
	if info.Currency == "" {
		t.Error("Expected non-empty Currency")
	}
}

func TestGetInfoInvalidSymbol(t *testing.T) {
	ticker := NewTicker("INVALID_SYMBOL_123")
	_, err := ticker.GetInfo()
	if err == nil {
		t.Error("Expected error for invalid symbol, got nil")
	}
}

func TestHistoryValidSymbol(t *testing.T) {
	ticker := NewTicker("AAPL")
	query := HistoryQuery{Range: "1mo", Interval: "1d"}
	data, err := ticker.History(query)
	if err != nil {
		t.Fatalf("History returned error: %v", err)
	}
	if len(data) == 0 {
		t.Error("History returned empty map for valid symbol")
	}
}

func TestHistoryInvalidSymbol(t *testing.T) {
	ticker := NewTicker("INVALID_SYMBOL_123")
	query := HistoryQuery{Range: "1mo", Interval: "1d"}
	_, err := ticker.History(query)
	if err == nil {
		t.Error("Expected error for invalid symbol, got nil")
	}
}

func TestOptionChain(t *testing.T) {
	ticker := NewTicker("AAPL")
	data := ticker.OptionChain()
	if len(data.Calls) == 0 && len(data.Puts) == 0 {
		t.Error("OptionChain returned empty calls and puts")
	}
}

func TestOptionChainByExpiration(t *testing.T) {
	ticker := NewTicker("AAPL")
	dates := ticker.ExpirationDates()
	if len(dates) == 0 {
		t.Skip("No expiration dates available for AAPL")
	}
	data := ticker.OptionChainByExpiration(dates[0])
	if len(data.Calls) == 0 && len(data.Puts) == 0 {
		t.Error("OptionChainByExpiration returned empty calls and puts")
	}
}

func TestExpirationDates(t *testing.T) {
	ticker := NewTicker("AAPL")
	dates := ticker.ExpirationDates()
	if len(dates) == 0 {
		t.Error("ExpirationDates returned empty slice")
	}
}
