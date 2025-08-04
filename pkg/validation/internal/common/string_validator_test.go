package common

import (
	"math/rand"
	"strings"
	"testing"
	"testing/quick"
)

func TestStringLengthValidator(t *testing.T) {
	t.Run("最小文字数境界で正しく検証される", func(t *testing.T) {
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
	
	t.Run("最大文字数境界で正しく検証される", func(t *testing.T) {
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
	
	t.Run("範囲内の文字数は有効と判定される", func(t *testing.T) {
		validator := NewStringLengthValidator(3, 10)
		
		result := validator.Validate("hello")
		if !result.IsValid {
			t.Errorf("expected validation to pass for string within valid range")
		}
	})
	
	t.Run("文字列以外の入力は無効と判定される", func(t *testing.T) {
		validator := NewStringLengthValidator(3, 10)
		
		result := validator.Validate(123)
		if result.IsValid {
			t.Errorf("expected validation to fail for non-string input")
		}
	})
}

func TestStringLengthValidatorProperties(t *testing.T) {
	t.Run("プロパティ_範囲内の文字数は必ず有効と判定される", func(t *testing.T) {
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
	
	t.Run("プロパティ_最小文字数未満は必ず無効と判定される", func(t *testing.T) {
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
	
	t.Run("プロパティ_最大文字数超過は必ず無効と判定される", func(t *testing.T) {
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
	
	t.Run("プロパティ_同じ入力に対する検証結果は常に同一である", func(t *testing.T) {
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