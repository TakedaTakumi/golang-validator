package validation

import (
	"regexp"
	"strings"
)

// EmailValidator validates email address format
type EmailValidator struct {
	emailRegex *regexp.Regexp
}

// NewEmailValidator creates a new EmailValidator
func NewEmailValidator() *EmailValidator {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return &EmailValidator{
		emailRegex: emailRegex,
	}
}

// Validate checks if the input string is a valid email address
func (v EmailValidator) Validate(value interface{}) ValidationResult {
	str, ok := value.(string)
	if !ok {
		return ValidationResult{
			IsValid: false,
			Message: "input must be a string",
		}
	}
	
	if str == "" {
		return ValidationResult{
			IsValid: false,
			Message: "email cannot be empty",
		}
	}
	
	if !v.emailRegex.MatchString(str) {
		return ValidationResult{
			IsValid: false,
			Message: "invalid email format",
		}
	}
	
	if strings.Contains(str, "..") {
		return ValidationResult{
			IsValid: false,
			Message: "consecutive dots not allowed",
		}
	}
	
	parts := strings.Split(str, "@")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return ValidationResult{
			IsValid: false,
			Message: "invalid email format",
		}
	}
	
	return ValidationResult{
		IsValid: true,
		Message: "valid",
	}
}