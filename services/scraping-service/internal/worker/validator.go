package worker

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// WorkerValidator holds the validator instance
type WorkerValidator struct {
	validator *validator.Validate
}

// NewWorkerValidator creates a new validator with custom rules registered
func NewWorkerValidator() *WorkerValidator {
	v := validator.New()

	// Custom validation for PageType
	v.RegisterValidation("page_type", func(fl validator.FieldLevel) bool {
		val := fl.Field().String()
		return val == "static" || val == "dynamic"
	})

	return &WorkerValidator{validator: v}
}

func (wv *WorkerValidator) ValidateJob(job Job) error {
	err := wv.validator.Struct(job)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Printf("Field '%s' failed validation rule '%s'\n", e.Field(), e.Tag())
		}
		return err
	}
	return nil
}
