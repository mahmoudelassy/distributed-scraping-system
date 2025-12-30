package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/adapters/fetcher"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/adapters/parser/goquery"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/logging"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/service"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/worker"
)

func main() {
	zLogger, _ := logging.NewZapLogger()
	c, _ := fetcher.NewChromeDPFetcherRemote("http://localhost:9222", 4*30*time.Second)
	h := &fetcher.HTTPStaticFetcher{
		Client: http.DefaultClient,
	}
	ff := service.NewFetcherFactory(h, c)

	svc := service.NewScrapingService(ff, &goquery.Parser{}, zLogger)

	w := worker.NewWorker(svc, zLogger)

	jsonData := []byte(`{
  "job_id": "job-123",
  "user_id": "user-456",
  "correlation_id": "corr-789",
  "requested_at": "2025-01-01T12:00:00Z",
  "url": "https://wuzzuf.net/a/IT-Software-Development-Jobs-in-Egypt?ref=browse-jobs",
  "fetcher": "http",
  "group_label": "wuzzuf-jobs",
  "queries": [
    {
      "selector": "#app > div > div > div > div > div > div > div > div > h2 > a",
      "label": "wuzzuf_link",
	  "all":false
    }
  ]
}`)

	var msg service.ScrapeJobMessage
	err := json.Unmarshal(jsonData, &msg)
	if err != nil {
		fmt.Println(err)
	}

	w.ProcessMessage(msg)

}
