package service

import (
	"time"

	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/worker"
)

// Request from queue
type ScrapeRequest struct {
	Metadata   worker.JobMetadata `json:"metadata"`
	URL        string             `json:"url"`
	Fetcher    string             `json:"fetcher"` // "http" or "chromedp"
	GroupLabel string
	Queries    []QueryDTO `json:"queries"`
}

type QueryDTO struct {
	Selector string `json:"selector"`
	Label    string `json:"label"`
	All      bool   `json:"all"`
}

// Response to queue
type ScrapeResponse struct {
	Success    bool               `json:"success"`
	Metadata   worker.JobMetadata `json:"metadata"`
	Results    []ResultDTO        `json:"results,omitempty"`
	Error      *ErrorDTO          `json:"error,omitempty"`
	DurationMS int64              `json:"duration_ms"`
	Timestamp  time.Time          `json:"timestamp"`
}

type ResultDTO struct {
	Selector string       `json:"selector"`
	Label    string       `json:"label"`
	Success  bool         `json:"success"`
	Elements []ElementDTO `json:"elements"`
}

type ElementDTO struct {
	Text       string            `json:"text"`
	Attributes map[string]string `json:"attributes"`
}

type ErrorDTO struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Stage   string `json:"stage,omitempty"`
}
