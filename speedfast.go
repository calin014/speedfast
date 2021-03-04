package speedfast

import (
	"fmt"
)

// Measurement contains the result of network speed measurement
type Measurement struct {
	Source           string
	Download, Upload float64
}

func (measurement Measurement) String() string {
	return fmt.Sprintf("Source: %v, Download speed: %f Mbps, Upload speed: %f Mbps", measurement.Source, measurement.Download, measurement.Upload)
}

// Not sure if by "1 exposed API" was ment 1 exposed function that takes a discriminator parameter, but I chose to expose one function per measurement type

// Measurer is the common interface for all the measurement functions
type Measurer interface {
	Measure() (Measurement, error)
}

// MeasurerFunc is used to create a Measurer from a function with the proper signature
type MeasurerFunc func() (Measurement, error)

// Measure is MeasurerFunc implementation of Measurer
func (f MeasurerFunc) Measure() (Measurement, error) {
	return f()
}
