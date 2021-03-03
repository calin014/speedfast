package speedfast

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gesquive/fast-cli/fast"
	"golang.org/x/sync/errgroup"
)

// MeasureWithFast naively attempts to measure network speed by using fast.com's api directly
// because fast-cli and go-fast libraries provide only download speed
func MeasureWithFast() (Measurement, error) {
	url := fast.GetDlUrls(1)[0]

	downloadSpeed, err := measureNetworkSpeed(download, url)
	uploadSpeed, err := measureNetworkSpeed(upload, url)

	if err != nil {
		return Measurement{}, err
	}

	return Measurement{"fast.com", downloadSpeed, uploadSpeed}, nil
}

var client = http.Client{}

const (
	workload      = 8
	payloadSizeMB = 25.0 // download payload is by debault 25MB, make upload 25MB also
)

func measureNetworkSpeed(operation func(url string) error, url string) (float64, error) {
	eg := errgroup.Group{}

	sTime := time.Now()
	for i := 0; i < workload; i++ {
		eg.Go(func() error {
			return operation(url)
		})
	}
	if err := eg.Wait(); err != nil {
		return 0, err
	}
	fTime := time.Now()

	return payloadSizeMB * 8 * float64(workload) / fTime.Sub(sTime).Seconds(), nil
}

func measureUploadSpeed(url string) float64 {
	return 0
}

func download(url string) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	ioutil.ReadAll(resp.Body)

	return nil
}

func upload(uri string) error {
	v := url.Values{}

	//10b * x MB / 10 = x MB
	v.Add("content", strings.Repeat("0123456789", payloadSizeMB*1024*1024/10))

	resp, err := client.PostForm(uri, v)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	ioutil.ReadAll(resp.Body)

	return nil
}
