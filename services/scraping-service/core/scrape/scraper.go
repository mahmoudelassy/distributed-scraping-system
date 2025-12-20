package scrape

import "github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/core/contracts"

type Scraper struct {
	Parser  contracts.HTMLParser
	Fetcher contracts.HTMLFetcher
}

func (s *Scraper) initDocument(url string) (contracts.Document, error) {
	page, ferr := s.Fetcher.Fetch(url)
	if ferr != nil {
		return nil, ferr
	}
	doc, perr := s.Parser.Parse(page)
	if perr != nil {
		return nil, perr
	}
	return doc, nil
}

func (s *Scraper) Scrape(url string, queries []Query) ([]Result, error) {

	doc, err := s.initDocument(url)
	if err != nil {
		return nil, err
	}

	results := make([]Result, 0, len(queries))

	for _, query := range queries {
		result := query.Select(doc)
		results = append(results, *result)
	}

	return results, nil
}
