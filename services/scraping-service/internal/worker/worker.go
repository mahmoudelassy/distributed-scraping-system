package worker

import (
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/logging"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/service"
)

type Worker struct {
	service   *service.ScrapingService
	logger    logging.Logger
	validator *WorkerValidator
}

func NewWorker(svc *service.ScrapingService, logger logging.Logger,
) *Worker {
	return &Worker{
		service:   svc,
		logger:    logger,
		validator: NewWorkerValidator(),
	}
}

func (w *Worker) ProcessMessage(msg service.ScrapeJobMessage) service.ScrapeResultMessage {
	w.logger.Info("received scrape job",
		logging.Field{Key: "job", Value: msg})

	// Validate
	if err := w.validator.ValidateJobMessage(msg); err != nil {
		w.logger.Error("job validation failed",
			logging.Field{Key: "job", Value: msg})

		// Return error response
		return service.ScrapeResultMessage{
			JobID:         msg.JobID,
			UserID:        msg.UserID,
			CorrelationID: msg.CorrelationID,
			RequestedAt:   msg.RequestedAt,
			URL:           msg.URL,
			GroupLabel:    msg.GroupLabel,
			Status:        "FAILED",
			Error: &service.ErrorDTO{
				Type:    "validation_error",
				Message: err.Error(),
			},
		}
	}

	// Process via service
	return w.service.ProcessMessage(msg)
}
