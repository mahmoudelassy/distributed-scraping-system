package goquery

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/core/contracts"
)

type Document struct {
	document *goquery.Document
}

func (d *Document) Get(selector string) contracts.HTMLElement {
	el := d.document.Find(selector).First()
	return NewElement(el)
}

func (d *Document) GetAll(selector string) []contracts.HTMLElement {
	elements := make([]contracts.HTMLElement, 0)
	d.document.Find(selector).Each(func(i int, s *goquery.Selection) {
		element := NewElement(s)
		elements = append(elements, element)
	})
	return elements
}
