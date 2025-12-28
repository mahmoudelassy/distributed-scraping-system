package goquery

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/domain"
)

type Document struct {
	document *goquery.Document
}

func (d *Document) Get(selector string) (domain.HTMLElement, error) {
	el := d.document.Find(selector).First()
	if el.Length() == 0 {
		return nil, fmt.Errorf("element not found: %s", selector)
	}
	return NewElement(el), nil
}

func (d *Document) GetAll(selector string) ([]domain.HTMLElement, error) {
	elements := make([]domain.HTMLElement, 0)
	d.document.Find(selector).Each(func(i int, s *goquery.Selection) {
		elements = append(elements, NewElement(s))
	})
	if len(elements) == 0 {
		return nil, fmt.Errorf("elements not found: %s", selector)
	}
	return elements, nil
}
