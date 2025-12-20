package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/core/scrape"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/fetcher"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/parser/goquery"
)

func main() {
	//f, _ := fetcher.NewChromeDPFetcher(true, 4*30*time.Second)
	s := scrape.Scraper{
		Fetcher: &fetcher.HTTPStaticFetcher{
			Client: http.DefaultClient,
		},
		Parser: &goquery.Parser{},
	}
	for i := range 10 {
		go func() {

			fmt.Println(i)
			results, _ := s.Scrape(
				"https://wuzzuf.net/a/IT-Software-Development-Jobs-in-Egypt?ref=browse-jobs",
				[]scrape.Query{{Selector: "#app > div > div > div > div > div > div > div > div > h2 > a", All: true}})
			for _, res := range results {
				for _, el := range res.Elements {
					fmt.Println(el.GetText())
				}
			}
		}()
	}
	time.Sleep(60 * time.Second)

}
