package fetcher

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/contracts"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/dom"
)

type HTTPStaticFetcher struct {
	Client *http.Client
}

func (fetcher *HTTPStaticFetcher) Fetch(url string, ctx context.Context) (contracts.HTMLPage, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := fetcher.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	s := string(body)

	return &dom.DOMPage{
		URL:    url,
		Source: &s,
	}, nil
}
