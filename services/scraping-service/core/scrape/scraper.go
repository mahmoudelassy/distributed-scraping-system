package scrape

import (
	"time"

	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/core/contracts"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/core/html"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/utils"
)

type Scraper struct {
	Parser  contracts.HTMLParser
	Fetcher contracts.HTMLFetcher
}

func (s *Scraper) initDocument(url string, retries int, delay time.Duration) (contracts.Document, error) {

	page, ferr := utils.Retry(func() (*html.Page, error) {
		return s.Fetcher.Fetch(url)
	}, retries, delay)

	if ferr != nil {
		return nil, ferr
	}

	doc, perr := utils.Retry(func() (contracts.Document, error) {
		return s.Parser.Parse(page)
	}, retries, delay)

	if perr != nil {
		return nil, perr
	}

	return doc, nil
}

func (s *Scraper) Scrape(url string, queries []Query) ([]*Result, error) {

	doc, err := s.initDocument(url, 3, 200)

	if err != nil {
		return nil, err
	}

	results := make([]*Result, 0, len(queries))

	for _, query := range queries {
		result := query.Select(doc)
		results = append(results, result)
	}

	return results, nil
}
