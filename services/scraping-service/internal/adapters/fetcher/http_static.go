package fetcher

import (
	"fmt"
	"io"
	"net/http"

	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/domain"
)

type HTTPStaticFetcher struct {
	Client *http.Client
}

func (f *HTTPStaticFetcher) Fetch(url string) (*domain.Page, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := f.Client.Do(req)
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

	return &domain.Page{
		URL:    url,
		Source: &s,
	}, nil
}
