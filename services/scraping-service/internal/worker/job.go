package worker

import "github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/domain"

type JobMetadata struct {
	JobID         string
	UserID        string
	CorrelationID string
}

type Job struct {
	Metadata JobMetadata
	URL      string
	PageType string
	Label    string
	Queries  []domain.Query
}
