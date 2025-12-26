package service

import (
	"context"
	"time"

	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/contracts"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/logging"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/scraper"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/scraper/ctxmeta"
)

type ScrapingService struct {
	fetcherFactory *FetcherFactory
	parser         contracts.HTMLParser
	logger         logging.Logger
}

func NewScrapingService(
	factory *FetcherFactory,
	parser contracts.HTMLParser,
	logger logging.Logger,
) *ScrapingService {
	return &ScrapingService{
		fetcherFactory: factory,
		parser:         parser,
		logger:         logger,
	}
}

func (s *ScrapingService) Process(req ScrapeRequest) ScrapeResponse {
	startTime := time.Now()

	s.logger.Info("processing scrape request",
		logging.Field{Key: "url", Value: req.URL},
		logging.Field{Key: "fetcher", Value: req.Fetcher})

	// Select fetcher
	fetcher, err := s.fetcherFactory.GetFetcher(req.Fetcher)
	if err != nil {
		return s.errorResponse(req, err, startTime)
	}

	// Create scraper with selected fetcher
	scraperInstance := &scraper.Scraper{
		Fetcher: fetcher,
		Parser:  s.parser,
		Logger:  s.logger,
	}

	// Convert DTOs to internal types
	queries := ToInternalQueries(req.Queries)

	// Create context with metadata
	ctx := ctxmeta.WithJobMetadata(
		context.Background(),
		req.Metadata.JobID,
		req.Metadata.UserID,
		req.Metadata.CorrelationID,
	)

	// Execute scraping
	results, err := scraperInstance.Scrape(ctx, req.URL, queries, nil)
	if err != nil {
		return s.errorResponse(req, err, startTime)
	}

	// Convert results to DTOs
	resultDTOs := ToResultDTOs(results)

	duration := time.Since(startTime)

	s.logger.Info("scrape request completed",
		logging.Field{Key: "duration_ms", Value: duration.Milliseconds()},
		logging.Field{Key: "result_count", Value: len(resultDTOs)})

	return ScrapeResponse{
		Success:    true,
		Metadata:   req.Metadata,
		Results:    resultDTOs,
		Error:      nil,
		DurationMS: duration.Milliseconds(),
		Timestamp:  time.Now(),
	}
}

func (s *ScrapingService) errorResponse(req ScrapeRequest, err error, startTime time.Time) ScrapeResponse {
	duration := time.Since(startTime)

	s.logger.Error("scrape request failed",
		logging.Field{Key: "error", Value: err.Error()},
		logging.Field{Key: "duration_ms", Value: duration.Milliseconds()})

	return ScrapeResponse{
		Success:    false,
		Metadata:   req.Metadata,
		Results:    nil,
		Error:      ToErrorDTO(err),
		DurationMS: duration.Milliseconds(),
		Timestamp:  time.Now(),
	}
}
