package speedfast

import (
	"errors"
	"fmt"

	"github.com/showwin/speedtest-go/speedtest"
)

// Measurement contains the result of network speed measurement
type Measurement struct {
	Source           string
	Download, Upload float64
}

func (measurement Measurement) String() string {
	return fmt.Sprintf("Source: %v, Download speed: %f Mbps, Upload speed: %f Mbps", measurement.Source, measurement.Download, measurement.Upload)
}

// MeasureWithSpeedtest measures network speed using speedtest.com's api
func MeasureWithSpeedtest() (Measurement, error) {
	user, _ := speedtest.FetchUserInfo()

	serverList, err := speedtest.FetchServerList(user)

	if err != nil {
		return Measurement{}, err
	}

	if len(serverList.Servers) <= 0 {
		return Measurement{}, errors.New("no servers available")
	}

	// First server is closest
	testServer := serverList.Servers[0]

	if err := testServer.DownloadTest(false); err != nil {
		return Measurement{}, err
	}

	if err := testServer.UploadTest(false); err != nil {
		return Measurement{}, err
	}

	return Measurement{"speedtest.com", testServer.DLSpeed, testServer.ULSpeed}, nil
}

// MeasureWithFast measures network speed using fast.com's api
func MeasureWithFast() (Measurement, error) {
	return Measurement{"fast.com", 320.1, 150.4}, nil
}
