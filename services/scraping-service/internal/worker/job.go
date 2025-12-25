package worker

import "github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/domain"

type JobMetadata struct {
	JobID         string `json:"job_id" validate:"required"`
	UserID        string `json:"user_id" validate:"required"`
	CorrelationID string `json:"correlation_id" validate:"required"`
}

type Job struct {
	Metadata   JobMetadata    `json:"metadata" validate:"required,dive"`
	URL        string         `json:"url" validate:"required,url"`
	PageType   string         `json:"page_type" validate:"required,page_type"` // custom validator
	GroupLabel string         `json:"group_label" validate:"required"`
	Queries    []domain.Query `json:"queries" validate:"required,min=1,dive"`
}
