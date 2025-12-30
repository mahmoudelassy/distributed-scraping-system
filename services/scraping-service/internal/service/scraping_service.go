package service

import (
	"context"
	"time"

	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/contracts"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/domain"
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

func (s *ScrapingService) ProcessMessage(msg ScrapeJobMessage) ScrapeResultMessage {
	startTime := time.Now()

	s.logger.Info("processing scrape job",
		logging.Field{Key: "job", Value: msg})

	// Select fetcher
	fetcher, err := s.fetcherFactory.GetFetcher(msg.Fetcher)
	if err != nil {
		return s.buildErrorResponse(msg, err, startTime)
	}

	// Create scraper
	scraperInstance := &scraper.Scraper{
		Fetcher: fetcher,
		Parser:  s.parser,
		Logger:  s.logger,
	}

	// Convert to internal types
	queries := ToInternalQueries(msg.Queries)

	// Create context with metadata
	ctx := ctxmeta.WithJobMetadata(
		context.Background(),
		msg.JobID,
		msg.UserID,
		msg.CorrelationID,
	)

	// Execute scraping
	results, err := scraperInstance.Scrape(ctx, msg.URL, queries, nil)
	if err != nil {
		return s.buildErrorResponse(msg, err, startTime)
	}

	// Build success response
	return s.buildSuccessResponse(msg, results, startTime)
}

// Build success response preserving request context
func (s *ScrapingService) buildSuccessResponse(msg ScrapeJobMessage, results []*domain.Result, startTime time.Time) ScrapeResultMessage {
	duration := time.Since(startTime)
	resultDTOs := ToResultDTOs(results)

	res := ScrapeResultMessage{
		// Preserve original request context
		JobID:         msg.JobID,
		UserID:        msg.UserID,
		CorrelationID: msg.CorrelationID,
		RequestedAt:   msg.RequestedAt,
		URL:           msg.URL,
		GroupLabel:    msg.GroupLabel,

		// Processing metadata
		Status:      "COMPLETED",
		ProcessedAt: time.Now(),
		DurationMS:  duration.Milliseconds(),

		// Results
		Results: resultDTOs,
		Error:   nil,
	}
	s.logger.Info("scrape job completed successfully",
		logging.Field{Key: "job", Value: msg},
		logging.Field{Key: "result", Value: res})

	return res
}

// Build error response preserving request context
func (s *ScrapingService) buildErrorResponse(
	msg ScrapeJobMessage,
	err error,
	startTime time.Time,
) ScrapeResultMessage {
	duration := time.Since(startTime)

	res := ScrapeResultMessage{
		// Preserve original request context
		JobID:         msg.JobID,
		UserID:        msg.UserID,
		CorrelationID: msg.CorrelationID,
		RequestedAt:   msg.RequestedAt,
		URL:           msg.URL,
		GroupLabel:    msg.GroupLabel,

		// Processing metadata
		Status:      "FAILED",
		ProcessedAt: time.Now(),
		DurationMS:  duration.Milliseconds(),

		// Error
		Results: nil,
		Error:   ToErrorDTO(err),
	}
	s.logger.Error("scrape job failed",
		logging.Field{Key: "job", Value: msg},
		logging.Field{Key: "result", Value: res})

	return res
}
