package validation

import "golang-validator/pkg/validation/internal/common"

// ValidationResult represents the result of a validation operation
type ValidationResult struct {
	IsValid bool
	Message string
}

// Validator is the interface that all validators must implement
type Validator interface {
	Validate(value interface{}) ValidationResult
}

// NewEmailValidator creates a new EmailValidator
func NewEmailValidator() Validator {
	return &emailValidatorWrapper{
		inner: common.NewEmailValidator(),
	}
}

// NewNumberRangeValidator creates a new NumberRangeValidator
func NewNumberRangeValidator(min, max float64) Validator {
	return &numberValidatorWrapper{
		inner: common.NewNumberRangeValidator(min, max),
	}
}

// NewStringLengthValidator creates a new StringLengthValidator
func NewStringLengthValidator(minLength, maxLength int) Validator {
	return &stringValidatorWrapper{
		inner: common.NewStringLengthValidator(minLength, maxLength),
	}
}

// emailValidatorWrapper wraps the internal EmailValidator
type emailValidatorWrapper struct {
	inner *common.EmailValidator
}

func (w *emailValidatorWrapper) Validate(value interface{}) ValidationResult {
	result := w.inner.Validate(value)
	return ValidationResult{
		IsValid: result.IsValid,
		Message: result.Message,
	}
}

// numberValidatorWrapper wraps the internal NumberRangeValidator
type numberValidatorWrapper struct {
	inner *common.NumberRangeValidator
}

func (w *numberValidatorWrapper) Validate(value interface{}) ValidationResult {
	result := w.inner.Validate(value)
	return ValidationResult{
		IsValid: result.IsValid,
		Message: result.Message,
	}
}

// stringValidatorWrapper wraps the internal StringLengthValidator
type stringValidatorWrapper struct {
	inner *common.StringLengthValidator
}

func (w *stringValidatorWrapper) Validate(value interface{}) ValidationResult {
	result := w.inner.Validate(value)
	return ValidationResult{
		IsValid: result.IsValid,
		Message: result.Message,
	}
}