package common

// StringLengthValidator validates string length within specified range
type StringLengthValidator struct {
	MinLength int
	MaxLength int
}

// NewStringLengthValidator creates a new StringLengthValidator
func NewStringLengthValidator(minLength, maxLength int) *StringLengthValidator {
	return &StringLengthValidator{
		MinLength: minLength,
		MaxLength: maxLength,
	}
}

// Validate checks if the input string length is within the specified range
func (v StringLengthValidator) Validate(value interface{}) ValidationResult {
	str, ok := value.(string)
	if !ok {
		return ValidationResult{
			IsValid: false,
			Message: "input must be a string",
		}
	}
	
	length := len(str)
	if length < v.MinLength {
		return ValidationResult{
			IsValid: false,
			Message: "string is too short",
		}
	}
	
	if length > v.MaxLength {
		return ValidationResult{
			IsValid: false,
			Message: "string is too long",
		}
	}
	
	return ValidationResult{
		IsValid: true,
		Message: "valid",
	}
}