package utils

import (
	"time"
)

func Retry[T any](
	fn func() (T, error),
	retries int,
	delay time.Duration,
) (T, error) {
	var result T
	var err error

	for attempt := 1; attempt <= retries; attempt++ {
		result, err = fn()
		if err == nil {
			return result, nil
		}
		time.Sleep(delay)
	}

	return result, err
}
