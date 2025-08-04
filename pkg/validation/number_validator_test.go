package validation

import (
	"testing"
	"testing/quick"
)

func TestNumberRangeValidator(t *testing.T) {
	t.Run("shouldValidateIntegerMinimum", func(t *testing.T) {
		validator := NewNumberRangeValidator(10, 100)
		
		result := validator.Validate(5)
		if result.IsValid {
			t.Errorf("expected validation to fail for integer below minimum")
		}
		
		result = validator.Validate(10)
		if !result.IsValid {
			t.Errorf("expected validation to pass for integer at minimum")
		}
	})
	
	t.Run("shouldValidateIntegerMaximum", func(t *testing.T) {
		validator := NewNumberRangeValidator(10, 100)
		
		result := validator.Validate(105)
		if result.IsValid {
			t.Errorf("expected validation to fail for integer above maximum")
		}
		
		result = validator.Validate(100)
		if !result.IsValid {
			t.Errorf("expected validation to pass for integer at maximum")
		}
	})
	
	t.Run("shouldValidateFloatMinimum", func(t *testing.T) {
		validator := NewNumberRangeValidator(10.5, 100.5)
		
		result := validator.Validate(9.5)
		if result.IsValid {
			t.Errorf("expected validation to fail for float below minimum")
		}
		
		result = validator.Validate(10.5)
		if !result.IsValid {
			t.Errorf("expected validation to pass for float at minimum")
		}
	})
	
	t.Run("shouldValidateFloatMaximum", func(t *testing.T) {
		validator := NewNumberRangeValidator(10.5, 100.5)
		
		result := validator.Validate(101.0)
		if result.IsValid {
			t.Errorf("expected validation to fail for float above maximum")
		}
		
		result = validator.Validate(100.5)
		if !result.IsValid {
			t.Errorf("expected validation to pass for float at maximum")
		}
	})
	
	t.Run("shouldHandleNonNumberInput", func(t *testing.T) {
		validator := NewNumberRangeValidator(10, 100)
		
		result := validator.Validate("not a number")
		if result.IsValid {
			t.Errorf("expected validation to fail for non-number input")
		}
	})
}

func TestNumberRangeValidatorProperties(t *testing.T) {
	t.Run("propertyIntegersWithinRangeShouldPass", func(t *testing.T) {
		validator := NewNumberRangeValidator(10, 100)
		
		property := func(num int) bool {
			if num < 10 || num > 100 {
				return true
			}
			
			result := validator.Validate(num)
			return result.IsValid
		}
		
		if err := quick.Check(property, nil); err != nil {
			t.Errorf("property failed: %v", err)
		}
	})
	
	t.Run("propertyFloatsWithinRangeShouldPass", func(t *testing.T) {
		validator := NewNumberRangeValidator(10.5, 100.5)
		
		property := func(num float64) bool {
			if num < 10.5 || num > 100.5 {
				return true
			}
			
			result := validator.Validate(num)
			return result.IsValid
		}
		
		if err := quick.Check(property, nil); err != nil {
			t.Errorf("property failed: %v", err)
		}
	})
	
	t.Run("propertyNumbersBelowMinShouldFail", func(t *testing.T) {
		validator := NewNumberRangeValidator(50, 100)
		
		property := func(num int) bool {
			if num >= 50 {
				return true
			}
			
			result := validator.Validate(num)
			return !result.IsValid
		}
		
		if err := quick.Check(property, nil); err != nil {
			t.Errorf("property failed: %v", err)
		}
	})
	
	t.Run("propertyNumbersAboveMaxShouldFail", func(t *testing.T) {
		validator := NewNumberRangeValidator(10, 50)
		
		property := func(num int) bool {
			if num <= 50 {
				return true
			}
			
			result := validator.Validate(num)
			return !result.IsValid
		}
		
		if err := quick.Check(property, nil); err != nil {
			t.Errorf("property failed: %v", err)
		}
	})
	
	t.Run("propertyValidationResultsAreConsistent", func(t *testing.T) {
		validator := NewNumberRangeValidator(10, 100)
		
		property := func(num int) bool {
			result1 := validator.Validate(num)
			result2 := validator.Validate(num)
			
			return result1.IsValid == result2.IsValid
		}
		
		if err := quick.Check(property, nil); err != nil {
			t.Errorf("property failed: %v", err)
		}
	})
}