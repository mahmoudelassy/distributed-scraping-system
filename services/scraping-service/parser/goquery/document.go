package goquery

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/core/contracts"
)

type Document struct {
	document *goquery.Document
}

func (d *Document) Get(selector string) (contracts.HTMLElement, error) {
	el := d.document.Find(selector).First()
	if el.Length() == 0 {
		return nil, fmt.Errorf("element not found: %s", selector)
	}
	return NewElement(el), nil
}

func (d *Document) GetAll(selector string) []contracts.HTMLElement {
	elements := make([]contracts.HTMLElement, 0)
	d.document.Find(selector).Each(func(i int, s *goquery.Selection) {
		elements = append(elements, NewElement(s))
	})
	return elements
}
