package yahoofinanceapi

import (
	"fmt"
	"testing"
)

func TestNewHistory(t *testing.T) {
	history := NewHistory()
	_, err := history.GetHistory("AAPL")
	if err != nil {
		panic(err)
	}
}

func TestTransformData(t *testing.T) {
	history := NewHistory()
	historyResponse, err := history.GetHistory("AAPL")
	if err != nil {
		panic(err)
	}
	transformedData := history.transformData(historyResponse)
	for key, data := range transformedData {
		fmt.Printf("%s -- %f$ -- vol %d\n", key, data.Close, data.Volume)
	}
}
