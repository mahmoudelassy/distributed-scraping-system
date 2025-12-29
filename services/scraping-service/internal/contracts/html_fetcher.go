package contracts

import "github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/domain"

type HTMLFetcher interface {
	Fetch(url string) (*domain.Page, error)
}
