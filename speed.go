package speedfast

import (
	"errors"

	"github.com/showwin/speedtest-go/speedtest"
)

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

	return Measurement{"speedtest.net", testServer.DLSpeed, testServer.ULSpeed}, nil
}
