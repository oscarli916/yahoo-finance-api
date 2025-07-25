package yahoofinanceapi

import (
	"testing"
	"time"
)

func TestNewHistory(t *testing.T) {
	history := NewHistory()
	if history == nil {
		t.Fatal("NewHistory returned nil")
	}
}

func TestGetHistoryValidSymbol(t *testing.T) {
	history := NewHistory()
	resp, err := history.GetHistory("AAPL")
	if err != nil {
		t.Fatalf("GetHistory returned error: %v", err)
	}
	if len(resp.Chart.Result) == 0 {
		t.Fatal("GetHistory returned empty result for valid symbol")
	}
}

func TestGetHistoryInvalidSymbol(t *testing.T) {
	history := NewHistory()
	_, err := history.GetHistory("INVALID_SYMBOL_123")
	if err == nil {
		t.Error("Expected error for invalid symbol, got nil")
	}
}

func TestTransformData(t *testing.T) {
	history := NewHistory()
	resp, err := history.GetHistory("AAPL")
	if err != nil {
		t.Fatalf("GetHistory returned error: %v", err)
	}
	transformed := history.transformData(resp)
	if len(transformed) == 0 {
		t.Error("transformData returned empty map")
	}
	for _, data := range transformed {
		if data.Close == 0 {
			t.Error("transformData returned PriceData with zero Close")
		}
	}
}

func TestTransformDataWithMinuteInterval(t *testing.T) {
	history := NewHistory()
	q := HistoryQuery{Range: "1d", Interval: "1m"}
	history.SetQuery(q)
	resp, err := history.GetHistory("AAPL")
	if err != nil {
		t.Fatalf("GetHistory returned error: %v", err)
	}
	transformed := history.transformData(resp)
	if len(transformed) == 0 {
		t.Error("transformData returned empty map")
	}
	for _, data := range transformed {
		if data.Close == 0 {
			t.Error("transformData returned PriceData with zero Close")
		}
	}
}

func TestSetQuery(t *testing.T) {
	history := NewHistory()
	q := HistoryQuery{Range: "5d", Interval: "1d"}
	history.SetQuery(q)
	if history.query.Range != "5d" || history.query.Interval != "1d" {
		t.Error("SetQuery did not set query fields correctly")
	}
}

func TestSetDefault(t *testing.T) {
	q := HistoryQuery{}
	q.SetDefault()
	if q.Range != "1mo" {
		t.Error("SetDefault did not set default Range")
	}
	if q.Interval != "1d" {
		t.Error("SetDefault did not set default Interval")
	}
	if q.UserAgent == "" {
		t.Error("SetDefault did not set UserAgent")
	}
}

func TestSetDefaultWithStartDate(t *testing.T) {
	q := HistoryQuery{Start: "2024-01-01"}
	q.SetDefault()
	_, err := time.Parse("2006-01-02", "2024-01-01")
	if err != nil {
		t.Fatal("Test setup error: invalid date")
	}
	if q.Start == "default" {
		t.Error("SetDefault failed to parse valid Start date")
	}
}

func TestSetDefaultWithInvalidStartDate(t *testing.T) {
	q := HistoryQuery{Start: "invalid-date"}
	q.SetDefault()
	if q.Start != "default" {
		t.Error("SetDefault did not set Start to 'default' for invalid date")
	}
}
