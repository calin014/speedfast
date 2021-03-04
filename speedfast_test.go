package speedfast_test

import (
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

		t.Log(result)
	}
}

func TestCustomMeasurer(t *testing.T) {
	someMeasurement := speedfast.Measurement{}

	measurer := func() (speedfast.Measurement, error) {
		return someMeasurement, nil
	}

	result, _ := speedfast.MeasurerFunc(measurer).Measure()

	if someMeasurement != result {
		t.Fatal("Result should be: ", someMeasurement, " but got: ", result)
	}
}
