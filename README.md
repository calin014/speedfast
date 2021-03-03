# speedfast
Small GO library that tests the download and upload speeds by using Ookla's https://www.speedtest.net/ and Netflix's https://fast.com/

## Dependency

```
go get github.com/calin014/speedfast
```

### API Usage
The code below finds closest available speedtest server and tests the latency, download, and upload speeds.
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
		result, err := m.Measure()
		fmt.Println(result)
	}
}
```