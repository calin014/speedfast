package speedfast

import (
	"errors"
	"testing"
	"time"
)

func TestCreateUploadPayload(t *testing.T) {
	payload := createUploadPayload()

	if len(payload) != 26214400 {
		t.Fatal("Payload is not 25MB")
	}
}

func TestMeasureNetworkSpeedWithError(t *testing.T) {
	_, err := measureNetworkSpeed(func(url string) error {
		return errors.New("Some Error")
	}, "http://some.url")

	if err == nil {
		t.Fatal("Should get an error")
	}
}

func TestMeasureNetworkSpeed(t *testing.T) {
	result, err := measureNetworkSpeed(func(url string) error {
		time.Sleep(100 * time.Millisecond)
		return nil
	}, "http://some.url")

	if err != nil {
		t.Fatal("Should not get an error")
	}

	if result == 0 {
		t.Fatal("Should have some result")
	}

	t.Log("Speed:", result)
}

func TestCalculateSpeed(t *testing.T) {
	sTime := time.Now()
	fTime := sTime.Add(1 * time.Second)

	res := calculateSpeed(sTime, fTime)

	if res != 1600 {
		t.Fatal("Speed should be 1600Mbps")
	}
}
