# speedfast
Small GO library that tests the download and upload speeds by using Ookla's https://www.speedtest.net/ and Netflix's https://fast.com/

Used libs: 
* Delegates to [showwin/speedtest-go](https://github.com/showwin/speedtest-go) for **speedtest.net**
* Inspired by [gesquive/fast-cli](https://github.com/gesquive/fast-cli) (reuses a small part of the exposed api) and [npotts/go-fast](https://github.com/npotts/go-fast) for **fast.com**
* Uses [chromedp/chromedp](https://github.com/chromedp/chromedp) for the **web scrape version of fast.com** 

## Dependency

```
go get github.com/calin014/speedfast
```

### API Usage

```go
package main

import (
	"fmt"

	"github.com/calin014/speedfast"
)

func main() {
	fmt.Println(speedfast.MeasureWithSpeedtest())
	fmt.Println(speedfast.MeasureWithFast())
	fmt.Println(speedfast.MeasureWithFastInHeadlessBrowser())

	// ...or use the Measurer interface
	measurers := []speedfast.Measurer{
		speedfast.MeasurerFunc(speedfast.MeasureWithSpeedtest),
		speedfast.MeasurerFunc(speedfast.MeasureWithFast),
		speedfast.MeasurerFunc(speedfast.MeasureWithFastInHeadlessBrowser),
	}

	for _, m := range measurers {
		fmt.Println(m.Measure())
	}
}
```