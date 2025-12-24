package scraper

import (
	"fmt"
	"time"

	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/adapters/fetcher"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/adapters/parser/goquery"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/domain"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/logging"
	scrape "github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/scraper"
)

func main() {
	zLogger, err := logging.NewZapLogger()
	if err != nil {
		panic(err) // handle error properly in production
	}
	f, _ := fetcher.NewChromeDPFetcherRemote("http://localhost:9222", 4*30*time.Second)

	s := scrape.Scraper{
		Fetcher: f,
		Parser:  &goquery.Parser{},
		Logger:  zLogger,
	}
	for i := range 10 {
		go func() {
			fmt.Println(i)
			results, _ := s.Scrape(
				"https://wuzzuf.net/a/IT-Software-Development-Jobs-in-Egypt?ref=browse-jobs",
				[]domain.Query{{Selector: "#app > div > div > div > div > div > div > div > div > h2 > a", All: false}}, nil)
			for _, res := range results {
				for _, el := range res.Elements {
					fmt.Println(el.GetText())
				}
			}
		}()
	}
	time.Sleep(60 * time.Second)

}
