package fetcher

import (
	"fmt"
	"io"
	"net/http"

	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/core/html"
)

type HTTPStaticFetcher struct {
	Client *http.Client
}

func (fetcher *HTTPStaticFetcher) Fetch(url string) (*html.Page, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
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

	return &html.Page{
		URL:    url,
		Source: &s,
	}, nil
}
