package speedfast_test

import (
	"fmt"
	"testing"

	"github.com/calin014/speedfast"
)

func TestMeasureWithSpeedtest(t *testing.T) {
	result, err := speedfast.MeasureWithSpeedtest()

	if err != nil {
		t.Fatal("Got an error!")
	}

	fmt.Println(result)
}
