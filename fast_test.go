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

func BenchmarkFastMeasurement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		speedfast.MeasureWithFast()
	}
}
