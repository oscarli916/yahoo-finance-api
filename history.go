package yahoofinanceapi

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

type YahooHistoryRespose struct {
	Chart YahooChart `json:"chart"`
}

type YahooChart struct {
	Result []YahooResult `json:"result"`
}

type YahooResult struct {
	Meta       YahooMeta      `json:"meta"`
	Timestamp  []int64        `json:"timestamp"`
	Indicators YahooIndicator `json:"indicators"`
}

type YahooMeta struct {
	Currency             string                 `json:"currency"`
	Symbol               string                 `json:"symbol"`
	ExchangeName         string                 `json:"exchangeName"`
	FullExchangeName     string                 `json:"fullExchangeName"`
	InstrumentType       string                 `json:"instrumentType"`
	FirstTradeDate       int64                  `json:"firstTradeDate"`
	RegularMarketTime    int64                  `json:"regularMarketTime"`
	HasPrePostMarketData bool                   `json:"hasPrePostMarketData"`
	GmtOffset            int                    `json:"gmtoffset"`
	Timezone             string                 `json:"timezone"`
	ExchangeTimezoneName string                 `json:"exchangeTimezoneName"`
	RegularMarketPrice   float64                `json:"regularMarketPrice"`
	FiftyTwoWeekHigh     float64                `json:"fiftyTwoWeekHigh"`
	FiftyTwoWeekLow      float64                `json:"fiftyTwoWeekLow"`
	RegularMarketDayHigh float64                `json:"regularMarketDayHigh"`
	RegularMarketDayLow  float64                `json:"regularMarketDayLow"`
	RegularMarketVolume  int64                  `json:"regularMarketVolume"`
	LongName             string                 `json:"longName"`
	ShortName            string                 `json:"shortName"`
	ChartPreviousClose   float64                `json:"chartPreviousClose"`
	PreviousClose        float64                `json:"previousClose"`
	Scale                int                    `json:"scale"`
	PriceHint            int                    `json:"priceHint"`
	CurrentTradingPeriod YahooTradingPeriod     `json:"currentTradingPeriod"`
	TradingPeriods       [][]YahooTradingPeriod `json:"tradingPeriods"`
	DataGranularity      string                 `json:"dataGranularity"`
	Range                string                 `json:"range"`
	ValidRanges          []string               `json:"validRanges"`
}

type YahooTradingPeriod struct {
	Timezone  string `json:"timezone"`
	End       int64  `json:"end"`
	Start     int64  `json:"start"`
	GmtOffset int    `json:"gmtoffset"`
}

type YahooIndicator struct {
	Quote []YahooQuote `json:"quote"`
}

type YahooQuote struct {
	Open   []float64 `json:"open"`
	High   []float64 `json:"high"`
	Low    []float64 `json:"low"`
	Close  []float64 `json:"close"`
	Volume []int64   `json:"volume"`
}

type PriceData struct {
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume int64
}

type HistoryQuery struct {
	Range     string
	Interval  string
	Start     string
	End       string
	UserAgent string
}

func (hq *HistoryQuery) SetDefault() {
	if hq.Range == "" {
		hq.Range = "1mo"
	}
	if hq.Interval == "" {
		hq.Interval = "1d"
	}
	if hq.UserAgent == "" {
		hq.UserAgent = USER_AGENTS[rand.Intn(len(USER_AGENTS))]
	}
}

type History struct {
	query *HistoryQuery
}

func NewHistory() *History {
	return &History{query: &HistoryQuery{}}
}

func (h *History) SetQuery(query HistoryQuery) {
	h.query = &query
}

func (h *History) GetHistory(symbol string) map[string]PriceData {
	h.query.SetDefault()

	params := url.Values{}
	params.Add("range", h.query.Range)
	params.Add("interval", h.query.Interval)

	url := fmt.Sprintf("%s/v8/finance/chart/%s?%s", BASE_URL, symbol, params.Encode())
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("User-Agent", h.query.UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to get history data from Yahoo Finance API: %v", err)
	}
	defer resp.Body.Close()

	var historyResponse YahooHistoryRespose
	if err := json.NewDecoder(resp.Body).Decode(&historyResponse); err != nil {
		log.Fatalf("Failed to decode history data JSON response: %v", err)
	}

	return h.transformData(historyResponse)
}

func (h *History) transformData(data YahooHistoryRespose) map[string]PriceData {
	d := make(map[string]PriceData)
	for i, result := range data.Chart.Result[0].Timestamp {
		t := time.Unix(result, 0)
		date := t.Format("2006-01-02 15:04:05")
		fmt.Println(date, result, i)
		d[date] = PriceData{
			Open:   data.Chart.Result[0].Indicators.Quote[0].Open[i],
			High:   data.Chart.Result[0].Indicators.Quote[0].High[i],
			Low:    data.Chart.Result[0].Indicators.Quote[0].Low[i],
			Close:  data.Chart.Result[0].Indicators.Quote[0].Close[i],
			Volume: data.Chart.Result[0].Indicators.Quote[0].Volume[i]}
	}

	return d
}
