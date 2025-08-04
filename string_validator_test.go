package main

import (
	"math/rand"
	"strings"
	"testing"
	"testing/quick"
)

func TestStringLengthValidator(t *testing.T) {
	t.Run("shouldValidateMinimumLength", func(t *testing.T) {
		validator := NewStringLengthValidator(3, 10)
		
		result := validator.Validate("ab")
		if result.IsValid {
			t.Errorf("expected validation to fail for string shorter than minimum length")
		}
		
		result = validator.Validate("abc")
		if !result.IsValid {
			t.Errorf("expected validation to pass for string at minimum length")
		}
	})
	
	t.Run("shouldValidateMaximumLength", func(t *testing.T) {
		validator := NewStringLengthValidator(3, 10)
		
		result := validator.Validate("12345678901")
		if result.IsValid {
			t.Errorf("expected validation to fail for string longer than maximum length")
		}
		
		result = validator.Validate("1234567890")
		if !result.IsValid {
			t.Errorf("expected validation to pass for string at maximum length")
		}
	})
	
	t.Run("shouldValidateWithinRange", func(t *testing.T) {
		validator := NewStringLengthValidator(3, 10)
		
		result := validator.Validate("hello")
		if !result.IsValid {
			t.Errorf("expected validation to pass for string within valid range")
		}
	})
	
	t.Run("shouldHandleNonStringInput", func(t *testing.T) {
		validator := NewStringLengthValidator(3, 10)
		
		result := validator.Validate(123)
		if result.IsValid {
			t.Errorf("expected validation to fail for non-string input")
		}
	})
}

func TestStringLengthValidatorProperties(t *testing.T) {
	t.Run("propertyValidStringsWithinRangeShouldPass", func(t *testing.T) {
		validator := NewStringLengthValidator(5, 15)
		
		property := func(length uint8) bool {
			if length < 5 || length > 15 {
				return true
			}
			
			testString := generateStringOfLength(int(length))
			result := validator.Validate(testString)
			return result.IsValid
		}
		
		if err := quick.Check(property, nil); err != nil {
			t.Errorf("property failed: %v", err)
		}
	})
	
	t.Run("propertyStringsShorterThanMinShouldFail", func(t *testing.T) {
		validator := NewStringLengthValidator(5, 15)
		
		property := func(length uint8) bool {
			if length >= 5 {
				return true
			}
			
			testString := generateStringOfLength(int(length))
			result := validator.Validate(testString)
			return !result.IsValid
		}
		
		if err := quick.Check(property, nil); err != nil {
			t.Errorf("property failed: %v", err)
		}
	})
	
	t.Run("propertyStringsLongerThanMaxShouldFail", func(t *testing.T) {
		validator := NewStringLengthValidator(5, 15)
		
		property := func(length uint8) bool {
			if length <= 15 {
				return true
			}
			
			testString := generateStringOfLength(int(length))
			result := validator.Validate(testString)
			return !result.IsValid
		}
		
		if err := quick.Check(property, nil); err != nil {
			t.Errorf("property failed: %v", err)
		}
	})
	
	t.Run("propertyValidationResultsAreConsistent", func(t *testing.T) {
		validator := NewStringLengthValidator(3, 10)
		
		property := func(input string) bool {
			result1 := validator.Validate(input)
			result2 := validator.Validate(input)
			
			return result1.IsValid == result2.IsValid
		}
		
		if err := quick.Check(property, nil); err != nil {
			t.Errorf("property failed: %v", err)
		}
	})
}

func generateStringOfLength(length int) string {
	if length == 0 {
		return ""
	}
	
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result strings.Builder
	result.Grow(length)
	
	for i := 0; i < length; i++ {
		result.WriteByte(chars[rand.Intn(len(chars))])
	}
	
	return result.String()
}