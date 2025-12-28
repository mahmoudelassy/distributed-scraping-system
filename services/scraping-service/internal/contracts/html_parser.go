package contracts

import "github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/domain"

type HTMLParser interface {
	Parse(page *domain.Page) (domain.Document, error)
}
