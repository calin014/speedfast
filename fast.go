package speedfast

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
)

// MeasureWithFastInHeadlessBrowser measures network speed by opening fast.com in a headless browser and scraping the page
func MeasureWithFastInHeadlessBrowser() (Measurement, error) {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 50*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var download, upload string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://fast.com`),
		chromedp.WaitVisible(`#show-more-details-link`),
		chromedp.Click(`#show-more-details-link`, chromedp.NodeVisible),
		chromedp.WaitVisible(`#upload-value.succeeded`),
		chromedp.Text(`#speed-value`, &download),
		chromedp.Text(`#upload-value`, &upload),
	)

	fmt.Println(err)
	if err != nil {
		return Measurement{}, err
	}

	downloadf, err := strconv.ParseFloat(download, 64)

	if err != nil {
		return Measurement{}, err
	}

	uploadf, err := strconv.ParseFloat(upload, 64)

	if err != nil {
		return Measurement{}, err
	}

	return Measurement{"fast.com", downloadf, uploadf}, nil
}
