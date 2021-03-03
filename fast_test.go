package speedfast_test

import (
	"testing"

	"github.com/calin014/speedfast"
)

func TestFastMeasurement(t *testing.T) {
	result, err := speedfast.MeasureWithFast()

	if err != nil {
		t.Fatal("Got an error:", err)
	}

	t.Log(result)
}
