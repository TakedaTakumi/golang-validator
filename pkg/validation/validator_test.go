package validation

import (
	"testing"
)

func TestValidatorInterface(t *testing.T) {
	t.Run("Validatorインターフェースを実装している", func(t *testing.T) {
		emailValidator := NewEmailValidator()
		numberValidator := NewNumberRangeValidator(0, 100)
		stringValidator := NewStringLengthValidator(1, 10)
		
		var _ Validator = emailValidator
		var _ Validator = numberValidator
		var _ Validator = stringValidator
	})
}

func TestValidationResult(t *testing.T) {
	t.Run("有効な検証結果が正しく作成される", func(t *testing.T) {
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
	
	t.Run("無効な検証結果が正しく作成される", func(t *testing.T) {
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