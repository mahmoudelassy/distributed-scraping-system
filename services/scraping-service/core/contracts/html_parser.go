package contracts

import "github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/core/html"

type HTMLParser interface {
	Parse(page *html.Page) (Document, error)
}
