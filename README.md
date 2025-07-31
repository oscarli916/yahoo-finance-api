# yahoo-finance-api

## Motivation

- I used to write Python programs and use Yahoo Finance data. The [yfinance](https://github.com/ranaroussi/yfinance) library is an awesome library which I enjoyed a lot.
- Could not find similar packages in Go.
- Learn Go
- Able to use this package on my other Go projects

## Installation

```
go get github.com/oscarli916/yahoo-finance-api
```

## Example

```go
package main

import (
	"fmt"

	yfa "github.com/oscarli916/yahoo-finance-api"
)

func main() {
	t := yfa.NewTicker("AAPL")

	// get the latest PriceData
	quote, err := t.Quote()
	if err != nil {
		fmt.Println("Error fetching quote:", err)
		return
	}
	fmt.Println(quote.Close)

	// history data
	history, err := t.History(yfa.HistoryQuery{Range: "1d", Interval: "1m"})
	if err != nil {
		fmt.Println("Error fetching history:", err)
		return
	}
	fmt.Println(history)

	// option chain
	e := t.ExpirationDates()
	oc := t.OptionChainByExpiration(e[2])
	fmt.Println(oc)
	
	// Ticker Information
	info, err := t.GetInfo()
	if err != nil {
		fmt.Println("GetInfo returned error:", err)
	}
	fmt.Println(info)
}

```

## Contributing

1. Fork the repository
2. Create your feature branch (git checkout -b feature/amazing-feature)
3. Commit your changes (git commit -m 'Add some amazing feature')
4. Push to the branch (git push origin feature/amazing-feature)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
