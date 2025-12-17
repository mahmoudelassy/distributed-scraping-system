package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/contracts"
	domgoquery "github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/dom/goquery"
)

type Goquery struct {
	parsedPage *goquery.Document
}

func NewParser(page contracts.HTMLPage) *Goquery {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(*page.GetSource()))
	if err != nil {
		panic("error in parsing")
	}
	return &Goquery{
		parsedPage: doc,
	}
}

func (p *Goquery) GetElement(selector string) contracts.HTMLElement {
	el := p.parsedPage.Find(selector).First()
	return domgoquery.NewElement(el)
}

func (p *Goquery) GetAllElements(selector string) []contracts.HTMLElement {
	elements := make([]contracts.HTMLElement, 0)
	p.parsedPage.Find(selector).Each(func(i int, s *goquery.Selection) {
		element := domgoquery.NewElement(s)
		elements = append(elements, element)

	})
	return elements
}
