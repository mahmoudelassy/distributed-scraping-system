package contracts

import "github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/core/html"

type HTMLFetcher interface {
	Fetch(url string) (*html.Page, error)
}
