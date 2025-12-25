package worker

import (
	"context"

	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/contracts"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/logging"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/scraper"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/scraper/ctxmeta"
)

type Worker struct {
	scraper   *scraper.Scraper
	logger    logging.Logger
	fetchers  map[string]contracts.HTMLFetcher
	validator *WorkerValidator
}

func NewWorker() *Worker {

}

func (w *Worker) ProcessJob(ctx context.Context, job Job) error {

	w.logger.Info("Processing job", logging.Field{
		Key:   "job",
		Value: job,
	})

	if err := w.validator.ValidateJob(job); err != nil {
		w.logger.Error("Job validation failed", logging.Field{
			Key:   "error",
			Value: err,
		})
		return err
	}
	ctx = context.WithValue(ctx, ctxmeta.JobIDKey, job.Metadata.JobID)
	ctx = context.WithValue(ctx, ctxmeta.UserIDKey, job.Metadata.UserID)
	ctx = context.WithValue(ctx, ctxmeta.CorrelationIDKey, job.Metadata.CorrelationID)
	//results, err := w.scraper.Scrape(ctx, job.URL, job.Queries, nil)
	return nil
}
