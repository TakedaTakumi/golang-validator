package validation

// ValidationResult represents the result of a validation operation
type ValidationResult struct {
	IsValid bool
	Message string
}

// Validator is the interface that all validators must implement
type Validator interface {
	Validate(value interface{}) ValidationResult
}