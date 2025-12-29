package goquery

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/domain"
)

type Element struct {
	Attrs domain.Attributes
	Text  string
}

func NewElement(selectedElement *goquery.Selection) domain.HTMLElement {
	attrs := make(domain.Attributes)
	if len(selectedElement.Nodes) > 0 {
		for _, attr := range selectedElement.Nodes[0].Attr {
			attrs[attr.Key] = attr.Val
		}
	}
	text := selectedElement.Text()

	return &Element{
		Attrs: attrs,
		Text:  text,
	}
}

func (el *Element) GetText() string {
	return el.Text
}

func (el *Element) GetAttributes() domain.Attributes {
	return el.Attrs
}
