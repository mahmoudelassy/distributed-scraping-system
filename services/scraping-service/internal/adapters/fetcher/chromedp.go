package fetcher

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/domain"
)

type ChromeDPFetcher struct {
	allocatorCtx context.Context
	cancel       context.CancelFunc
	timeout      time.Duration
}

func NewChromeDPFetcherRemote(wsURL string, timeout time.Duration) (*ChromeDPFetcher, error) {
	allocCtx, cancel := chromedp.NewRemoteAllocator(context.Background(), wsURL)
	return &ChromeDPFetcher{
		allocatorCtx: allocCtx,
		cancel:       cancel,
		timeout:      timeout,
	}, nil
}

func (f *ChromeDPFetcher) Fetch(url string) (*domain.Page, error) {
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

	return &domain.Page{
		URL:    url,
		Source: &source,
	}, nil
}
