package goquery

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/domain"
)

type Parser struct{}

func (p *Parser) Parse(page *domain.Page) (domain.Document, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(*page.Source))
	if err != nil {
		return nil, err
	}
	return &Document{
		document: doc,
	}, nil
}
