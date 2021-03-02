package speedfast

import (
	"fmt"
)

// Measurement contains the result of network speed measurement
type Measurement struct {
	download, upload float32
}

func (measurement Measurement) String() string {
	return fmt.Sprintf("Download speed: %v, Upload speed: %v", measurement.download, measurement.upload)
}

// MeasureWithFast measures network speed with fast.com
func MeasureWithFast() Measurement {
	return Measurement{320.1, 150.4}
}

// MeasureWithSpeedtest measures network speed with speedtest.com
func MeasureWithSpeedtest() Measurement {
	return Measurement{350.1, 140.2}
}
