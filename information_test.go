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
	data, err := info.GetTickerInfo("AAPL")
	log.Println(data, err)
	if err != nil {
		t.Fatalf("GetTickerInfo returned error: %v", err)
	}
	if data.Symbol != "AAPL" {
		t.Errorf("Expected symbol 'AAPL', got '%s'", data.Symbol)
	}
	if data.ShortName == "" {
		t.Error("Expected non-empty ShortName")
	}
	if data.Currency == "" {
		t.Error("Expected non-empty Currency")
	}
}

func TestGetTickerInfoInvalidSymbol(t *testing.T) {
	info := newInformation()
	_, err := info.GetTickerInfo("INVALID_SYMBOL_123")
	if err == nil {
		t.Error("Expected error for invalid symbol, got nil")
	}
}
