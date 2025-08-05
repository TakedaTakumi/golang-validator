package validation

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
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
		
		assert.True(t, result.IsValid, "expected IsValid to be true")
		
		assert.Equal(t, "valid input", result.Message, "expected Message to be 'valid input'")
	})
	
	t.Run("無効な検証結果が正しく作成される", func(t *testing.T) {
		result := ValidationResult{
			IsValid: false,
			Message: "invalid input",
		}
		
		assert.False(t, result.IsValid, "expected IsValid to be false")
		
		assert.Equal(t, "invalid input", result.Message, "expected Message to be 'invalid input'")
	})
}