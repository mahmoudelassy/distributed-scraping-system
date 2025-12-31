package service

import (
	"time"
)

type ScrapeJobMessage struct {
	JobID         string     `json:"job_id"`
	UserID        string     `json:"user_id"`
	CorrelationID string     `json:"correlation_id"`
	RequestedAt   time.Time  `json:"requested_at"`
	URL           string     `json:"url"`
	Fetcher       string     `json:"fetcher"`
	GroupLabel    string     `json:"group_label"`
	Queries       []QueryDTO `json:"queries"`
}

type ScrapeResultMessage struct {
	JobID         string      `json:"job_id"`
	UserID        string      `json:"user_id"`
	CorrelationID string      `json:"correlation_id"`
	RequestedAt   time.Time   `json:"requested_at"`
	URL           string      `json:"url"`
	GroupLabel    string      `json:"group_label"`
	Status        string      `json:"status"` // "MATCHED" | "EMPTY"
	ProcessedAt   time.Time   `json:"processed_at"`
	DurationMS    int64       `json:"duration_ms"`
	Results       []ResultDTO `json:"results,omitempty"`
	Error         *ErrorDTO   `json:"error,omitempty"`
}

type QueryDTO struct {
	Selector string `json:"selector"`
	Label    string `json:"label"`
	All      bool   `json:"all"`
}

type ResultDTO struct {
	Selector string       `json:"selector"`
	Label    string       `json:"label"`
	Status   string       `json:"status"`
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
