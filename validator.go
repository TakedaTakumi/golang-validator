package main

import (
	"regexp"
	"strings"
)

type ValidationResult struct {
	IsValid bool
	Message string
}

type Validator interface {
	Validate(value interface{}) ValidationResult
}

type StringLengthValidator struct {
	MinLength int
	MaxLength int
}

func NewStringLengthValidator(minLength, maxLength int) *StringLengthValidator {
	return &StringLengthValidator{
		MinLength: minLength,
		MaxLength: maxLength,
	}
}

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

type NumberRangeValidator struct {
	Min float64
	Max float64
}

func NewNumberRangeValidator(min, max float64) *NumberRangeValidator {
	return &NumberRangeValidator{
		Min: min,
		Max: max,
	}
}

func (v NumberRangeValidator) Validate(value interface{}) ValidationResult {
	var num float64
	var ok bool
	
	switch val := value.(type) {
	case int:
		num = float64(val)
		ok = true
	case float64:
		num = val
		ok = true
	case float32:
		num = float64(val)
		ok = true
	default:
		ok = false
	}
	
	if !ok {
		return ValidationResult{
			IsValid: false,
			Message: "input must be a number",
		}
	}
	
	if num < v.Min {
		return ValidationResult{
			IsValid: false,
			Message: "number is too small",
		}
	}
	
	if num > v.Max {
		return ValidationResult{
			IsValid: false,
			Message: "number is too large",
		}
	}
	
	return ValidationResult{
		IsValid: true,
		Message: "valid",
	}
}

type EmailValidator struct {
	emailRegex *regexp.Regexp
}

func NewEmailValidator() *EmailValidator {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return &EmailValidator{
		emailRegex: emailRegex,
	}
}

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