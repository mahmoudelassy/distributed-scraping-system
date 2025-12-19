package fetcher

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/core/html"
)

type ChromeDPFetcher struct {
	allocatorCtx context.Context
	cancel       context.CancelFunc
	timeout      time.Duration
}

func NewChromeDPFetcher(headless bool, timeout time.Duration) (*ChromeDPFetcher, error) {
	chromedp.ExecPath("/snap/bin/chromium")
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", headless),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)

	return &ChromeDPFetcher{
		allocatorCtx: allocCtx,
		cancel:       cancel,
		timeout:      timeout,
	}, nil
}

func (f *ChromeDPFetcher) Fetch(url string) (*html.Page, error) {
	ctx, cancel := chromedp.NewContext(f.allocatorCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, f.timeout)
	defer cancel()

	var source string

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.OuterHTML("html", &source, chromedp.ByQuery),
	)
	if err != nil {
		return nil, err
	}

	return &html.Page{
		URL:    url,
		Source: &source,
	}, nil
}
