package speedfast_test

import (
	"fmt"
	"testing"

	"github.com/calin014/speedfast"
)

func TestAllMeasurers(t *testing.T) {

	measurers := []speedfast.Measurer{
		speedfast.MeasurerFunc(speedfast.MeasureWithSpeedtest),
		speedfast.MeasurerFunc(speedfast.MeasureWithFast),
		speedfast.MeasurerFunc(speedfast.MeasureWithFastInHeadlessBrowser),
	}

	for _, m := range measurers {
		result, err := m.Measure()

		if err != nil {
			t.Fatal("Got an error!")
		}

		fmt.Println(result)
	}
}
