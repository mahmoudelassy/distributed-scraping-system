package scrape

import (
	"time"

	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/contracts"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/domain"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/logging"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/utils"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/worker"
)

type Scraper struct {
	Parser  contracts.HTMLParser
	Fetcher contracts.HTMLFetcher
	Logger  logging.Logger
}

type ScraperOptions struct {
	Retries int
	Delay   time.Duration
	Ctx     *worker.JobMetadata
}

func (s *Scraper) initDocument(url string, opts *ScraperOptions) (domain.Document, error) {
	opts = s.ensureOptions(opts)

	meta := map[string]interface{}{
		"url":            url,
		"job_id":         opts.Ctx.JobID,
		"correlation_id": opts.Ctx.CorrelationID,
		"user_id":        opts.Ctx.UserID,
	}

	s.Logger.Info(
		"starting fetch",
		logging.Field{Key: "meta", Value: meta},
	)

	page, ferr := utils.Retry(func() (*domain.Page, error) {
		return s.Fetcher.Fetch(url)
	}, opts.Retries, opts.Delay)

	if ferr != nil {
		s.Logger.Error(
			"failed to fetch URL",
			logging.Field{Key: "meta", Value: meta},
			logging.Field{Key: "error", Value: ferr},
		)
		return nil, &domain.ScraperError{Stage: "fetch", Err: ferr}
	}

	s.Logger.Info(
		"successfully fetched URL",
		logging.Field{Key: "meta", Value: meta},
	)

	doc, perr := utils.Retry(func() (domain.Document, error) {
		return s.Parser.Parse(page)
	}, opts.Retries, opts.Delay)

	if perr != nil {
		s.Logger.Error(
			"failed to parse document",
			logging.Field{Key: "meta", Value: meta},
			logging.Field{Key: "error", Value: perr},
		)
		return nil, &domain.ScraperError{Stage: "parse", Err: perr}
	}

	s.Logger.Info(
		"successfully parsed document",
		logging.Field{Key: "meta", Value: meta},
	)

	return doc, nil
}

func (s *Scraper) Scrape(url string, queries []domain.Query, opts *ScraperOptions) ([]*domain.Result, error) {
	opts = s.ensureOptions(opts)

	meta := map[string]interface{}{
		"url":            url,
		"job_id":         opts.Ctx.JobID,
		"correlation_id": opts.Ctx.CorrelationID,
		"user_id":        opts.Ctx.UserID,
		"num_queries":    len(queries),
	}

	s.Logger.Info("starting scrape", logging.Field{Key: "meta", Value: meta})

	doc, err := s.initDocument(url, opts)
	if err != nil {
		s.Logger.Error(
			"scrape failed",
			logging.Field{Key: "meta", Value: meta},
			logging.Field{Key: "error", Value: err},
		)
		return nil, err
	}

	results := s.executeQueries(doc, queries, meta)

	s.Logger.Info("scrape completed", logging.Field{Key: "meta", Value: meta})

	return results, nil
}

func (s *Scraper) executeQueries(doc domain.Document, queries []domain.Query, meta map[string]interface{}) []*domain.Result {
	results := make([]*domain.Result, 0, len(queries))

	for _, query := range queries {
		result := query.Select(doc)
		results = append(results, result)

		s.Logger.Info(
			"query executed",
			logging.Field{Key: "selector", Value: query.Selector},
			logging.Field{Key: "label", Value: query.Label},
			logging.Field{Key: "elements_found", Value: len(result.Elements)},
			logging.Field{Key: "meta", Value: meta},
		)
	}

	return results
}

func (s *Scraper) ensureOptions(opts *ScraperOptions) *ScraperOptions {
	if opts == nil {
		opts = &ScraperOptions{}
	}

	if opts.Ctx == nil {
		opts.Ctx = &worker.JobMetadata{}
	}

	if opts.Retries == 0 {
		opts.Retries = 3
	}

	if opts.Delay == 0 {
		opts.Delay = 200 * time.Millisecond
	}

	return opts
}
