package goquery

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/core/contracts"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/core/html"
)

type Parser struct{}

func (p *Parser) Parse(page *html.Page) (contracts.Document, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(*page.Source))
	if err != nil {
		return nil, err
	}
	return &Document{
		document: doc,
	}, nil
}
