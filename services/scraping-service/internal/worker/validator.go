package worker

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/mahmoudelassy/distributed-scraping-system/services/scraping-service/internal/service"
)

// WorkerValidator holds the validator instance
type WorkerValidator struct {
	validator *validator.Validate
}

func NewWorkerValidator() *WorkerValidator {
	v := validator.New()

	v.RegisterValidation("fetcher", func(fl validator.FieldLevel) bool {
		val := fl.Field().String()
		return val == "http" || val == "chromedp"
	})

	return &WorkerValidator{validator: v}
}

// Validate Kafka message
func (wv *WorkerValidator) ValidateJobMessage(msg service.ScrapeJobMessage) error {
	// Basic validation
	if msg.JobID == "" {
		return fmt.Errorf("job_id is required")
	}

	if msg.UserID == "" {
		return fmt.Errorf("user_id is required")
	}

	if msg.CorrelationID == "" {
		return fmt.Errorf("correlation_id is required")
	}

	if msg.URL == "" {
		return fmt.Errorf("url is required")
	}

	// Validate fetcher type
	if msg.Fetcher != "http" && msg.Fetcher != "chromedp" {
		return fmt.Errorf("fetcher must be 'http' or 'chromedp', got '%s'", msg.Fetcher)
	}

	// Validate queries
	if len(msg.Queries) == 0 {
		return fmt.Errorf("at least one query is required")
	}

	for i, query := range msg.Queries {
		if query.Selector == "" {
			return fmt.Errorf("query[%d].selector is required", i)
		}
		if query.Label == "" {
			return fmt.Errorf("query[%d].label is required", i)
		}
	}

	return nil
}

// Alternative: Use struct tags if you prefer
func (wv *WorkerValidator) ValidateJobMessageWithTags(msg service.ScrapeJobMessage) error {
	err := wv.validator.Struct(msg)
	if err != nil {
		// Format validation errors nicely
		validationErrors := err.(validator.ValidationErrors)
		for _, e := range validationErrors {
			fmt.Printf("Field '%s' failed validation rule '%s'\n", e.Field(), e.Tag())
		}
		return fmt.Errorf("validation failed: %v", err)
	}
	return nil
}
