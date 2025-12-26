package worker

import "github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/domain"

type JobStatus string

const (
	JobSuccess JobStatus = "SUCCESS"
	JobFailed  JobStatus = "FAILED"
)

type JobMetadata struct {
	JobID         string
	UserID        string
	CorrelationID string
}

type JobConfig struct {
	URL        string `json:"url" validate:"required,url"`
	Fetcher    string `json:"fetcher" validate:"required,fetcher"`
	GroupLabel string `json:"group_label" validate:"required"`
}
type Job struct {
	Metadata JobMetadata `json:"metadata" validate:"required,dive"`
	Config   JobConfig
	Queries  []domain.Query `json:"queries" validate:"required,min=1,dive"`
}

type JobResult struct {
	MetaData *JobMetadata
	Config   *JobConfig
	Status   JobStatus
	Results  []domain.Result
}
