# yahoo-finance-api

## Motivation

- I used to write Python programs and use Yahoo Finance data. The [yfinance](https://github.com/ranaroussi/yfinance) library is an awesome library which I enjoyed a lot.
- Could not find similar packages in Go.
- Learn Go
- Able to use this package on my other Go projects

## Contributing

1. Fork the repository
2. Create your feature branch (git checkout -b feature/amazing-feature)
3. Commit your changes (git commit -m 'Add some amazing feature')
4. Push to the branch (git push origin feature/amazing-feature)
5. Open a Pull Request

## Getting Started with Ticker
~~~
import (
    "fmt"
    finance "github.com/oscarli916/yahoo-finance-api/"
)

func GetQuote(){

    // create the ticker object
	ticker := finance.NewTicker("AAPL)

    // get the latest PriceData
	quote, _ := ticker.Quote()
	fmt.Println(quote.Close)

}

~~~
## Getting Started with History


import (
    "fmt"
    finance "github.com/oscarli916/yahoo-finance-api/"
)

func GetHistory(){

    // create the history manager
	history := finance.NewHistory()
	appleHistoryResponse, err := history.GetHistory("AAPL")
	if err != nil {
		panic(err)
	}

    // then transform it into a map of PriceData
	applePriceData := history.transformData(appleHistoryResponse)

    // print the Closing price and Volume of that day
	for date, data := range transformedData {
		fmt.Printf("%s -- %f$ -- vol %d \n", data, data.Close, data.Volume)
	}

}
## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
