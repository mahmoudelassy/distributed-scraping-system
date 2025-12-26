package service

import (
	"fmt"

	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/contracts"
)

type FetcherFactory struct {
	httpFetcher     contracts.HTMLFetcher
	chromedpFetcher contracts.HTMLFetcher
}

func NewFetcherFactory(
	httpFetcher contracts.HTMLFetcher,
	chromedpFetcher contracts.HTMLFetcher,
) *FetcherFactory {
	return &FetcherFactory{
		httpFetcher:     httpFetcher,
		chromedpFetcher: chromedpFetcher,
	}
}

func (f *FetcherFactory) GetFetcher(fetcherType string) (contracts.HTMLFetcher, error) {
	switch fetcherType {
	case "http":
		if f.httpFetcher == nil {
			return nil, fmt.Errorf("http fetcher not configured")
		}
		return f.httpFetcher, nil

	case "chromedp":
		if f.chromedpFetcher == nil {
			return nil, fmt.Errorf("chromedp fetcher not configured")
		}
		return f.chromedpFetcher, nil

	default:
		return nil, fmt.Errorf("unknown fetcher type: %s", fetcherType)
	}
}
