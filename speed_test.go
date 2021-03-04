package speedfast_test

import (
	"testing"

	"github.com/calin014/speedfast"
)

func TestSpeedtestMeasurement(t *testing.T) {
	result, err := speedfast.MeasureWithSpeedtest()

	if err != nil {
		t.Fatal("Got an error:", err)
	}

	t.Log(result)
}

func BenchmarkSpeedtestMeasurement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		speedfast.MeasureWithSpeedtest()
	}
}
