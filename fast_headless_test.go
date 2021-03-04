package speedfast_test

import (
	"testing"

	"github.com/calin014/speedfast"
)

func TestFastHeadlessMeasurement(t *testing.T) {
	result, err := speedfast.MeasureWithFastInHeadlessBrowser()

	if err != nil {
		t.Fatal("Got an error:", err)
	}

	t.Log(result)
}

func BenchmarkFastHeadlessMeasurement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		speedfast.MeasureWithFastInHeadlessBrowser()
	}
}
