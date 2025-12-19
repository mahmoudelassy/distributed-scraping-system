package scrape

import "github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/core/contracts"

type Result struct {
	Query    *Query
	Elements []contracts.HTMLElement
}
