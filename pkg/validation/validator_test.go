package validation

import (
	"testing"
)

func TestValidatorInterface(t *testing.T) {
	t.Run("shouldImplementValidatorInterface", func(t *testing.T) {
		var _ Validator = (*StringLengthValidator)(nil)
	})
}

func TestValidationResult(t *testing.T) {
	t.Run("shouldCreateValidationResult", func(t *testing.T) {
		result := ValidationResult{
			IsValid: true,
			Message: "valid input",
		}
		
		if !result.IsValid {
			t.Errorf("expected IsValid to be true, got %v", result.IsValid)
		}
		
		if result.Message != "valid input" {
			t.Errorf("expected Message to be 'valid input', got %v", result.Message)
		}
	})
	
	t.Run("shouldCreateInvalidValidationResult", func(t *testing.T) {
		result := ValidationResult{
			IsValid: false,
			Message: "invalid input",
		}
		
		if result.IsValid {
			t.Errorf("expected IsValid to be false, got %v", result.IsValid)
		}
		
		if result.Message != "invalid input" {
			t.Errorf("expected Message to be 'invalid input', got %v", result.Message)
		}
	})
}