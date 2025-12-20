package scrape

import "github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/core/contracts"

type Query struct {
	Selector string
	All      bool
}

func (q *Query) Select(document contracts.Document) *Result {
	result := Result{
		Query:    q,
		Elements: nil,
	}
	if q.All {
		result.Elements = document.GetAll(q.Selector)
	} else {
		result.Elements = []contracts.HTMLElement{document.Get(q.Selector)}
	}
	return &result
}
