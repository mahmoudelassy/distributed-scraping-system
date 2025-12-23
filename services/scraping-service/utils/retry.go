package utils

import (
	"fmt"
	"time"
)

func Retry[T any](fn func() (T, error), retries int, delay time.Duration) (T, error) {
	var result T
	var err error
	for attempt := 1; attempt <= retries; attempt++ {
		result, err = fn()
		if err == nil {
			return result, nil
		}
		fmt.Printf("Attempt %d failed: %v\n", attempt, err)
		time.Sleep(delay)
	}
	return result, fmt.Errorf("all %d attempts failed: %w", retries, err)

}
