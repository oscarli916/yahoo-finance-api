package yahoofinanceapi

import (
	"log"
	"testing"
)

func TestNewInformation(t *testing.T) {
	info := newInformation()
	if info == nil {
		t.Fatal("newInformation returned nil")
	}
}

func TestGetTickerInfoValidSymbol(t *testing.T) {
	info := newInformation()
	data, err := info.GetInfo("AAPL")
	log.Println(data, err)
	if err != nil {
		t.Fatalf("GetTickerInfo returned error: %v", err)
	}
	if data.Symbol != "AAPL" {
		t.Errorf("Expected symbol 'AAPL', got '%s'", data.Symbol)
	}
	if data.ShortName != "Apple Inc." {
		t.Errorf("Expected ShortName 'Apple Inc.', got '%s'", data.ShortName)
	}
	if data.Currency != "USD" {
		t.Errorf("Expected Currency 'USD', got '%s'", data.Currency)
	}
}

func TestGetTickerInfoInvalidSymbol(t *testing.T) {
	info := newInformation()
	_, err := info.GetInfo("INVALID_SYMBOL_123")
	if err == nil {
		t.Error("Expected error for invalid symbol, got nil")
	}
}
