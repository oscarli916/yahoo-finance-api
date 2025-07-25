package yahoofinanceapi

import (
	"fmt"
	"testing"
)

func TestNewTicker(t *testing.T) {
	ticker := NewTicker("AAPL")

	if ticker.Symbol != "AAPL" {
		panic("Something went wrong creating a ticker")
	}
}

func TestGetQuote(t *testing.T) {
	ticker := NewTicker("AAPL")
	quote, _ := ticker.Quote()
	fmt.Println(quote.Close)
}
