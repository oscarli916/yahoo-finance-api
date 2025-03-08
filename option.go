package yahoofinanceapi

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/url"
	"time"
)

type YahooOptionResponse struct {
	OptionChain YahooOptionChain `json:"optionChain"`
}

type YahooOptionChain struct {
	Result []YahooOptionResult `json:"result"`
	Error  any                 `json:"error"`
}

type YahooOptionResult struct {
	UnderlyingSymbol string           `json:"underlyingSymbol"`
	ExpirationDates  []int64          `json:"expirationDates"`
	Strikes          []float64        `json:"strikes"`
	HasMiniOptions   bool             `json:"hasMiniOptions"`
	Quote            YahooOptionQuote `json:"quote"`
	Options          []YahooOptions   `json:"options"`
}

type YahooOptionQuote struct {
	Language                          string  `json:"language"`
	Region                            string  `json:"region"`
	QuoteType                         string  `json:"quoteType"`
	TypeDisp                          string  `json:"typeDisp"`
	QuoteSourceName                   string  `json:"quoteSourceName"`
	Triggerable                       bool    `json:"triggerable"`
	CustomPriceAlertConfidence        string  `json:"customPriceAlertConfidence"`
	MarketState                       string  `json:"marketState"`
	RegularMarketChangePercent        float64 `json:"regularMarketChangePercent"`
	RegularMarketPrice                float64 `json:"regularMarketPrice"`
	ShortName                         string  `json:"shortName"`
	LongName                          string  `json:"longName"`
	Exchange                          string  `json:"exchange"`
	MessageBoardId                    string  `json:"messageBoardId"`
	ExchangeTimezoneName              string  `json:"exchangeTimezoneName"`
	ExchangeTimezoneShortName         string  `json:"exchangeTimezoneShortName"`
	GmtOffSetMilliseconds             int64   `json:"gmtOffSetMilliseconds"`
	Market                            string  `json:"market"`
	EsgPopulated                      bool    `json:"esgPopulated"`
	Currency                          string  `json:"currency"`
	HasPrePostMarketData              bool    `json:"hasPrePostMarketData"`
	FirstTradeDateMilliseconds        int64   `json:"firstTradeDateMilliseconds"`
	PriceHint                         int     `json:"priceHint"`
	PostMarketChangePercent           float64 `json:"postMarketChangePercent"`
	PostMarketPrice                   float64 `json:"postMarketPrice"`
	PostMarketChange                  float64 `json:"postMarketChange"`
	RegularMarketChange               float64 `json:"regularMarketChange"`
	RegularMarketDayHigh              float64 `json:"regularMarketDayHigh"`
	RegularMarketDayRange             string  `json:"regularMarketDayRange"`
	RegularMarketDayLow               float64 `json:"regularMarketDayLow"`
	RegularMarketVolume               int64   `json:"regularMarketVolume"`
	RegularMarketPreviousClose        float64 `json:"regularMarketPreviousClose"`
	Bid                               float64 `json:"bid"`
	Ask                               float64 `json:"ask"`
	BidSize                           int     `json:"bidSize"`
	AskSize                           int     `json:"askSize"`
	FullExchangeName                  string  `json:"fullExchangeName"`
	FinancialCurrency                 string  `json:"financialCurrency"`
	RegularMarketOpen                 float64 `json:"regularMarketOpen"`
	AverageDailyVolume3Month          int64   `json:"averageDailyVolume3Month"`
	AverageDailyVolume10Day           int64   `json:"averageDailyVolume10Day"`
	FiftyTwoWeekLowChange             float64 `json:"fiftyTwoWeekLowChange"`
	FiftyTwoWeekLowChangePercent      float64 `json:"fiftyTwoWeekLowChangePercent"`
	FiftyTwoWeekRange                 string  `json:"fiftyTwoWeekRange"`
	FiftyTwoWeekHighChange            float64 `json:"fiftyTwoWeekHighChange"`
	FiftyTwoWeekHighChangePercent     float64 `json:"fiftyTwoWeekHighChangePercent"`
	FiftyTwoWeekLow                   float64 `json:"fiftyTwoWeekLow"`
	FiftyTwoWeekHigh                  float64 `json:"fiftyTwoWeekHigh"`
	FiftyTwoWeekChangePercent         float64 `json:"fiftyTwoWeekChangePercent"`
	DividendDate                      int64   `json:"dividendDate"`
	EarningsTimestamp                 int64   `json:"earningsTimestamp"`
	EarningsTimestampStart            int64   `json:"earningsTimestampStart"`
	EarningsTimestampEnd              int64   `json:"earningsTimestampEnd"`
	EarningsCallTimestampStart        int64   `json:"earningsCallTimestampStart"`
	EarningsCallTimestampEnd          int64   `json:"earningsCallTimestampEnd"`
	IsEarningsDateEstimate            bool    `json:"isEarningsDateEstimate"`
	TrailingAnnualDividendRate        float64 `json:"trailingAnnualDividendRate"`
	TrailingPE                        float64 `json:"trailingPE"`
	DividendRate                      float64 `json:"dividendRate"`
	TrailingAnnualDividendYield       float64 `json:"trailingAnnualDividendYield"`
	DividendYield                     float64 `json:"dividendYield"`
	EpsTrailingTwelveMonths           float64 `json:"epsTrailingTwelveMonths"`
	EpsForward                        float64 `json:"epsForward"`
	EpsCurrentYear                    float64 `json:"epsCurrentYear"`
	PriceEpsCurrentYear               float64 `json:"priceEpsCurrentYear"`
	SharesOutstanding                 int64   `json:"sharesOutstanding"`
	BookValue                         float64 `json:"bookValue"`
	FiftyDayAverage                   float64 `json:"fiftyDayAverage"`
	FiftyDayAverageChange             float64 `json:"fiftyDayAverageChange"`
	FiftyDayAverageChangePercent      float64 `json:"fiftyDayAverageChangePercent"`
	TwoHundredDayAverage              float64 `json:"twoHundredDayAverage"`
	TwoHundredDayAverageChange        float64 `json:"twoHundredDayAverageChange"`
	TwoHundredDayAverageChangePercent float64 `json:"twoHundredDayAverageChangePercent"`
	MarketCap                         int64   `json:"marketCap"`
	ForwardPE                         float64 `json:"forwardPE"`
	PriceToBook                       float64 `json:"priceToBook"`
	SourceInterval                    int     `json:"sourceInterval"`
	ExchangeDataDelayedBy             int     `json:"exchangeDataDelayedBy"`
	AverageAnalystRating              string  `json:"averageAnalystRating"`
	Tradeable                         bool    `json:"tradeable"`
	CryptoTradeable                   bool    `json:"cryptoTradeable"`
	CorporateActions                  []any   `json:"corporateActions"`
	PostMarketTime                    int64   `json:"postMarketTime"`
	RegularMarketTime                 int64   `json:"regularMarketTime"`
	DisplayName                       string  `json:"displayName"`
	Symbol                            string  `json:"symbol"`
}

type YahooOptions struct {
	ExpirationDate int64         `json:"expirationDate"`
	HasMiniOptions bool          `json:"hasMiniOptions"`
	Calls          []YahooOption `json:"calls"`
	Puts           []YahooOption `json:"puts"`
}

type YahooOption struct {
	ContractSymbol    string  `json:"contractSymbol"`
	Strike            float64 `json:"strike"`
	Currency          string  `json:"currency"`
	LastPrice         float64 `json:"lastPrice"`
	Change            float64 `json:"change"`
	PercentChange     float64 `json:"percentChange"`
	Volume            int64   `json:"volume"`
	OpenInterest      int64   `json:"openInterest"`
	Bid               float64 `json:"bid"`
	Ask               float64 `json:"ask"`
	ContractSize      string  `json:"contractSize"`
	Expiration        int64   `json:"expiration"`
	LastTradeDate     int64   `json:"lastTradeDate"`
	ImpliedVolatility float64 `json:"impliedVolatility"`
	InTheMoney        bool    `json:"inTheMoney"`
}

type OptionData struct {
	ExpirationDate string         `json:"expirationDate"`
	HasMiniOptions bool           `json:"hasMiniOptions"`
	Calls          []OptionDetail `json:"calls"`
	Puts           []OptionDetail `json:"puts"`
}

type OptionDetail struct {
	ContractSymbol    string  `json:"contractSymbol"`
	Strike            float64 `json:"strike"`
	Currency          string  `json:"currency"`
	LastPrice         float64 `json:"lastPrice"`
	Change            float64 `json:"change"`
	PercentChange     float64 `json:"percentChange"`
	Volume            int64   `json:"volume"`
	OpenInterest      int64   `json:"openInterest"`
	Bid               float64 `json:"bid"`
	Ask               float64 `json:"ask"`
	ContractSize      string  `json:"contractSize"`
	Expiration        string  `json:"expiration"`
	LastTradeDate     string  `json:"lastTradeDate"`
	ImpliedVolatility float64 `json:"impliedVolatility"`
	InTheMoney        bool    `json:"inTheMoney"`
}

type Option struct {
	client *Client
}

func NewOption() *Option {
	return &Option{client: GetClient()}
}

func (o *Option) GetOptionChain(symbol string) YahooOptionResponse {
	endpoint := fmt.Sprintf("%s/v7/finance/options/%s", BASE_URL, symbol)
	resp, err := o.client.Get(endpoint, url.Values{})
	if err != nil {
		slog.Error("Failed to get option chain", "err", err)
		return YahooOptionResponse{}
	}
	defer resp.Body.Close()

	var optionResponse YahooOptionResponse
	if err := json.NewDecoder(resp.Body).Decode(&optionResponse); err != nil {
		slog.Error("Failed to decode option data JSON response", "err", err)
	}

	return optionResponse
}

func (o *Option) GetOptionChainByExpiration(symbol string, expirationDate string) YahooOptionResponse {
	t, err := time.Parse("2006-01-02", expirationDate)
	if err != nil {
		slog.Error("Failed to parse expiration date", "err", err)
		return YahooOptionResponse{}
	}
	endpoint := fmt.Sprintf("%s/v7/finance/options/%s", BASE_URL, symbol)
	params := url.Values{}
	params.Add("date", fmt.Sprintf("%d", t.Unix()))
	resp, err := o.client.Get(endpoint, params)
	if err != nil {
		slog.Error("Failed to get option chain by expiration", "err", err)
		return YahooOptionResponse{}
	}
	defer resp.Body.Close()

	var optionResponse YahooOptionResponse
	if err := json.NewDecoder(resp.Body).Decode(&optionResponse); err != nil {
		slog.Error("Failed to decode option data JSON response", "err", err)
	}

	return optionResponse
}

func (o *Option) transformData(data YahooOptionResponse) OptionData {
	date := time.Unix(data.OptionChain.Result[0].Options[0].ExpirationDate, 0).Format("2006-01-02")
	var calls []OptionDetail
	var puts []OptionDetail
	for _, call := range data.OptionChain.Result[0].Options[0].Calls {
		calls = append(calls, OptionDetail{
			ContractSymbol:    call.ContractSymbol,
			Strike:            call.Strike,
			Currency:          call.Currency,
			LastPrice:         call.LastPrice,
			Change:            call.Change,
			PercentChange:     call.PercentChange,
			Volume:            call.Volume,
			OpenInterest:      call.OpenInterest,
			Bid:               call.Bid,
			Ask:               call.Ask,
			ContractSize:      call.ContractSize,
			Expiration:        time.Unix(call.Expiration, 0).Format("2006-01-02"),
			LastTradeDate:     time.Unix(call.LastTradeDate, 0).Format("2006-01-02"),
			ImpliedVolatility: call.ImpliedVolatility,
			InTheMoney:        call.InTheMoney,
		})
	}
	for _, put := range data.OptionChain.Result[0].Options[0].Puts {
		puts = append(puts, OptionDetail{
			ContractSymbol:    put.ContractSymbol,
			Strike:            put.Strike,
			Currency:          put.Currency,
			LastPrice:         put.LastPrice,
			Change:            put.Change,
			PercentChange:     put.PercentChange,
			Volume:            put.Volume,
			OpenInterest:      put.OpenInterest,
			Bid:               put.Bid,
			Ask:               put.Ask,
			ContractSize:      put.ContractSize,
			Expiration:        time.Unix(put.Expiration, 0).Format("2006-01-02"),
			LastTradeDate:     time.Unix(put.LastTradeDate, 0).Format("2006-01-02"),
			ImpliedVolatility: put.ImpliedVolatility,
			InTheMoney:        put.InTheMoney,
		})
	}
	return OptionData{
		ExpirationDate: date,
		HasMiniOptions: data.OptionChain.Result[0].HasMiniOptions,
		Calls:          calls,
		Puts:           puts,
	}
}

func (o *Option) GetExpirationDates(symbol string) []string {
	optionChain := o.GetOptionChain(symbol)
	var expirationDates []string
	for _, date := range optionChain.OptionChain.Result[0].ExpirationDates {
		expirationDates = append(expirationDates, time.Unix(date, 0).Format("2006-01-02"))
	}
	return expirationDates
}
