// internal/worker/worker.go
package worker

import (
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/domain"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/logging"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/service"
)

type Worker struct {
	service   *service.ScrapingService
	logger    logging.Logger
	validator *WorkerValidator
}

func NewWorker(
	service *service.ScrapingService,
	logger logging.Logger,
) *Worker {
	return &Worker{
		service:   service,
		logger:    logger,
		validator: NewWorkerValidator(),
	}
}

func (w *Worker) ProcessJob(job Job) JobResult {
	w.logger.Info("Processing job", logging.Field{
		Key:   "job_id",
		Value: job.Metadata.JobID,
	})

	// Validate
	if err := w.validator.ValidateJob(job); err != nil {
		w.logger.Error("Job validation failed", logging.Field{
			Key:   "error",
			Value: err,
		})
		return JobResult{
			MetaData: &job.Metadata,
			Config:   &job.Config,
			Status:   JobFailed,
			Results:  nil,
		}
	}

	// Convert Job to Service Request
	req := service.ScrapeRequest{
		Metadata: job.Metadata,
		URL:      job.Config.URL,
		Fetcher:  job.Config.Fetcher,
		Queries:  ToQueryDTOs(job.Queries),
	}

	// Call Service
	response := w.service.Process(req)

	// Convert response to JobResult
	if !response.Success {
		return JobResult{
			MetaData: &job.Metadata,
			Config:   &job.Config,
			Status:   JobFailed,
			Results:  nil,
		}
	}

	// Convert DTOs back to domain Results
	domainResults := make([]domain.Result, len(response.Results))
	for i, dto := range response.Results {
		domainResults[i] = domain.Result{
			// Map fields from DTO
		}
	}

	return JobResult{
		MetaData: &job.Metadata,
		Config:   &job.Config,
		Status:   JobSuccess,
		Results:  domainResults,
	}
}

func ToQueryDTOs(queries []domain.Query) []service.QueryDTO {
	dtos := make([]service.QueryDTO, len(queries))
	for i, q := range queries {
		dtos[i] = service.QueryDTO{
			Selector: q.Selector,
			Label:    q.Label,
			All:      q.All,
		}
	}
	return dtos
}
